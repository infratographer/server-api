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

func TestCreate_servermotherboardtype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vendor := gofakeit.Company()
	model := gofakeit.CarModel()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerMotherboardTypeInput
		ExpectedMT *ent.ServerMotherboardType
		errorMsg   string
	}{
		{
			TestName: "creates server motherboard type ",
			Input:    graphclient.CreateServerMotherboardTypeInput{Vendor: vendor, Model: model},
			ExpectedMT: &ent.ServerMotherboardType{
				Vendor: vendor,
				Model:  model,
			},
		},
		{
			TestName: "fails to create server motherboard type with empty vendor",
			Input:    graphclient.CreateServerMotherboardTypeInput{Vendor: "", Model: model},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server motherboard type with empty model",
			Input:    graphclient.CreateServerMotherboardTypeInput{Vendor: vendor, Model: ""},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerMotherboardTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMotherboardTypeCreate)

			createdMT := resp.ServerMotherboardTypeCreate.ServerMotherboardType
			assert.Equal(t, tt.ExpectedMT.Vendor, createdMT.Vendor)
			assert.Equal(t, tt.ExpectedMT.Model, createdMT.Model)
			assert.Equal(t, "srvrmbt", createdMT.ID.Prefix())
		})
	}
}

func TestUpdate_servermotherboardtype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	updateString := gofakeit.UUID()
	emptyString := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerMotherboardTypeInput
		ExpectedMT *ent.ServerMotherboardType
		errorMsg   string
	}{
		{
			TestName: "updates server motherboard type - vendor",
			Input:    graphclient.UpdateServerMotherboardTypeInput{Vendor: &updateString},
			ExpectedMT: &ent.ServerMotherboardType{
				Vendor: updateString,
			},
		},
		{
			TestName: "fails to update to empty vendor",
			Input:    graphclient.UpdateServerMotherboardTypeInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server motherboard type - model",
			Input:    graphclient.UpdateServerMotherboardTypeInput{Model: &updateString},
			ExpectedMT: &ent.ServerMotherboardType{
				Model: updateString,
			},
		},
		{
			TestName: "fails to update to empty model",
			Input:    graphclient.UpdateServerMotherboardTypeInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			mt := (&ServerMotherboardTypeBuilder{
				Vendor: gofakeit.CarMaker(),
				Model:  gofakeit.CarModel(),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedMT, mt)

			resp, err := graphTestClient().ServerMotherboardTypeUpdate(ctx, mt.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMotherboardTypeUpdate)

			updatedMT := resp.ServerMotherboardTypeUpdate.ServerMotherboardType
			assert.Equal(t, tt.ExpectedMT.Vendor, updatedMT.Vendor)
			assert.Equal(t, tt.ExpectedMT.Model, updatedMT.Model)
			assert.Equal(t, mt.ID, updatedMT.ID)
		})
	}
}

func TestDelete_servermotherboardtype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	mt := (&ServerMotherboardTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server motherboard type",
			Input:      mt.ID,
			ExpectedID: mt.ID,
		},
		{
			TestName: "fails to delete server motherboard type that does not exist",
			Input:    gidx.PrefixedID("srvrmbt-1234"),
			errorMsg: "server_motherboard_type not found",
		},
		{
			TestName: "fails to delete empty server motherboard type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_motherboard_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerMotherboardTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMotherboardTypeDelete)

			deletedMT := resp.ServerMotherboardTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedMT.DeletedID)
		})
	}
}

func TestFullServerMotherboardTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvmt := (&ServerMotherboardTypeBuilder{}).MustNew(ctx)

	// create the Chassis
	createdSrvMTypeResp, err := graphTestClient().ServerMotherboardTypeCreate(ctx, graphclient.CreateServerMotherboardTypeInput{
		Vendor: srvmt.Vendor,
		Model:  srvmt.Model,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvMTypeResp)
	require.NotNil(t, createdSrvMTypeResp.ServerMotherboardTypeCreate.ServerMotherboardType)

	createdMT := createdSrvMTypeResp.ServerMotherboardTypeCreate.ServerMotherboardType
	require.NotNil(t, createdMT.ID)
	assert.Equal(t, "srvrmbt", createdMT.ID.Prefix())
	assert.Equal(t, srvmt.Vendor, createdMT.Vendor)
	assert.Equal(t, srvmt.Model, createdMT.Model)

	// Update the Motherboard Type
	newModel := gofakeit.DomainName()
	updatedMotherboardTypeResp, err := graphTestClient().ServerMotherboardTypeUpdate(ctx, createdMT.ID, graphclient.UpdateServerMotherboardTypeInput{Model: &newModel})

	require.NoError(t, err)
	require.NotNil(t, updatedMotherboardTypeResp)
	require.NotNil(t, updatedMotherboardTypeResp.ServerMotherboardTypeUpdate.ServerMotherboardType)

	updatedMT := updatedMotherboardTypeResp.ServerMotherboardTypeUpdate.ServerMotherboardType
	require.EqualValues(t, createdMT.ID, updatedMT.ID)
	require.Equal(t, newModel, updatedMT.Model)

	// Query the Motherboard
	queryMT, err := graphTestClient().GetServerMotherboardType(ctx, gidx.PrefixedID(createdMT.ID))
	require.NoError(t, err)
	require.NotNil(t, queryMT)
	require.NotNil(t, queryMT.ServerMotherboardType)
	require.Equal(t, newModel, queryMT.ServerMotherboardType.Model)

	// Delete the Motherboard type
	deletedResp, err := graphTestClient().ServerMotherboardTypeDelete(ctx, createdMT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerMotherboardTypeDelete)
	require.EqualValues(t, createdMT.ID, deletedResp.ServerMotherboardTypeDelete.DeletedID.String())

	// Query the Motherboard to ensure it's no longer available
	deletedCT, err := graphTestClient().GetServerMotherboardType(ctx, createdMT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_motherboard_type not found")
}
