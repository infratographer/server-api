package graphapi_test

import (
	"context"
	"testing"

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

func TestCreate_servercomponenttype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	name := gofakeit.DomainName()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerComponentTypeInput
		ExpectedCT *ent.ServerComponentType
		errorMsg   string
	}{
		{
			TestName: "creates server component type ",
			Input:    graphclient.CreateServerComponentTypeInput{Name: name},
			ExpectedCT: &ent.ServerComponentType{
				Name: name,
			},
		},
		{
			TestName: "fails to create server component type with empty name",
			Input:    graphclient.CreateServerComponentTypeInput{Name: ""},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerComponentTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerComponentTypeCreate)

			createdCT := resp.ServerComponentTypeCreate.ServerComponentType
			assert.Equal(t, tt.ExpectedCT.Name, createdCT.Name)
			assert.Equal(t, "srvrcpt", createdCT.ID.Prefix())
		})
	}
}

func TestUpdate_servercomponenttype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ch := (&ServerComponentTypeBuilder{}).MustNew(ctx)

	updateString := gofakeit.UUID()
	emptyString := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerComponentTypeInput
		ExpectedCT *ent.ServerComponentType
		errorMsg   string
	}{
		{
			TestName: "updates server component type - name",
			Input:    graphclient.UpdateServerComponentTypeInput{Name: &updateString},
			ExpectedCT: &ent.ServerComponentType{
				Name: updateString,
			},
		},
		{
			TestName: "fails to update to empty name",
			Input:    graphclient.UpdateServerComponentTypeInput{Name: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			resp, err := graphTestClient().ServerComponentTypeUpdate(ctx, ch.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerComponentTypeUpdate)

			updatedCT := resp.ServerComponentTypeUpdate.ServerComponentType
			assert.Equal(t, tt.ExpectedCT.Name, updatedCT.Name)
			assert.Equal(t, ch.ID, updatedCT.ID)
		})
	}
}

func TestDelete_servercomponenttype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ct := (&ServerComponentTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server component type",
			Input:      ct.ID,
			ExpectedID: ct.ID,
		},
		{
			TestName: "fails to delete server chassis that does not exist",
			Input:    gidx.PrefixedID("srvrcpt-1234"),
			errorMsg: "server_component_type not found",
		},
		{
			TestName: "fails to delete empty server chassis ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_component_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerComponentTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerComponentTypeDelete)

			deletedCT := resp.ServerComponentTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerComponentTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	// srvct := (&ServerComponentTypeBuilder{}).MustNew(ctx)
	name := gofakeit.UUID()

	exp := &ent.ServerComponentType{
		Name: name,
	}

	// create the Chassis
	createdSrvComponentTypeResp, err := graphTestClient().ServerComponentTypeCreate(ctx, graphclient.CreateServerComponentTypeInput{
		Name: exp.Name,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvComponentTypeResp)
	require.NotNil(t, createdSrvComponentTypeResp.ServerComponentTypeCreate.ServerComponentType)

	createdCT := createdSrvComponentTypeResp.ServerComponentTypeCreate.ServerComponentType
	require.NotNil(t, createdCT.ID)
	assert.Equal(t, "srvrcpt", createdCT.ID.Prefix())
	assert.Equal(t, exp.Name, createdCT.Name)

	// Update the Chassis
	newName := gofakeit.DomainName()
	updatedComponentTypeResp, err := graphTestClient().ServerComponentTypeUpdate(ctx, createdCT.ID, graphclient.UpdateServerComponentTypeInput{Name: &newName})

	require.NoError(t, err)
	require.NotNil(t, updatedComponentTypeResp)
	require.NotNil(t, updatedComponentTypeResp.ServerComponentTypeUpdate.ServerComponentType)

	updatedCT := updatedComponentTypeResp.ServerComponentTypeUpdate.ServerComponentType
	require.EqualValues(t, createdCT.ID, updatedCT.ID)
	require.Equal(t, newName, updatedCT.Name)

	// Query the Chassis
	queryCT, err := graphTestClient().GetServerComponentType(ctx, gidx.PrefixedID(createdCT.ID))
	require.NoError(t, err)
	require.NotNil(t, queryCT)
	require.NotNil(t, queryCT.ServerComponentType)
	require.Equal(t, newName, queryCT.ServerComponentType.Name)

	// Delete the Server Chassis
	deletedResp, err := graphTestClient().ServerComponentTypeDelete(ctx, createdCT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerComponentTypeDelete)
	require.EqualValues(t, createdCT.ID, deletedResp.ServerComponentTypeDelete.DeletedID.String())

	// Query the CPU to ensure it's no longer available
	deletedCT, err := graphTestClient().GetServerComponentType(ctx, createdCT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_component_type not found")
}
