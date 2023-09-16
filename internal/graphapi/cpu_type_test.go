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

func TestCreate_servercputype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	vendor := gofakeit.CarMaker()
	model := gofakeit.CarModel()
	clock := gofakeit.HackerAdjective()
	count := int64(4)

	testCases := []struct {
		TestName   string
		Input      graphclient.CreateServerCPUTypeInput
		ExpectedCT *ent.ServerCPUType
		errorMsg   string
	}{
		{
			TestName: "creates server cpu type",
			Input:    graphclient.CreateServerCPUTypeInput{Vendor: vendor, Model: model, ClockSpeed: clock, CoreCount: count},
			ExpectedCT: &ent.ServerCPUType{
				Vendor:     vendor,
				Model:      model,
				ClockSpeed: clock,
				CoreCount:  count,
			},
		},
		{
			TestName: "fails to create server cpu type with empty vendor",
			Input:    graphclient.CreateServerCPUTypeInput{Vendor: "", Model: model, ClockSpeed: clock, CoreCount: count},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server cpu type with empty model",
			Input:    graphclient.CreateServerCPUTypeInput{Vendor: vendor, Model: "", ClockSpeed: clock, CoreCount: count},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server cpu type with empty clock speed",
			Input:    graphclient.CreateServerCPUTypeInput{Vendor: vendor, Model: model, ClockSpeed: "", CoreCount: count},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "fails to create server cpu type with empty core count",
			Input:    graphclient.CreateServerCPUTypeInput{Vendor: vendor, Model: model, ClockSpeed: clock, CoreCount: 0},
			errorMsg: "value out of range",
		},
		{
			TestName: "fails to create server cpu type with negative core count",
			Input:    graphclient.CreateServerCPUTypeInput{Vendor: vendor, Model: model, ClockSpeed: clock, CoreCount: -1},
			errorMsg: "value out of range",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			tt := tt
			t.Parallel()

			resp, err := graphTestClient().ServerCPUTypeCreate(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCPUTypeCreate)

			createdCT := resp.ServerCPUTypeCreate.ServerCPUType
			assert.Equal(t, tt.ExpectedCT.Vendor, createdCT.Vendor)
			assert.Equal(t, tt.ExpectedCT.Model, createdCT.Model)
			assert.Equal(t, "srvrcpt", createdCT.ID.Prefix())
			assert.Equal(t, tt.ExpectedCT.ClockSpeed, createdCT.ClockSpeed)
			assert.Equal(t, tt.ExpectedCT.CoreCount, createdCT.CoreCount)
		})
	}
}

func TestUpdate_servercputype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// // Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	updateString := gofakeit.DomainName()
	emptyString := ""

	updateCount := int64(5)
	emptyCount := int64(0)
	negativeCount := int64(-1)

	testCases := []struct {
		TestName   string
		Input      graphclient.UpdateServerCPUTypeInput
		ExpectedCT *ent.ServerCPUType
		errorMsg   string
	}{
		{
			TestName: "updates server cpu type - model",
			Input:    graphclient.UpdateServerCPUTypeInput{Model: &updateString},
			ExpectedCT: &ent.ServerCPUType{
				Model: updateString,
			},
		},
		{
			TestName: "fails to update model to empty",
			Input:    graphclient.UpdateServerCPUTypeInput{Model: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server cpu type - vendor",
			Input:    graphclient.UpdateServerCPUTypeInput{Vendor: &updateString},
			ExpectedCT: &ent.ServerCPUType{
				Vendor: updateString,
			},
		},
		{
			TestName: "fails to update vendor to empty",
			Input:    graphclient.UpdateServerCPUTypeInput{Vendor: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server cpu type - clock speed",
			Input:    graphclient.UpdateServerCPUTypeInput{ClockSpeed: &updateString},
			ExpectedCT: &ent.ServerCPUType{
				ClockSpeed: updateString,
			},
		},
		{
			TestName: "fails to update clock speed to empty",
			Input:    graphclient.UpdateServerCPUTypeInput{ClockSpeed: &emptyString},
			errorMsg: "value is less than the required length",
		},
		{
			TestName: "updates server cpu type - core count",
			Input:    graphclient.UpdateServerCPUTypeInput{CoreCount: &updateCount},
			ExpectedCT: &ent.ServerCPUType{
				CoreCount: updateCount,
			},
		},
		{
			TestName: "fails to update core count to empty",
			Input:    graphclient.UpdateServerCPUTypeInput{CoreCount: &emptyCount},
			errorMsg: "value out of range",
		},
		{
			TestName: "fails to update core count to negative",
			Input:    graphclient.UpdateServerCPUTypeInput{CoreCount: &negativeCount},
			errorMsg: "value out of range",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {

			ct := (&ServerCPUTypeBuilder{
				Vendor:     gofakeit.CarMaker(),
				Model:      gofakeit.CarModel(),
				ClockSpeed: gofakeit.HackerAdjective(),
				CoreCount:  int64(4),
			}).MustNew(ctx)

			_ = mergo.Merge(tt.ExpectedCT, ct)

			resp, err := graphTestClient().ServerCPUTypeUpdate(ctx, ct.ID, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCPUTypeUpdate)

			updatedCT := resp.ServerCPUTypeUpdate.ServerCPUType
			assert.Equal(t, tt.ExpectedCT.Model, updatedCT.Model)
			assert.Equal(t, tt.ExpectedCT.Vendor, updatedCT.Vendor)
			assert.Equal(t, tt.ExpectedCT.ClockSpeed, updatedCT.ClockSpeed)
			assert.Equal(t, tt.ExpectedCT.CoreCount, updatedCT.CoreCount)
			assert.Equal(t, ct.ID, updatedCT.ID)
		})
	}
}

func TestDelete_servercputype(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ct := (&ServerCPUTypeBuilder{}).MustNew(ctx)

	testCases := []struct {
		TestName   string
		Input      gidx.PrefixedID
		ExpectedID gidx.PrefixedID
		errorMsg   string
	}{
		{
			TestName:   "deletes servertype",
			Input:      ct.ID,
			ExpectedID: ct.ID,
		},
		{
			TestName: "fails to delete server cpu type that does not exist",
			Input:    gidx.PrefixedID("srvrcpt-1234"),
			errorMsg: "server_cpu_type not found",
		},
		{
			TestName: "fails to delete empty server cpu type ID",
			Input:    gidx.PrefixedID(""),
			errorMsg: "server_cpu_type not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().ServerCPUTypeDelete(ctx, tt.Input)

			if tt.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.ServerCPUTypeDelete)

			deletedCT := resp.ServerCPUTypeDelete
			assert.EqualValues(t, tt.ExpectedID, deletedCT.DeletedID)
		})
	}
}

func TestFullServerCPUTypeLifecycle(t *testing.T) {
	ctx := context.Background()
	perms := new(mockpermissions.MockPermissions)
	perms.On("CreateAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	perms.On("DeleteAuthRelationships", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx = perms.ContextWithHandler(ctx)

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	exp := &ent.ServerCPUType{
		Vendor:     gofakeit.CarMaker(),
		Model:      gofakeit.CarModel(),
		ClockSpeed: gofakeit.HackerAdjective(),
		CoreCount:  int64(4),
	}

	// create the CT
	createdSrvCPUTypeResp, err := graphTestClient().ServerCPUTypeCreate(ctx, graphclient.CreateServerCPUTypeInput{
		Vendor:     exp.Vendor,
		Model:      exp.Model,
		ClockSpeed: exp.ClockSpeed,
		CoreCount:  exp.CoreCount,
	})

	require.NoError(t, err)
	require.NotNil(t, createdSrvCPUTypeResp)
	require.NotNil(t, createdSrvCPUTypeResp.ServerCPUTypeCreate.ServerCPUType)

	createdCT := createdSrvCPUTypeResp.ServerCPUTypeCreate.ServerCPUType
	require.NotNil(t, createdCT.ID)
	assert.Equal(t, "srvrcpt", createdCT.ID.Prefix())
	assert.Equal(t, exp.Vendor, createdCT.Vendor)
	assert.Equal(t, exp.Model, createdCT.Model)
	assert.Equal(t, exp.ClockSpeed, createdCT.ClockSpeed)
	assert.Equal(t, exp.CoreCount, createdCT.CoreCount)

	// Update the Server
	newVendor := gofakeit.DomainName()
	updatedCTResp, err := graphTestClient().ServerCPUTypeUpdate(ctx, createdCT.ID, graphclient.UpdateServerCPUTypeInput{Vendor: &newVendor})

	require.NoError(t, err)
	require.NotNil(t, updatedCTResp)
	require.NotNil(t, updatedCTResp.ServerCPUTypeUpdate.ServerCPUType)

	updatedCT := updatedCTResp.ServerCPUTypeUpdate.ServerCPUType
	require.EqualValues(t, createdCT.ID, updatedCT.ID)
	require.Equal(t, newVendor, updatedCT.Vendor)

	// Query the CT
	queryCT, err := graphTestClient().GetServerCPUType(ctx, createdCT.ID)
	require.NoError(t, err)
	require.NotNil(t, queryCT)
	require.NotNil(t, queryCT.ServerCPUType)
	require.Equal(t, newVendor, queryCT.ServerCPUType.Vendor)

	// Delete the Server CPU Type
	deletedResp, err := graphTestClient().ServerCPUTypeDelete(ctx, createdCT.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.ServerCPUTypeDelete)
	require.EqualValues(t, createdCT.ID, deletedResp.ServerCPUTypeDelete.DeletedID.String())

	// Query the CT to ensure it's no longer available
	deletedCT, err := graphTestClient().GetServerCPUType(ctx, createdCT.ID)
	require.Error(t, err)
	require.Nil(t, deletedCT)
	require.ErrorContains(t, err, "server_cpu_type not found")
}
