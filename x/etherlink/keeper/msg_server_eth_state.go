package keeper

import (
	"context"

	"etherlink/x/etherlink/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateEthState(goCtx context.Context, msg *types.MsgCreateEthState) (*types.MsgCreateEthStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetEthState(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var ethState = types.EthState{
		Creator: msg.Creator,
	}

	k.SetEthState(
		ctx,
		ethState,
	)
	return &types.MsgCreateEthStateResponse{}, nil
}

func (k msgServer) UpdateEthState(goCtx context.Context, msg *types.MsgUpdateEthState) (*types.MsgUpdateEthStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEthState(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var ethState = types.EthState{
		Creator: msg.Creator,
	}

	k.SetEthState(ctx, ethState)

	return &types.MsgUpdateEthStateResponse{}, nil
}

func (k msgServer) DeleteEthState(goCtx context.Context, msg *types.MsgDeleteEthState) (*types.MsgDeleteEthStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEthState(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEthState(ctx)

	return &types.MsgDeleteEthStateResponse{}, nil
}
