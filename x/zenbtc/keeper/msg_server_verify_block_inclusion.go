package keeper

import (
	"context"
	"errors"

	// "net/rpc"
	// "github.com/Zenrock-Foundation/zrchain/v4/sidecar/proto/api"
	// "github.com/btcsuite/btcd/chaincfg/chainhash"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Zenrock-Foundation/zrchain/v4/bitcoin"
	treasurytypes "github.com/Zenrock-Foundation/zrchain/v4/x/treasury/types"
	"github.com/Zenrock-Foundation/zrchain/v4/x/zenbtc/types"
)

func (k msgServer) VerifyDepositBlockInclusion(goCtx context.Context, msg *types.MsgVerifyDepositBlockInclusion) (*types.MsgVerifyDepositBlockInclusionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	blockHeader, err := k.validationKeeper.BtcBlockHeaders.Get(ctx, msg.BlockHeight)
	//try and get missing Blockheader over RPC - WARNING for debugging only!!!!
	// if err != nil {
	// 	bh, _ := debugRetrieveBlockHeaderViaRPC(msg.ChainName, msg.BlockHeight)
	// 	if bh != nil {
	// 		err = nil
	// 		blockHeader = *bh
	// 	}
	// }
	// END of debug code

	if err != nil {
		return nil, err
	}

	found := false
	outputs, _, err := bitcoin.VerifyBTCLockTransaction(msg.RawTx, msg.ChainName, int(msg.Index), msg.Proof, &blockHeader)
	if err != nil {
		return nil, err
	}
	for _, output := range outputs {
		if output.Address == msg.DepositAddr && output.Amount == msg.Amount {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("zenBTC deposit not found in outputs from provided lock tx")
	}

	q, err := k.treasuryKeeper.KeyByAddress(ctx, &treasurytypes.QueryKeyByAddressRequest{
		Address:     msg.DepositAddr,
		KeyringAddr: k.validationKeeper.GetZenBTCDepositKeyringAddr(ctx),
		KeyType:     treasurytypes.KeyType_KEY_TYPE_BITCOIN_SECP256K1,
		WalletType:  treasurytypes.WalletType_WALLET_TYPE_BTC_TESTNET,
	})
	if err != nil {
		return nil, err
	}

	found = false
	for _, wallet := range q.Response.Wallets {
		if wallet.Address == msg.DepositAddr {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("zenBTC deposit address not found in key wallets")
	}

	// Deposit/lock txs are stored in zenBTC module so they can't be used to mint zenBTC tokens more than once
	txExists, err := k.LockTransactionStore.Has(ctx, msg.RawTx)
	if err != nil {
		return nil, err
	}
	if txExists {
		return nil, errors.New("lock tx was already previously used to mint zenBTC tokens")
	}

	pendingTxs, err := k.validationKeeper.PendingMintTransactions.Get(ctx)
	if err != nil {
		return nil, err
	}
	pendingTxs.Txs = append(pendingTxs.Txs, &treasurytypes.PendingMintTransaction{
		ChainId:          q.Response.Key.ZenbtcMetadata.ChainId,
		ChainType:        q.Response.Key.ZenbtcMetadata.ChainType,
		RecipientAddress: q.Response.Key.ZenbtcMetadata.RecipientAddr,
		Amount:           msg.Amount,
		Creator:          msg.Creator,
		KeyId:            k.validationKeeper.GetZenBTCMinterKeyID(ctx),
	})
	if err := k.validationKeeper.PendingMintTransactions.Set(ctx, pendingTxs); err != nil {
		return nil, err
	}

	if err := k.LockTransactionStore.Set(ctx, msg.RawTx); err != nil {
		return nil, err
	}

	return &types.MsgVerifyDepositBlockInclusionResponse{}, nil
}

// func debugRetrieveBlockHeaderViaRPC(chainName string, blockHeight int64) (*api.BTCBlockHeader, error) {
// 	type GetBlockHeaderByHeightArgs struct {
// 		ChainName string
// 		Height    int64
// 	}

// 	type GetBlockHeaderByHeightReply struct {
// 		BlockHeader *api.BTCBlockHeader
// 		BlockHash   *chainhash.Hash
// 		Height      int32
// 	}

// 	args := GetBlockHeaderByHeightArgs{
// 		ChainName: chainName,
// 		Height:    blockHeight,
// 	}
// 	var resp GetBlockHeaderByHeightReply
// 	client, err := rpc.DialHTTP("tcp", "localhost"+":12345")
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = client.Call("NeutrinoServer.BlockHeaderByHeight", args, &resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.BlockHeader, nil
// }
