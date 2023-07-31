package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEthInput = "create_eth_input"
	TypeMsgUpdateEthInput = "update_eth_input"
	TypeMsgDeleteEthInput = "delete_eth_input"
)

var _ sdk.Msg = &MsgCreateEthInput{}

func NewMsgCreateEthInput(creator string) *MsgCreateEthInput {
	return &MsgCreateEthInput{
		Creator: creator,
	}
}

func (msg *MsgCreateEthInput) Route() string {
	return RouterKey
}

func (msg *MsgCreateEthInput) Type() string {
	return TypeMsgCreateEthInput
}

func (msg *MsgCreateEthInput) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEthInput) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEthInput) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEthInput{}

func NewMsgUpdateEthInput(creator string) *MsgUpdateEthInput {
	return &MsgUpdateEthInput{
		Creator: creator,
	}
}

func (msg *MsgUpdateEthInput) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEthInput) Type() string {
	return TypeMsgUpdateEthInput
}

func (msg *MsgUpdateEthInput) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEthInput) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEthInput) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEthInput{}

func NewMsgDeleteEthInput(creator string) *MsgDeleteEthInput {
	return &MsgDeleteEthInput{
		Creator: creator,
	}
}
func (msg *MsgDeleteEthInput) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEthInput) Type() string {
	return TypeMsgDeleteEthInput
}

func (msg *MsgDeleteEthInput) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEthInput) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEthInput) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
