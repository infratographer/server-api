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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/server-api/internal/ent/generated/serverchassis"
)

// ServerChassisUpdate is the builder for updating ServerChassis entities.
type ServerChassisUpdate struct {
	config
	hooks    []Hook
	mutation *ServerChassisMutation
}

// Where appends a list predicates to the ServerChassisUpdate builder.
func (scu *ServerChassisUpdate) Where(ps ...predicate.ServerChassis) *ServerChassisUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetSerial sets the "serial" field.
func (scu *ServerChassisUpdate) SetSerial(s string) *ServerChassisUpdate {
	scu.mutation.SetSerial(s)
	return scu
}

// Mutation returns the ServerChassisMutation object of the builder.
func (scu *ServerChassisUpdate) Mutation() *ServerChassisMutation {
	return scu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *ServerChassisUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *ServerChassisUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *ServerChassisUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *ServerChassisUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *ServerChassisUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := serverchassis.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *ServerChassisUpdate) check() error {
	if v, ok := scu.mutation.Serial(); ok {
		if err := serverchassis.SerialValidator(v); err != nil {
			return &ValidationError{Name: "serial", err: fmt.Errorf(`generated: validator failed for field "ServerChassis.serial": %w`, err)}
		}
	}
	if _, ok := scu.mutation.ServerID(); scu.mutation.ServerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "ServerChassis.server"`)
	}
	if _, ok := scu.mutation.ServerChassisTypeID(); scu.mutation.ServerChassisTypeCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "ServerChassis.server_chassis_type"`)
	}
	return nil
}

func (scu *ServerChassisUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(serverchassis.Table, serverchassis.Columns, sqlgraph.NewFieldSpec(serverchassis.FieldID, field.TypeString))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.SetField(serverchassis.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := scu.mutation.Serial(); ok {
		_spec.SetField(serverchassis.FieldSerial, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serverchassis.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// ServerChassisUpdateOne is the builder for updating a single ServerChassis entity.
type ServerChassisUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServerChassisMutation
}

// SetSerial sets the "serial" field.
func (scuo *ServerChassisUpdateOne) SetSerial(s string) *ServerChassisUpdateOne {
	scuo.mutation.SetSerial(s)
	return scuo
}

// Mutation returns the ServerChassisMutation object of the builder.
func (scuo *ServerChassisUpdateOne) Mutation() *ServerChassisMutation {
	return scuo.mutation
}

// Where appends a list predicates to the ServerChassisUpdate builder.
func (scuo *ServerChassisUpdateOne) Where(ps ...predicate.ServerChassis) *ServerChassisUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *ServerChassisUpdateOne) Select(field string, fields ...string) *ServerChassisUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated ServerChassis entity.
func (scuo *ServerChassisUpdateOne) Save(ctx context.Context) (*ServerChassis, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *ServerChassisUpdateOne) SaveX(ctx context.Context) *ServerChassis {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *ServerChassisUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *ServerChassisUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *ServerChassisUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := serverchassis.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *ServerChassisUpdateOne) check() error {
	if v, ok := scuo.mutation.Serial(); ok {
		if err := serverchassis.SerialValidator(v); err != nil {
			return &ValidationError{Name: "serial", err: fmt.Errorf(`generated: validator failed for field "ServerChassis.serial": %w`, err)}
		}
	}
	if _, ok := scuo.mutation.ServerID(); scuo.mutation.ServerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "ServerChassis.server"`)
	}
	if _, ok := scuo.mutation.ServerChassisTypeID(); scuo.mutation.ServerChassisTypeCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "ServerChassis.server_chassis_type"`)
	}
	return nil
}

func (scuo *ServerChassisUpdateOne) sqlSave(ctx context.Context) (_node *ServerChassis, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(serverchassis.Table, serverchassis.Columns, sqlgraph.NewFieldSpec(serverchassis.FieldID, field.TypeString))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ServerChassis.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serverchassis.FieldID)
		for _, f := range fields {
			if !serverchassis.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != serverchassis.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.SetField(serverchassis.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := scuo.mutation.Serial(); ok {
		_spec.SetField(serverchassis.FieldSerial, field.TypeString, value)
	}
	_node = &ServerChassis{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serverchassis.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
