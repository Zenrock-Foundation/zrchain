package types

import (
	"context"

	"github.com/Zenrock-Foundation/zrchain/v5/x/identity/types"
	treasurytypes "github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	GetModuleAddress(name string) sdk.AccAddress
	HasAccount(context.Context, sdk.AccAddress) bool
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	BurnCoins(ctx context.Context, moduleAccount string, amounts sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// TreasuryKeeper defines the expected interface for the Treasury module.
type TreasuryKeeper interface {
	GetKey(ctx sdk.Context, keyID uint64) (*treasurytypes.Key, error)
}

type IdentityKeeper interface {
	GetWorkspaces(goCtx context.Context, user string) ([]*types.Workspace, error)
}
