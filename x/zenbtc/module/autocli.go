package zenbtc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/Zenrock-Foundation/zrchain/v4/api/zrchain/zenbtc"
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
					RpcMethod:      "LockTransactions",
					Use:            "lock-transactions",
					Short:          "Query LockTransactions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ConfirmedUnlockTransactions",
					Use:            "confirmed-unlock-transactions",
					Short:          "Query ConfirmedUnlockTransactions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod: "VerifyDepositBlockInclusion",
					Use:       "verify-block-inclusion [chain_name] [block_height] [raw_tx] [index] [proof] [deposit_addr] [amount]",
					Short:     "Send a VerifyDepositBlockInclusion tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "chain_name"},
						{ProtoField: "block_height"},
						{ProtoField: "raw_tx"},
						{ProtoField: "index"},
						{ProtoField: "proof"},
						{ProtoField: "deposit_addr"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "SubmitUnlockTransaction",
					Use:       "submit-unlock-transaction",
					Short:     "Send a SubmitUnlockTransaction tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "chain"},
						{ProtoField: "txID"},
						{ProtoField: "withdrawalAddr"},
						{ProtoField: "amount"},
					},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
