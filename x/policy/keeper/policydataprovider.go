package keeper

import (
	"crypto/sha256"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/policy/types"
)

type PolicyDataProvider interface {
	GetData(addr string, act *types.Action, cdc codec.BinaryCodec) (string, []byte, error)
}

type PasskeyPolicyDataProvider struct{}

func (p *PasskeyPolicyDataProvider) GetData(addr string, act *types.Action, cdc codec.BinaryCodec) (string, []byte, error) {
	msgbz, err := cdc.Marshal(act)
	if err != nil {
		return "", nil, err
	}
	challengePayload := append(msgbz, []byte(addr)...)
	challenge := sha256.Sum256(challengePayload)
	return "challenge-" + addr, challenge[:], nil
}
