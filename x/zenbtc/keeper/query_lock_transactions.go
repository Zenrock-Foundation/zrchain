package keeper

import (
	"context"

	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LockTransactions(goCtx context.Context, req *types.QueryLockTransactionsRequest) (*types.QueryLockTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	lockTransactions := []string{}
	if err := k.LockTransactionStore.Walk(ctx, nil, func(key string) (bool, error) {
		lockTransactions = append(lockTransactions, key)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &types.QueryLockTransactionsResponse{LockTransactions: lockTransactions}, nil
}
