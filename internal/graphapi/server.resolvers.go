package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"database/sql"

	"go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// ServerCreate is the resolver for the serverCreate field.
func (r *mutationResolver) ServerCreate(ctx context.Context, input generated.CreateServerInput) (*ServerCreatePayload, error) {
	// if err := permissions.CheckAccess(ctx, input.OwnerID, actionServerCreate); err != nil {
	// 	return nil, err
	// }

	srv, err := r.client.Server.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerCreatePayload{Server: srv}, nil
}

// ServerUpdate is the resolver for the serverUpdate field.
func (r *mutationResolver) ServerUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateServerInput) (*ServerUpdatePayload, error) {
	// TODO: check permissions

	srv, err := r.client.Server.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	srv, err = srv.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerUpdatePayload{Server: srv}, nil
}

// ServerDelete is the resolver for the serverDelete field.
func (r *mutationResolver) ServerDelete(ctx context.Context, id gidx.PrefixedID) (*ServerDeletePayload, error) {
	//TODO: check permissions

	tx, err := r.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// TODO: theres a whole lotta things we got do some cascading deletes for

	if err := tx.Server.DeleteOneID(id).Exec(ctx); err != nil {
		r.logger.Errorw("failed to commit transaction", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete server")
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

	return &ServerDeletePayload{DeletedID: id}, nil
}

// Server is the resolver for the server field.
func (r *queryResolver) Server(ctx context.Context, id gidx.PrefixedID) (*generated.Server, error) {
	//TODO: check permissions

	return r.client.Server.Get(ctx, id)
}
