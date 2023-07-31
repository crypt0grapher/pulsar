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

func TestEthStateMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EtherlinkKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateEthState{Creator: creator}
	_, err := srv.CreateEthState(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetEthState(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestEthStateMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateEthState
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateEthState{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateEthState{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EtherlinkKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateEthState{Creator: creator}
			_, err := srv.CreateEthState(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateEthState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetEthState(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestEthStateMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteEthState
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteEthState{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteEthState{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EtherlinkKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateEthState(wctx, &types.MsgCreateEthState{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteEthState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetEthState(ctx)
				require.False(t, found)
			}
		})
	}
}
