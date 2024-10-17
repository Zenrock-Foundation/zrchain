package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

func TestParamsQuery(t *testing.T) {
	keepers := keepertest.NewTest(t)
	pk := keepers.PolicyKeeper
	ctx := keepers.Ctx
	params := types.DefaultParams()
	require.NoError(t, pk.ParamStore.Set(ctx, params))

	response, err := pk.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
