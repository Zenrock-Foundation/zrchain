package eigenlayer

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/Zenrock-Foundation/zrchain/zenbtc-rewards/internal/chain"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	contractrewardscoordinator "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/IRewardsCoordinator"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/claimgen"
	"github.com/Layr-Labs/eigenlayer-rewards-proofs/pkg/proofDataFetcher/httpProofDataFetcher"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	eigensdkLogger "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type elChainReader interface {
	GetRootIndexFromHash(opts *bind.CallOpts, hash [32]byte) (uint32, error)
	GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (rewardscoordinator.IRewardsCoordinatorDistributionRoot, error)
}

type EigenlayerClient interface {
	ClaimRewards(earnerAddress string, broadcast bool) (*types.Receipt, error)
}

type eigenlayerClient struct {
	chainId     *big.Int
	networkname string
	environment string
	ethAccount  *chain.EthAccount
	ethClient   *ethclient.Client
	logger      eigensdkLogger.Logger
}

func NewEigenlayerClient(logger eigensdkLogger.Logger, ethClient *ethclient.Client, ethAccount *chain.EthAccount) (EigenlayerClient, error) {
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	networkname := "holesky"
	environment := "testnet"
	if chainId.Int64() == MainnetChainId {
		networkname = "mainnet"
		environment = "prod"
	}

	return &eigenlayerClient{
		chainId:     chainId,
		ethClient:   ethClient,
		ethAccount:  ethAccount,
		logger:      logger,
		networkname: networkname,
		environment: environment,
	}, nil
}

type ChainMetadata struct {
	ELRewardsCoordinatorAddress string
	ProofStoreBaseURL           string
}

var (
	MainnetChainId int64 = 1
	HoleskyChainId int64 = 17000

	ChainMetadataMap = map[int64]ChainMetadata{
		MainnetChainId: {
			ELRewardsCoordinatorAddress: "0x7750d328b314EfFa365A0402CcfD489B80B0adda",
			ProofStoreBaseURL:           "https://eigenlabs-rewards-mainnet-ethereum.s3.amazonaws.com",
		},
		HoleskyChainId: {
			ELRewardsCoordinatorAddress: "0xAcc1fb458a1317E886dB376Fc8141540537E68fE",
			ProofStoreBaseURL:           "https://eigenlabs-rewards-testnet-holesky.s3.amazonaws.com",
		},
	}
)

func (c *eigenlayerClient) ClaimRewards(earnerAddress string, broadcast bool) (*types.Receipt, error) {
	ctx := context.Background()

	rewardCoordinatorAddress := common.HexToAddress(ChainMetadataMap[c.chainId.Int64()].ELRewardsCoordinatorAddress)
	earner := common.HexToAddress(earnerAddress)
	recvAddr := c.ethAccount.GetAddress()

	rc, err := rewardscoordinator.NewContractIRewardsCoordinator(rewardCoordinatorAddress, c.ethClient)
	if err != nil {
		return nil, errors.Wrap(err, "could not create reward coordinator client")
	}

	claimerFor, err := rc.ClaimerFor(&bind.CallOpts{}, earner)
	if err != nil {
		return nil, errors.Wrap(err, "could not get claimerFor")
	}
	if claimerFor.Cmp(recvAddr) != 0 {
		return nil, fmt.Errorf("claimer for earner %s doesnt match signer %s", claimerFor, recvAddr)
	}

	elReader, err := elcontracts.NewReaderFromConfig(
		elcontracts.Config{
			RewardsCoordinatorAddress: rewardCoordinatorAddress,
		},
		c.ethClient,
		c.logger,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not create elreader")
	}

	df := httpProofDataFetcher.NewHttpProofDataFetcher(
		ChainMetadataMap[c.chainId.Int64()].ProofStoreBaseURL,
		c.environment,
		c.networkname,
		http.DefaultClient,
	)

	claimDate, rootIndex, err := getClaimDistributionRoot(ctx, elReader, c.logger)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get claim distribution root")
	}

	proofData, err := df.FetchClaimAmountsForDate(ctx, claimDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch claim amounts for date")
	}

	tokenAddressesMap, present := proofData.Distribution.GetTokensForEarner(earner)
	if !present {
		return nil, errors.Wrap(err, "earner address not found in distribution")
	}

	tokens := []common.Address{}
	for pair := tokenAddressesMap.Oldest(); pair != nil; pair = pair.Next() {
		c.logger.Infof("found reward token: %s: %d")
		tokens = append(tokens, pair.Key)
	}

	cg := claimgen.NewClaimgen(proofData.Distribution)
	_, claim, err := cg.GenerateClaimProofForEarner(
		earner,
		tokens,
		rootIndex,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate claim proof for earner")
	}

	elClaim := rewardscoordinator.IRewardsCoordinatorRewardsMerkleClaim{
		RootIndex:       claim.RootIndex,
		EarnerIndex:     claim.EarnerIndex,
		EarnerTreeProof: claim.EarnerTreeProof,
		EarnerLeaf: rewardscoordinator.IRewardsCoordinatorEarnerTreeMerkleLeaf{
			Earner:          claim.EarnerLeaf.Earner,
			EarnerTokenRoot: claim.EarnerLeaf.EarnerTokenRoot,
		},
		TokenIndices:    claim.TokenIndices,
		TokenTreeProofs: claim.TokenTreeProofs,
		TokenLeaves:     convertClaimTokenLeaves(claim.TokenLeaves),
	}

	c.logger.Info("Validating claim proof...")
	ok, err := elReader.CheckClaim(&bind.CallOpts{Context: ctx}, elClaim)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check claim")
	}
	if !ok {
		return nil, errors.New("failed to validate claim")
	}
	c.logger.Info("Claim proof validated successfully")

	txMgr, err := c.getTxMgr()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tx mgr")
	}

	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get no send tx opts")
	}

	c.logger.Info("Preparing tx")
	tx, err := rc.ProcessClaim(noSendTxOpts, elClaim, recvAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ProcessClaim tx")
	}

	if broadcast {
		c.logger.Info("Broadcasting tx")
		receipt, err := txMgr.Send(ctx, tx, true)
		if err != nil {
			return nil, errors.Wrap(err, "failed broadcast ProcessClaim")
		}

		c.logger.Infof("Claim transaction submitted successfully, %s, %d", receipt.TxHash.String(), c.chainId)
		return receipt, nil
	}

	return nil, nil
}

func getClaimDistributionRoot(
	ctx context.Context,
	elReader elChainReader,
	logger eigensdkLogger.Logger,
) (string, uint32, error) {
	latestClaimableRoot, err := elReader.GetCurrentClaimableDistributionRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", 0, err
	}

	rootIndex, err := elReader.GetRootIndexFromHash(&bind.CallOpts{Context: ctx}, latestClaimableRoot.Root)
	if err != nil {
		return "", 0, err
	}

	ts := time.Unix(int64(latestClaimableRoot.RewardsCalculationEndTimestamp), 0).UTC().Format(time.DateOnly)
	logger.Debugf("Latest rewards snapshot timestamp: %s, root index: %d", ts, rootIndex)

	return ts, rootIndex, nil
}

func convertClaimTokenLeaves(
	claimTokenLeaves []contractrewardscoordinator.IRewardsCoordinatorTokenTreeMerkleLeaf,
) []rewardscoordinator.IRewardsCoordinatorTokenTreeMerkleLeaf {
	var tokenLeaves []rewardscoordinator.IRewardsCoordinatorTokenTreeMerkleLeaf
	for _, claimTokenLeaf := range claimTokenLeaves {
		tokenLeaves = append(tokenLeaves, rewardscoordinator.IRewardsCoordinatorTokenTreeMerkleLeaf{
			Token:              claimTokenLeaf.Token,
			CumulativeEarnings: claimTokenLeaf.CumulativeEarnings,
		})
	}
	return tokenLeaves
}

func (c *eigenlayerClient) getTxMgr() (*txmgr.SimpleTxManager, error) {
	chainId, err := c.ethClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	signerCfg := signerv2.Config{
		PrivateKey: c.ethAccount.GetPrivKey(),
	}
	sgn, sender, err := signerv2.SignerFromConfig(signerCfg, chainId)
	if err != nil {
		return nil, err
	}
	keyWallet, err := wallet.NewPrivateKeyWallet(c.ethClient, sgn, sender, c.logger)
	if err != nil {
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(keyWallet, c.ethClient, c.logger, sender)
	return txMgr, nil
}
