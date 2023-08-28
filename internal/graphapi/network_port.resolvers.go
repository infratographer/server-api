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

// ServerNetworkPort is the resolver for the serverNetworkPort field.
func (r *mutationResolver) ServerNetworkPort(ctx context.Context, input generated.CreateServerNetworkPortInput) (*ServerNetworkPortCreatePayload, error) {
	// TODO: check permissions

	np, err := r.client.ServerNetworkPort.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerNetworkPortCreatePayload{ServerNetworkPort: np}, nil
}

// ServerNetworkPortUpdate is the resolver for the serverNetworkPortUpdate field.
func (r *mutationResolver) ServerNetworkPortUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateServerNetworkPortInput) (*ServerNetworkPortUpdatePayload, error) {
	// TODO: check permissions

	np, err := r.client.ServerNetworkPort.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	np, err = np.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerNetworkPortUpdatePayload{ServerNetworkPort: np}, nil
}

// ServerNetworkPortDelete is the resolver for the serverNetworkPortDelete field.
func (r *mutationResolver) ServerNetworkPortDelete(ctx context.Context, id gidx.PrefixedID) (*ServerNetworkPortDeletePayload, error) {
	//TODO: check permissions

	tx, err := r.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	if err := tx.ServerNetworkPort.DeleteOneID(id).Exec(ctx); err != nil {
		r.logger.Errorw("failed to commit transaction", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete network port")
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

	return &ServerNetworkPortDeletePayload{DeletedID: id}, nil
}

// ServerNetworkPort is the resolver for the serverNetworkPort field.
func (r *queryResolver) ServerNetworkPort(ctx context.Context, id gidx.PrefixedID) (*generated.ServerNetworkPort, error) {
	return r.client.ServerNetworkPort.Get(ctx, id)
}
