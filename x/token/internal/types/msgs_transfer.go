package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/contract"
)

var _ contract.Msg = (*MsgTransfer)(nil)

type MsgTransfer struct {
	From       sdk.AccAddress `json:"from"`
	ContractID string         `json:"contract_id"`
	To         sdk.AccAddress `json:"to"`
	Amount     sdk.Int        `json:"amount"`
}

func NewMsgTransfer(from sdk.AccAddress, to sdk.AccAddress, contractID string, amount sdk.Int) MsgTransfer {
	return MsgTransfer{From: from, To: to, ContractID: contractID, Amount: amount}
}

func (msg MsgTransfer) Route() string { return RouterKey }

func (msg MsgTransfer) Type() string { return "transfer_ft" }

func (msg MsgTransfer) ValidateBasic() sdk.Error {
	if err := contract.ValidateContractIDBasic(msg); err != nil {
		return err
	}

	if msg.From.Empty() {
		return sdk.ErrInvalidAddress("from cannot be empty")
	}

	if msg.To.Empty() {
		return sdk.ErrInvalidAddress("to cannot be empty")
	}

	if !msg.Amount.IsPositive() {
		return sdk.ErrInsufficientCoins("send amount must be positive")
	}
	return nil
}

func (msg MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgTransfer) GetContractID() string {
	return msg.ContractID
}
