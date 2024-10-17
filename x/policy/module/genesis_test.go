package policy_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/testutil/nullify"
	policy "github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/module"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	keepers := keepertest.NewTest(t)
	ik := keepers.PolicyKeeper
	ctx := keepers.Ctx
	policy.InitGenesis(ctx, *ik, genesisState)
	got := policy.ExportGenesis(ctx, *ik)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
