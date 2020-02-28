package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/token/internal/types"
)

func (k Keeper) IssueToken(ctx sdk.Context, token types.Token, amount sdk.Int, owner sdk.AccAddress) sdk.Error {
	if !types.ValidateImageURI(token.GetImageURI()) {
		return types.ErrInvalidImageURILength(types.DefaultCodespace, token.GetImageURI())
	}
	err := k.SetToken(ctx, token)
	if err != nil {
		return err
	}

	err = k.MintSupply(ctx, token.GetContractID(), owner, amount)
	if err != nil {
		return err
	}

	modifyTokenURIPermission := types.NewModifyPermission(token.GetContractID())
	k.AddPermission(ctx, owner, modifyTokenURIPermission)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeIssueToken,
			sdk.NewAttribute(types.AttributeKeyContractID, token.GetContractID()),
			sdk.NewAttribute(types.AttributeKeyName, token.GetName()),
			sdk.NewAttribute(types.AttributeKeySymbol, token.GetSymbol()),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyMintable, strconv.FormatBool(token.GetMintable())),
			sdk.NewAttribute(types.AttributeKeyDecimals, token.GetDecimals().String()),
		),
		sdk.NewEvent(
			types.EventTypeGrantPermToken,
			sdk.NewAttribute(types.AttributeKeyTo, owner.String()),
			sdk.NewAttribute(types.AttributeKeyResource, modifyTokenURIPermission.GetResource()),
			sdk.NewAttribute(types.AttributeKeyAction, modifyTokenURIPermission.GetAction()),
		),
	})

	if token.GetMintable() {
		mintPerm := types.NewMintPermission(token.GetContractID())
		k.AddPermission(ctx, owner, mintPerm)
		burnPerm := types.NewBurnPermission(token.GetContractID())
		k.AddPermission(ctx, owner, burnPerm)
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeGrantPermToken,
				sdk.NewAttribute(types.AttributeKeyTo, owner.String()),
				sdk.NewAttribute(types.AttributeKeyResource, mintPerm.GetResource()),
				sdk.NewAttribute(types.AttributeKeyAction, mintPerm.GetAction()),
			),
			sdk.NewEvent(
				types.EventTypeGrantPermToken,
				sdk.NewAttribute(types.AttributeKeyTo, owner.String()),
				sdk.NewAttribute(types.AttributeKeyResource, burnPerm.GetResource()),
				sdk.NewAttribute(types.AttributeKeyAction, burnPerm.GetAction()),
			),
		})
	}

	return nil
}
