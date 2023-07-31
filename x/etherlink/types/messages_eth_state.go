package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateEthState = "create_eth_state"
	TypeMsgUpdateEthState = "update_eth_state"
	TypeMsgDeleteEthState = "delete_eth_state"
)

var _ sdk.Msg = &MsgCreateEthState{}

func NewMsgCreateEthState(creator string) *MsgCreateEthState {
	return &MsgCreateEthState{
		Creator: creator,
	}
}

func (msg *MsgCreateEthState) Route() string {
	return RouterKey
}

func (msg *MsgCreateEthState) Type() string {
	return TypeMsgCreateEthState
}

func (msg *MsgCreateEthState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEthState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEthState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateEthState{}

func NewMsgUpdateEthState(creator string) *MsgUpdateEthState {
	return &MsgUpdateEthState{
		Creator: creator,
	}
}

func (msg *MsgUpdateEthState) Route() string {
	return RouterKey
}

func (msg *MsgUpdateEthState) Type() string {
	return TypeMsgUpdateEthState
}

func (msg *MsgUpdateEthState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateEthState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateEthState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteEthState{}

func NewMsgDeleteEthState(creator string) *MsgDeleteEthState {
	return &MsgDeleteEthState{
		Creator: creator,
	}
}
func (msg *MsgDeleteEthState) Route() string {
	return RouterKey
}

func (msg *MsgDeleteEthState) Type() string {
	return TypeMsgDeleteEthState
}

func (msg *MsgDeleteEthState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteEthState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteEthState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
