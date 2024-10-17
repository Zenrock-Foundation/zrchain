package keeper

import (
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/zenbtc/types"
)

var _ types.QueryServer = Keeper{}
