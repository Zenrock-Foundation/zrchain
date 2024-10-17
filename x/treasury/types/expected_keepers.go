package types

import (
	"context"

	policykeeper "github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	idtypes "github.com/zenrocklabs/zenrock/zrchain/v4/x/identity/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

type IdentityKeeper interface {
	GetWorkspace(ctx sdk.Context, addr string) *idtypes.Workspace
	GetKeyring(ctx sdk.Context, addr string) *idtypes.Keyring
	GetZrSignWorkspace(goCtx context.Context, ethAddress, walletType string) (string, error)
	GetZrSignWorkspaces(goCtx context.Context, ethAddress, walletType string) (map[string]string, error)
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

type PolicyKeeper interface {
	policykeeper.ExportedKeeper
}
