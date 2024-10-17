package keeper

import (
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

var _ types.QueryServer = Keeper{}
