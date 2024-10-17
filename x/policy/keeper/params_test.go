package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

func TestGetParams(t *testing.T) {
	keepers := keepertest.NewTest(t)
	pk := keepers.PolicyKeeper
	ctx := keepers.Ctx
	params := types.DefaultParams()

	require.NoError(t, pk.ParamStore.Set(ctx, params))
	got, err := pk.ParamStore.Get(ctx)
	require.NoError(t, err)
	require.EqualValues(t, params, got)
}
