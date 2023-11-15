package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/server-api/internal/ent/generated/serverharddrivetype"
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

type ServerCPUBuilder struct {
	Serial        string
	Server        *ent.Server
	ServerCPUType *ent.ServerCPUType
}

func (p ServerCPUBuilder) MustNew(ctx context.Context) *ent.ServerCPU {
	if p.Serial == "" {
		p.Serial = gofakeit.UUID()
	}

	if p.Server == nil {
		provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
		srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		p.Server = (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	}

	if p.ServerCPUType == nil {
		p.ServerCPUType = (&ServerCPUTypeBuilder{}).MustNew(ctx)
	}

	return EntClient.ServerCPU.Create().SetSerial(p.Serial).SetServer(p.Server).SetServerCPUType(p.ServerCPUType).SaveX(ctx)
}

type ServerChassisTypeBuilder struct {
	Model               string
	Vendor              string
	Height              string
	IsFullDepth         bool
	ParentChassisTypeID gidx.PrefixedID
}

func (p ServerChassisTypeBuilder) MustNew(ctx context.Context) *ent.ServerChassisType {
	if p.Model == "" {
		p.Model = gofakeit.CarModel()
	}

	if p.Vendor == "" {
		p.Vendor = gofakeit.CarMaker()
	}

	if p.Height == "" {
		p.Height = "1U"
	}

	return EntClient.ServerChassisType.Create().SetModel(p.Model).SetVendor(p.Vendor).SetHeight(p.Height).SetParentChassisTypeID(p.ParentChassisTypeID).SetIsFullDepth(p.IsFullDepth).SaveX(ctx)
}

type ServerChassisBuilder struct {
	Serial string
	Server *ent.Server
	// ParentChassis     *ent.ServerChassis
	ParentChassisTypeID gidx.PrefixedID
	ServerChassisType   *ent.ServerChassisType
}

func (p ServerChassisBuilder) MustNew(ctx context.Context) *ent.ServerChassis {
	if p.Serial == "" {
		p.Serial = gofakeit.UUID()
	}

	if p.Server == nil {
		provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
		srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		p.Server = (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	}

	if p.ServerChassisType == nil {
		p.ServerChassisType = (&ServerChassisTypeBuilder{IsFullDepth: true}).MustNew(ctx)
	}

	// if p.ParentChassis == nil {
	// 	provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
	// 	srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
	// 	server := (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	// 	p.ParentChassis = (&ServerChassisBuilder{Server: server, ServerChassisType: p.ServerChassisType}).MustNew(ctx)
	// }

	// return EntClient.ServerChassis.Create().SetServer(p.Server).SetSerial(p.Serial).SetServerChassisType(p.ServerChassisType).SetParentChassisID(p.ParentChassis.ID).SaveX(ctx)
	return EntClient.ServerChassis.Create().SetServer(p.Server).SetSerial(p.Serial).SetServerChassisType(p.ServerChassisType).SetParentChassisID(p.ParentChassisTypeID).SaveX(ctx)
}

type ServerComponentTypeBuilder struct {
	Name string
}

func (p ServerComponentTypeBuilder) MustNew(ctx context.Context) *ent.ServerComponentType {
	if p.Name == "" {
		p.Name = gofakeit.DomainName()
	}

	return EntClient.ServerComponentType.Create().SetName(p.Name).SaveX(ctx)
}

type ServerComponentBuilder struct {
	Name                string
	Model               string
	Vendor              string
	Serial              string
	Server              *ent.Server
	ServerComponentType *ent.ServerComponentType
}

func (p ServerComponentBuilder) MustNew(ctx context.Context) *ent.ServerComponent {
	if p.Name == "" {
		p.Name = gofakeit.DomainName()
	}

	if p.Model == "" {
		p.Model = gofakeit.CarModel()
	}

	if p.Vendor == "" {
		p.Vendor = gofakeit.CarMaker()
	}

	if p.Serial == "" {
		p.Serial = gofakeit.UUID()
	}

	if p.Server == nil {
		provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
		srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		p.Server = (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	}

	if p.ServerComponentType == nil {
		p.ServerComponentType = (&ServerComponentTypeBuilder{}).MustNew(ctx)
	}

	return EntClient.ServerComponent.Create().SetName(p.Name).SetModel(p.Model).SetVendor(p.Vendor).SetSerial(p.Serial).SetServer(p.Server).SetComponentType(p.ServerComponentType).SaveX(ctx)
}

type ServerHardDriveTypeBuilder struct {
	Model        string
	Vendor       string
	Speed        string
	Capacity     string
	Type         serverharddrivetype.Type
	HardDriveIDs []gidx.PrefixedID
}

func (p ServerHardDriveTypeBuilder) MustNew(ctx context.Context) *ent.ServerHardDriveType {
	if p.Model == "" {
		p.Model = gofakeit.CarModel()
	}

	if p.Vendor == "" {
		p.Vendor = gofakeit.CarMaker()
	}

	if p.Capacity == "" {
		p.Capacity = gofakeit.DomainName()
	}

	if p.Speed == "" {
		p.Speed = gofakeit.DomainName()
	}

	if p.Type == "" {
		p.Type = serverharddrivetype.TypeHdd
	}

	return EntClient.ServerHardDriveType.Create().SetVendor(p.Vendor).SetModel(p.Model).SetCapacity(p.Capacity).SetSpeed(p.Speed).SetType(p.Type).SaveX(ctx)
}

type ServerHardDriveBuilder struct {
	Serial              string
	Server              *ent.Server
	ServerHardDriveType *ent.ServerHardDriveType
}

func (p ServerHardDriveBuilder) MustNew(ctx context.Context) *ent.ServerHardDrive {
	if p.Serial == "" {
		p.Serial = gofakeit.UUID()
	}

	if p.Server == nil {
		provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
		srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		p.Server = (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	}

	if p.ServerHardDriveType == nil {
		p.ServerHardDriveType = (&ServerHardDriveTypeBuilder{}).MustNew(ctx)
	}

	return EntClient.ServerHardDrive.Create().SetSerial(p.Serial).SetServer(p.Server).SetServerHardDriveTypeID(p.ServerHardDriveType.ID).SaveX(ctx)
}

type ServerMemoryTypeBuilder struct {
	Vendor string
	Model  string
	Size   string
	Speed  string
}

func (p ServerMemoryTypeBuilder) MustNew(ctx context.Context) *ent.ServerMemoryType {
	if p.Vendor == "" {
		p.Vendor = gofakeit.CarMaker()
	}

	if p.Model == "" {
		p.Model = gofakeit.CarModel()
	}

	if p.Speed == "" {
		p.Speed = gofakeit.Animal()
	}

	if p.Size == "" {
		p.Size = gofakeit.AnimalType()
	}

	return EntClient.ServerMemoryType.Create().SetModel(p.Model).SetSize(p.Size).SetSpeed(p.Speed).SetVendor(p.Vendor).SaveX(ctx)
}

type ServerMemoryBuilder struct {
	Serial           string
	Server           *ent.Server
	ServerMemoryType *ent.ServerMemoryType
}

func (p ServerMemoryBuilder) MustNew(ctx context.Context) *ent.ServerMemory {
	if p.Serial == "" {
		p.Serial = gofakeit.UUID()
	}

	if p.Server == nil {
		provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
		srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		p.Server = (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	}

	if p.ServerMemoryType == nil {
		p.ServerMemoryType = (&ServerMemoryTypeBuilder{}).MustNew(ctx)
	}

	return EntClient.ServerMemory.Create().SetSerial(p.Serial).SetServer(p.Server).SetServerMemoryType(p.ServerMemoryType).SaveX(ctx)
}

type ServerMotherboardTypeBuilder struct {
	Vendor string
	Model  string
}

func (p ServerMotherboardTypeBuilder) MustNew(ctx context.Context) *ent.ServerMotherboardType {
	if p.Vendor == "" {
		p.Vendor = gofakeit.CarMaker()
	}

	if p.Model == "" {
		p.Model = gofakeit.CarModel()
	}

	return EntClient.ServerMotherboardType.Create().SetModel(p.Model).SetVendor(p.Vendor).SaveX(ctx)
}

type ServerMotherboardBuilder struct {
	Serial                string
	Server                *ent.Server
	ServerMotherBoardType *ent.ServerMotherboardType
}

func (p ServerMotherboardBuilder) MustNew(ctx context.Context) *ent.ServerMotherboard {
	if p.Serial == "" {
		p.Serial = gofakeit.UUID()
	}

	if p.Server == nil {
		provider := (&ProviderBuilder{ResourceProviderID: gidx.MustNewID(resourceProviderPrefix)}).MustNew(ctx)
		srvtype := (&ServerTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		p.Server = (&ServerBuilder{ProviderID: provider.ID, ServerTypeID: srvtype.ID}).MustNew(ctx)
	}

	if p.ServerMotherBoardType == nil {
		p.ServerMotherBoardType = (&ServerMotherboardTypeBuilder{}).MustNew(ctx)
	}

	return EntClient.ServerMotherboard.Create().SetSerial(p.Serial).SetServer(p.Server).SetServerMotherboardType(p.ServerMotherBoardType).SaveX(ctx)
}
