package graphapi_test

import (
	"context"
	"testing"

	"dario.cat/mergo"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/permissions-api/pkg/permissions/mockpermissions"
	ent "go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/server-api/internal/graphclient"
	"go.infratographer.com/x/gidx"
)

func TestCreate_serverharddrive(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srv := (&ServerBuilder{}).MustNew(ctx)
	srvhdt := (&ServerHardDriveTypeBuilder{}).MustNew(ctx)
	serial := gofakeit.UUID()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerHardDriveInput
		ExpectedHD *ent.ServerHardDrive
		errorMsg   string
	}{
		{
			TestName: "creates server hard drive",
			Input:    graphclient.CreateServerHardDriveInput{Serial: serial, ServerID: srv.ID, ServerHardDriveTypeID: srvhdt.ID},
			ExpectedHD: &ent.ServerHardDrive{
				Serial:                serial,
				ServerID:              srv.ID,
				ServerHardDriveTypeID: srvhdt.ID,
			},
		},
		{
			TestName: "fails to create server component type with empty serial",
			Input:    graphclient.CreateServerHardDriveInput{Serial: "", ServerID: srv.ID, ServerHardDriveTypeID: srvhdt.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty server id",
			Input:    graphclient.CreateServerHardDriveInput{Serial: serial, ServerID: "", ServerHardDriveTypeID: srvhdt.ID},
			errorMsg: "constraint failed",
		},
		{
			TestName: "fails to create server component type with empty hard drive type id",
			Input:    graphclient.CreateServerHardDriveInput{Serial: serial, ServerID: srv.ID, ServerHardDriveTypeID: ""},
			errorMsg: "constraint failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerHardDriveCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerHardDriveCreate)

			createdHD := resp.ServerHardDriveCreate.ServerHardDrive
			assert.Equal(t, tt.ExpectedHD.Serial, createdHD.Serial)
			assert.Equal(t, tt.ExpectedHD.ServerID, createdHD.Server.ID)
			assert.Equal(t, tt.ExpectedHD.ServerHardDriveTypeID, createdHD.ServerHardDriveType.ID)
			assert.Equal(t, "srvrshd", createdHD.ID.Prefix())
		})
	}
}

func TestUpdate_serverharddrive(t *testing.T) {
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
		Input      graphclient.UpdateServerHardDriveInput
		ExpectedHD *ent.ServerHardDrive
		errorMsg   string
	}{
		{
			TestName: "updates server hard drive type - serial",
			Input:    graphclient.UpdateServerHardDriveInput{Serial: &updateString},
			ExpectedHD: &ent.ServerHardDrive{
				Serial: updateString,
			},
		},
		{
			TestName: "updates server hard drive - empty serial",
			Input:    graphclient.UpdateServerHardDriveInput{Serial: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			hd := (&ServerHardDriveBuilder{
				Serial:              gofakeit.UUID(),
				Server:              (&ServerBuilder{}).MustNew(ctx),
				ServerHardDriveType: (&ServerHardDriveTypeBuilder{}).MustNew(ctx),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedHD, hd)

			resp, err := graphTestClient().ServerHardDriveUpdate(ctx, hd.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerHardDriveUpdate)

			updatedHD := resp.ServerHardDriveUpdate.ServerHardDrive
			assert.Equal(t, tt.ExpectedHD.Serial, updatedHD.Serial)
			assert.Equal(t, tt.ExpectedHD.ServerID, updatedHD.Server.ID)
			assert.Equal(t, tt.ExpectedHD.ServerHardDriveTypeID, updatedHD.ServerHardDriveType.ID)
			assert.Equal(t, hd.ID, updatedHD.ID)
		})
	}
}

func TestDelete_serverharddrive(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	hd := (&ServerHardDriveBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server hard drive type",
			Input:      hd.ID,
			ExpectedID: hd.ID,
		},
		{
			TestName: "fails to delete server hard drive that does not exist",
			Input:    gidx.PrefixedID("srvrshd-1234"),
			errorMsg: "server_hard_drive not found",
		},
		{
			TestName: "fails to delete empty server hard drive ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_hard_drive not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerHardDriveDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerHardDriveDelete)

			deletedHD := resp.ServerHardDriveDelete
			assert.EqualValues(t, tt.ExpectedID, deletedHD.DeletedID)
		})
	}
}
