package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/server-api/internal/ent/generated"
)

type ServerTypeBuilder struct {
	Name    string
	OwnerID gidx.PrefixedID
}

func (p ServerTypeBuilder) MustNew(ctx context.Context) *ent.ServerType {
	if p.Name == "" {
		p.Name = gofakeit.JobTitle()
	}

	if p.OwnerID == "" {
		p.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	return EntClient.ServerType.Create().SetName(p.Name).SetOwnerID(p.OwnerID).SaveX(ctx)
}

type ProviderBuilder struct {
	Name               string
	ResourceProviderID gidx.PrefixedID
}

func (p ProviderBuilder) MustNew(ctx context.Context) *ent.Provider {
	if p.Name == "" {
		p.Name = gofakeit.JobTitle()
	}

	if p.ResourceProviderID == "" {
		p.ResourceProviderID = gidx.MustNewID(ownerPrefix)
	}

	return EntClient.Provider.Create().SetResourceProviderID(p.ResourceProviderID).SetName(p.Name).SaveX(ctx)
}

type ServerBuilder struct {
	Name         string
	Description  string
	OwnerID      gidx.PrefixedID
	LocationID   gidx.PrefixedID
	ProviderID   gidx.PrefixedID
	Provider     *ent.Provider
	ServerType   *ent.ServerType
	ServerTypeID gidx.PrefixedID
}

func (b ServerBuilder) MustNew(ctx context.Context) *ent.Server {
	if b.Provider == nil {
		b.Provider = (&ProviderBuilder{ResourceProviderID: b.ProviderID}).MustNew(ctx)
	}

	if b.ServerType == nil {
		b.ServerType = (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
	}

	if b.ProviderID == "" {
		b.ProviderID = b.Provider.ID
	}

	if b.OwnerID == "" {
		b.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	if b.ServerTypeID == "" {
		b.ServerTypeID = b.ServerType.ID
	}

	if b.Description == "" {
		b.Description = gofakeit.Sentence(10)
	}

	if b.Name == "" {
		b.Name = gofakeit.AppName()
	}

	if b.LocationID == "" {
		b.LocationID = gidx.MustNewID(locationPrefix)
	}

	// if b.ServerTypeID == "" {
	// 	b.ServerTypeID = gidx.MustNewID(schema.ServerTypePrefix)
	// }

	return EntClient.Server.Create().SetName(b.Name).SetOwnerID(b.OwnerID).SetLocationID(b.LocationID).SetProviderID(b.ProviderID).SetServerTypeID(b.ServerTypeID).SetDescription(b.Description).SaveX(ctx)
}

type ServerCPUTypeBuilder struct {
	Model      string
	Vendor     string
	CoreCount  int64
	ClockSpeed string
}

func (p ServerCPUTypeBuilder) MustNew(ctx context.Context) *ent.ServerCPUType {
	if p.Model == "" {
		p.Model = gofakeit.CarModel()
	}

	if p.Vendor == "" {
		p.Vendor = gofakeit.CarMaker()
	}

	if p.CoreCount == 0 {
		p.CoreCount = 4
	}

	if p.ClockSpeed == "" {
		p.ClockSpeed = "2.5GHz"
	}

	return EntClient.ServerCPUType.Create().SetVendor(p.Vendor).SetModel(p.Model).SetCoreCount(p.CoreCount).SetClockSpeed(p.ClockSpeed).SaveX(ctx)
}
