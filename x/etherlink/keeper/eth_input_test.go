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

func createTestEthInput(keeper *keeper.Keeper, ctx sdk.Context) types.EthInput {
	item := types.EthInput{}
	keeper.SetEthInput(ctx, item)
	return item
}

func TestEthInputGet(t *testing.T) {
	keeper, ctx := keepertest.EtherlinkKeeper(t)
	item := createTestEthInput(keeper, ctx)
	rst, found := keeper.GetEthInput(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestEthInputRemove(t *testing.T) {
	keeper, ctx := keepertest.EtherlinkKeeper(t)
	createTestEthInput(keeper, ctx)
	keeper.RemoveEthInput(ctx)
	_, found := keeper.GetEthInput(ctx)
	require.False(t, found)
}
