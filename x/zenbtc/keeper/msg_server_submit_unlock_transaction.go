package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	validationtypes "github.com/Zenrock-Foundation/zrchain/v4/x/validation/types"
	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitUnlockTransaction(goCtx context.Context, msg *types.MsgSubmitUnlockTransaction) (*types.MsgSubmitUnlockTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	found, err := k.validationKeeper.UnconfirmedUnlockTxs.Has(ctx, collections.Join(msg.Chain, msg.TxID))
	if err != nil {
		return nil, err
	}
	if found {
		return nil, fmt.Errorf("unconfirmed unlock transaction on chain %s with txID %s already exists", msg.Chain, msg.TxID)
	}

	found, err = k.validationKeeper.ConfirmedUnlockTxs.Has(ctx, collections.Join(msg.Chain, msg.TxID))
	if err != nil {
		return nil, err
	}
	if found {
		return nil, fmt.Errorf("confirmed unlock transaction on chain %s with txID %s already exists", msg.Chain, msg.TxID)
	}

	if err := k.validationKeeper.UnconfirmedUnlockTxs.Set(ctx, collections.Join(msg.Chain, msg.TxID), validationtypes.WithdrawalInfo{
		WithdrawalAddress: msg.WithdrawalAddr,
		Amount:            msg.Amount,
		RetryCount:        0,
	}); err != nil {
		return nil, err
	}

	return &types.MsgSubmitUnlockTransactionResponse{}, nil
}
