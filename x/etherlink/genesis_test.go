package etherlink_test

import (
	"testing"

	keepertest "etherlink/testutil/keeper"
	"etherlink/testutil/nullify"
	"etherlink/x/etherlink"
	"etherlink/x/etherlink/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EthState: &types.EthState{},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EtherlinkKeeper(t)
	etherlink.InitGenesis(ctx, *k, genesisState)
	got := etherlink.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.EthState, got.EthState)
	// this line is used by starport scaffolding # genesis/test/assert
}
