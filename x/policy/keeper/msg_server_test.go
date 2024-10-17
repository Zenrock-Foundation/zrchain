package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	keepers := keepertest.NewTest(t)
	pk := keepers.PolicyKeeper
	ctx := keepers.Ctx
	return *pk, keeper.NewMsgServerImpl(*pk), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
