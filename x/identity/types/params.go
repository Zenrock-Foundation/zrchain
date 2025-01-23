package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(keyringCreationFee, minimumBtl, defaultBtl, minMPCbtl uint64) Params {
	return Params{
		KeyringCreationFee: keyringCreationFee,
		MpcMinimumBtl:      minMPCbtl,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(10000000000, 10, 1000, 20)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
