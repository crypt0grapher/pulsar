package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "etherlink/testutil/keeper"
	"etherlink/testutil/nullify"
	"etherlink/x/etherlink/types"
)

func TestEthStateQuery(t *testing.T) {
	keeper, ctx := keepertest.EtherlinkKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestEthState(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetEthStateRequest
		response *types.QueryGetEthStateResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEthStateRequest{},
			response: &types.QueryGetEthStateResponse{EthState: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EthState(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
