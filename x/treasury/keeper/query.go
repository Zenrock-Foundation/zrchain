package keeper

import (
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/types"
)

var _ types.QueryServer = Keeper{}
