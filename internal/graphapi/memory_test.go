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

func TestCreate_servermemory(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	serial := gofakeit.UUID()
	srv := (&ServerBuilder{}).MustNew(ctx)
	smt := (&ServerMemoryTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName  string
		Input     graphclient.CreateServerMemoryInput
		ExpectedM *ent.ServerMemory
		errorMsg  string
	}{
		{
			TestName: "creates server memory ",
			Input:    graphclient.CreateServerMemoryInput{Serial: serial, ServerID: srv.ID, ServerMemoryTypeID: smt.ID},
			ExpectedM: &ent.ServerMemory{
				Serial:             serial,
				ServerID:           srv.ID,
				ServerMemoryTypeID: smt.ID,
			},
		},
		{
			TestName: "fails to create server memory with empty serial",
			Input:    graphclient.CreateServerMemoryInput{Serial: "", ServerID: srv.ID, ServerMemoryTypeID: smt.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server memory with empty server",
			Input:    graphclient.CreateServerMemoryInput{Serial: serial, ServerID: "", ServerMemoryTypeID: smt.ID},
			errorMsg: "constraint failed",
		},
		{
			TestName: "fails to create server memory with empty memory type",
			Input:    graphclient.CreateServerMemoryInput{Serial: serial, ServerID: srv.ID, ServerMemoryTypeID: ""},
			errorMsg: "constraint failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerMemoryCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMemoryCreate)

			createdM := resp.ServerMemoryCreate.ServerMemory
			assert.Equal(t, tt.ExpectedM.Serial, createdM.Serial)
			assert.Equal(t, tt.ExpectedM.ServerID, createdM.Server.ID)
			assert.Equal(t, tt.ExpectedM.ServerMemoryTypeID, createdM.ServerMemoryType.ID)
			assert.Equal(t, "srvrmem", createdM.ID.Prefix())
		})
	}
}

func TestUpdate_servermemory(t *testing.T) {
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
		Input      graphclient.UpdateServerMemoryInput
		ExpectedMT *ent.ServerMemory
		errorMsg   string
	}{
		{
			TestName: "updates server memory - serial",
			Input:    graphclient.UpdateServerMemoryInput{Serial: &updateString},
			ExpectedMT: &ent.ServerMemory{
				Serial: updateString,
			},
		},
		{
			TestName: "fails to update to empty serial",
			Input:    graphclient.UpdateServerMemoryInput{Serial: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			mt := (&ServerMemoryBuilder{}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedMT, mt)

			resp, err := graphTestClient().ServerMemoryUpdate(ctx, mt.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMemoryUpdate)

			updatedMT := resp.ServerMemoryUpdate.ServerMemory
			assert.Equal(t, tt.ExpectedMT.Serial, updatedMT.Serial)
			assert.Equal(t, tt.ExpectedMT.ServerID, updatedMT.Server.ID)
			assert.Equal(t, tt.ExpectedMT.ServerMemoryTypeID, updatedMT.ServerMemoryType.ID)
			assert.Equal(t, mt.ID, updatedMT.ID)
		})
	}
}

func TestDelete_servermemory(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	m := (&ServerMemoryBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server memory",
			Input:      m.ID,
			ExpectedID: m.ID,
		},
		{
			TestName: "fails to delete server memory that does not exist",
			Input:    gidx.PrefixedID("srvrmem-1234"),
			errorMsg: "server_memory not found",
		},
		{
			TestName: "fails to delete empty server memory ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_memory not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerMemoryDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerMemoryDelete)

			deletedM := resp.ServerMemoryDelete
			assert.EqualValues(t, tt.ExpectedID, deletedM.DeletedID)
		})
	}
}

func TestFullServerMemoryLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	serial := gofakeit.UUID()
	srvmt := (&ServerMemoryTypeBuilder{}).MustNew(ctx)
	srv := (&ServerBuilder{}).MustNew(ctx)

	// create the Chassis
	createdSrvMResp, err := graphTestClient().ServerMemoryCreate(ctx, graphclient.CreateServerMemoryInput{
		Serial:             serial,
		ServerID:           srv.ID,
		ServerMemoryTypeID: srvmt.ID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvMResp)
	require.NotNil(t, createdSrvMResp.ServerMemoryCreate.ServerMemory)

	createdM := createdSrvMResp.ServerMemoryCreate.ServerMemory
	require.NotNil(t, createdM.ID)
	assert.Equal(t, "srvrmem", createdM.ID.Prefix())
	assert.Equal(t, serial, createdM.Serial)
	assert.Equal(t, srvmt.ID, createdM.ServerMemoryType.ID)
	assert.Equal(t, srv.ID, createdM.Server.ID)

	// Update the Memory
	newSerial := gofakeit.UUID()
	updatedMemoryResp, err := graphTestClient().ServerMemoryUpdate(ctx, createdM.ID, graphclient.UpdateServerMemoryInput{Serial: &newSerial})

	require.NoError(t, err)
	require.NotNil(t, updatedMemoryResp)
	require.NotNil(t, updatedMemoryResp.ServerMemoryUpdate.ServerMemory)

	updatedM := updatedMemoryResp.ServerMemoryUpdate.ServerMemory
	require.EqualValues(t, createdM.ID, updatedM.ID)
	require.Equal(t, newSerial, updatedM.Serial)

	// Query the Memory
	queryM, err := graphTestClient().GetServerMemory(ctx, gidx.PrefixedID(createdM.ID))
	require.NoError(t, err)
	require.NotNil(t, queryM)
	require.NotNil(t, queryM.ServerMemory)
	require.Equal(t, newSerial, queryM.ServerMemory.Serial)

	// Delete the Server Chassis
	deletedResp, err := graphTestClient().ServerMemoryDelete(ctx, createdM.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerMemoryDelete)
	require.EqualValues(t, createdM.ID, deletedResp.ServerMemoryDelete.DeletedID.String())

	// Query the Memory to ensure it's no longer available
	deletedM, err := graphTestClient().GetServerMemory(ctx, createdM.ID)
	require.Error(t, err)
	require.Nil(t, deletedM)
	require.ErrorContains(t, err, "server_memory not found")
}
