package keeper_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zenrocklabs/zenrock/zrchain/v4/testutil/keeper"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/identity/keeper"
	identity "github.com/zenrocklabs/zenrock/zrchain/v4/x/identity/module"
	"github.com/zenrocklabs/zenrock/zrchain/v4/x/identity/types"
)

func Test_msgServer_RemoveKeyringParty(t *testing.T) {

	var defaultKrWithParties = types.Keyring{
		Address:     "keyring1pfnq7r04rept47gaf5cpdew2",
		Creator:     "testCreator",
		Description: "testDescription",
		Admins:      []string{"testCreator"},
		Parties:     []string{"party1", "party2", "party3"},
		KeyReqFee:   0,
		SigReqFee:   0,
		IsActive:    true,
	}

	type args struct {
		keyring *types.Keyring
		msg     *types.MsgRemoveKeyringParty
	}
	tests := []struct {
		name        string
		args        args
		want        *types.MsgRemoveKeyringPartyResponse
		wantKeyring *types.Keyring
		wantErr     bool
	}{
		{
			name: "PASS: remove keyring party",
			args: args{
				keyring: &defaultKrWithParties,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "keyring1pfnq7r04rept47gaf5cpdew2", "party3"),
			},
			want: &types.MsgRemoveKeyringPartyResponse{},
			wantKeyring: &types.Keyring{
				Address:     "keyring1pfnq7r04rept47gaf5cpdew2",
				Creator:     "testCreator",
				Description: "testDescription",
				Admins:      []string{"testCreator"},
				Parties:     []string{"party1", "party2"},
				KeyReqFee:   0,
				SigReqFee:   0,
				IsActive:    true,
			},
		},
		{
			name: "PASS: remove single party",
			args: args{
				keyring: &types.Keyring{
					Address:        "keyring1pfnq7r04rept47gaf5cpdew2",
					Creator:        "testCreator",
					Description:    "testDescription",
					Admins:         []string{"testCreator"},
					Parties:        []string{"party1", "party2"},
					PartyThreshold: 2,
					KeyReqFee:      0,
					SigReqFee:      0,
					IsActive:       true,
				},
				msg: types.NewMsgRemoveKeyringParty("testCreator", "keyring1pfnq7r04rept47gaf5cpdew2", "party2"),
			},
			want: &types.MsgRemoveKeyringPartyResponse{},
			wantKeyring: &types.Keyring{
				Address:        "keyring1pfnq7r04rept47gaf5cpdew2",
				Creator:        "testCreator",
				Description:    "testDescription",
				Admins:         []string{"testCreator"},
				Parties:        []string{"party1"},
				PartyThreshold: 1,
				KeyReqFee:      0,
				SigReqFee:      0,
				IsActive:       true,
			},
		},
		{
			name: "FAIL: keyring is nil or not found",
			args: args{
				keyring: &defaultKrWithParties,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "notAKeyring", "party1"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: removed party is no party",
			args: args{
				keyring: &defaultKrWithParties,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "keyring1pfnq7r04rept47gaf5cpdew2", "noParty"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: creator is no admin owner",
			args: args{
				keyring: &defaultKrWithParties,
				msg:     types.NewMsgRemoveKeyringParty("noAdmin", "keyring1pfnq7r04rept47gaf5cpdew2", "party1"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyring is not active",
			args: args{
				keyring: &types.Keyring{
					Address:     "keyring1pfnq7r04rept47gaf5cpdew2",
					Creator:     "testCreator",
					Description: "testDescription",
					Admins:      []string{"testCreator", "admin1", "admin2", "admin3"},
					Parties:     []string{"party1", "party2", "party3"},
					KeyReqFee:   0,
					SigReqFee:   0,
					IsActive:    false,
				},
				msg: types.NewMsgRemoveKeyringParty("noAdmin", "keyring1pfnq7r04rept47gaf5cpdew2", "party1"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx
			msgSer := keeper.NewMsgServerImpl(*ik)

			genesis := types.GenesisState{
				PortId:   types.PortID,
				Keyrings: []types.Keyring{*tt.args.keyring},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			got, err := msgSer.RemoveKeyringParty(ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RemoveKeyringParty() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RemoveKeyringParty() got = %v, want %v", got, tt.want)
				}

				gotKeyring, err := ik.KeyringStore.Get(ctx, tt.args.keyring.Address)
				require.NoError(t, err)

				require.Equal(t, tt.wantKeyring, &gotKeyring)
			}
		})
	}
}
