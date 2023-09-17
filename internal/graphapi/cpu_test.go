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

func TestCreate_servercpu(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvct := (&ServerCPUTypeBuilder{}).MustNew(ctx)
	srv := (&ServerBuilder{}).MustNew(ctx)

	serial := gofakeit.UUID()

	testCases := []struct {
		TestName    string
		Input       graphclient.CreateServerCPUInput
		ExpectedCPU *ent.ServerCPU
		errorMsg    string
	}{
		{
			TestName: "creates server cpu ",
			Input:    graphclient.CreateServerCPUInput{Serial: serial, ServerID: srv.ID, ServerCPUTypeID: srvct.ID},
			ExpectedCPU: &ent.ServerCPU{
				Serial:          serial,
				ServerID:        srv.ID,
				ServerCPUTypeID: srvct.ID,
			},
		},
		{
			TestName: "fails to create server cpu with empty serial",
			Input:    graphclient.CreateServerCPUInput{Serial: "", ServerID: srv.ID, ServerCPUTypeID: srvct.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server cpu with empty cpu type",
			Input:    graphclient.CreateServerCPUInput{Serial: serial, ServerID: srv.ID, ServerCPUTypeID: ""},
			errorMsg: "constraint failed",
		},
		{
			TestName: "fails to create server cpu type with empty clock speed",
			Input:    graphclient.CreateServerCPUInput{Serial: serial, ServerID: "", ServerCPUTypeID: srvct.ID},
			errorMsg: "constraint failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerCPUCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCPUCreate)

			createdCPU := resp.ServerCPUCreate.ServerCPU
			assert.Equal(t, tt.ExpectedCPU.Serial, createdCPU.Serial)
			assert.Equal(t, "srvrcpu", createdCPU.ID.Prefix())
			assert.Equal(t, tt.ExpectedCPU.ServerID, createdCPU.Server.ID)
			assert.Equal(t, tt.ExpectedCPU.ServerCPUTypeID, createdCPU.ServerCPUType.ID)
		})
	}
}

func TestUpdate_servercpu(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	cpu := (&ServerCPUBuilder{}).MustNew(ctx)

	updateString := gofakeit.UUID()
	emptyString := ""

	testCases := []struct {
		TestName    string
		Input       graphclient.UpdateServerCPUInput
		ExpectedCPU *ent.ServerCPU
		errorMsg    string
	}{
		{
			TestName: "updates server cpu - model",
			Input:    graphclient.UpdateServerCPUInput{Serial: &updateString},
			ExpectedCPU: &ent.ServerCPU{
				Serial: updateString,
			},
		},
		{
			TestName: "fails to update to empty serial",
			Input:    graphclient.UpdateServerCPUInput{Serial: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			resp, err := graphTestClient().ServerCPUUpdate(ctx, cpu.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCPUUpdate)

			updatedCPU := resp.ServerCPUUpdate.ServerCPU
			assert.Equal(t, tt.ExpectedCPU.Serial, updatedCPU.Serial)
			assert.Equal(t, cpu.ID, updatedCPU.ID)
		})
	}
}

func TestDelete_servercpu(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	cpu := (&ServerCPUBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server cpu",
			Input:      cpu.ID,
			ExpectedID: cpu.ID,
		},
		{
			TestName: "fails to delete server cpu type that does not exist",
			Input:    gidx.PrefixedID("srvrcpt-1234"),
			errorMsg: "server_cpu not found",
		},
		{
			TestName: "fails to delete empty server cpu type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_cpu not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerCPUDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCPUDelete)

			deletedCT := resp.ServerCPUDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerCPULifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srv := (&ServerBuilder{}).MustNew(ctx)
	srvct := (&ServerCPUTypeBuilder{}).MustNew(ctx)
	serial := gofakeit.UUID()

	exp := &ent.ServerCPU{
		Serial:          serial,
		ServerCPUTypeID: srvct.ID,
		ServerID:        srv.ID,
	}

	// create the CPU
	createdSrvCPUResp, err := graphTestClient().ServerCPUCreate(ctx, graphclient.CreateServerCPUInput{
		Serial:          exp.Serial,
		ServerID:        exp.ServerID,
		ServerCPUTypeID: exp.ServerCPUTypeID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvCPUResp)
	require.NotNil(t, createdSrvCPUResp.ServerCPUCreate.ServerCPU)

	createdCPU := createdSrvCPUResp.ServerCPUCreate.ServerCPU
	require.NotNil(t, createdCPU.ID)
	assert.Equal(t, "srvrcpu", createdCPU.ID.Prefix())
	assert.Equal(t, exp.Serial, createdCPU.Serial)
	assert.Equal(t, exp.ServerID, createdCPU.Server.ID)
	assert.Equal(t, exp.ServerCPUTypeID, createdCPU.ServerCPUType.ID)

	// Update the CPU
	newSerial := gofakeit.DomainName()
	updatedCPUResp, err := graphTestClient().ServerCPUUpdate(ctx, createdCPU.ID, graphclient.UpdateServerCPUInput{Serial: &newSerial})

	require.NoError(t, err)
	require.NotNil(t, updatedCPUResp)
	require.NotNil(t, updatedCPUResp.ServerCPUUpdate.ServerCPU)

	updatedCPU := updatedCPUResp.ServerCPUUpdate.ServerCPU
	require.EqualValues(t, createdCPU.ID, updatedCPU.ID)
	require.Equal(t, newSerial, updatedCPU.Serial)

	// Query the CPU
	queryCPU, err := graphTestClient().GetServerCPU(ctx, createdCPU.ID)
	require.NoError(t, err)
	require.NotNil(t, queryCPU)
	require.NotNil(t, queryCPU.ServerCPU)
	require.Equal(t, newSerial, queryCPU.ServerCPU.Serial)

	// Delete the Server CPU
	deletedResp, err := graphTestClient().ServerCPUDelete(ctx, createdCPU.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerCPUDelete)
	require.EqualValues(t, createdCPU.ID, deletedResp.ServerCPUDelete.DeletedID.String())

	// Query the CPU to ensure it's no longer available
	deletedCPU, err := graphTestClient().GetServerCPU(ctx, createdCPU.ID)
	require.Error(t, err)
	require.Nil(t, deletedCPU)
	require.ErrorContains(t, err, "server_cpu not found")
}
