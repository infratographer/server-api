package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// FindServerByID is the resolver for the findServerByID field.
func (r *entityResolver) FindServerByID(ctx context.Context, id gidx.PrefixedID) (*generated.Server, error) {
	panic(fmt.Errorf("not implemented: FindServerByID - findServerByID"))
}

// FindServerCPUByID is the resolver for the findServerCPUByID field.
func (r *entityResolver) FindServerCPUByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerCPU, error) {
	panic(fmt.Errorf("not implemented: FindServerCPUByID - findServerCPUByID"))
}

// FindServerCPUTypeByID is the resolver for the findServerCPUTypeByID field.
func (r *entityResolver) FindServerCPUTypeByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerCPUType, error) {
	panic(fmt.Errorf("not implemented: FindServerCPUTypeByID - findServerCPUTypeByID"))
}

// FindServerChassisByID is the resolver for the findServerChassisByID field.
func (r *entityResolver) FindServerChassisByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerChassis, error) {
	panic(fmt.Errorf("not implemented: FindServerChassisByID - findServerChassisByID"))
}

// FindServerChassisTypeByID is the resolver for the findServerChassisTypeByID field.
func (r *entityResolver) FindServerChassisTypeByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerChassisType, error) {
	panic(fmt.Errorf("not implemented: FindServerChassisTypeByID - findServerChassisTypeByID"))
}

// FindServerComponentByID is the resolver for the findServerComponentByID field.
func (r *entityResolver) FindServerComponentByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerComponent, error) {
	panic(fmt.Errorf("not implemented: FindServerComponentByID - findServerComponentByID"))
}

// FindServerComponentTypeByID is the resolver for the findServerComponentTypeByID field.
func (r *entityResolver) FindServerComponentTypeByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerComponentType, error) {
	panic(fmt.Errorf("not implemented: FindServerComponentTypeByID - findServerComponentTypeByID"))
}

// FindServerProviderByID is the resolver for the findServerProviderByID field.
func (r *entityResolver) FindServerProviderByID(ctx context.Context, id gidx.PrefixedID) (*generated.Provider, error) {
	panic(fmt.Errorf("not implemented: FindServerProviderByID - findServerProviderByID"))
}

// FindServerTypeByID is the resolver for the findServerTypeByID field.
func (r *entityResolver) FindServerTypeByID(ctx context.Context, id gidx.PrefixedID) (*generated.ServerType, error) {
	panic(fmt.Errorf("not implemented: FindServerTypeByID - findServerTypeByID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
