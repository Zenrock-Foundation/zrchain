package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/types"
)

func SimulateMsgTransferFromKeyring(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTransferFromKeyring{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TransferFromKeyring simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "TransferFromKeyring simulation not implemented"), nil, nil
	}
}
