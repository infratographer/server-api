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
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkcard"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkcardtype"
	"go.infratographer.com/x/gidx"
)

// ServerNetworkCardTypeUpdate is the builder for updating ServerNetworkCardType entities.
type ServerNetworkCardTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ServerNetworkCardTypeMutation
}

// Where appends a list predicates to the ServerNetworkCardTypeUpdate builder.
func (snctu *ServerNetworkCardTypeUpdate) Where(ps ...predicate.ServerNetworkCardType) *ServerNetworkCardTypeUpdate {
	snctu.mutation.Where(ps...)
	return snctu
}

// SetVendor sets the "vendor" field.
func (snctu *ServerNetworkCardTypeUpdate) SetVendor(s string) *ServerNetworkCardTypeUpdate {
	snctu.mutation.SetVendor(s)
	return snctu
}

// SetModel sets the "model" field.
func (snctu *ServerNetworkCardTypeUpdate) SetModel(s string) *ServerNetworkCardTypeUpdate {
	snctu.mutation.SetModel(s)
	return snctu
}

// SetPortCount sets the "port_count" field.
func (snctu *ServerNetworkCardTypeUpdate) SetPortCount(i int) *ServerNetworkCardTypeUpdate {
	snctu.mutation.ResetPortCount()
	snctu.mutation.SetPortCount(i)
	return snctu
}

// AddPortCount adds i to the "port_count" field.
func (snctu *ServerNetworkCardTypeUpdate) AddPortCount(i int) *ServerNetworkCardTypeUpdate {
	snctu.mutation.AddPortCount(i)
	return snctu
}

// AddNetworkCardIDs adds the "network_card" edge to the ServerNetworkCard entity by IDs.
func (snctu *ServerNetworkCardTypeUpdate) AddNetworkCardIDs(ids ...gidx.PrefixedID) *ServerNetworkCardTypeUpdate {
	snctu.mutation.AddNetworkCardIDs(ids...)
	return snctu
}

// AddNetworkCard adds the "network_card" edges to the ServerNetworkCard entity.
func (snctu *ServerNetworkCardTypeUpdate) AddNetworkCard(s ...*ServerNetworkCard) *ServerNetworkCardTypeUpdate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snctu.AddNetworkCardIDs(ids...)
}

// Mutation returns the ServerNetworkCardTypeMutation object of the builder.
func (snctu *ServerNetworkCardTypeUpdate) Mutation() *ServerNetworkCardTypeMutation {
	return snctu.mutation
}

// ClearNetworkCard clears all "network_card" edges to the ServerNetworkCard entity.
func (snctu *ServerNetworkCardTypeUpdate) ClearNetworkCard() *ServerNetworkCardTypeUpdate {
	snctu.mutation.ClearNetworkCard()
	return snctu
}

// RemoveNetworkCardIDs removes the "network_card" edge to ServerNetworkCard entities by IDs.
func (snctu *ServerNetworkCardTypeUpdate) RemoveNetworkCardIDs(ids ...gidx.PrefixedID) *ServerNetworkCardTypeUpdate {
	snctu.mutation.RemoveNetworkCardIDs(ids...)
	return snctu
}

// RemoveNetworkCard removes "network_card" edges to ServerNetworkCard entities.
func (snctu *ServerNetworkCardTypeUpdate) RemoveNetworkCard(s ...*ServerNetworkCard) *ServerNetworkCardTypeUpdate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snctu.RemoveNetworkCardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (snctu *ServerNetworkCardTypeUpdate) Save(ctx context.Context) (int, error) {
	snctu.defaults()
	return withHooks(ctx, snctu.sqlSave, snctu.mutation, snctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (snctu *ServerNetworkCardTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := snctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (snctu *ServerNetworkCardTypeUpdate) Exec(ctx context.Context) error {
	_, err := snctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snctu *ServerNetworkCardTypeUpdate) ExecX(ctx context.Context) {
	if err := snctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (snctu *ServerNetworkCardTypeUpdate) defaults() {
	if _, ok := snctu.mutation.UpdatedAt(); !ok {
		v := servernetworkcardtype.UpdateDefaultUpdatedAt()
		snctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (snctu *ServerNetworkCardTypeUpdate) check() error {
	if v, ok := snctu.mutation.Vendor(); ok {
		if err := servernetworkcardtype.VendorValidator(v); err != nil {
			return &ValidationError{Name: "vendor", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkCardType.vendor": %w`, err)}
		}
	}
	if v, ok := snctu.mutation.Model(); ok {
		if err := servernetworkcardtype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkCardType.model": %w`, err)}
		}
	}
	if v, ok := snctu.mutation.PortCount(); ok {
		if err := servernetworkcardtype.PortCountValidator(v); err != nil {
			return &ValidationError{Name: "port_count", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkCardType.port_count": %w`, err)}
		}
	}
	return nil
}

func (snctu *ServerNetworkCardTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := snctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(servernetworkcardtype.Table, servernetworkcardtype.Columns, sqlgraph.NewFieldSpec(servernetworkcardtype.FieldID, field.TypeString))
	if ps := snctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := snctu.mutation.UpdatedAt(); ok {
		_spec.SetField(servernetworkcardtype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := snctu.mutation.Vendor(); ok {
		_spec.SetField(servernetworkcardtype.FieldVendor, field.TypeString, value)
	}
	if value, ok := snctu.mutation.Model(); ok {
		_spec.SetField(servernetworkcardtype.FieldModel, field.TypeString, value)
	}
	if value, ok := snctu.mutation.PortCount(); ok {
		_spec.SetField(servernetworkcardtype.FieldPortCount, field.TypeInt, value)
	}
	if value, ok := snctu.mutation.AddedPortCount(); ok {
		_spec.AddField(servernetworkcardtype.FieldPortCount, field.TypeInt, value)
	}
	if snctu.mutation.NetworkCardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   servernetworkcardtype.NetworkCardTable,
			Columns: []string{servernetworkcardtype.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snctu.mutation.RemovedNetworkCardIDs(); len(nodes) > 0 && !snctu.mutation.NetworkCardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   servernetworkcardtype.NetworkCardTable,
			Columns: []string{servernetworkcardtype.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snctu.mutation.NetworkCardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   servernetworkcardtype.NetworkCardTable,
			Columns: []string{servernetworkcardtype.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, snctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servernetworkcardtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	snctu.mutation.done = true
	return n, nil
}

// ServerNetworkCardTypeUpdateOne is the builder for updating a single ServerNetworkCardType entity.
type ServerNetworkCardTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServerNetworkCardTypeMutation
}

// SetVendor sets the "vendor" field.
func (snctuo *ServerNetworkCardTypeUpdateOne) SetVendor(s string) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.SetVendor(s)
	return snctuo
}

// SetModel sets the "model" field.
func (snctuo *ServerNetworkCardTypeUpdateOne) SetModel(s string) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.SetModel(s)
	return snctuo
}

// SetPortCount sets the "port_count" field.
func (snctuo *ServerNetworkCardTypeUpdateOne) SetPortCount(i int) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.ResetPortCount()
	snctuo.mutation.SetPortCount(i)
	return snctuo
}

// AddPortCount adds i to the "port_count" field.
func (snctuo *ServerNetworkCardTypeUpdateOne) AddPortCount(i int) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.AddPortCount(i)
	return snctuo
}

// AddNetworkCardIDs adds the "network_card" edge to the ServerNetworkCard entity by IDs.
func (snctuo *ServerNetworkCardTypeUpdateOne) AddNetworkCardIDs(ids ...gidx.PrefixedID) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.AddNetworkCardIDs(ids...)
	return snctuo
}

// AddNetworkCard adds the "network_card" edges to the ServerNetworkCard entity.
func (snctuo *ServerNetworkCardTypeUpdateOne) AddNetworkCard(s ...*ServerNetworkCard) *ServerNetworkCardTypeUpdateOne {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snctuo.AddNetworkCardIDs(ids...)
}

// Mutation returns the ServerNetworkCardTypeMutation object of the builder.
func (snctuo *ServerNetworkCardTypeUpdateOne) Mutation() *ServerNetworkCardTypeMutation {
	return snctuo.mutation
}

// ClearNetworkCard clears all "network_card" edges to the ServerNetworkCard entity.
func (snctuo *ServerNetworkCardTypeUpdateOne) ClearNetworkCard() *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.ClearNetworkCard()
	return snctuo
}

// RemoveNetworkCardIDs removes the "network_card" edge to ServerNetworkCard entities by IDs.
func (snctuo *ServerNetworkCardTypeUpdateOne) RemoveNetworkCardIDs(ids ...gidx.PrefixedID) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.RemoveNetworkCardIDs(ids...)
	return snctuo
}

// RemoveNetworkCard removes "network_card" edges to ServerNetworkCard entities.
func (snctuo *ServerNetworkCardTypeUpdateOne) RemoveNetworkCard(s ...*ServerNetworkCard) *ServerNetworkCardTypeUpdateOne {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return snctuo.RemoveNetworkCardIDs(ids...)
}

// Where appends a list predicates to the ServerNetworkCardTypeUpdate builder.
func (snctuo *ServerNetworkCardTypeUpdateOne) Where(ps ...predicate.ServerNetworkCardType) *ServerNetworkCardTypeUpdateOne {
	snctuo.mutation.Where(ps...)
	return snctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (snctuo *ServerNetworkCardTypeUpdateOne) Select(field string, fields ...string) *ServerNetworkCardTypeUpdateOne {
	snctuo.fields = append([]string{field}, fields...)
	return snctuo
}

// Save executes the query and returns the updated ServerNetworkCardType entity.
func (snctuo *ServerNetworkCardTypeUpdateOne) Save(ctx context.Context) (*ServerNetworkCardType, error) {
	snctuo.defaults()
	return withHooks(ctx, snctuo.sqlSave, snctuo.mutation, snctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (snctuo *ServerNetworkCardTypeUpdateOne) SaveX(ctx context.Context) *ServerNetworkCardType {
	node, err := snctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (snctuo *ServerNetworkCardTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := snctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (snctuo *ServerNetworkCardTypeUpdateOne) ExecX(ctx context.Context) {
	if err := snctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (snctuo *ServerNetworkCardTypeUpdateOne) defaults() {
	if _, ok := snctuo.mutation.UpdatedAt(); !ok {
		v := servernetworkcardtype.UpdateDefaultUpdatedAt()
		snctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (snctuo *ServerNetworkCardTypeUpdateOne) check() error {
	if v, ok := snctuo.mutation.Vendor(); ok {
		if err := servernetworkcardtype.VendorValidator(v); err != nil {
			return &ValidationError{Name: "vendor", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkCardType.vendor": %w`, err)}
		}
	}
	if v, ok := snctuo.mutation.Model(); ok {
		if err := servernetworkcardtype.ModelValidator(v); err != nil {
			return &ValidationError{Name: "model", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkCardType.model": %w`, err)}
		}
	}
	if v, ok := snctuo.mutation.PortCount(); ok {
		if err := servernetworkcardtype.PortCountValidator(v); err != nil {
			return &ValidationError{Name: "port_count", err: fmt.Errorf(`generated: validator failed for field "ServerNetworkCardType.port_count": %w`, err)}
		}
	}
	return nil
}

func (snctuo *ServerNetworkCardTypeUpdateOne) sqlSave(ctx context.Context) (_node *ServerNetworkCardType, err error) {
	if err := snctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(servernetworkcardtype.Table, servernetworkcardtype.Columns, sqlgraph.NewFieldSpec(servernetworkcardtype.FieldID, field.TypeString))
	id, ok := snctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "ServerNetworkCardType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := snctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servernetworkcardtype.FieldID)
		for _, f := range fields {
			if !servernetworkcardtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != servernetworkcardtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := snctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := snctuo.mutation.UpdatedAt(); ok {
		_spec.SetField(servernetworkcardtype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := snctuo.mutation.Vendor(); ok {
		_spec.SetField(servernetworkcardtype.FieldVendor, field.TypeString, value)
	}
	if value, ok := snctuo.mutation.Model(); ok {
		_spec.SetField(servernetworkcardtype.FieldModel, field.TypeString, value)
	}
	if value, ok := snctuo.mutation.PortCount(); ok {
		_spec.SetField(servernetworkcardtype.FieldPortCount, field.TypeInt, value)
	}
	if value, ok := snctuo.mutation.AddedPortCount(); ok {
		_spec.AddField(servernetworkcardtype.FieldPortCount, field.TypeInt, value)
	}
	if snctuo.mutation.NetworkCardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   servernetworkcardtype.NetworkCardTable,
			Columns: []string{servernetworkcardtype.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snctuo.mutation.RemovedNetworkCardIDs(); len(nodes) > 0 && !snctuo.mutation.NetworkCardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   servernetworkcardtype.NetworkCardTable,
			Columns: []string{servernetworkcardtype.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := snctuo.mutation.NetworkCardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   servernetworkcardtype.NetworkCardTable,
			Columns: []string{servernetworkcardtype.NetworkCardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ServerNetworkCardType{config: snctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, snctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servernetworkcardtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	snctuo.mutation.done = true
	return _node, nil
}
