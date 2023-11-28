package graphapi_test

import (
	"context"
	"testing"

	"dario.cat/mergo"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/permissions-api/pkg/permissions/mockpermissions"
	ent "go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/server-api/internal/graphclient"
	"go.infratographer.com/x/gidx"
)

func TestCreate_servernetworkcardtype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vendor := gofakeit.Company()
	model := gofakeit.CarModel()

	testCases := []struct {
		TestName    string
		Input       graphclient.CreateServerNetworkCardTypeInput
		ExpectedNCT *ent.ServerNetworkCardType
		errorMsg    string
	}{
		{
			TestName: "creates server network card type ",
			Input:    graphclient.CreateServerNetworkCardTypeInput{Vendor: vendor, Model: model, PortCount: 4},
			ExpectedNCT: &ent.ServerNetworkCardType{
				Vendor:    vendor,
				Model:     model,
				PortCount: 4,
			},
		},
		{
			TestName: "fails to create server network card type with empty vendor",
			Input:    graphclient.CreateServerNetworkCardTypeInput{Vendor: "", Model: model, PortCount: 2},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server network card type with empty model",
			Input:    graphclient.CreateServerNetworkCardTypeInput{Vendor: vendor, Model: "", PortCount: 1},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server network card type with zero port count",
			Input:    graphclient.CreateServerNetworkCardTypeInput{Vendor: vendor, Model: model, PortCount: 0},
			errorMsg: "value out of range",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerNetworkCardTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerNetworkCardTypeCreate)

			createdNCT := resp.ServerNetworkCardTypeCreate.ServerNetworkCardType
			assert.Equal(t, tt.ExpectedNCT.Vendor, createdNCT.Vendor)
			assert.Equal(t, tt.ExpectedNCT.Model, createdNCT.Model)
			assert.Equal(t, "srvrnct", createdNCT.ID.Prefix())
		})
	}
}

func TestUpdate_servernetworkcardtype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	updateString := gofakeit.UUID()
	emptyString := ""

	testCases := []struct {
		TestName    string
		Input       graphclient.UpdateServerNetworkCardTypeInput
		ExpectedNCT *ent.ServerNetworkCardType
		errorMsg    string
	}{
		{
			TestName: "updates server motherboard type - vendor",
			Input:    graphclient.UpdateServerNetworkCardTypeInput{Vendor: &updateString},
			ExpectedNCT: &ent.ServerNetworkCardType{
				Vendor: updateString,
			},
		},
		{
			TestName: "fails to update to empty vendor",
			Input:    graphclient.UpdateServerNetworkCardTypeInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server network card type - model",
			Input:    graphclient.UpdateServerNetworkCardTypeInput{Model: &updateString},
			ExpectedNCT: &ent.ServerNetworkCardType{
				Model: updateString,
			},
		},
		{
			TestName: "fails to update to empty model",
			Input:    graphclient.UpdateServerNetworkCardTypeInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			nct := (&NetworkCardTypeBuilder{
				Vendor: gofakeit.CarMaker(),
				Model:  gofakeit.CarModel(),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedNCT, nct)

			resp, err := graphTestClient().ServerNetworkCardTypeUpdate(ctx, nct.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerNetworkCardTypeUpdate)

			updatedNCT := resp.ServerNetworkCardTypeUpdate.ServerNetworkCardType
			assert.Equal(t, tt.ExpectedNCT.Vendor, updatedNCT.Vendor)
			assert.Equal(t, tt.ExpectedNCT.Model, updatedNCT.Model)
			assert.Equal(t, nct.ID, updatedNCT.ID)
		})
	}
}

func TestDelete_servernetworkcardtype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	nct := (&NetworkCardTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server network card type",
			Input:      nct.ID,
			ExpectedID: nct.ID,
		},
		{
			TestName: "fails to delete server network card type that does not exist",
			Input:    gidx.PrefixedID("srvrnct-1234"),
			errorMsg: "server_network_card_type not found",
		},
		{
			TestName: "fails to delete empty server network card type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_network_card_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerNetworkCardTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerNetworkCardTypeDelete)

			deletedMT := resp.ServerNetworkCardTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedMT.DeletedID)
		})
	}
}
