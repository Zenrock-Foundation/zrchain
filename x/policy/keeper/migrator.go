package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/migrations/v1"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) *Migrator {
	return &Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	ctx.Logger().With("module", types.ModuleName).Info("starting migration to v2")
	v1.UpdateParams(ctx, m.keeper.ParamStore)

	// ...

	return nil
}
