package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/zenrocklabs/zenrock/zrchain/v4/testutil/sample"
)

func TestMsgSubmitSolanaUnlockTransaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitSolanaUnlockTransaction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmitSolanaUnlockTransaction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSubmitSolanaUnlockTransaction{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
