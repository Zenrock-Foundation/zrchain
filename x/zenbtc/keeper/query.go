package keeper

import (
	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
)

var _ types.QueryServer = Keeper{}
