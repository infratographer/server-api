package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"database/sql"

	"go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// ServerComponentCreate is the resolver for the serverComponentCreate field.
func (r *mutationResolver) ServerComponentCreate(ctx context.Context, input generated.CreateServerComponentInput) (*ServerComponentCreatePayload, error) {
	// TODO: check permissions

	sc, err := r.client.ServerComponent.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerComponentCreatePayload{ServerComponent: sc}, nil
}

// ServerComponentUpdate is the resolver for the serverComponentUpdate field.
func (r *mutationResolver) ServerComponentUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateServerComponentInput) (*ServerComponentUpdatePayload, error) {
	// TODO: check permissions

	sc, err := r.client.ServerComponent.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	sc, err = sc.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerComponentUpdatePayload{ServerComponent: sc}, nil
}

// ServerComponentDelete is the resolver for the serverComponentDelete field.
func (r *mutationResolver) ServerComponentDelete(ctx context.Context, id gidx.PrefixedID) (*ServerComponentDeletePayload, error) {

	//TODO: check permissions

	tx, err := r.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	if err := tx.ServerComponent.DeleteOneID(id).Exec(ctx); err != nil {
		r.logger.Errorw("failed to commit transaction", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete server component")
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		r.logger.Errorw("failed to commit transaction", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "commit transaction")
		}
		return nil, err
	}

	return &ServerComponentDeletePayload{DeletedID: id}, nil

}

// ServerComponent is the resolver for the serverComponent field.
func (r *queryResolver) ServerComponent(ctx context.Context, id gidx.PrefixedID) (*generated.ServerComponent, error) {
	return r.client.ServerComponent.Get(ctx, id)

}
