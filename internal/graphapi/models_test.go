package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/server-api/internal/ent/generated"
)

type ProviderBuilder struct {
	Name               string
	ResourceProviderID gidx.PrefixedID
}

func (p ProviderBuilder) MustNew(ctx context.Context) *ent.Provider {
	if p.Name == "" {
		p.Name = gofakeit.JobTitle()
	}

	if p.ResourceProviderID == "" {
		p.ResourceProviderID = gidx.MustNewID(resourceProviderPrefix)
	}

	return EntClient.Provider.Create().SetName(p.Name).SetResourceProviderID(gidx.MustNewID("testprv")).SaveX(ctx)
}

type ServerBuilder struct {
	Name         string
	Description  string
	OwnerID      gidx.PrefixedID
	LocationID   gidx.PrefixedID
	ProviderID   gidx.PrefixedID
	ServerTypeID gidx.PrefixedID
}

func (b ServerBuilder) MustNew(ctx context.Context) *ent.Server {
	if b.ProviderID == "" {
		b.ProviderID = gidx.MustNewID(resourceProviderPrefix)
	}

	if b.Name == "" {
		b.Name = gofakeit.AppName()
	}

	if b.Description == "" {
		b.Description = gofakeit.HackerPhrase()
	}

	if b.OwnerID == "" {
		b.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	if b.LocationID == "" {
		b.LocationID = gidx.MustNewID(locationPrefix)
	}

	if b.ServerTypeID == "" {
		b.ServerTypeID = gidx.MustNewID("testst")
	}

	return EntClient.Server.Create().SetName(b.Name).SetOwnerID(b.OwnerID).SetLocationID(b.LocationID).SetProviderID(b.ProviderID).SetDescription(b.Description).SetServerTypeID(b.ServerTypeID).SaveX(ctx)
}

// type PortBuilder struct {
// 	Name           string
// 	LoadBalancerID gidx.PrefixedID
// 	Number         int
// }

// func (p PortBuilder) MustNew(ctx context.Context) *ent.Port {
// 	if p.Name == "" {
// 		p.Name = gofakeit.AppName()
// 	}

// 	if p.LoadBalancerID == "" {
// 		p.LoadBalancerID = gidx.MustNewID(lbPrefix)
// 	}

// 	if p.Number == 0 {
// 		p.Number = gofakeit.Number(1, 65535)
// 	}

// 	return EntClient.Port.Create().SetName(p.Name).SetLoadBalancerID(p.LoadBalancerID).SetNumber(p.Number).SaveX(ctx)
// }

// type PoolBuilder struct {
// 	Name     string
// 	OwnerID  gidx.PrefixedID
// 	Protocol pool.Protocol
// }

// func (p *PoolBuilder) MustNew(ctx context.Context) *ent.Pool {
// 	if p.Name == "" {
// 		p.Name = gofakeit.AppName()
// 	}

// 	if p.OwnerID == "" {
// 		p.OwnerID = gidx.MustNewID(ownerPrefix)
// 	}

// 	if p.Protocol == "" {
// 		p.Protocol = pool.Protocol(gofakeit.RandomString([]string{"tcp", "udp"}))
// 	}

// 	return EntClient.Pool.Create().SetName(p.Name).SetOwnerID(p.OwnerID).SetProtocol(p.Protocol).SaveX(ctx)
// }

// type OriginBuilder struct {
// 	Name       string
// 	Target     string
// 	PortNumber int
// 	Active     bool
// 	PoolID     gidx.PrefixedID
// }

// func (o *OriginBuilder) MustNew(ctx context.Context) *ent.Origin {
// 	if o.Name == "" {
// 		o.Name = gofakeit.AppName()
// 	}

// 	if o.Target == "" {
// 		o.Target = gofakeit.IPv4Address()
// 	}

// 	if o.PortNumber == 0 {
// 		o.PortNumber = gofakeit.Number(1, 65535)
// 	}

// 	if o.PoolID == "" {
// 		pb := &PoolBuilder{}
// 		o.PoolID = pb.MustNew(ctx).ID
// 	}

// 	return EntClient.Origin.Create().SetName(o.Name).SetTarget(o.Target).SetPortNumber(o.PortNumber).SetActive(o.Active).SetPoolID(o.PoolID).SaveX(ctx)
// }
