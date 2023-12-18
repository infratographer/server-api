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

func TestCreate_serverchassistype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ct := (&ServerChassisTypeBuilder{IsFullDepth: true}).MustNew(ctx)
	depth := false

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerChassisTypeInput
		ExpectedCT *ent.ServerChassisType
		errorMsg   string
	}{
		{
			TestName: "creates server chassis type",
			Input:    graphclient.CreateServerChassisTypeInput{Model: ct.Model, Vendor: ct.Vendor, Height: ct.Height},
			ExpectedCT: &ent.ServerChassisType{
				Model:               ct.Model,
				Vendor:              ct.Vendor,
				Height:              ct.Height,
				IsFullDepth:         ct.IsFullDepth,
				ParentChassisTypeID: ct.ParentChassisTypeID,
			},
		},
		{
			TestName: "creates server chassis type with parent chassis type",
			Input:    graphclient.CreateServerChassisTypeInput{Model: ct.Model, Vendor: ct.Vendor, Height: ct.Height, ParentChassisTypeID: gidx.MustNewID("srvrsct")},
			ExpectedCT: &ent.ServerChassisType{
				Model:               ct.Model,
				Vendor:              ct.Vendor,
				Height:              ct.Height,
				IsFullDepth:         ct.IsFullDepth,
				ParentChassisTypeID: ct.ParentChassisTypeID,
			},
		},
		{
			TestName: "creates server chassis type with non full depth",
			Input:    graphclient.CreateServerChassisTypeInput{Model: ct.Model, Vendor: ct.Vendor, Height: ct.Height, ParentChassisTypeID: gidx.MustNewID("srvrsct"), IsFullDepth: &depth},
			ExpectedCT: &ent.ServerChassisType{
				Model:               ct.Model,
				Vendor:              ct.Vendor,
				Height:              ct.Height,
				IsFullDepth:         depth,
				ParentChassisTypeID: ct.ParentChassisTypeID,
			},
		},
		{
			TestName: "fails to create server chassis type with empty vendor",
			Input:    graphclient.CreateServerChassisTypeInput{Vendor: "", Model: ct.Model, Height: ct.Height, IsFullDepth: &ct.IsFullDepth},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server chassis type with empty model",
			Input:    graphclient.CreateServerChassisTypeInput{Vendor: ct.Vendor, Model: "", Height: ct.Height, IsFullDepth: &ct.IsFullDepth},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server chassis type with empty height",
			Input:    graphclient.CreateServerChassisTypeInput{Vendor: ct.Vendor, Model: ct.Model, Height: "", IsFullDepth: &ct.IsFullDepth},
			errorMsg: "value is less than the required length",
		},
		// {
		// 	TestName: "creates server chassis type with invalid parent chassis type",
		// 	Input:    graphclient.CreateServerChassisTypeInput{Model: ct.Model, Vendor: ct.Vendor, Height: ct.Height, ParentChassisTypeID: gidx.MustNewID("srvrsct-12341234"), IsFullDepth: &depth},
		// 	errorMsg: "expected prefix length",
		// },
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerChassisTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerChassisTypeCreate)

			createdCT := resp.ServerChassisTypeCreate.ServerChassisType
			assert.Equal(t, tt.ExpectedCT.Vendor, createdCT.Vendor)
			assert.Equal(t, tt.ExpectedCT.Model, createdCT.Model)
			assert.Equal(t, tt.ExpectedCT.Height, createdCT.Height)
			assert.Equal(t, tt.ExpectedCT.IsFullDepth, createdCT.IsFullDepth)
			// TODO: why can't we compare the parent chassis type ID?
			// assert.Equal(t, tt.ExpectedCT.ParentChassisTypeID, createdCT.)
			assert.Equal(t, "srvrsct", createdCT.ID.Prefix())
		})
	}
}

func TestUpdate_serverchassistype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	updateString := gofakeit.CarMaker()
	emptyString := ""

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerChassisTypeInput
		ExpectedCT *ent.ServerChassisType
		errorMsg   string
	}{
		{
			TestName: "updates server chassis type - model",
			Input:    graphclient.UpdateServerChassisTypeInput{Model: &updateString},
			ExpectedCT: &ent.ServerChassisType{
				Model: updateString,
			},
		},
		{
			TestName: "fails to update to empty model",
			Input:    graphclient.UpdateServerChassisTypeInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server chassis type - vendor",
			Input:    graphclient.UpdateServerChassisTypeInput{Vendor: &updateString},
			ExpectedCT: &ent.ServerChassisType{
				Vendor: updateString,
			},
		},
		{
			TestName: "fails to update to empty vendor",
			Input:    graphclient.UpdateServerChassisTypeInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server chassis type - height",
			Input:    graphclient.UpdateServerChassisTypeInput{Height: &updateString},
			ExpectedCT: &ent.ServerChassisType{
				Height: updateString,
			},
		},
		{
			TestName: "fails to update to empty height",
			Input:    graphclient.UpdateServerChassisTypeInput{Height: &emptyString},
			errorMsg: "value is less than the required length",
		},
		// TODO: flaky test
		// {
		// 	TestName: "updates server chassis type - is full depth",
		// 	Input:    graphclient.UpdateServerChassisTypeInput{IsFullDepth: newBool(false)},
		// 	ExpectedCT: &ent.ServerChassisType{
		// 		IsFullDepth: false,
		// 	},
		// },
		// TODO: why can't we access parentchassistypeid
		// {
		// 	TestName: "updates server chassis type - parent chassis type ID",
		// 	Input:    graphclient.UpdateServerChassisTypeInput{ParentChassisTypeID: gidx.MustNewID("srvrsct")},
		// 	ExpectedCT: &ent.ServerChassisType{
		// 		ParentChassisTypeID: gidx.MustNewID("srvrsct"),
		// 	},
		// },
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			ct := (&ServerChassisTypeBuilder{
				Vendor:              gofakeit.CarMaker(),
				Model:               gofakeit.CarModel(),
				Height:              gofakeit.HackerVerb(),
				IsFullDepth:         gofakeit.Bool(),
				ParentChassisTypeID: gidx.MustNewID("srvrsct"),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedCT, ct)

			resp, err := graphTestClient().ServerChassisTypeUpdate(ctx, ct.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerChassisTypeUpdate)

			updatedCT := resp.ServerChassisTypeUpdate.ServerChassisType
			assert.Equal(t, tt.ExpectedCT.Vendor, updatedCT.Vendor)
			assert.Equal(t, tt.ExpectedCT.Model, updatedCT.Model)
			assert.Equal(t, tt.ExpectedCT.Height, updatedCT.Height)
			assert.Equal(t, tt.ExpectedCT.IsFullDepth, updatedCT.IsFullDepth)
			// assert.Equal(t, tt.ExpectedCT.ParentChassisTypeID, updatedCT.ParentChassisTypeID)
			assert.Equal(t, ct.ID, updatedCT.ID)
		})
	}
}

func TestDelete_serverchassistype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ct := (&ServerChassisTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes server chassis type",
			Input:      ct.ID,
			ExpectedID: ct.ID,
		},
		{
			TestName: "fails to delete server chassis type that does not exist",
			Input:    gidx.PrefixedID("srvrcpt-1234"),
			errorMsg: "server_chassis_type not found",
		},
		{
			TestName: "fails to delete empty server chassis type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_chassis_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerChassisTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerChassisTypeDelete)

			deletedCT := resp.ServerChassisTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerChassisTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	srvct := (&ServerChassisTypeBuilder{IsFullDepth: false}).MustNew(ctx)

	exp := &ent.ServerChassisType{
		Vendor:              srvct.Vendor,
		Model:               srvct.Model,
		IsFullDepth:         srvct.IsFullDepth,
		Height:              srvct.Height,
		ParentChassisTypeID: "",
	}

	// create the Chassis type
	createdCTResp, err := graphTestClient().ServerChassisTypeCreate(ctx, graphclient.CreateServerChassisTypeInput{
		Vendor:      exp.Vendor,
		Model:       exp.Model,
		IsFullDepth: &exp.IsFullDepth,
		Height:      exp.Height,
	})

	require.NoError(t, err)
	require.NotNil(t, createdCTResp)
	require.NotNil(t, createdCTResp.ServerChassisTypeCreate.ServerChassisType)

	createdCT := createdCTResp.ServerChassisTypeCreate.ServerChassisType
	require.NotNil(t, createdCT.ID)
	assert.Equal(t, "srvrsct", createdCT.ID.Prefix())
	assert.Equal(t, exp.Vendor, createdCT.Vendor)
	assert.Equal(t, exp.Model, createdCT.Model)
	assert.Equal(t, exp.Height, createdCT.Height)
	assert.Equal(t, exp.IsFullDepth, createdCT.IsFullDepth)
	// assert.Equal(t, exp.ParentChassisTypeID, createdCT.ParentChassisTypeID)

	// Update the Chassis Type
	newVendor := gofakeit.CarMaker()
	updatedCTResp, err := graphTestClient().ServerChassisTypeUpdate(ctx, createdCT.ID, graphclient.UpdateServerChassisTypeInput{Vendor: &newVendor})

	require.NoError(t, err)
	require.NotNil(t, updatedCTResp)
	require.NotNil(t, updatedCTResp.ServerChassisTypeUpdate.ServerChassisType)

	updatedCT := updatedCTResp.ServerChassisTypeUpdate.ServerChassisType
	require.EqualValues(t, createdCT.ID, updatedCT.ID)
	require.Equal(t, newVendor, updatedCT.Vendor)

	// Query the chassis type
	queryCT, err := graphTestClient().GetServerChassisType(ctx, createdCT.ID)
	require.NoError(t, err)
	require.NotNil(t, queryCT)
	require.NotNil(t, queryCT.ServerChassisType)
	require.Equal(t, newVendor, queryCT.ServerChassisType.Vendor)

	// Delete the Server chassis type
	deletedResp, err := graphTestClient().ServerChassisTypeDelete(ctx, createdCT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerChassisTypeDelete)
	require.EqualValues(t, createdCT.ID, deletedResp.ServerChassisTypeDelete.DeletedID.String())

	// Query the Chassis type to ensure it's no longer available
	deletedCPU, err := graphTestClient().GetServerChassisType(ctx, createdCT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCPU)
	require.ErrorContains(t, err, "server_chassis_type not found")
}

// TODO: write tests for cascading deletes

// TODO: write tests to confirm additional subjects

// TODO: write tests to confirm soft deletes
