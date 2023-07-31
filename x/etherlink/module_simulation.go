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
	opWeightMsgCreateEthInput = "op_weight_msg_eth_input"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEthInput int = 100

	opWeightMsgUpdateEthInput = "op_weight_msg_eth_input"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEthInput int = 100

	opWeightMsgDeleteEthInput = "op_weight_msg_eth_input"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEthInput int = 100

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

	var weightMsgCreateEthInput int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEthInput, &weightMsgCreateEthInput, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEthInput = defaultWeightMsgCreateEthInput
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEthInput,
		etherlinksimulation.SimulateMsgCreateEthInput(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEthInput int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateEthInput, &weightMsgUpdateEthInput, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEthInput = defaultWeightMsgUpdateEthInput
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEthInput,
		etherlinksimulation.SimulateMsgUpdateEthInput(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEthInput int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteEthInput, &weightMsgDeleteEthInput, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEthInput = defaultWeightMsgDeleteEthInput
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEthInput,
		etherlinksimulation.SimulateMsgDeleteEthInput(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEthInput,
			defaultWeightMsgCreateEthInput,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherlinksimulation.SimulateMsgCreateEthInput(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEthInput,
			defaultWeightMsgUpdateEthInput,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherlinksimulation.SimulateMsgUpdateEthInput(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEthInput,
			defaultWeightMsgDeleteEthInput,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				etherlinksimulation.SimulateMsgDeleteEthInput(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
