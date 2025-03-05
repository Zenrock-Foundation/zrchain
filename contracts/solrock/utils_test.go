package solrock

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/Zenrock-Foundation/zrchain/v5/contracts/solrock/generated/zenbtc_spl_token"
	"github.com/gagliardetto/solana-go"
	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

var endpoint = "api.devnet.solana.com" //rpc.LocalNet.RPC
var authority, _ = LoadLocalWallet()
var programID = solana.MustPublicKeyFromBase58("3dvdvGsDiHnJrCWBpuR1yDmozY2BBruEzPBcf52MWviX")

var client = rpc.New("https://api.devnet.solana.com")

var tokenParams = Token{
	Name:     "Zenrock BTC",
	Symbol:   "zenBTC",
	Decimals: 8,
	Uri:      "https://www.zenrocklabs.io/metadata.json",
}

var mintAddress, _ = GetMintAddress(programID, tokenParams.Symbol) // You can also just use the token address

var userWallet = solana.NewWallet()
var feeWallet = authority.PublicKey()

func TestInitialize(t *testing.T) {
	signer := authority

	err := RequestAirdrop(context.Background(), client, userWallet.PublicKey(), 1000000000)
	require.NoError(t, err)

	recent, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentConfirmed)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			Initialize(
				zenbtc_spl_token.InitializeArgs{
					GlobalAuthority: authority.PublicKey(),
				},
				programID,
				signer.PublicKey(),
			),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	require.NoError(t, err)

	_, err = SignTransaction(tx, signer)
	require.NoError(t, err)

	signature, err := SendTransaction(client, context.Background(), tx)
	require.NoError(t, err)

	confirmedTx, err := WaitForTransactionConfirmation(
		context.Background(),
		client,
		signature,
		rpc.CommitmentConfirmed,
	)
	require.NoError(t, err)

	require.Empty(t, confirmedTx.Meta.Err) // The transaction succeeded if there's no error
}

func TestDeployToken(t *testing.T) {
	signer := authority

	recent, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentConfirmed)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			DeployToken(
				zenbtc_spl_token.DeployTokenArgs{
					MintAuthorities: []solana.PublicKey{authority.PublicKey()},
					FeeAuthorities:  []solana.PublicKey{authority.PublicKey()},
					FeeWallet:       authority.PublicKey(),
					BurnFeeBps:      2,
					TokenName:       tokenParams.Name,
					TokenSymbol:     tokenParams.Symbol,
					TokenDecimals:   tokenParams.Decimals,
					TokenUri:        tokenParams.Uri,
				},
				programID,
				mintAddress,
				signer.PublicKey(),
			),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	require.NoError(t, err)

	_, err = SignTransaction(tx, signer)
	require.NoError(t, err)

	signature, err := SendTransaction(client, context.Background(), tx)
	require.NoError(t, err)

	confirmedTx, err := WaitForTransactionConfirmation(
		context.Background(),
		client,
		signature,
		rpc.CommitmentConfirmed,
	)
	require.NoError(t, err)
	require.Empty(t, confirmedTx.Meta.Err) // The transaction succeeded if there's no error
}

func TestWrap(t *testing.T) {
	signer := authority

	receiver := userWallet.PublicKey()
	value := uint64(10000)
	fee := uint64(20)

	recent, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentConfirmed)
	require.NoError(t, err)

	instructions := []solana.Instruction{}

	feeWalletAta, _, err := solana.FindAssociatedTokenAddress(feeWallet, mintAddress)
	require.NoError(t, err)

	_, err = GetTokenAccount(context.Background(), client, feeWalletAta)

	if err.Error() == "not found" {
		instructions = append(
			instructions,
			ata.NewCreateInstruction(
				signer.PublicKey(),
				feeWallet,
				mintAddress,
			).Build(),
		)
	} else {
		require.NoError(t, err)
	}

	receiverAta, _, err := solana.FindAssociatedTokenAddress(receiver, mintAddress)
	require.NoError(t, err)

	_, err = GetTokenAccount(context.Background(), client, receiverAta)

	if err.Error() == "not found" {
		instructions = append(
			instructions,
			ata.NewCreateInstruction(
				signer.PublicKey(),
				receiver,
				mintAddress,
			).Build(),
		)
	}

	instructions = append(instructions, Wrap(
		zenbtc_spl_token.WrapArgs{
			Value: value,
			Fee:   fee,
		},
		programID,
		mintAddress,
		signer.PublicKey(),
		feeWallet,
		feeWalletAta,
		receiver,
		receiverAta,
	))

	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	require.NoError(t, err)

	_, err = SignTransaction(tx, signer)
	require.NoError(t, err)

	// // Pretty print the transaction:
	// fmt.Println(tx.String())

	signature, err := SendTransaction(client, context.Background(), tx)
	require.NoError(t, err)

	confirmedTx, err := WaitForTransactionConfirmation(
		context.Background(),
		client,
		signature,
		rpc.CommitmentConfirmed,
	)
	require.NoError(t, err)
	require.Empty(t, confirmedTx.Meta.Err) // The transaction succeeded if there's no error
}

func TestUnwrap(t *testing.T) {
	signer := userWallet.PrivateKey

	value := uint64(5000)
	destAddr := [25]uint8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
		22, 23, 24, 25,
	}

	recent, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentConfirmed)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			Unwrap(
				zenbtc_spl_token.UnwrapArgs{
					Value:    value,
					DestAddr: destAddr,
				},
				programID,
				mintAddress,
				signer.PublicKey(),
				feeWallet,
			),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	require.NoError(t, err)

	_, err = SignTransaction(tx, signer)
	require.NoError(t, err)

	signature, err := SendTransaction(client, context.Background(), tx)
	require.NoError(t, err)

	confirmedTx, err := WaitForTransactionConfirmation(
		context.Background(),
		client,
		signature,
		rpc.CommitmentConfirmed,
	)
	require.NoError(t, err)
	require.Empty(t, confirmedTx.Meta.Err) // The transaction succeeded if there's no error
}

func TestGetTokenBalance(t *testing.T) {
	userAta, _, err := solana.FindAssociatedTokenAddress(userWallet.PublicKey(), mintAddress)
	require.NoError(t, err)

	accountAtaInfo, err := client.GetTokenAccountBalance(context.Background(), userAta, rpc.CommitmentConfirmed)
	require.NoError(t, err)

	balance := accountAtaInfo.Value.Amount
	require.NotEmpty(t, balance)
}

func TestGetMint(t *testing.T) {
	mint, err := GetMint(context.Background(), client, mintAddress)
	require.NoError(t, err)

	globalConfigPDA, err := GetGlobalConfigPDA(programID)
	require.NoError(t, err)

	require.Equal(t, *mint.MintAuthority, globalConfigPDA)
	require.Greater(t, mint.Decimals, uint8(0))
}

func TestGetGlobalConfig(t *testing.T) {
	globalConfig, err := GetGlobalConfig(context.Background(), client, programID)

	require.NoError(t, err)
	require.NotEmpty(t, globalConfig.GlobalAuthority.String())
}

func TestGetTokenConfig(t *testing.T) {
	tokenConfig, err := GetTokenConfig(context.Background(), client, programID, mintAddress)
	require.NoError(t, err)

	require.NotEmpty(t, tokenConfig.BurnFeeBps)
	require.NotEmpty(t, tokenConfig.FeeWallet)
	require.GreaterOrEqual(t, len(tokenConfig.MintAuthorities), 1)
	require.GreaterOrEqual(t, len(tokenConfig.FeeAuthorities), 1)
}

func TestGetTokenRedemptionEvents(t *testing.T) {
	limit := 1000

	signatures, err := client.GetSignaturesForAddressWithOpts(context.Background(), programID, &rpc.GetSignaturesForAddressOpts{
		Limit:      &limit,
		Commitment: rpc.CommitmentConfirmed,
	})
	require.NoError(t, err)
	require.NotEmpty(t, signatures)

	tokenRedemptionEvents := []TokenRedemptionEvent{}

	for _, signature := range signatures {
		tx, err := client.GetTransaction(context.Background(), signature.Signature, &rpc.GetTransactionOpts{
			Commitment: rpc.CommitmentConfirmed,
		})
		require.NoError(t, err)
		require.NotEmpty(t, tx)

		events, err := zenbtc_spl_token.DecodeEvents(tx, programID)
		require.NoError(t, err)

		for _, event := range events {
			if event.Name == "TokenRedemption" {
				eventData := event.Data.(*zenbtc_spl_token.TokenRedemptionEventData)

				tokenRedemptionEvents = append(tokenRedemptionEvents, TokenRedemptionEvent{
					Signature: signature.Signature.String(),
					Slot:      tx.Slot,
					Date:      tx.BlockTime.Time(),
					Redeemer:  eventData.Redeemer,
					Value:     eventData.Value,
					DestAddr:  eventData.DestAddr,
					Fee:       eventData.Fee,
					Mint:      eventData.Mint,
					Id:        eventData.Id.BigInt(),
				})
			}
		}
	}

	require.NotEmpty(t, tokenRedemptionEvents)
	require.Len(t, tokenRedemptionEvents, 1)
	require.Equal(t, tokenRedemptionEvents[0].Id.Cmp(big.NewInt(0)), 0)
}

func TestDurableNonces(t *testing.T) {
	nonceAuth, err := solana.WalletFromPrivateKeyBase58("3E2YL7Sf5EZS6kXrHEd34byyF5Kzh7PCPH9Q4tRyRaTJfcHmbzr9nPUWgJ9t5Xy9LBg9BXJFdbuJwTtpdNoJS1Yh") //solana.NewWallet()
	fmt.Printf("pubkey: %s\n", nonceAuth.PublicKey().String())
	require.NoError(t, err)
	nonceWallet, err := solana.WalletFromPrivateKeyBase58("7hLFPdmGFZkzPjdi75qsYfFg9gpaZttwa1bqvJmzKydGodV5PQqQ1oi7HFaT54F8UmcZb4vjuBa4RWMwFJuKy7k") //solana.NewWallet()
	require.NoError(t, err)
	fmt.Printf("nonceWallet: %s\n", nonceWallet.PublicKey().String())
	sender := solana.NewWallet()
	receiver := solana.NewWallet().PublicKey()

	//err := RequestAirdrop(context.Background(), client, nonceAuth.PublicKey(), 1000*solana.LAMPORTS_PER_SOL)
	//require.NoError(t, err)
	//
	//err = RequestAirdrop(context.Background(), client, sender.PublicKey(), 1000*solana.LAMPORTS_PER_SOL)
	//require.NoError(t, err)
	recentHash, err := solana.HashFromBase58("HmhsXVRocBP2NuTkVoSgqfpEhuygfVTW6nYVKmshXPCA")

	t.Run("Creates a nonce account", func(t *testing.T) {
		//recent, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentConfirmed)
		require.NoError(t, err)

		tx, err := solana.NewTransaction(
			[]solana.Instruction{
				system.NewCreateAccountInstruction(
					uint64(0.0015*float64(solana.LAMPORTS_PER_SOL)),
					NONCE_ACCOUNT_LENGTH,
					solana.SystemProgramID,
					nonceAuth.PublicKey(),
					nonceWallet.PublicKey(),
				).Build(),

				system.NewInitializeNonceAccountInstruction(
					nonceAuth.PublicKey(),
					nonceWallet.PublicKey(),
					solana.SysVarRecentBlockHashesPubkey,
					solana.SysVarRentPubkey,
				).Build(),
			},
			recentHash,
			solana.TransactionPayer(nonceAuth.PublicKey()),
		)
		require.NoError(t, err)

		_, err = SignTransaction(tx, nonceAuth.PrivateKey)
		require.NoError(t, err)

		_, err = SignTransaction(tx, nonceWallet.PrivateKey)

		require.NoError(t, err)
		bin, err := tx.MarshalBinary()
		require.NoError(t, err)
		fmt.Printf("tx: %s\n", string(bin))
		signature, err := SendTransaction(client, context.Background(), tx)
		require.NoError(t, err)

		confirmedTx, err := WaitForTransactionConfirmation(
			context.Background(),
			client,
			signature,
			rpc.CommitmentConfirmed,
		)
		require.NoError(t, err)

		require.Empty(t, confirmedTx.Meta.Err) // The transaction succeeded if there's no error

		nonceAccount, err := GetNonceAccount(context.Background(), client, nonceWallet.PublicKey())
		require.NoError(t, err)

		require.Equal(t, nonceAccount.AuthorizedPubkey, nonceAuth.PublicKey())
	})

	t.Run("Performs a transfer using a durable nonce", func(t *testing.T) {
		noncePubkey := nonceWallet.PublicKey()

		nonceAccountBefore, err := GetNonceAccount(context.Background(), client, noncePubkey)
		require.NoError(t, err)

		tx, err := solana.NewTransaction(
			[]solana.Instruction{
				system.NewAdvanceNonceAccountInstruction(
					noncePubkey,
					solana.SysVarRecentBlockHashesPubkey,
					nonceAuth.PublicKey(),
				).Build(),
				system.NewTransferInstruction(
					uint64(0.01*float64(solana.LAMPORTS_PER_SOL)),
					sender.PublicKey(),
					receiver,
				).Build(),
			},
			solana.Hash(nonceAccountBefore.Nonce),
			solana.TransactionPayer(sender.PublicKey()),
		)
		require.NoError(t, err)

		_, err = SignTransaction(tx, nonceAuth.PrivateKey)
		require.NoError(t, err)

		_, err = SignTransaction(tx, sender.PrivateKey)
		require.NoError(t, err)

		signature, err := SendTransaction(client, context.Background(), tx)
		require.NoError(t, err)

		confirmedTx, err := WaitForTransactionConfirmation(
			context.Background(),
			client,
			signature,
			rpc.CommitmentConfirmed,
		)
		require.NoError(t, err)

		require.Empty(t, confirmedTx.Meta.Err) // The transaction succeeded if there's no error

		nonceAccountAfter, err := GetNonceAccount(context.Background(), client, noncePubkey)
		require.NoError(t, err)

		require.NotEqual(t, nonceAccountBefore.Nonce, nonceAccountAfter.Nonce)
	})
}

func TestCreateDurableNonceAccount(t *testing.T) {

	var client = rpc.New("https://api.devnet.solana.com")

	nonceAuthPubKey := solana.MustPublicKeyFromBase58("2zE66k3xtC7NXZ7Qb9YXYvZYRLhtPQaS9m23rWvXTHch")
	nonceAccPubKey := solana.MustPublicKeyFromBase58("FbZM91qUMLLj7ro8iAmdKSaCSiySeMzXQB7HJDnXA5zv")
	recent, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentConfirmed)
	require.NoError(t, err)
	require.NoError(t, err)
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewCreateAccountInstruction(
				uint64(0.0015*float64(solana.LAMPORTS_PER_SOL)),
				NONCE_ACCOUNT_LENGTH,
				solana.SystemProgramID,
				nonceAuthPubKey,
				nonceAccPubKey,
			).Build(),

			system.NewInitializeNonceAccountInstruction(
				nonceAuthPubKey,
				nonceAccPubKey,
				solana.SysVarRecentBlockHashesPubkey,
				solana.SysVarRentPubkey,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(nonceAuthPubKey),
	)
	require.NoError(t, err)
	bin, err := tx.Message.MarshalBinary()
	require.NoError(t, err)

	fmt.Printf("transaction : %s", hex.EncodeToString(bin))
}

func TestParseDurableNonceTransaction(t *testing.T) {
	txB64 := "AAIAAwXm3dWSnU+s1iGpvDycQw5mTg4smpU65kusWo4VRYRamexoOnfU4JeVvQ7ggopVP8es3TyNU3+VlFV5tngMQp03BqfVFxksVo7gioRfc9KXiM8DXDFFshqzRNgGLqlAAAAGp9UXGSxcUSGMyUw9SvF/WNruCJuh/UTj29mKAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAifT0s57LkkWQSpOS/6DVc9hstj1+edfJ0NXqI9eVk7UCBAIAATQAAAAAYOMWAAAAAABQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAMBAgMkBgAAAObd1ZKdT6zWIam8PJxDDmZODiyalTrmS6xajhVFhFqZ"

	//rawTx, err := base64.StdEncoding.DecodeString(txB64)
	//require.NoError(t, err)

	tx := &solana.Transaction{
		Message: solana.Message{},
	}

	err := tx.UnmarshalBase64(txB64)
	for _, i := range tx.Message.AccountKeys {
		fmt.Println(i.String())
	}
	require.NoError(t, err)
	//err = tx.Message.UnmarshalWithDecoder(bin.NewBinDecoder(rawTx))
	//require.NoError(t, err)
}
