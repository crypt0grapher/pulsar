package etherlink

import (
	"etherlink/x/etherlink/keeper"
	"etherlink/x/etherlink/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.EthState != nil {
		k.SetEthState(ctx, *genState.EthState)
	}
	// Set if defined
	if genState.EthInput != nil {
		k.SetEthInput(ctx, *genState.EthInput)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all ethState
	ethState, found := k.GetEthState(ctx)
	if found {
		genesis.EthState = &ethState
	}
	// Get all ethInput
	ethInput, found := k.GetEthInput(ctx)
	if found {
		genesis.EthInput = &ethInput
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
