package keeper_test

import (
	"testing"

	testkeeper "etherlink/testutil/keeper"
	"etherlink/x/etherlink/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EtherlinkKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
