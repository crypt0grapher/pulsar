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

func TestEthInputQuery(t *testing.T) {
	keeper, ctx := keepertest.EtherlinkKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestEthInput(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetEthInputRequest
		response *types.QueryGetEthInputResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEthInputRequest{},
			response: &types.QueryGetEthInputResponse{EthInput: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EthInput(wctx, tc.request)
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
