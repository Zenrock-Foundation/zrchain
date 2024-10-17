package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/identity/types"
)

func (k *Keeper) ExportState(ctx sdk.Context, genState *types.GenesisState) error {
	wsStore, err := k.WorkspaceStore.Iterate(ctx, nil)
	if err != nil {
		return err
	}
	genState.Workspaces, err = wsStore.Values()
	if err != nil {
		return err
	}

	krStore, err := k.KeyringStore.Iterate(ctx, nil)
	if err != nil {
		return err
	}
	genState.Keyrings, err = krStore.Values()
	if err != nil {
		return err
	}

	return nil
}
