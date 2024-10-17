package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	validationtypes "github.com/zenrocklabs/zenrock/zrchain/v4/x/validation/types"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/zenbtc/types"
)

func (k Keeper) ConfirmedSolanaUnlockTransactions(goCtx context.Context, req *types.QueryConfirmedSolanaUnlockTransactionsRequest) (*types.QueryConfirmedSolanaUnlockTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	txSignatures := make([]string, 0)
	if err := k.validationKeeper.ConfirmedSolanaUnlockTxs.Walk(ctx, nil, func(txSignature string, withdrawalInfo validationtypes.WithdrawalInfo) (bool, error) {
		txSignatures = append(txSignatures, txSignature)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &types.QueryConfirmedSolanaUnlockTransactionsResponse{UnlockTransactions: txSignatures}, nil
}
