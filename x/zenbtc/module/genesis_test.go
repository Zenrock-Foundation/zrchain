package zenbtc_test

import (
	"testing"

	keepertest "github.com/Zenrock-Foundation/zrchain/v4/testutil/keeper"
	"github.com/Zenrock-Foundation/zrchain/v4/testutil/nullify"
	zenbtc "github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/module"
	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZenbtcKeeper(t)
	zenbtc.InitGenesis(ctx, k, genesisState)
	got := zenbtc.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
