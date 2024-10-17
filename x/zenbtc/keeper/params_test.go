package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/zenbtc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ZenbtcKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
