package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"database/sql"

	"go.infratographer.com/server-api/internal/ent/generated"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/server-api/internal/ent/generated/servermotherboard"
	"go.infratographer.com/x/gidx"
)

// ServerMotherboardTypeCreate is the resolver for the serverMotherboardTypeCreate field.
func (r *mutationResolver) ServerMotherboardTypeCreate(ctx context.Context, input generated.CreateServerMotherboardTypeInput) (*ServerMotherboardTypeCreatePayload, error) {
	// TODO: check permissions

	mb, err := r.client.ServerMotherboardType.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerMotherboardTypeCreatePayload{ServerMotherboardType: mb}, nil
}

// ServerMotherboardTypeUpdate is the resolver for the serverMotherboardTypeUpdate field.
func (r *mutationResolver) ServerMotherboardTypeUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateServerMotherboardTypeInput) (*ServerMotherboardTypeUpdatePayload, error) {
	// TODO: check permissions

	mb, err := r.client.ServerMotherboardType.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	mb, err = mb.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &ServerMotherboardTypeUpdatePayload{ServerMotherboardType: mb}, nil
}

// ServerMotherboardTypeDelete is the resolver for the serverMotherboardTypeDelete field.
func (r *mutationResolver) ServerMotherboardTypeDelete(ctx context.Context, id gidx.PrefixedID) (*ServerMotherboardTypeDeletePayload, error) {
	//TODO: check permissions

	tx, err := r.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// cleanup chassis associated with type
	mb, err := tx.ServerMotherboard.Query().Where(predicate.ServerMotherboard(servermotherboard.ServerMotherboardTypeIDEQ(id))).All(ctx)
	if err != nil {
		r.logger.Errorw("failed to query motherboards", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "query motherboards")
		}
		return nil, err
	}

	for _, m := range mb {
		if err = tx.ServerMotherboard.DeleteOne(m).Exec(ctx); err != nil {
			r.logger.Errorw("failed to delete server", "port", m.ID, "error", err)
			if rerr := tx.Rollback(); rerr != nil {
				r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete motherboard")
			}
		}
	}

	if err := tx.ServerMotherboardType.DeleteOneID(id).Exec(ctx); err != nil {
		r.logger.Errorw("failed to commit transaction", "error", err)
		if rerr := tx.Rollback(); rerr != nil {
			r.logger.Errorw("failed to rollback transaction", "error", rerr, "stage", "delete motherboard type")
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

	return &ServerMotherboardTypeDeletePayload{DeletedID: id}, nil
}

// ServerMotherboardType is the resolver for the serverMotherboardType field.
func (r *queryResolver) ServerMotherboardType(ctx context.Context, id gidx.PrefixedID) (*generated.ServerMotherboardType, error) {
	return r.client.ServerMotherboardType.Get(ctx, id)
}
