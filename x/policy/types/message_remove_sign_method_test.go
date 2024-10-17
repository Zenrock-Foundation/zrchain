package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/zenrocklabs/zenrock/zrchain/v4/testutil/sample"
)

func TestMsgRemoveSignMethod_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveSignMethod
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRemoveSignMethod{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid id",
			msg: MsgRemoveSignMethod{
				Creator: sample.AccAddress(),
				Id:      sample.StringLen(513),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid address",
			msg: MsgRemoveSignMethod{
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
