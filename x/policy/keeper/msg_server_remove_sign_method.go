package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

func (k msgServer) RemoveSignMethod(goCtx context.Context, msg *types.MsgRemoveSignMethod) (*types.MsgRemoveSignMethodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signMethod, err := k.GetSignMethod(ctx, msg.GetCreator(), msg.GetId())
	if err != nil {
		return nil, err
	}

	if signMethod == nil {
		return nil, fmt.Errorf("signmethod not found")
	}

	signMethod.SetActive(false)

	err = k.SetSignMethod(ctx, msg.GetCreator(), msg.GetId(), signMethod)

	return &types.MsgRemoveSignMethodResponse{}, err
}
