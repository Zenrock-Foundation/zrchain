package keeper

import (
	"context"

	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

func (k msgServer) RemoveMultiGrant(goCtx context.Context, msg *types.MsgRemoveMultiGrant) (*types.MsgRemoveMultiGrantResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	for _, m := range msg.Msgs {
		_, err := k.authzKeeper.Revoke(goCtx, &authztypes.MsgRevoke{
			Granter:    msg.Creator,
			Grantee:    msg.Grantee,
			MsgTypeUrl: m,
		})
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgRemoveMultiGrantResponse{}, nil
}
