package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "etherlink/testutil/keeper"
	"etherlink/testutil/nullify"
	"etherlink/x/etherlink/keeper"
	"etherlink/x/etherlink/types"
)

func createTestEthState(keeper *keeper.Keeper, ctx sdk.Context) types.EthState {
	item := types.EthState{}
	keeper.SetEthState(ctx, item)
	return item
}

func TestEthStateGet(t *testing.T) {
	keeper, ctx := keepertest.EtherlinkKeeper(t)
	item := createTestEthState(keeper, ctx)
	rst, found := keeper.GetEthState(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestEthStateRemove(t *testing.T) {
	keeper, ctx := keepertest.EtherlinkKeeper(t)
	createTestEthState(keeper, ctx)
	keeper.RemoveEthState(ctx)
	_, found := keeper.GetEthState(ctx)
	require.False(t, found)
}
