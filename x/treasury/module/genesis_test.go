package treasury_test

import (
	"testing"

	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/testutil/nullify"
	treasury "github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/module"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	keepers := keepertest.NewTest(t)
	tk := keepers.TreasuryKeeper
	ctx := keepers.Ctx
	treasury.InitGenesis(ctx, *tk, genesisState)
	got := treasury.ExportGenesis(ctx, *tk)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
