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
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/server"
	"go.infratographer.com/server-api/internal/ent/generated/serverchassis"
	"go.infratographer.com/server-api/internal/ent/generated/serverchassistype"
	"go.infratographer.com/x/gidx"
)

// ServerChassisCreate is the builder for creating a ServerChassis entity.
type ServerChassisCreate struct {
	config
	mutation *ServerChassisMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *ServerChassisCreate) SetCreatedAt(t time.Time) *ServerChassisCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *ServerChassisCreate) SetNillableCreatedAt(t *time.Time) *ServerChassisCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetUpdatedAt sets the "updated_at" field.
func (scc *ServerChassisCreate) SetUpdatedAt(t time.Time) *ServerChassisCreate {
	scc.mutation.SetUpdatedAt(t)
	return scc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (scc *ServerChassisCreate) SetNillableUpdatedAt(t *time.Time) *ServerChassisCreate {
	if t != nil {
		scc.SetUpdatedAt(*t)
	}
	return scc
}

// SetServerChassisTypeID sets the "server_chassis_type_id" field.
func (scc *ServerChassisCreate) SetServerChassisTypeID(gi gidx.PrefixedID) *ServerChassisCreate {
	scc.mutation.SetServerChassisTypeID(gi)
	return scc
}

// SetParentChassisID sets the "parent_chassis_id" field.
func (scc *ServerChassisCreate) SetParentChassisID(gi gidx.PrefixedID) *ServerChassisCreate {
	scc.mutation.SetParentChassisID(gi)
	return scc
}

// SetServerID sets the "server_id" field.
func (scc *ServerChassisCreate) SetServerID(gi gidx.PrefixedID) *ServerChassisCreate {
	scc.mutation.SetServerID(gi)
	return scc
}

// SetSerial sets the "serial" field.
func (scc *ServerChassisCreate) SetSerial(s string) *ServerChassisCreate {
	scc.mutation.SetSerial(s)
	return scc
}

// SetID sets the "id" field.
func (scc *ServerChassisCreate) SetID(gi gidx.PrefixedID) *ServerChassisCreate {
	scc.mutation.SetID(gi)
	return scc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (scc *ServerChassisCreate) SetNillableID(gi *gidx.PrefixedID) *ServerChassisCreate {
	if gi != nil {
		scc.SetID(*gi)
	}
	return scc
}

// SetServer sets the "server" edge to the Server entity.
func (scc *ServerChassisCreate) SetServer(s *Server) *ServerChassisCreate {
	return scc.SetServerID(s.ID)
}

// SetServerChassisType sets the "server_chassis_type" edge to the ServerChassisType entity.
func (scc *ServerChassisCreate) SetServerChassisType(s *ServerChassisType) *ServerChassisCreate {
	return scc.SetServerChassisTypeID(s.ID)
}

// Mutation returns the ServerChassisMutation object of the builder.
func (scc *ServerChassisCreate) Mutation() *ServerChassisMutation {
	return scc.mutation
}

// Save creates the ServerChassis in the database.
func (scc *ServerChassisCreate) Save(ctx context.Context) (*ServerChassis, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *ServerChassisCreate) SaveX(ctx context.Context) *ServerChassis {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *ServerChassisCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *ServerChassisCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *ServerChassisCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := serverchassis.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		v := serverchassis.DefaultUpdatedAt()
		scc.mutation.SetUpdatedAt(v)
	}
	if _, ok := scc.mutation.ID(); !ok {
		v := serverchassis.DefaultID()
		scc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *ServerChassisCreate) check() error {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "ServerChassis.created_at"`)}
	}
	if _, ok := scc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "ServerChassis.updated_at"`)}
	}
	if _, ok := scc.mutation.ServerChassisTypeID(); !ok {
		return &ValidationError{Name: "server_chassis_type_id", err: errors.New(`generated: missing required field "ServerChassis.server_chassis_type_id"`)}
	}
	if _, ok := scc.mutation.ParentChassisID(); !ok {
		return &ValidationError{Name: "parent_chassis_id", err: errors.New(`generated: missing required field "ServerChassis.parent_chassis_id"`)}
	}
	if _, ok := scc.mutation.ServerID(); !ok {
		return &ValidationError{Name: "server_id", err: errors.New(`generated: missing required field "ServerChassis.server_id"`)}
	}
	if _, ok := scc.mutation.Serial(); !ok {
		return &ValidationError{Name: "serial", err: errors.New(`generated: missing required field "ServerChassis.serial"`)}
	}
	if v, ok := scc.mutation.Serial(); ok {
		if err := serverchassis.SerialValidator(v); err != nil {
			return &ValidationError{Name: "serial", err: fmt.Errorf(`generated: validator failed for field "ServerChassis.serial": %w`, err)}
		}
	}
	if _, ok := scc.mutation.ServerID(); !ok {
		return &ValidationError{Name: "server", err: errors.New(`generated: missing required edge "ServerChassis.server"`)}
	}
	if _, ok := scc.mutation.ServerChassisTypeID(); !ok {
		return &ValidationError{Name: "server_chassis_type", err: errors.New(`generated: missing required edge "ServerChassis.server_chassis_type"`)}
	}
	return nil
}

func (scc *ServerChassisCreate) sqlSave(ctx context.Context) (*ServerChassis, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *ServerChassisCreate) createSpec() (*ServerChassis, *sqlgraph.CreateSpec) {
	var (
		_node = &ServerChassis{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(serverchassis.Table, sqlgraph.NewFieldSpec(serverchassis.FieldID, field.TypeString))
	)
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.SetField(serverchassis.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.UpdatedAt(); ok {
		_spec.SetField(serverchassis.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := scc.mutation.ParentChassisID(); ok {
		_spec.SetField(serverchassis.FieldParentChassisID, field.TypeString, value)
		_node.ParentChassisID = value
	}
	if value, ok := scc.mutation.Serial(); ok {
		_spec.SetField(serverchassis.FieldSerial, field.TypeString, value)
		_node.Serial = value
	}
	if nodes := scc.mutation.ServerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serverchassis.ServerTable,
			Columns: []string{serverchassis.ServerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ServerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := scc.mutation.ServerChassisTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serverchassis.ServerChassisTypeTable,
			Columns: []string{serverchassis.ServerChassisTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(serverchassistype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ServerChassisTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServerChassisCreateBulk is the builder for creating many ServerChassis entities in bulk.
type ServerChassisCreateBulk struct {
	config
	builders []*ServerChassisCreate
}

// Save creates the ServerChassis entities in the database.
func (sccb *ServerChassisCreateBulk) Save(ctx context.Context) ([]*ServerChassis, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*ServerChassis, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerChassisMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *ServerChassisCreateBulk) SaveX(ctx context.Context) []*ServerChassis {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *ServerChassisCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *ServerChassisCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}
