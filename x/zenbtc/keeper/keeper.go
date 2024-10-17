package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	treasury "github.com/zenrocklabs/zenrock/zrchain/v4/x/treasury/keeper"
	validation "github.com/zenrocklabs/zenrock/zrchain/v4/x/validation/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/zenbtc/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		validationKeeper *validation.Keeper
		treasuryKeeper   *treasury.Keeper

		Schema               collections.Schema
		LockTransactionStore collections.KeySet[string]
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	validationKeeper *validation.Keeper,
	treasuryKeeper *treasury.Keeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:                  cdc,
		storeService:         storeService,
		authority:            authority,
		logger:               logger,
		validationKeeper:     validationKeeper,
		treasuryKeeper:       treasuryKeeper,
		LockTransactionStore: collections.NewKeySet(sb, types.AssetPricesKey, types.AssetPricesIndex, collections.StringKey),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
