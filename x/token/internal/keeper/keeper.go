package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/contract"
	"github.com/line/link/x/token/internal/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	iamKeeper      types.IamKeeper
	storeKey       sdk.StoreKey
	contractKeeper contract.Keeper
	cdc            *codec.Codec
}

func NewKeeper(cdc *codec.Codec, iamKeeper types.IamKeeper, contractKeeper contract.Keeper, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		iamKeeper:      iamKeeper.WithPrefix(types.ModuleName),
		storeKey:       storeKey,
		contractKeeper: contractKeeper,
		cdc:            cdc,
	}
}

func (k Keeper) NewContractID(ctx sdk.Context) string {
	return k.contractKeeper.NewContractID(ctx)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) UnmarshalJSON(bz []byte, ptr interface{}) error {
	return k.cdc.UnmarshalJSON(bz, ptr)
}

func (k Keeper) MarshalJSON(o interface{}) ([]byte, error) {
	return k.cdc.MarshalJSON(o)
}

func (k Keeper) MarshalJSONIndent(o interface{}) ([]byte, error) {
	return k.cdc.MarshalJSONIndent(o, "", "  ")
}
