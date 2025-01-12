package rune

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "torram/api/torram/rune"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetStakedRune",
					Use:            "get-staked-rune [rune-id]",
					Short:          "Query get-staked-rune",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "runeId"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "StakeRune",
					Use:            "stake-rune [rune-id] [owner] [amount]",
					Short:          "Send a stake-rune tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "runeId"}, {ProtoField: "owner"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "UnstakeRune",
					Use:            "unstake-rune [rune-id] [owner] [amount]",
					Short:          "Send a unstake-rune tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "runeId"}, {ProtoField: "owner"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "UpdateRuneStatus",
					Use:            "update-rune-status [rune-id] [status]",
					Short:          "Send a update-rune-status tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "runeId"}, {ProtoField: "status"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
