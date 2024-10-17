package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/types"
)

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidArgument, "invalid arguments: request is nil")
	}

	params, err := k.ParamStore.Get(goCtx)
	if err != nil {
		return nil, nil
	}

	return &types.QueryParamsResponse{Params: params}, nil
}
