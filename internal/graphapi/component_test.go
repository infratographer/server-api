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

func TestCreate_servercomponent(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvprv := (&ProviderBuilder{}).MustNew(ctx)
	srvtyp := (&ServerTypeBuilder{}).MustNew(ctx)
	srv := (&ServerBuilder{ProviderID: srvprv.ID, ServerTypeID: srvtyp.ID}).MustNew(ctx)

	srvcpt := (&ServerComponentTypeBuilder{}).MustNew(ctx)

	// exp := &ent.ServerComponent{
	name := gofakeit.DomainName()
	vendor := gofakeit.Company()
	model := gofakeit.CarModel()
	serial := gofakeit.UUID()
	serverID := srv.ID
	componentTypeID := srvcpt.ID
	// }

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerComponentInput
		ExpectedCT *ent.ServerComponent
		errorMsg   string
	}{
		{
			TestName: "creates server component",
			Input:    graphclient.CreateServerComponentInput{ServerID: serverID, ComponentTypeID: componentTypeID, Name: name, Vendor: vendor, Model: model, Serial: serial},
			ExpectedCT: &ent.ServerComponent{
				Name:            name,
				ServerID:        serverID,
				ComponentTypeID: componentTypeID,
				Vendor:          vendor,
				Model:           model,
				Serial:          serial,
			},
		},
		{
			TestName: "fails to create server component with empty name",
			Input:    graphclient.CreateServerComponentInput{ServerID: serverID, ComponentTypeID: componentTypeID, Name: "", Vendor: vendor, Model: model, Serial: serial},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty vendor",
			Input:    graphclient.CreateServerComponentInput{ServerID: serverID, ComponentTypeID: componentTypeID, Name: name, Vendor: "", Model: model, Serial: serial},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty model",
			Input:    graphclient.CreateServerComponentInput{ServerID: serverID, ComponentTypeID: componentTypeID, Name: name, Vendor: vendor, Model: "", Serial: serial},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty serial",
			Input:    graphclient.CreateServerComponentInput{ServerID: serverID, ComponentTypeID: componentTypeID, Name: name, Vendor: vendor, Model: model, Serial: ""},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty server",
			Input:    graphclient.CreateServerComponentInput{ServerID: "", ComponentTypeID: componentTypeID, Name: name, Vendor: "", Model: model, Serial: serial},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with empty component type",
			Input:    graphclient.CreateServerComponentInput{ServerID: serverID, ComponentTypeID: "", Name: name, Vendor: "", Model: model, Serial: serial},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server component type with invalid server",
			Input:    graphclient.CreateServerComponentInput{ServerID: "srvrsrv-123412341234123412341", ComponentTypeID: componentTypeID, Name: name, Vendor: vendor, Model: model, Serial: serial},
			errorMsg: "constraint failed",
		},
		{
			TestName: "fails to create server component type with invalid component type",
			Input:    graphclient.CreateServerComponentInput{ServerID: gidx.MustNewID("srvrcpt"), ComponentTypeID: componentTypeID, Name: name, Vendor: vendor, Model: model, Serial: serial},
			errorMsg: "constraint failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerComponentCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerComponentCreate)

			createdCT := resp.ServerComponentCreate.ServerComponent
			assert.Equal(t, tt.ExpectedCT.Name, createdCT.Name)
			assert.Equal(t, tt.ExpectedCT.Model, createdCT.Model)
			assert.Equal(t, tt.ExpectedCT.Vendor, createdCT.Vendor)
			assert.Equal(t, tt.ExpectedCT.Serial, createdCT.Serial)
			assert.Equal(t, tt.ExpectedCT.ServerID, createdCT.Server.ID)
			assert.Equal(t, tt.ExpectedCT.ComponentTypeID, createdCT.ServerComponentType.ID)
			assert.Equal(t, "srvrcmp", createdCT.ID.Prefix())
		})
	}
}

func TestUpdate_servercomponent(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	// ct := (&ServerComponentBuilder{}).MustNew(ctx)

	updateString := gofakeit.UUID()
	emptyString := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerComponentInput
		ExpectedCT *ent.ServerComponent
		errorMsg   string
	}{
		{
			TestName: "updates server component - name",
			Input:    graphclient.UpdateServerComponentInput{Name: &updateString},
			ExpectedCT: &ent.ServerComponent{
				Name: updateString,
			},
		},
		{
			TestName: "failed updating name to empty string",
			Input:    graphclient.UpdateServerComponentInput{Name: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server component - model",
			Input:    graphclient.UpdateServerComponentInput{Model: &updateString},
			ExpectedCT: &ent.ServerComponent{
				Model: updateString,
			},
		},
		{
			TestName: "failed updating model to empty string",
			Input:    graphclient.UpdateServerComponentInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server component - vendor",
			Input:    graphclient.UpdateServerComponentInput{Vendor: &updateString},
			ExpectedCT: &ent.ServerComponent{
				Vendor: updateString,
			},
		},
		{
			TestName: "failed updating vendor to empty string",
			Input:    graphclient.UpdateServerComponentInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server component - serial",
			Input:    graphclient.UpdateServerComponentInput{Serial: &updateString},
			ExpectedCT: &ent.ServerComponent{
				Serial: updateString,
			},
		},
		{
			TestName: "failed updating serial to empty string",
			Input:    graphclient.UpdateServerComponentInput{Serial: &emptyString},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			srv := (&ServerBuilder{}).MustNew(ctx)
			srvct := (&ServerComponentTypeBuilder{}).MustNew(ctx)

			ct := (&ServerComponentBuilder{
				Vendor:              gofakeit.CarMaker(),
				Model:               gofakeit.CarModel(),
				Serial:              gofakeit.UUID(),
				Name:                gofakeit.DomainName(),
				Server:              srv,
				ServerComponentType: srvct,
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedCT, ct)

			resp, err := graphTestClient().ServerComponentUpdate(ctx, ct.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerComponentUpdate)

			updatedCT := resp.ServerComponentUpdate.ServerComponent
			assert.Equal(t, tt.ExpectedCT.Name, updatedCT.Name)
			assert.Equal(t, tt.ExpectedCT.Model, updatedCT.Model)
			assert.Equal(t, tt.ExpectedCT.Vendor, updatedCT.Vendor)
			assert.Equal(t, tt.ExpectedCT.Serial, updatedCT.Serial)
			assert.Equal(t, tt.ExpectedCT.ServerID, updatedCT.Server.ID)
			assert.Equal(t, tt.ExpectedCT.ComponentTypeID, updatedCT.ServerComponentType.ID)
			assert.Equal(t, ct.ID, updatedCT.ID)
		})
	}
}

func TestDelete_servercomponent(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ct := (&ServerComponentBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server component",
			Input:      ct.ID,
			ExpectedID: ct.ID,
		},
		{
			TestName: "fails to delete server component that does not exist",
			Input:    gidx.PrefixedID("srvrcmp-1234"),
			errorMsg: "server_component not found",
		},
		{
			TestName: "fails to delete empty server component ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_component not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerComponentDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerComponentDelete)

			deletedCT := resp.ServerComponentDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerComponentLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srv := (&ServerBuilder{}).MustNew(ctx)
	srvct := (&ServerComponentTypeBuilder{}).MustNew(ctx)

	exp := &ent.ServerComponent{
		Name:            gofakeit.Name(),
		Vendor:          gofakeit.Company(),
		Model:           gofakeit.CarModel(),
		Serial:          gofakeit.UUID(),
		ServerID:        srv.ID,
		ComponentTypeID: srvct.ID,
	}

	// create the Component
	createdSrvComponentResp, err := graphTestClient().ServerComponentCreate(ctx, graphclient.CreateServerComponentInput{
		Name:            exp.Name,
		Vendor:          exp.Vendor,
		Model:           exp.Model,
		Serial:          exp.Serial,
		ServerID:        exp.ServerID,
		ComponentTypeID: exp.ComponentTypeID,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvComponentResp)
	require.NotNil(t, createdSrvComponentResp.ServerComponentCreate.ServerComponent)

	createdCT := createdSrvComponentResp.ServerComponentCreate.ServerComponent
	require.NotNil(t, createdCT.ID)
	assert.Equal(t, "srvrcmp", createdCT.ID.Prefix())
	assert.Equal(t, exp.Name, createdCT.Name)

	// Update the Component
	newName := gofakeit.Name()
	updatedComponentResp, err := graphTestClient().ServerComponentUpdate(ctx, createdCT.ID, graphclient.UpdateServerComponentInput{Name: &newName})

	require.NoError(t, err)
	require.NotNil(t, updatedComponentResp)
	require.NotNil(t, updatedComponentResp.ServerComponentUpdate.ServerComponent)

	updatedCT := updatedComponentResp.ServerComponentUpdate.ServerComponent
	require.EqualValues(t, createdCT.ID, updatedCT.ID)
	require.Equal(t, newName, updatedCT.Name)

	// Query the Component
	queryCT, err := graphTestClient().GetServerComponent(ctx, gidx.PrefixedID(createdCT.ID))
	require.NoError(t, err)
	require.NotNil(t, queryCT)
	require.NotNil(t, queryCT.ServerComponent)
	require.Equal(t, newName, queryCT.ServerComponent.Name)

	// Delete the Server Component
	deletedResp, err := graphTestClient().ServerComponentDelete(ctx, createdCT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerComponentDelete)
	require.EqualValues(t, createdCT.ID, deletedResp.ServerComponentDelete.DeletedID.String())

	// Query the Component to ensure it's no longer available
	deletedCT, err := graphTestClient().GetServerComponent(ctx, createdCT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_component not found")
}
