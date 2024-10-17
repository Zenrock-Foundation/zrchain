package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	validationtypes "github.com/zenrocklabs/zenrock/zrchain/v4/x/validation/types"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/zenbtc/types"
)

func (k msgServer) SubmitSolanaUnlockTransaction(goCtx context.Context, msg *types.MsgSubmitSolanaUnlockTransaction) (*types.MsgSubmitSolanaUnlockTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	found, err := k.validationKeeper.UnconfirmedSolanaUnlockTxs.Has(ctx, msg.TxSignature)
	if err != nil {
		return nil, err
	}
	if found {
		return nil, fmt.Errorf("unconfirmed solana unlock transaction with signature %s already exists", msg.TxSignature)
	}

	found, err = k.validationKeeper.ConfirmedSolanaUnlockTxs.Has(ctx, msg.TxSignature)
	if err != nil {
		return nil, err
	}
	if found {
		return nil, fmt.Errorf("confirmed solana unlock transaction with signature %s already exists", msg.TxSignature)
	}

	if err := k.validationKeeper.UnconfirmedSolanaUnlockTxs.Set(ctx, msg.TxSignature, validationtypes.WithdrawalInfo{
		WithdrawalAddress: msg.WithdrawalAddr,
		Amount:            msg.Amount,
		RetryCount:        0,
	}); err != nil {
		return nil, err
	}

	return &types.MsgSubmitSolanaUnlockTransactionResponse{}, nil
}
