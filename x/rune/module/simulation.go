package rune

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"torram/testutil/sample"
	runesimulation "torram/x/rune/simulation"
	"torram/x/rune/types"
)

// avoid unused import issue
var (
	_ = runesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgStakeRune = "op_weight_msg_stake_rune"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStakeRune int = 100

	opWeightMsgUnstakeRune = "op_weight_msg_unstake_rune"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnstakeRune int = 100

	opWeightMsgUpdateRuneStatus = "op_weight_msg_update_rune_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRuneStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	runeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&runeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgStakeRune int
	simState.AppParams.GetOrGenerate(opWeightMsgStakeRune, &weightMsgStakeRune, nil,
		func(_ *rand.Rand) {
			weightMsgStakeRune = defaultWeightMsgStakeRune
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStakeRune,
		runesimulation.SimulateMsgStakeRune(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnstakeRune int
	simState.AppParams.GetOrGenerate(opWeightMsgUnstakeRune, &weightMsgUnstakeRune, nil,
		func(_ *rand.Rand) {
			weightMsgUnstakeRune = defaultWeightMsgUnstakeRune
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnstakeRune,
		runesimulation.SimulateMsgUnstakeRune(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRuneStatus int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateRuneStatus, &weightMsgUpdateRuneStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRuneStatus = defaultWeightMsgUpdateRuneStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRuneStatus,
		runesimulation.SimulateMsgUpdateRuneStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgStakeRune,
			defaultWeightMsgStakeRune,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				runesimulation.SimulateMsgStakeRune(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnstakeRune,
			defaultWeightMsgUnstakeRune,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				runesimulation.SimulateMsgUnstakeRune(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRuneStatus,
			defaultWeightMsgUpdateRuneStatus,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				runesimulation.SimulateMsgUpdateRuneStatus(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
