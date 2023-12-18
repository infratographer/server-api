// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkport"
)

// ServerNetworkPortDelete is the builder for deleting a ServerNetworkPort entity.
type ServerNetworkPortDelete struct {
	config
	hooks    []Hook
	mutation *ServerNetworkPortMutation
}

// Where appends a list predicates to the ServerNetworkPortDelete builder.
func (snpd *ServerNetworkPortDelete) Where(ps ...predicate.ServerNetworkPort) *ServerNetworkPortDelete {
	snpd.mutation.Where(ps...)
	return snpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (snpd *ServerNetworkPortDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, snpd.sqlExec, snpd.mutation, snpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (snpd *ServerNetworkPortDelete) ExecX(ctx context.Context) int {
	n, err := snpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (snpd *ServerNetworkPortDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(servernetworkport.Table, sqlgraph.NewFieldSpec(servernetworkport.FieldID, field.TypeString))
	if ps := snpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, snpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	snpd.mutation.done = true
	return affected, err
}

// ServerNetworkPortDeleteOne is the builder for deleting a single ServerNetworkPort entity.
type ServerNetworkPortDeleteOne struct {
	snpd *ServerNetworkPortDelete
}

// Where appends a list predicates to the ServerNetworkPortDelete builder.
func (snpdo *ServerNetworkPortDeleteOne) Where(ps ...predicate.ServerNetworkPort) *ServerNetworkPortDeleteOne {
	snpdo.snpd.mutation.Where(ps...)
	return snpdo
}

// Exec executes the deletion query.
func (snpdo *ServerNetworkPortDeleteOne) Exec(ctx context.Context) error {
	n, err := snpdo.snpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{servernetworkport.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (snpdo *ServerNetworkPortDeleteOne) ExecX(ctx context.Context) {
	if err := snpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
