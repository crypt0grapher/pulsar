package etherlink

import (
	"math/rand"

	"etherlink/testutil/sample"
	etherlinksimulation "etherlink/x/etherlink/simulation"
	"etherlink/x/etherlink/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = etherlinksimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateEthState = "op_weight_msg_eth_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEthState int = 100

	opWeightMsgUpdateEthState = "op_weight_msg_eth_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEthState int = 100

	opWeightMsgDeleteEthState = "op_weight_msg_eth_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEthState int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	etherlinkGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&etherlinkGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateEthState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEthState, &weightMsgCreateEthState, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEthState = defaultWeightMsgCreateEthState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEthState,
		etherlinksimulation.SimulateMsgCreateEthState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEthState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateEthState, &weightMsgUpdateEthState, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEthState = defaultWeightMsgUpdateEthState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEthState,
		etherlinksimulation.SimulateMsgUpdateEthState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEthState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteEthState, &weightMsgDeleteEthState, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEthState = defaultWeightMsgDeleteEthState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEthState,
		etherlinksimulation.SimulateMsgDeleteEthState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEthState,
			defaultWeightMsgCreateEthState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherlinksimulation.SimulateMsgCreateEthState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEthState,
			defaultWeightMsgUpdateEthState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherlinksimulation.SimulateMsgUpdateEthState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEthState,
			defaultWeightMsgDeleteEthState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherlinksimulation.SimulateMsgDeleteEthState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
