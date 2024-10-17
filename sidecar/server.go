package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/Zenrock-Foundation/zrchain/v4/sidecar/proto/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"google.golang.org/grpc"
)

func startGRPCServer(oracle *Oracle, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterSidecarServiceServer(s, &oracleService{oracle: oracle})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type oracleService struct {
	api.UnimplementedSidecarServiceServer
	oracle *Oracle
}

func NewOracleService(oracle *Oracle) *oracleService {
	return &oracleService{
		oracle: oracle,
	}
}

func (s *oracleService) GetSidecarState(ctx context.Context, req *api.SidecarStateRequest) (*api.SidecarStateResponse, error) {
	currentState := s.oracle.currentState.Load().(*OracleState)

	contractState, err := json.Marshal(currentState.Delegations)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal state: %w", err)
	}

	return &api.SidecarStateResponse{
		Delegations:    contractState,
		ROCKUSDPrice:   fmt.Sprint(currentState.ROCKUSDPrice),
		ETHUSDPrice:    fmt.Sprint(currentState.ETHUSDPrice),
		EthBlockHeight: currentState.EthBlockHeight,
		EthBlockHash:   currentState.EthBlockHash,
		EthGasPrice:    currentState.EthGasPrice,
		EthGasLimit:    currentState.EthGasLimit,
	}, nil
}

func (s *oracleService) GetSidecarStateByEthHeight(ctx context.Context, req *api.SidecarStateByEthHeightRequest) (*api.SidecarStateResponse, error) {
	state, err := s.oracle.getStateByEthHeight(req.EthBlockHeight)
	if err != nil {
		return nil, err
	}

	contractState, err := json.Marshal(state.Delegations)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal state: %w", err)
	}

	return &api.SidecarStateResponse{
		Delegations:    contractState,
		ROCKUSDPrice:   fmt.Sprint(state.ROCKUSDPrice),
		ETHUSDPrice:    fmt.Sprint(state.ETHUSDPrice),
		EthBlockHeight: state.EthBlockHeight,
		EthBlockHash:   state.EthBlockHash,
		EthGasPrice:    state.EthGasPrice,
		EthGasLimit:    state.EthGasLimit,
	}, nil
}

func (s *oracleService) GetBitcoinBlockHeaderByHeight(ctx context.Context, req *api.BitcoinBlockHeaderByHeightRequest) (*api.BitcoinBlockHeaderResponse, error) {
	blockheader, blockhash, tipHeight, err := s.oracle.neutrinoServer.GetBlockHeaderByHeight(req.ChainName, req.BlockHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to GetBlockHeaderByHeight: %w", err)
	}

	bh := &api.BTCBlockHeader{
		Version:    int64(blockheader.Version),
		PrevBlock:  blockheader.PrevBlock.String(),
		MerkleRoot: blockheader.MerkleRoot.String(),
		TimeStamp:  blockheader.Timestamp.Unix(),
		Bits:       int64(blockheader.Bits),
		Nonce:      int64(blockheader.Nonce),
		BlockHash:  blockhash.String(),
	}

	return &api.BitcoinBlockHeaderResponse{
		BlockHeader: bh,
		BlockHeight: req.BlockHeight,
		TipHeight:   int64(tipHeight),
	}, nil
}

func (s *oracleService) GetLatestBitcoinBlockHeader(ctx context.Context, req *api.LatestBitcoinBlockHeaderRequest) (*api.BitcoinBlockHeaderResponse, error) {
	blockheader, blockhash, tipHeight, err := s.oracle.neutrinoServer.GetLatestBlockHeader(req.ChainName)
	if err != nil {
		return nil, fmt.Errorf("failed to GetBlockHeaderByHeight: %w", err)
	}

	bh := &api.BTCBlockHeader{
		Version:    int64(blockheader.Version),
		PrevBlock:  blockheader.PrevBlock.String(),
		MerkleRoot: blockheader.MerkleRoot.String(),
		TimeStamp:  blockheader.Timestamp.Unix(),
		Bits:       int64(blockheader.Bits),
		Nonce:      int64(blockheader.Nonce),
		BlockHash:  blockhash.String(),
	}

	return &api.BitcoinBlockHeaderResponse{
		BlockHeader: bh,
		BlockHeight: int64(tipHeight),
		TipHeight:   int64(tipHeight),
	}, nil
}

func (s *oracleService) GetSolanaTransaction(ctx context.Context, req *api.SolanaTransactionRequest) (*api.SolanaTransactionResponse, error) {
	out, err := s.oracle.solanaClient.GetTransaction(
		ctx,
		solana.MustSignatureFromBase58(req.TxSignature),
		&rpc.GetTransactionOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return &api.SolanaTransactionResponse{
		TxSlot: out.Slot,
	}, nil
}

type DebugSolanaTransactionResponse struct {
	Tx *rpc.GetTransactionResult
}

func (s *oracleService) DebugGetSolanaTransaction(ctx context.Context, req *api.SolanaTransactionRequest) (*DebugSolanaTransactionResponse, error) {
	out, err := s.oracle.solanaClient.GetTransaction(
		ctx,
		solana.MustSignatureFromBase58(req.TxSignature),
		&rpc.GetTransactionOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return &DebugSolanaTransactionResponse{
		Tx: out,
	}, nil
}

func (s *oracleService) GetEthereumNonceAtHeight(ctx context.Context, req *api.EthereumNonceAtHeightRequest) (*api.EthereumNonceAtHeightResponse, error) {
	nonce, err := s.oracle.EthClient.NonceAt(ctx, common.HexToAddress(req.Address), big.NewInt(int64(req.Height)))
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	return &api.EthereumNonceAtHeightResponse{
		Nonce: nonce,
	}, nil
}
