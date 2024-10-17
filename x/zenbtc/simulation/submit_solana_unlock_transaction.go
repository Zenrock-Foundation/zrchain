package simulation

import (
	"math/rand"

	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/keeper"
	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSubmitSolanaUnlockTransaction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmitSolanaUnlockTransaction{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SubmitSolanaUnlockTransaction simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "SubmitSolanaUnlockTransaction simulation not implemented"), nil, nil
	}
}
