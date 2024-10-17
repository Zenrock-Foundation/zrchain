package v2

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/zenrocklabs/zenrock/zrchain/v4/app/upgrades"
)

const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        storetypes.StoreUpgrades{},
}
