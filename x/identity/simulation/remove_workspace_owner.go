package simulation

import (
	"math/rand"

	"github.com/Zenrock-Foundation/zrchain/v6/x/identity/keeper"
	"github.com/Zenrock-Foundation/zrchain/v6/x/identity/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRemoveWorkspaceOwner(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRemoveWorkspaceOwner{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RemoveWorkspaceOwner simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "RemoveWorkspaceOwner simulation not implemented"), nil, nil
	}
}
