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

func TestCreate_servertype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	// srvtype := (&ServerTypeBuilder{}).MustNew(ctx)
	ownerID := gidx.MustNewID(ownerPrefix)
	name := gofakeit.DomainName()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerTypeInput
		ExpectedLB *ent.ServerType
		errorMsg   string
	}{
		{
			TestName: "creates server type",
			Input:    graphclient.CreateServerTypeInput{Name: name, OwnerID: ownerID},
			ExpectedLB: &ent.ServerType{
				Name:    name,
				OwnerID: ownerID,
			},
		},
		{
			TestName: "fails to create server with empty name",
			Input:    graphclient.CreateServerTypeInput{Name: "", OwnerID: ownerID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server with empty ownerID",
			Input:    graphclient.CreateServerTypeInput{Name: name, OwnerID: ""},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerTypeCreate)

			createdType := resp.ServerTypeCreate.ServerType
			assert.Equal(t, tt.ExpectedLB.Name, createdType.Name)
			assert.Equal(t, "srvrtyp", createdType.ID.Prefix())
			assert.Equal(t, ownerID, createdType.Owner.ID)
		})
	}
}

func TestUpdate_servertype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	st := (&ServerTypeBuilder{}).MustNew(ctx)
	updateName := gofakeit.DomainName()
	emptyName := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerTypeInput
		ExpectedLB *ent.ServerType
		errorMsg   string
	}{
		{
			TestName: "updates server type",
			Input:    graphclient.UpdateServerTypeInput{Name: &updateName},
			ExpectedLB: &ent.ServerType{
				Name:    updateName,
				OwnerID: st.OwnerID,
			},
		},
		{
			TestName: "fails to update name to empty",
			Input:    graphclient.UpdateServerTypeInput{Name: &emptyName},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerTypeUpdate(ctx, st.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerTypeUpdate)

			updatedST := resp.ServerTypeUpdate.ServerType
			assert.Equal(t, tt.ExpectedLB.Name, updatedST.Name)
			assert.Equal(t, st.ID, updatedST.ID)
		})
	}
}

func TestDelete_servertype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	st1 := (&ServerTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes servertype",
			Input:      st1.ID,
			ExpectedID: st1.ID,
		},
		{
			TestName: "fails to delete server type that does not exist",
			Input:    gidx.PrefixedID("srvrtyp-1234"),
			errorMsg: "server_type not found",
		},
		{
			TestName: "fails to delete empty server type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerTypeDelete)

			deletedLB := resp.ServerTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedLB.DeletedID)
		})
	}
}

func TestFullServerTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	// srvType := (&ServerTypeBuilder{}).MustNew(ctx)
	ownerID := gidx.MustNewID(ownerPrefix)
	name := gofakeit.DomainName()

	// create the LB
	createdSrvTypeResp, err := graphTestClient().ServerTypeCreate(ctx, graphclient.CreateServerTypeInput{
		Name:    name,
		OwnerID: ownerID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvTypeResp)
	require.NotNil(t, createdSrvTypeResp.ServerTypeCreate.ServerType)

	createdST := createdSrvTypeResp.ServerTypeCreate.ServerType
	require.NotNil(t, createdST.ID)
	require.Equal(t, name, createdST.Name)
	assert.Equal(t, "srvrtyp", createdST.ID.Prefix())
	assert.Equal(t, ownerID, createdST.Owner.ID)

	// createdPortResp, err := graphTestClient().LoadBalancerPortCreate(ctx, graphclient.CreateLoadBalancerPortInput{
	// 	Name:           gofakeit.DomainName(),
	// 	Number:         8080,
	// 	LoadBalancerID: createdLB.ID,
	// })

	// require.NoError(t, err)
	// require.NotNil(t, createdPortResp)
	// require.NotNil(t, createdPortResp.LoadBalancerPortCreate.LoadBalancerPort)

	// Update the Server
	newName := gofakeit.DomainName()
	updatedSTResp, err := graphTestClient().ServerTypeUpdate(ctx, createdST.ID, graphclient.UpdateServerTypeInput{Name: &newName})

	require.NoError(t, err)
	require.NotNil(t, updatedSTResp)
	require.NotNil(t, updatedSTResp.ServerTypeUpdate.ServerType)

	updatedST := updatedSTResp.ServerTypeUpdate.ServerType
	require.EqualValues(t, createdST.ID, updatedST.ID)
	require.Equal(t, newName, updatedST.Name)

	// Query the LB
	queryST, err := graphTestClient().GetServerType(ctx, createdST.ID)
	require.NoError(t, err)
	require.NotNil(t, queryST)
	require.NotNil(t, queryST.ServerType)
	require.Equal(t, newName, queryST.ServerType.Name)

	// Delete the Server
	deletedResp, err := graphTestClient().ServerTypeDelete(ctx, createdST.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerTypeDelete)
	require.EqualValues(t, createdST.ID, deletedResp.ServerTypeDelete.DeletedID.String())

	// Query the LB to ensure it's no longer available
	deletedLB, err := graphTestClient().GetServerType(ctx, createdST.ID)
	require.Error(t, err)
	require.Nil(t, deletedLB)
	require.ErrorContains(t, err, "server_type not found")
}
