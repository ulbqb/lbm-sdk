package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/line/link/version"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/line/link/app"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/line/link/client"
	"github.com/line/link/types"
)

// linkd custom flags
const (
	flagInvCheckPeriod = "inv-check-period"
	flagTestnet        = "testnet"
)

var invCheckPeriod uint
var testnet bool

func main() {
	cdc := app.MakeCodec()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "linkd",
		Short:             "Link Daemon (server)",
		PersistentPreRunE: LinkPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx, cdc))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		auth.GenesisAccountIterator{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(testnetCmd(ctx, cdc, app.ModuleBasics, auth.GenesisAccountIterator{}))
	rootCmd.AddCommand(replayCmd())
	rootCmd.AddCommand(version.Cmd)

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "GA", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	rootCmd.PersistentFlags().BoolVar(&testnet, flagTestnet, testnet, "Run with testnet mode. The address prefix becomes tlink if this flag is set.")

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func LinkPreRunEFn(context *server.Context) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		f := server.PersistentPreRunEFn(context)
		err := f(cmd, args)

		if cmd.Name() == version.Cmd.Name() {
			return nil
		}
		testnet := viper.GetBool(flagTestnet)
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(types.Bech32PrefixAcc(testnet), types.Bech32PrefixAccPub(testnet))
		config.SetBech32PrefixForValidator(types.Bech32PrefixValAddr(testnet), types.Bech32PrefixValPub(testnet))
		config.SetBech32PrefixForConsensusNode(types.Bech32PrefixConsAddr(testnet), types.Bech32PrefixConsPub(testnet))
		config.SetCoinType(types.CoinType)
		config.SetFullFundraiserPath(types.FullFundraiserPath)
		config.Seal()

		if cmd.Name() == server.StartCmd(nil, nil).Name() {
			var networkMode string
			if testnet {
				networkMode = "testnet"
			} else {
				networkMode = "mainnet"
			}
			context.Logger.Info(fmt.Sprintf("Network mode is %s", networkMode))
			printDBBackend(context)
		}
		return err
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewLinkApp(
		logger, db, traceStore, true, map[int64]bool{}, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		gApp := app.NewLinkApp(logger, db, traceStore, false, map[int64]bool{}, uint(1))
		err := gApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return gApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	gApp := app.NewLinkApp(logger, db, traceStore, true, map[int64]bool{}, uint(1))
	return gApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

func printDBBackend(context *server.Context) {
	var linkDBBackend dbm.BackendType
	if sdk.DBBackend == "" {
		linkDBBackend = dbm.GoLevelDBBackend
	} else {
		linkDBBackend = dbm.BackendType(sdk.DBBackend)
	}
	context.Logger.Info(fmt.Sprintf("LINK DB Backend is %s", linkDBBackend))
	context.Logger.Info(fmt.Sprintf("Tendermint DB Backend is %s", context.Config.DBBackend))
}
