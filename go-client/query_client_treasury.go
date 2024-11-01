package client

import (
	"context"

	"github.com/Zenrock-Foundation/zrchain/v5/x/treasury/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc"
)

type PageRequest = query.PageRequest

// TreasuryQueryClient is the client for the treasury module.
type TreasuryQueryClient struct {
	client types.QueryClient
}

// NewTreasuryQueryClient returns a TreasuryQueryClient
func NewTreasuryQueryClient(c *grpc.ClientConn) *TreasuryQueryClient {
	return &TreasuryQueryClient{
		client: types.NewQueryClient(c),
	}
}

// PendingKeyRequests executes a paginated pending key request query with context. zenrockd will return a slice of pending
// key requests for the supplied keyring address.
func (t *TreasuryQueryClient) PendingKeyRequests(ctx context.Context, page *PageRequest, keyringAddr string) ([]*types.KeyReqResponse, error) {
	res, err := t.client.KeyRequests(ctx, &types.QueryKeyRequestsRequest{
		Pagination:  page,
		KeyringAddr: keyringAddr,
		Status:      types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
	})
	if err != nil {
		return nil, err
	}

	return res.KeyRequests, nil
}

func (t *TreasuryQueryClient) GetKeys(ctx context.Context, offset uint64, pagesize uint64) (*types.QueryKeysResponse, error) {
	res, err := t.client.Keys(ctx, &types.QueryKeysRequest{
		Pagination: &query.PageRequest{Offset: offset, Limit: pagesize},
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *TreasuryQueryClient) GetKey(ctx context.Context, keyId uint64) (*types.QueryKeyByIDResponse, error) {
	res, err := t.client.KeyByID(ctx, &types.QueryKeyByIDRequest{
		Id: keyId,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetKeyRequest returns the key request corresponding to the specific request ID.
func (t *TreasuryQueryClient) GetKeyRequest(ctx context.Context, requestID uint64) (*types.KeyReqResponse, error) {
	res, err := t.client.KeyRequestByID(ctx, &types.QueryKeyRequestByIDRequest{
		Id: requestID,
	})
	if err != nil {
		return nil, err
	}

	return res.KeyRequest, nil
}

// PendingSignatureRequests executes a paginated pending signature request query with context. zenrockd will return a slice of pending
// signature requests for the supplied keyring address.
func (t *TreasuryQueryClient) PendingSignatureRequests(ctx context.Context, page *PageRequest, keyringAddr string) ([]*types.SignReqResponse, error) {
	res, err := t.client.SignatureRequests(ctx, &types.QuerySignatureRequestsRequest{
		Pagination:  page,
		KeyringAddr: keyringAddr,
		Status:      types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
	})
	if err != nil {
		return nil, err
	}

	return res.SignRequests, nil
}

func (t *TreasuryQueryClient) FulfilledSignatureRequests(ctx context.Context, offset uint64, pagesize uint64) (*types.QuerySignatureRequestsResponse, error) {
	res, err := t.client.SignatureRequests(ctx, &types.QuerySignatureRequestsRequest{
		Pagination: &query.PageRequest{Offset: offset, Limit: pagesize},
		Status:     types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetSignatureRequest returns the signature request corresponding to the specific request ID.
func (t *TreasuryQueryClient) GetSignatureRequest(ctx context.Context, requestID uint64) (*types.SignReqResponse, error) {
	res, err := t.client.SignatureRequestByID(ctx, &types.QuerySignatureRequestByIDRequest{
		Id: requestID,
	})
	if err != nil {
		return nil, err
	}

	return res.SignRequest, nil
}

// SignedTransactions returns a paginated set of fulfilled signature requests for the supplied wallet type.
func (t *TreasuryQueryClient) SignedTransactions(ctx context.Context, page *PageRequest, walletType types.WalletType) (*types.QuerySignTransactionRequestsResponse, error) {
	res, err := t.client.SignTransactionRequests(ctx, &types.QuerySignTransactionRequestsRequest{
		Pagination: page,
		WalletType: walletType,
		Status:     types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *TreasuryQueryClient) ZrSignKeys(ctx context.Context, page *PageRequest, address, walletType string) (*types.QueryZrSignKeysResponse, error) {
	res, err := t.client.ZrSignKeys(ctx, &types.QueryZrSignKeysRequest{
		Address:    address,
		WalletType: walletType,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}