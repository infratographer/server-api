package graphapi_test

import (
	"context"
	"testing"

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

func TestCreate_serverchassis(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvct := (&ServerChassisTypeBuilder{}).MustNew(ctx)
	srv := (&ServerBuilder{}).MustNew(ctx)

	serial := gofakeit.UUID()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerChassisInput
		ExpectedCh *ent.ServerChassis
		errorMsg   string
	}{
		{
			TestName: "creates server chassis ",
			Input:    graphclient.CreateServerChassisInput{Serial: serial, ServerID: srv.ID, ServerChassisTypeID: srvct.ID},
			ExpectedCh: &ent.ServerChassis{
				Serial:              serial,
				ServerID:            srv.ID,
				ServerChassisTypeID: srvct.ID,
			},
		},
		{
			TestName: "fails to create chassis with empty serial",
			Input:    graphclient.CreateServerChassisInput{Serial: "", ServerID: srv.ID, ServerChassisTypeID: srvct.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create chassis with empty chassis type",
			Input:    graphclient.CreateServerChassisInput{Serial: serial, ServerID: srv.ID, ServerChassisTypeID: ""},
			errorMsg: "constraint failed",
		},
		{
			TestName: "fails to create chassis with empty server",
			Input:    graphclient.CreateServerChassisInput{Serial: serial, ServerID: "", ServerChassisTypeID: srvct.ID},
			errorMsg: "constraint failed",
		},
		// TODO: test for parent chassis
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerChassisCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerChassisCreate)

			createdCh := resp.ServerChassisCreate.ServerChassis
			assert.Equal(t, tt.ExpectedCh.Serial, createdCh.Serial)
			assert.Equal(t, "srvrsch", createdCh.ID.Prefix())
			assert.Equal(t, tt.ExpectedCh.ServerID, createdCh.Server.ID)
			assert.Equal(t, tt.ExpectedCh.ServerChassisTypeID, createdCh.ServerChassisType.ID)
			// assert.Equal(t, tt.ExpectedCh.ParentChassisID, createdCh.ParentChassisID)
		})
	}
}

func TestUpdate_serverchassis(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ch := (&ServerChassisBuilder{}).MustNew(ctx)

	updateString := gofakeit.UUID()
	emptyString := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerChassisInput
		ExpectedCh *ent.ServerChassis
		errorMsg   string
	}{
		{
			TestName: "updates server chassis - serial",
			Input:    graphclient.UpdateServerChassisInput{Serial: &updateString},
			ExpectedCh: &ent.ServerChassis{
				Serial:              updateString,
				ServerID:            ch.ServerID,
				ServerChassisTypeID: ch.ServerChassisTypeID,
				ParentChassisID:     ch.ParentChassisID,
			},
		},
		{
			TestName: "fails to update to empty serial",
			Input:    graphclient.UpdateServerChassisInput{Serial: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			resp, err := graphTestClient().ServerChassisUpdate(ctx, ch.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerChassisUpdate)

			updatedCh := resp.ServerChassisUpdate.ServerChassis
			assert.Equal(t, tt.ExpectedCh.Serial, updatedCh.Serial)
			assert.Equal(t, ch.ID, updatedCh.ID)
		})
	}
}

func TestDelete_serverchassis(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ch := (&ServerChassisBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server chassis",
			Input:      ch.ID,
			ExpectedID: ch.ID,
		},
		{
			TestName: "fails to delete server chassis that does not exist",
			Input:    gidx.PrefixedID("srvrsch-1234"),
			errorMsg: "server_chassis not found",
		},
		{
			TestName: "fails to delete empty server chassis ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_chassis not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerChassisDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerChassisDelete)

			deletedCT := resp.ServerChassisDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerChassisLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvct := (&ServerChassisTypeBuilder{}).MustNew(ctx)
	srv := (&ServerBuilder{}).MustNew(ctx)
	serial := gofakeit.UUID()

	exp := &ent.ServerChassis{
		Serial:              serial,
		ServerChassisTypeID: srvct.ID,
		ServerID:            srv.ID,
	}

	// create the Chassis
	createdSrvChassisResp, err := graphTestClient().ServerChassisCreate(ctx, graphclient.CreateServerChassisInput{
		Serial:              exp.Serial,
		ServerID:            exp.ServerID,
		ServerChassisTypeID: exp.ServerChassisTypeID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvChassisResp)
	require.NotNil(t, createdSrvChassisResp.ServerChassisCreate.ServerChassis)

	createdChassis := createdSrvChassisResp.ServerChassisCreate.ServerChassis
	require.NotNil(t, createdChassis.ID)
	assert.Equal(t, "srvrsch", createdChassis.ID.Prefix())
	assert.Equal(t, exp.Serial, createdChassis.Serial)
	assert.Equal(t, exp.ServerID, createdChassis.Server.ID)
	assert.Equal(t, exp.ServerChassisTypeID, createdChassis.ServerChassisType.ID)

	// Update the Chassis
	newSerial := gofakeit.DomainName()
	updatedChassisResp, err := graphTestClient().ServerChassisUpdate(ctx, createdChassis.ID, graphclient.UpdateServerChassisInput{Serial: &newSerial})

	require.NoError(t, err)
	require.NotNil(t, updatedChassisResp)
	require.NotNil(t, updatedChassisResp.ServerChassisUpdate.ServerChassis)

	updatedChassis := updatedChassisResp.ServerChassisUpdate.ServerChassis
	require.EqualValues(t, createdChassis.ID, updatedChassis.ID)
	require.Equal(t, newSerial, updatedChassis.Serial)

	// Query the Chassis
	queryChassis, err := graphTestClient().GetServerChassis(ctx, createdChassis.ID)
	require.NoError(t, err)
	require.NotNil(t, queryChassis)
	require.NotNil(t, queryChassis.ServerChassis)
	require.Equal(t, newSerial, queryChassis.ServerChassis.Serial)

	// Delete the Server Chassis
	deletedResp, err := graphTestClient().ServerChassisDelete(ctx, createdChassis.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerChassisDelete)
	require.EqualValues(t, createdChassis.ID, deletedResp.ServerChassisDelete.DeletedID.String())

	// Query the CPU to ensure it's no longer available
	deletedChassis, err := graphTestClient().GetServerChassis(ctx, createdChassis.ID)
	require.Error(t, err)
	require.Nil(t, deletedChassis)
	require.ErrorContains(t, err, "server_chassis not found")
}
