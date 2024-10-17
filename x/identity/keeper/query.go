package keeper

import (
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/identity/types"
)

var _ types.QueryServer = Keeper{}
