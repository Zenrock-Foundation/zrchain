package testutil

import (
	_ "github.com/Zenrock-Foundation/zrchain/v6/x/mint" // import as blank for app wiring
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	_ "github.com/cosmos/cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/distribution"   // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/slashing"       // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/x/staking"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.GenutilModule(),
	configurator.MintModule(),
	configurator.DistributionModule(),
)
