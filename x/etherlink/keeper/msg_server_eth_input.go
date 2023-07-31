package keeper

import (
	"context"

	"etherlink/x/etherlink/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateEthInput(goCtx context.Context, msg *types.MsgCreateEthInput) (*types.MsgCreateEthInputResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetEthInput(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var ethInput = types.EthInput{
		Creator: msg.Creator,
	}

	k.SetEthInput(
		ctx,
		ethInput,
	)
	return &types.MsgCreateEthInputResponse{}, nil
}

func (k msgServer) UpdateEthInput(goCtx context.Context, msg *types.MsgUpdateEthInput) (*types.MsgUpdateEthInputResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEthInput(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var ethInput = types.EthInput{
		Creator: msg.Creator,
	}

	k.SetEthInput(ctx, ethInput)

	return &types.MsgUpdateEthInputResponse{}, nil
}

func (k msgServer) DeleteEthInput(goCtx context.Context, msg *types.MsgDeleteEthInput) (*types.MsgDeleteEthInputResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetEthInput(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveEthInput(ctx)

	return &types.MsgDeleteEthInputResponse{}, nil
}
