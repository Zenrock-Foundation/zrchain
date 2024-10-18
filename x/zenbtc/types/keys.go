package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "zenbtc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zenbtc"
)

var (
	AssetPricesKey = collections.NewPrefix(0)

	AssetPricesIndex = "asset_prices"

	ParamsKey = []byte("p_zenbtc")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
