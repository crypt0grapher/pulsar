package types

import (
	"testing"

	"etherlink/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateEthInput_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateEthInput
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateEthInput{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateEthInput{
				Creator: sample.AccAddress(),
				Address: "test-address",
				Slot:    "test-slot",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateEthInput_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateEthInput
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateEthInput{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateEthInput{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteEthInput_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteEthInput
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteEthInput{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteEthInput{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
