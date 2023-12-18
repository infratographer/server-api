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

func TestCreate_servermotherboard(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	serial := gofakeit.UUID()
	srv := (&ServerBuilder{}).MustNew(ctx)
	srvmt := (&ServerMotherboardTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerMotherboardInput
		ExpectedMB *ent.ServerMotherboard
		errorMsg   string
	}{
		{
			TestName: "creates server motherboard ",
			Input:    graphclient.CreateServerMotherboardInput{Serial: serial, ServerID: srv.ID, ServerMotherboardTypeID: srvmt.ID},
			ExpectedMB: &ent.ServerMotherboard{
				Serial:                  serial,
				ServerID:                srv.ID,
				ServerMotherboardTypeID: srvmt.ID,
			},
		},
		{
			TestName: "fails to create server motherboard with empty serial",
			Input:    graphclient.CreateServerMotherboardInput{Serial: "", ServerID: srv.ID, ServerMotherboardTypeID: srvmt.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server motherboard with nil server ID",
			Input:    graphclient.CreateServerMotherboardInput{Serial: serial, ServerID: "", ServerMotherboardTypeID: srvmt.ID},
			errorMsg: "constraint failed",
		},
		{
			TestName: "fails to create server motherboard with nil server motherboard type ID",
			Input:    graphclient.CreateServerMotherboardInput{Serial: serial, ServerID: srv.ID, ServerMotherboardTypeID: ""},
			errorMsg: "constraint failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerMotherboardCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMotherboardCreate)

			createdMB := resp.ServerMotherboardCreate.ServerMotherboard
			assert.Equal(t, tt.ExpectedMB.Serial, createdMB.Serial)
			assert.Equal(t, tt.ExpectedMB.ServerID, createdMB.Server.ID)
			assert.Equal(t, "srvrmbd", createdMB.ID.Prefix())
		})
	}
}

func TestUpdate_servermotherboard(t *testing.T) {
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
		Input      graphclient.UpdateServerMotherboardInput
		ExpectedMB *ent.ServerMotherboard
		errorMsg   string
	}{
		{
			TestName: "updates server motherboard - serial",
			Input:    graphclient.UpdateServerMotherboardInput{Serial: &updateString},
			ExpectedMB: &ent.ServerMotherboard{
				Serial: updateString,
			},
		},
		{
			TestName: "fails to update to empty serial",
			Input:    graphclient.UpdateServerMotherboardInput{Serial: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			mt := (&ServerMotherboardBuilder{
				Serial: gofakeit.UUID(),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedMB, mt)

			resp, err := graphTestClient().ServerMotherboardUpdate(ctx, mt.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMotherboardUpdate)

			updatedMB := resp.ServerMotherboardUpdate.ServerMotherboard
			assert.Equal(t, tt.ExpectedMB.Serial, updatedMB.Serial)
			assert.Equal(t, tt.ExpectedMB.ServerID, updatedMB.Server.ID)
			assert.Equal(t, tt.ExpectedMB.ServerMotherboardTypeID, updatedMB.ServerMotherboardType.ID)
			assert.Equal(t, mt.ID, updatedMB.ID)
		})
	}
}

func TestDelete_servermotherboard(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	mb := (&ServerMotherboardBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server motherboard",
			Input:      mb.ID,
			ExpectedID: mb.ID,
		},
		{
			TestName: "fails to delete server motherboard that does not exist",
			Input:    gidx.PrefixedID("srvrmbd-1234"),
			errorMsg: "server_motherboard not found",
		},
		{
			TestName: "fails to delete empty server motherboard ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_motherboard not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerMotherboardDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMotherboardDelete)

			deletedMT := resp.ServerMotherboardDelete
			assert.EqualValues(t, tt.ExpectedID, deletedMT.DeletedID)
		})
	}
}

func TestFullServerMotherboardLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvmb := (&ServerMotherboardBuilder{}).MustNew(ctx)

	// create the Chassis
	createdSrvMTypeResp, err := graphTestClient().ServerMotherboardCreate(ctx, graphclient.CreateServerMotherboardInput{
		Serial:                  srvmb.Serial,
		ServerID:                srvmb.ServerID,
		ServerMotherboardTypeID: srvmb.ServerMotherboardTypeID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvMTypeResp)
	require.NotNil(t, createdSrvMTypeResp.ServerMotherboardCreate.ServerMotherboard)

	createdMB := createdSrvMTypeResp.ServerMotherboardCreate.ServerMotherboard
	require.NotNil(t, createdMB.ID)
	assert.Equal(t, "srvrmbd", createdMB.ID.Prefix())
	assert.Equal(t, srvmb.Serial, createdMB.Serial)
	assert.Equal(t, srvmb.ServerID, createdMB.Server.ID)
	assert.Equal(t, srvmb.ServerMotherboardTypeID, createdMB.ServerMotherboardType.ID)

	// Update the Motherboard
	newSerial := gofakeit.UUID()
	updatedMotherboardResp, err := graphTestClient().ServerMotherboardUpdate(ctx, createdMB.ID, graphclient.UpdateServerMotherboardInput{Serial: &newSerial})

	require.NoError(t, err)
	require.NotNil(t, updatedMotherboardResp)
	require.NotNil(t, updatedMotherboardResp.ServerMotherboardUpdate.ServerMotherboard)

	updatedMB := updatedMotherboardResp.ServerMotherboardUpdate.ServerMotherboard
	require.EqualValues(t, createdMB.ID, updatedMB.ID)
	require.Equal(t, newSerial, updatedMB.Serial)

	// Query the Motherboard
	queryMB, err := graphTestClient().GetServerMotherboard(ctx, gidx.PrefixedID(createdMB.ID))
	require.NoError(t, err)
	require.NotNil(t, queryMB)
	require.NotNil(t, queryMB.ServerMotherboard)
	require.Equal(t, newSerial, queryMB.ServerMotherboard.Serial)

	// Delete the Motherboard
	deletedResp, err := graphTestClient().ServerMotherboardDelete(ctx, createdMB.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerMotherboardDelete)
	require.EqualValues(t, createdMB.ID, deletedResp.ServerMotherboardDelete.DeletedID.String())

	// Query the Motherboard to ensure it's no longer available
	deletedCT, err := graphTestClient().GetServerMotherboard(ctx, createdMB.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_motherboard not found")
}
