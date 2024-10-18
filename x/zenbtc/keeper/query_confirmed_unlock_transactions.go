package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	validationtypes "github.com/Zenrock-Foundation/zrchain/v4/x/validation/types"
	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
)

func (k Keeper) ConfirmedUnlockTransactions(goCtx context.Context, req *types.QueryConfirmedUnlockTransactionsRequest) (*types.QueryConfirmedUnlockTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	txIDs := make([]string, 0)
	if err := k.validationKeeper.ConfirmedUnlockTxs.Walk(ctx, nil, func(key collections.Pair[string, string], withdrawalInfo validationtypes.WithdrawalInfo) (bool, error) {
		if req.Chain == "" || key.K1() == req.Chain {
			txIDs = append(txIDs, key.K2())
		}
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &types.QueryConfirmedUnlockTransactionsResponse{UnlockTransactions: txIDs}, nil
}
