package keeper

import (
	"context"

	"etherlink/x/etherlink/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EthInput(goCtx context.Context, req *types.QueryGetEthInputRequest) (*types.QueryGetEthInputResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEthInput(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEthInputResponse{EthInput: val}, nil
}
