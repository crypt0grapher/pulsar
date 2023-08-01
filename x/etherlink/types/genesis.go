package types

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EthState: nil,
		EthInput: &EthInput{
			Creator:    "genesis",
			EthAddress: "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
			EthSlot:    "0x0",
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
