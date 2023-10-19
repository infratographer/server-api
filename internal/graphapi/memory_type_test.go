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

func TestCreate_servermemorytype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vendor := gofakeit.Company()
	model := gofakeit.CarModel()
	size := gofakeit.Animal()
	speed := gofakeit.AnimalType()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerMemoryTypeInput
		ExpectedMT *ent.ServerMemoryType
		errorMsg   string
	}{
		{
			TestName: "creates server memory type ",
			Input:    graphclient.CreateServerMemoryTypeInput{Vendor: vendor, Model: model, Size: size, Speed: speed},
			ExpectedMT: &ent.ServerMemoryType{
				Vendor: vendor,
				Model:  model,
				Size:   size,
				Speed:  speed,
			},
		},
		{
			TestName: "fails to create server memory type with empty vendor",
			Input:    graphclient.CreateServerMemoryTypeInput{Vendor: "", Model: model, Size: size, Speed: speed},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server memory type with empty model",
			Input:    graphclient.CreateServerMemoryTypeInput{Vendor: vendor, Model: "", Size: size, Speed: speed},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server memory type with empty size",
			Input:    graphclient.CreateServerMemoryTypeInput{Vendor: vendor, Model: model, Size: "", Speed: speed},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server memory type with empty speed",
			Input:    graphclient.CreateServerMemoryTypeInput{Vendor: vendor, Model: model, Size: size, Speed: ""},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerMemoryTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMemoryTypeCreate)

			createdMT := resp.ServerMemoryTypeCreate.ServerMemoryType
			assert.Equal(t, tt.ExpectedMT.Vendor, createdMT.Vendor)
			assert.Equal(t, tt.ExpectedMT.Model, createdMT.Model)
			assert.Equal(t, tt.ExpectedMT.Size, createdMT.Size)
			assert.Equal(t, tt.ExpectedMT.Speed, createdMT.Speed)
			assert.Equal(t, "srvrmmt", createdMT.ID.Prefix())
		})
	}
}

func TestUpdate_servermemorytype(t *testing.T) {
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
		Input      graphclient.UpdateServerMemoryTypeInput
		ExpectedMT *ent.ServerMemoryType
		errorMsg   string
	}{
		{
			TestName: "updates server memory type - vendor",
			Input:    graphclient.UpdateServerMemoryTypeInput{Vendor: &updateString},
			ExpectedMT: &ent.ServerMemoryType{
				Vendor: updateString,
			},
		},
		{
			TestName: "fails to update to empty vendor",
			Input:    graphclient.UpdateServerMemoryTypeInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server memory type - model",
			Input:    graphclient.UpdateServerMemoryTypeInput{Model: &updateString},
			ExpectedMT: &ent.ServerMemoryType{
				Model: updateString,
			},
		},
		{
			TestName: "fails to update to empty model",
			Input:    graphclient.UpdateServerMemoryTypeInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server memory type - size",
			Input:    graphclient.UpdateServerMemoryTypeInput{Size: &updateString},
			ExpectedMT: &ent.ServerMemoryType{
				Size: updateString,
			},
		},
		{
			TestName: "fails to update to empty size",
			Input:    graphclient.UpdateServerMemoryTypeInput{Size: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server hard drive type - speed",
			Input:    graphclient.UpdateServerMemoryTypeInput{Speed: &updateString},
			ExpectedMT: &ent.ServerMemoryType{
				Speed: updateString,
			},
		},
		{
			TestName: "fails to update to empty speed",
			Input:    graphclient.UpdateServerMemoryTypeInput{Speed: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			mt := (&ServerMemoryTypeBuilder{
				Vendor: gofakeit.CarMaker(),
				Model:  gofakeit.CarModel(),
				Size:   gofakeit.DomainName(),
				Speed:  gofakeit.DomainName(),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedMT, mt)

			resp, err := graphTestClient().ServerMemoryTypeUpdate(ctx, mt.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMemoryTypeUpdate)

			updatedMT := resp.ServerMemoryTypeUpdate.ServerMemoryType
			assert.Equal(t, tt.ExpectedMT.Vendor, updatedMT.Vendor)
			assert.Equal(t, tt.ExpectedMT.Model, updatedMT.Model)
			assert.Equal(t, tt.ExpectedMT.Size, updatedMT.Size)
			assert.Equal(t, tt.ExpectedMT.Speed, updatedMT.Speed)
			assert.Equal(t, mt.ID, updatedMT.ID)
		})
	}
}

func TestDelete_servermemorytype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	mt := (&ServerMemoryTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server hard drive type",
			Input:      mt.ID,
			ExpectedID: mt.ID,
		},
		{
			TestName: "fails to delete server memory type that does not exist",
			Input:    gidx.PrefixedID("srvrmmt-1234"),
			errorMsg: "server_memory_type not found",
		},
		{
			TestName: "fails to delete empty server memory type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_memory_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerMemoryTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMemoryTypeDelete)

			deletedMT := resp.ServerMemoryTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedMT.DeletedID)
		})
	}
}

func TestFullServerMemoryTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvmt := (&ServerMemoryTypeBuilder{}).MustNew(ctx)

	// create the Chassis
	createdSrvMTypeResp, err := graphTestClient().ServerMemoryTypeCreate(ctx, graphclient.CreateServerMemoryTypeInput{
		Vendor: srvmt.Vendor,
		Model:  srvmt.Model,
		Speed:  srvmt.Speed,
		Size:   srvmt.Size,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvMTypeResp)
	require.NotNil(t, createdSrvMTypeResp.ServerMemoryTypeCreate.ServerMemoryType)

	createdMT := createdSrvMTypeResp.ServerMemoryTypeCreate.ServerMemoryType
	require.NotNil(t, createdMT.ID)
	assert.Equal(t, "srvrmmt", createdMT.ID.Prefix())
	assert.Equal(t, srvmt.Vendor, createdMT.Vendor)
	assert.Equal(t, srvmt.Model, createdMT.Model)
	assert.Equal(t, srvmt.Size, createdMT.Size)
	assert.Equal(t, srvmt.Speed, createdMT.Speed)

	// Update the Hard Drive Type
	newModel := gofakeit.DomainName()
	updatedMemoryTypeResp, err := graphTestClient().ServerMemoryTypeUpdate(ctx, createdMT.ID, graphclient.UpdateServerMemoryTypeInput{Model: &newModel})

	require.NoError(t, err)
	require.NotNil(t, updatedMemoryTypeResp)
	require.NotNil(t, updatedMemoryTypeResp.ServerMemoryTypeUpdate.ServerMemoryType)

	updatedMT := updatedMemoryTypeResp.ServerMemoryTypeUpdate.ServerMemoryType
	require.EqualValues(t, createdMT.ID, updatedMT.ID)
	require.Equal(t, newModel, updatedMT.Model)

	// Query the Chassis
	queryMT, err := graphTestClient().GetServerMemoryType(ctx, gidx.PrefixedID(createdMT.ID))
	require.NoError(t, err)
	require.NotNil(t, queryMT)
	require.NotNil(t, queryMT.ServerMemoryType)
	require.Equal(t, newModel, queryMT.ServerMemoryType.Model)

	// Delete the Server Chassis
	deletedResp, err := graphTestClient().ServerMemoryTypeDelete(ctx, createdMT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerMemoryTypeDelete)
	require.EqualValues(t, createdMT.ID, deletedResp.ServerMemoryTypeDelete.DeletedID.String())

	// Query the CPU to ensure it's no longer available
	deletedCT, err := graphTestClient().GetServerMemoryType(ctx, createdMT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_memory_type not found")
}
