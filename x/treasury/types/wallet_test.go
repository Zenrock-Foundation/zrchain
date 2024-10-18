package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

// Test that every WalletType defined in protobuf is covered by NewWalletI.
func Test_NewWalletI_Exhaustive(t *testing.T) {
	for walletType, name := range WalletType_name {
		if walletType == int32(WalletType_WALLET_TYPE_UNSPECIFIED) {
			continue
		}

		t.Run(name, func(t *testing.T) {
			_, err := NewWallet(
				&Key{
					Id:            0,
					WorkspaceAddr: "workspace14a2hpadpsy9h4auve2z8lw",
					Type:          0,
					PublicKey:     hexutil.MustDecode("0x025cd45a6614df5348692ea4d0f7c16255b75a6b6f67bea5013621fe84af8031f0"),
				},
				WalletType(walletType),
			)
			require.NotErrorIs(t, err, ErrUnknownWalletType)
		})
	}
}
