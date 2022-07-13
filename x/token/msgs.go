package token

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

var _ sdk.Msg = (*MsgSend)(nil)

// ValidateBasic implements Msg.
func (m MsgSend) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgSend) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.From)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgOperatorSend)(nil)

// ValidateBasic implements Msg.
func (m MsgOperatorSend) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgOperatorSend) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgTransferFrom)(nil)

// ValidateBasic implements Msg.
func (m MsgTransferFrom) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Proxy); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proxy address: %s", m.Proxy)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgTransferFrom) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Proxy)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgAuthorizeOperator)(nil)

// ValidateBasic implements Msg.
func (m MsgAuthorizeOperator) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid holder address: %s", m.Holder)
	}
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgAuthorizeOperator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgRevokeOperator)(nil)

// ValidateBasic implements Msg.
func (m MsgRevokeOperator) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid holder address: %s", m.Holder)
	}
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgRevokeOperator) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgApprove)(nil)

// ValidateBasic implements Msg.
func (m MsgApprove) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Approver); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid approver address: %s", m.Approver)
	}
	if _, err := sdk.AccAddressFromBech32(m.Proxy); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proxy address: %s", m.Proxy)
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgApprove) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Approver)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgIssue)(nil)

// ValidateBasic implements Msg.
func (m MsgIssue) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", m.Owner)
	}

	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateName(m.Name); err != nil {
		return err
	}

	if err := validateSymbol(m.Symbol); err != nil {
		return err
	}

	if err := validateImageURI(m.ImageUri); err != nil {
		return err
	}

	if err := validateMeta(m.Meta); err != nil {
		return err
	}

	if err := validateDecimals(m.Decimals); err != nil {
		return err
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg.
func (m MsgIssue) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgGrant)(nil)

// ValidateBasic implements Msg.
func (m MsgGrant) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Granter); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid granter address: %s", m.Granter)
	}
	if _, err := sdk.AccAddressFromBech32(m.Grantee); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.Grantee)
	}

	if err := validatePermission(m.Permission); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgGrant) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Granter)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgAbandon)(nil)

// ValidateBasic implements Msg.
func (m MsgAbandon) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Grantee); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.Grantee)
	}

	if err := validatePermission(m.Permission); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgAbandon) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Grantee)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgGrantPermission)(nil)

// ValidateBasic implements Msg.
func (m MsgGrantPermission) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid granter address: %s", m.From)
	}
	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.To)
	}

	if err := validatePermission(m.Permission); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgGrantPermission) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.From)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgRevokePermission)(nil)

// ValidateBasic implements Msg.
func (m MsgRevokePermission) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validatePermission(m.Permission); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgRevokePermission) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.From)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgMint)(nil)

// ValidateBasic implements Msg.
func (m MsgMint) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.From)
	}

	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", m.To)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgMint) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.From)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgBurn)(nil)

// ValidateBasic implements Msg.
func (m MsgBurn) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgBurn) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.From)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgOperatorBurn)(nil)

// ValidateBasic implements Msg.
func (m MsgOperatorBurn) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid operator address: %s", m.Operator)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgOperatorBurn) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Operator)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgBurnFrom)(nil)

// ValidateBasic implements Msg.
func (m MsgBurnFrom) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(m.Proxy); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid proxy address: %s", m.Proxy)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", m.From)
	}

	if err := validateAmount(m.Amount); err != nil {
		return err
	}

	return nil
}

// GetSigners implements Msg
func (m MsgBurnFrom) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Proxy)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = (*MsgModify)(nil)

// ValidateBasic implements Msg.
func (m MsgModify) ValidateBasic() error {
	if err := ValidateContractID(m.ContractId); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", m.Owner)
	}

	checkedFields := map[string]bool{}
	for _, change := range m.Changes {
		if checkedFields[change.Field] {
			return sdkerrors.ErrInvalidRequest.Wrapf("duplicate fields: %s", change.Field)
		}
		checkedFields[change.Field] = true

		if err := validateChange(change); err != nil {
			return err
		}
	}
	if len(checkedFields) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("no field provided")
	}

	return nil
}

// GetSigners implements Msg
func (m MsgModify) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil { // should never happen as valid basic rejects invalid addresses
		panic(err.Error())
	}
	return []sdk.AccAddress{signer}
}
