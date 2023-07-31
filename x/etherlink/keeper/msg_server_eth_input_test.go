package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "etherlink/testutil/keeper"
	"etherlink/x/etherlink/keeper"
	"etherlink/x/etherlink/types"
)

func TestEthInputMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EtherlinkKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateEthInput{Creator: creator}
	_, err := srv.CreateEthInput(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetEthInput(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestEthInputMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateEthInput
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateEthInput{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEthInput{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EtherlinkKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateEthInput{Creator: creator}
			_, err := srv.CreateEthInput(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateEthInput(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetEthInput(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestEthInputMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteEthInput
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteEthInput{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteEthInput{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EtherlinkKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateEthInput(wctx, &types.MsgCreateEthInput{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteEthInput(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetEthInput(ctx)
				require.False(t, found)
			}
		})
	}
}
