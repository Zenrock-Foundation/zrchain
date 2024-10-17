package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/types"
)

func SimulateMsgUpdateKeyPolicy(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateKeyPolicy{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateKeyPolicy simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "UpdateKeyPolicy simulation not implemented"), nil, nil
	}
}
