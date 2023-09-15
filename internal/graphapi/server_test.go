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
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/server-api/internal/graphclient"
)

func TestQuery_server(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srv1 := (&ServerBuilder{}).MustNew(ctx)
	srv2 := (&ServerBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		QueryID    gidx.PrefixedID
		ExpectedLB *ent.Server
		errorMsg   string
	}{
		{
			TestName:   "Happy Path - srv1",
			QueryID:    srv1.ID,
			ExpectedLB: srv1,
		},
		{
			TestName:   "Happy Path - srv2",
			QueryID:    srv2.ID,
			ExpectedLB: srv2,
		},
		{
			TestName: "No server found with ID",
			QueryID:  gidx.MustNewID("testing"),
			errorMsg: "server not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().GetServer(ctx, tt.QueryID)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Server)
			assert.EqualValues(t, tt.ExpectedLB.Name, resp.Server.Name)
		})
	}
}

func TestCreate_loadBalancer(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	prov := (&ProviderBuilder{}).MustNew(ctx)
	srvtype := (&ServerTypeBuilder{}).MustNew(ctx)
	ownerID := gidx.MustNewID(ownerPrefix)
	locationID := gidx.MustNewID(locationPrefix)
	name := gofakeit.DomainName()

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerInput
		ExpectedLB *ent.Server
		errorMsg   string
	}{
		{
			TestName: "creates server",
			Input:    graphclient.CreateServerInput{Name: name, ProviderID: prov.ID, OwnerID: ownerID, LocationID: locationID, ServerTypeID: srvtype.ID},
			ExpectedLB: &ent.Server{
				Name:         name,
				ProviderID:   prov.ID,
				OwnerID:      ownerID,
				LocationID:   locationID,
				ServerTypeID: srvtype.ID,
			},
		},
		{
			TestName: "fails to create server with empty name",
			Input:    graphclient.CreateServerInput{Name: "", ProviderID: prov.ID, OwnerID: ownerID, LocationID: locationID, ServerTypeID: srvtype.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server with empty ownerID",
			Input:    graphclient.CreateServerInput{Name: name, ProviderID: prov.ID, OwnerID: "", LocationID: locationID, ServerTypeID: srvtype.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create loadbalancer with empty locationID",
			Input:    graphclient.CreateServerInput{Name: name, ProviderID: prov.ID, OwnerID: ownerID, LocationID: "", ServerTypeID: srvtype.ID},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create loadbalancer with empty serverTypeID",
			Input:    graphclient.CreateServerInput{Name: name, ProviderID: prov.ID, OwnerID: ownerID, LocationID: locationID, ServerTypeID: ""},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCreate)

			createdLB := resp.ServerCreate.Server
			assert.Equal(t, tt.ExpectedLB.Name, createdLB.Name)
			assert.Equal(t, "srvrsrv", createdLB.ID.Prefix())
			assert.Equal(t, prov.ID, createdLB.ServerProvider.ID)
			assert.Equal(t, locationID, createdLB.Location.ID)
			assert.Equal(t, ownerID, createdLB.Owner.ID)
		})
	}
}

func TestUpdate_server(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	lb := (&ServerBuilder{}).MustNew(ctx)
	updateName := gofakeit.DomainName()
	emptyName := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerInput
		ExpectedLB *ent.Server
		errorMsg   string
	}{
		{
			TestName: "updates server",
			Input:    graphclient.UpdateServerInput{Name: &updateName},
			ExpectedLB: &ent.Server{
				Name:       updateName,
				ProviderID: lb.ProviderID,
				OwnerID:    lb.OwnerID,
				LocationID: lb.LocationID,
			},
		},
		{
			TestName: "fails to update name to empty",
			Input:    graphclient.UpdateServerInput{Name: &emptyName},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerUpdate(ctx, lb.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerUpdate)

			updatedLB := resp.ServerUpdate.Server
			assert.Equal(t, tt.ExpectedLB.Name, updatedLB.Name)
			assert.Equal(t, lb.ID, updatedLB.ID)
		})
	}
}

func TestDelete_server(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srv1 := (&ServerBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server",
			Input:      srv1.ID,
			ExpectedID: srv1.ID,
		},
		{
			TestName: "fails to delete server that does not exist",
			Input:    gidx.PrefixedID("srvrsrv-1234"),
			errorMsg: "server not found",
		},
		{
			TestName: "fails to delete empty server ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerDelete)

			deletedLB := resp.ServerDelete
			assert.EqualValues(t, tt.ExpectedID, deletedLB.DeletedID)
		})
	}
}

func TestFullServerLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvType := (&ServerTypeBuilder{}).MustNew(ctx)
	prov := (&ProviderBuilder{}).MustNew(ctx)
	ownerID := gidx.MustNewID(ownerPrefix)
	locationID := gidx.MustNewID(locationPrefix)
	name := gofakeit.DomainName()

	// create the LB
	createdSrvResp, err := graphTestClient().ServerCreate(ctx, graphclient.CreateServerInput{
		Name:         name,
		ProviderID:   prov.ID,
		OwnerID:      ownerID,
		LocationID:   locationID,
		ServerTypeID: srvType.ID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvResp)
	require.NotNil(t, createdSrvResp.ServerCreate.Server)

	createdLB := createdSrvResp.ServerCreate.Server
	require.NotNil(t, createdLB.ID)
	require.Equal(t, name, createdLB.Name)
	assert.Equal(t, "srvrsrv", createdLB.ID.Prefix())
	assert.Equal(t, prov.ID, createdLB.ServerProvider.ID)
	assert.Equal(t, locationID, createdLB.Location.ID)
	assert.Equal(t, ownerID, createdLB.Owner.ID)

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
	updatedSrvResp, err := graphTestClient().ServerUpdate(ctx, createdLB.ID, graphclient.UpdateServerInput{Name: &newName})

	require.NoError(t, err)
	require.NotNil(t, updatedSrvResp)
	require.NotNil(t, updatedSrvResp.ServerUpdate.Server)

	updatedSrv := updatedSrvResp.ServerUpdate.Server
	require.EqualValues(t, createdLB.ID, updatedSrv.ID)
	require.Equal(t, newName, updatedSrv.Name)

	// Query the LB
	querySrv, err := graphTestClient().GetServer(ctx, createdLB.ID)
	require.NoError(t, err)
	require.NotNil(t, querySrv)
	require.NotNil(t, querySrv.Server)
	require.Equal(t, newName, querySrv.Server.Name)

	// Delete the Server
	deletedResp, err := graphTestClient().ServerDelete(ctx, createdLB.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerDelete)
	require.EqualValues(t, createdLB.ID, deletedResp.ServerDelete.DeletedID.String())

	// Query the LB to ensure it's no longer available
	deletedLB, err := graphTestClient().GetServer(ctx, createdLB.ID)
	require.Error(t, err)
	require.Nil(t, deletedLB)
	require.ErrorContains(t, err, "server not found")
}
