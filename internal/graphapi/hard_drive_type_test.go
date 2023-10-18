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
	"go.infratographer.com/server-api/internal/ent/generated/serverharddrivetype"
	"go.infratographer.com/server-api/internal/graphclient"
	"go.infratographer.com/x/gidx"
)

func TestCreate_serverharddrivetype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vendor := gofakeit.Company()
	model := gofakeit.CarModel()
	capacity := "100 GB"
	speed := "7200 RPM"
	stype := graphclient.ServerHardDriveTypeTypeSsd

	testCases := []struct {
		TestName    string
		Input       graphclient.CreateServerHardDriveTypeInput
		ExpectedHDT *ent.ServerHardDriveType
		errorMsg    string
	}{
		{
			TestName: "creates server hard drive type ",
			Input:    graphclient.CreateServerHardDriveTypeInput{Vendor: vendor, Model: model, Capacity: capacity, Speed: speed, Type: stype},
			ExpectedHDT: &ent.ServerHardDriveType{
				Vendor:   vendor,
				Model:    model,
				Capacity: capacity,
				Speed:    speed,
				Type:     serverharddrivetype.TypeSsd,
			},
		},
		{
			TestName: "fails to create server component type with empty vendor",
			Input:    graphclient.CreateServerHardDriveTypeInput{Vendor: "", Model: model, Capacity: capacity, Speed: speed, Type: stype},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty model",
			Input:    graphclient.CreateServerHardDriveTypeInput{Vendor: vendor, Model: "", Capacity: capacity, Speed: speed, Type: stype},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty capacity",
			Input:    graphclient.CreateServerHardDriveTypeInput{Vendor: vendor, Model: model, Capacity: "", Speed: speed, Type: stype},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty speed",
			Input:    graphclient.CreateServerHardDriveTypeInput{Vendor: vendor, Model: model, Capacity: capacity, Speed: "", Type: stype},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty invalid type",
			Input:    graphclient.CreateServerHardDriveTypeInput{Vendor: vendor, Model: model, Capacity: capacity, Speed: speed, Type: ""},
			errorMsg: "is not a valid ServerHardDriveTypeType",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerHardDriveTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerHardDriveTypeCreate)

			createdHDT := resp.ServerHardDriveTypeCreate.ServerHardDriveType
			assert.Equal(t, tt.ExpectedHDT.Vendor, createdHDT.Vendor)
			assert.Equal(t, tt.ExpectedHDT.Model, createdHDT.Model)
			assert.Equal(t, tt.ExpectedHDT.Capacity, createdHDT.Capacity)
			assert.Equal(t, tt.ExpectedHDT.Speed, createdHDT.Speed)
			assert.Equal(t, tt.ExpectedHDT.Type.String(), createdHDT.Type.String())
			assert.Equal(t, "srvrhdt", createdHDT.ID.Prefix())
		})
	}
}

func TestUpdate_serverharddrivetype(t *testing.T) {
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
		Input       graphclient.UpdateServerHardDriveTypeInput
		ExpectedHDT *ent.ServerHardDriveType
		errorMsg    string
	}{
		{
			TestName: "updates server hard drive type - vendor",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Vendor: &updateString},
			ExpectedHDT: &ent.ServerHardDriveType{
				Vendor: updateString,
			},
		},
		{
			TestName: "fails to update to empty vendor",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server hard drive type - model",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Model: &updateString},
			ExpectedHDT: &ent.ServerHardDriveType{
				Model: updateString,
			},
		},
		{
			TestName: "fails to update to empty model",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server hard drive type - capacity",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Capacity: &updateString},
			ExpectedHDT: &ent.ServerHardDriveType{
				Capacity: updateString,
			},
		},
		{
			TestName: "fails to update to empty capacity",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Capacity: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server hard drive type - speed",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Speed: &updateString},
			ExpectedHDT: &ent.ServerHardDriveType{
				Speed: updateString,
			},
		},
		{
			TestName: "fails to update to empty speed",
			Input:    graphclient.UpdateServerHardDriveTypeInput{Speed: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			hdt := (&ServerHardDriveTypeBuilder{
				Vendor:   gofakeit.CarMaker(),
				Model:    gofakeit.CarModel(),
				Capacity: gofakeit.DomainName(),
				Speed:    gofakeit.DomainName(),
				Type:     serverharddrivetype.TypeHdd,
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedHDT, hdt)

			resp, err := graphTestClient().ServerHardDriveTypeUpdate(ctx, hdt.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerHardDriveTypeUpdate)

			updatedHDT := resp.ServerHardDriveTypeUpdate.ServerHardDriveType
			assert.Equal(t, tt.ExpectedHDT.Vendor, updatedHDT.Vendor)
			assert.Equal(t, tt.ExpectedHDT.Model, updatedHDT.Model)
			assert.Equal(t, tt.ExpectedHDT.Capacity, updatedHDT.Capacity)
			assert.Equal(t, tt.ExpectedHDT.Speed, updatedHDT.Speed)
			assert.Equal(t, hdt.ID, updatedHDT.ID)
		})
	}
}

func TestDelete_serverharddrivetype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	hdt := (&ServerHardDriveTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server hard drive type",
			Input:      hdt.ID,
			ExpectedID: hdt.ID,
		},
		{
			TestName: "fails to delete server hard drive type that does not exist",
			Input:    gidx.PrefixedID("srvrhdt-1234"),
			errorMsg: "server_hard_drive_type not found",
		},
		{
			TestName: "fails to delete empty server hard drive type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_hard_drive_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerHardDriveTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerHardDriveTypeDelete)

			deletedCT := resp.ServerHardDriveTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerHardDriveTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvhdt := (&ServerHardDriveTypeBuilder{}).MustNew(ctx)

	// create the Chassis
	createdSrvHDTypeResp, err := graphTestClient().ServerHardDriveTypeCreate(ctx, graphclient.CreateServerHardDriveTypeInput{
		Vendor:   srvhdt.Vendor,
		Model:    srvhdt.Model,
		Speed:    srvhdt.Speed,
		Type:     graphclient.ServerHardDriveTypeTypeSsd,
		Capacity: srvhdt.Capacity,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvHDTypeResp)
	require.NotNil(t, createdSrvHDTypeResp.ServerHardDriveTypeCreate.ServerHardDriveType)

	createdHDT := createdSrvHDTypeResp.ServerHardDriveTypeCreate.ServerHardDriveType
	require.NotNil(t, createdHDT.ID)
	assert.Equal(t, "srvrhdt", createdHDT.ID.Prefix())
	assert.Equal(t, srvhdt.Vendor, createdHDT.Vendor)
	assert.Equal(t, srvhdt.Model, createdHDT.Model)
	assert.Equal(t, srvhdt.Capacity, createdHDT.Capacity)
	assert.Equal(t, srvhdt.Speed, createdHDT.Speed)
	assert.Equal(t, serverharddrivetype.TypeSsd.String(), createdHDT.Type.String())

	// Update the Hard Drive Type
	newModel := gofakeit.DomainName()
	updatedHardDriveTypeResp, err := graphTestClient().ServerHardDriveTypeUpdate(ctx, createdHDT.ID, graphclient.UpdateServerHardDriveTypeInput{Model: &newModel})

	require.NoError(t, err)
	require.NotNil(t, updatedHardDriveTypeResp)
	require.NotNil(t, updatedHardDriveTypeResp.ServerHardDriveTypeUpdate.ServerHardDriveType)

	updatedHDT := updatedHardDriveTypeResp.ServerHardDriveTypeUpdate.ServerHardDriveType
	require.EqualValues(t, createdHDT.ID, updatedHDT.ID)
	require.Equal(t, newModel, updatedHDT.Model)

	// Query the Chassis
	queryHDT, err := graphTestClient().GetHardDriveType(ctx, gidx.PrefixedID(createdHDT.ID))
	require.NoError(t, err)
	require.NotNil(t, queryHDT)
	require.NotNil(t, queryHDT.ServerHardDriveType)
	require.Equal(t, newModel, queryHDT.ServerHardDriveType.Model)

	// Delete the Server Chassis
	deletedResp, err := graphTestClient().ServerHardDriveTypeDelete(ctx, createdHDT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerHardDriveTypeDelete)
	require.EqualValues(t, createdHDT.ID, deletedResp.ServerHardDriveTypeDelete.DeletedID.String())

	// Query the CPU to ensure it's no longer available
	deletedCT, err := graphTestClient().GetHardDriveType(ctx, createdHDT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_hard_drive_type not found")
}
