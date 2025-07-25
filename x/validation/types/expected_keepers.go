package types

import (
	context "context"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	zentptypes "github.com/Zenrock-Foundation/zrchain/v6/x/zentp/types"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"

	treasurytypes "github.com/Zenrock-Foundation/zrchain/v6/x/treasury/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	AddressCodec() address.Codec

	IterateAccounts(ctx context.Context, process func(sdk.AccountI) (stop bool))
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI // only used for simulation

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(context.Context, sdk.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins

	GetSupply(ctx context.Context, denom string) sdk.Coin

	SendCoinsFromModuleToModule(ctx context.Context, senderPool, recipientPool string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, name string, amt sdk.Coins) error
}

// ValidatorSet expected properties for the set of all validators (noalias)
type ValidatorSet interface {
	// iterate through validators by operator address, execute func for each validator
	IterateValidators(context.Context,
		func(index int64, validator stakingtypes.ValidatorI) (stop bool)) error

	// iterate through bonded validators by operator address, execute func for each validator
	IterateBondedValidatorsByPower(context.Context,
		func(index int64, validator stakingtypes.ValidatorI) (stop bool)) error

	// iterate through the consensus validator set of the last block by operator address, execute func for each validator
	IterateLastValidators(context.Context,
		func(index int64, validator ValidatorI) (stop bool)) error

	Validator(context.Context, sdk.ValAddress) (stakingtypes.ValidatorI, error)            // get a particular validator by operator address
	ValidatorByConsAddr(context.Context, sdk.ConsAddress) (stakingtypes.ValidatorI, error) // get a particular validator by consensus address
	TotalBondedTokens(context.Context) (math.Int, error)                                   // total bonded tokens within the validator set
	StakingTokenSupply(context.Context) (math.Int, error)                                  // total staking token supply

	// slash the validator and delegators of the validator, specifying offense height, offense power, and slash fraction
	Slash(context.Context, sdk.ConsAddress, int64, int64, math.LegacyDec) (math.Int, error)
	SlashWithInfractionReason(context.Context, sdk.ConsAddress, int64, int64, math.LegacyDec, stakingtypes.Infraction) (math.Int, error)
	Jail(context.Context, sdk.ConsAddress) error   // jail a validator
	Unjail(context.Context, sdk.ConsAddress) error // unjail a validator

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(context.Context, sdk.AccAddress, sdk.ValAddress) (stakingtypes.DelegationI, error)

	// MaxValidators returns the maximum amount of bonded validators
	MaxValidators(context.Context) (uint32, error)

	// GetPubKeyByConsAddr returns the consensus public key for a validator. Used in vote
	// extension validation.
	GetPubKeyByConsAddr(context.Context, sdk.ConsAddress) (cmtprotocrypto.PublicKey, error)
}

// DelegationSet expected properties for the set of all delegations for a particular (noalias)
type DelegationSet interface {
	GetValidatorSet() ValidatorSet // validator set for which delegation set is based upon

	// iterate through all delegations from one delegator by validator-AccAddress,
	//   execute func for each validator
	IterateDelegations(ctx context.Context, delegator sdk.AccAddress,
		fn func(index int64, delegation stakingtypes.DelegationI) (stop bool)) error
}

// Event Hooks
// These can be utilized to communicate between a staking keeper and another
// keeper which must take particular actions when validators/delegators change
// state. The second keeper must implement this interface, which then the
// staking keeper can call.

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error                           // Must be called when a validator is created
	BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error                         // Must be called when a validator's state changes
	AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error // Must be called when a validator is deleted

	AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error         // Must be called when a validator is bonded
	AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error // Must be called when a validator begins unbonding

	BeforeDelegationCreated(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error        // Must be called when a delegation is created
	BeforeDelegationSharesModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error // Must be called when a delegation's shares are modified
	BeforeDelegationRemoved(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error        // Must be called when a delegation is removed
	AfterDelegationModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error
	BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction math.LegacyDec) error
	AfterUnbondingInitiated(ctx context.Context, id uint64) error
}

// StakingHooksWrapper is a wrapper for modules to inject StakingHooks using depinject.
type StakingHooksWrapper struct{ StakingHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (StakingHooksWrapper) IsOnePerModuleType() {}

type ZentpKeeper interface {
	GetSignerKeyID(ctx context.Context) uint64
	GetSolanaParams(ctx context.Context) *zentptypes.Solana
	GetMintsWithStatus(goCtx context.Context, status zentptypes.BridgeStatus) ([]*zentptypes.Bridge, error)
	GetMintsWithStatusPending(goCtx context.Context) ([]*zentptypes.Bridge, error)
	UpdateMint(ctx context.Context, id uint64, mint *zentptypes.Bridge) error
	GetBurns(goCtx context.Context, address, chainID, txHash string) ([]*zentptypes.Bridge, error)
	AddBurn(ctx context.Context, burn *zentptypes.Bridge) error
	GetBridgeFeeParams(ctx context.Context) (sdk.AccAddress, math.LegacyDec, error)
	GetBridgeFeeAmount(ctx context.Context, amount uint64, bridgeFee math.LegacyDec) (sdk.Coins, error)
	GetSolanaROCKSupply(ctx context.Context) (math.Int, error)
	SetSolanaROCKSupply(ctx context.Context, supply math.Int) error
	CheckROCKSupplyCap(ctx sdk.Context, newAmount math.Int) error
	CheckCanBurnFromSolana(ctx sdk.Context, burnAmount math.Int) error
	GetTotalROCKSupply(ctx sdk.Context) (math.Int, error)
}

type TreasuryKeeper interface {
	GetSignTransactionRequest(ctx sdk.Context, id uint64) (*treasurytypes.SignTransactionRequest, error)
	GetSignRequest(ctx sdk.Context, id uint64) (*treasurytypes.SignRequest, error)
	GetAddressByWalletType(ctx sdk.Context, id uint64, walletType treasurytypes.WalletType, prefixes []string) (string, error)
	HandleSignTransactionRequest(ctx sdk.Context, msg *treasurytypes.MsgNewSignTransactionRequest, data []byte) (*treasurytypes.MsgNewSignTransactionRequestResponse, error)
	GetKey(ctx sdk.Context, keyID uint64) (*treasurytypes.Key, error)
}

type SlashingKeeper interface {
	GetValidatorSigningInfo(ctx context.Context, address sdk.ConsAddress) (slashingtypes.ValidatorSigningInfo, error)
	SetValidatorSigningInfo(ctx context.Context, address sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo) error
}
