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
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/server-api/internal/ent/generated/servermemory"
	"go.infratographer.com/server-api/internal/ent/generated/servermemorytype"
	"go.infratographer.com/x/gidx"
)

// ServerMemoryTypeQuery is the builder for querying ServerMemoryType entities.
type ServerMemoryTypeQuery struct {
	config
	ctx             *QueryContext
	order           []servermemorytype.OrderOption
	inters          []Interceptor
	predicates      []predicate.ServerMemoryType
	withMemory      *ServerMemoryQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*ServerMemoryType) error
	withNamedMemory map[string]*ServerMemoryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServerMemoryTypeQuery builder.
func (smtq *ServerMemoryTypeQuery) Where(ps ...predicate.ServerMemoryType) *ServerMemoryTypeQuery {
	smtq.predicates = append(smtq.predicates, ps...)
	return smtq
}

// Limit the number of records to be returned by this query.
func (smtq *ServerMemoryTypeQuery) Limit(limit int) *ServerMemoryTypeQuery {
	smtq.ctx.Limit = &limit
	return smtq
}

// Offset to start from.
func (smtq *ServerMemoryTypeQuery) Offset(offset int) *ServerMemoryTypeQuery {
	smtq.ctx.Offset = &offset
	return smtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smtq *ServerMemoryTypeQuery) Unique(unique bool) *ServerMemoryTypeQuery {
	smtq.ctx.Unique = &unique
	return smtq
}

// Order specifies how the records should be ordered.
func (smtq *ServerMemoryTypeQuery) Order(o ...servermemorytype.OrderOption) *ServerMemoryTypeQuery {
	smtq.order = append(smtq.order, o...)
	return smtq
}

// QueryMemory chains the current query on the "memory" edge.
func (smtq *ServerMemoryTypeQuery) QueryMemory() *ServerMemoryQuery {
	query := (&ServerMemoryClient{config: smtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servermemorytype.Table, servermemorytype.FieldID, selector),
			sqlgraph.To(servermemory.Table, servermemory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, servermemorytype.MemoryTable, servermemorytype.MemoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(smtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServerMemoryType entity from the query.
// Returns a *NotFoundError when no ServerMemoryType was found.
func (smtq *ServerMemoryTypeQuery) First(ctx context.Context) (*ServerMemoryType, error) {
	nodes, err := smtq.Limit(1).All(setContextOp(ctx, smtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{servermemorytype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) FirstX(ctx context.Context) *ServerMemoryType {
	node, err := smtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServerMemoryType ID from the query.
// Returns a *NotFoundError when no ServerMemoryType ID was found.
func (smtq *ServerMemoryTypeQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = smtq.Limit(1).IDs(setContextOp(ctx, smtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{servermemorytype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := smtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServerMemoryType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServerMemoryType entity is found.
// Returns a *NotFoundError when no ServerMemoryType entities are found.
func (smtq *ServerMemoryTypeQuery) Only(ctx context.Context) (*ServerMemoryType, error) {
	nodes, err := smtq.Limit(2).All(setContextOp(ctx, smtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{servermemorytype.Label}
	default:
		return nil, &NotSingularError{servermemorytype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) OnlyX(ctx context.Context) *ServerMemoryType {
	node, err := smtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServerMemoryType ID in the query.
// Returns a *NotSingularError when more than one ServerMemoryType ID is found.
// Returns a *NotFoundError when no entities are found.
func (smtq *ServerMemoryTypeQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = smtq.Limit(2).IDs(setContextOp(ctx, smtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{servermemorytype.Label}
	default:
		err = &NotSingularError{servermemorytype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := smtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServerMemoryTypes.
func (smtq *ServerMemoryTypeQuery) All(ctx context.Context) ([]*ServerMemoryType, error) {
	ctx = setContextOp(ctx, smtq.ctx, "All")
	if err := smtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ServerMemoryType, *ServerMemoryTypeQuery]()
	return withInterceptors[[]*ServerMemoryType](ctx, smtq, qr, smtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) AllX(ctx context.Context) []*ServerMemoryType {
	nodes, err := smtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServerMemoryType IDs.
func (smtq *ServerMemoryTypeQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if smtq.ctx.Unique == nil && smtq.path != nil {
		smtq.Unique(true)
	}
	ctx = setContextOp(ctx, smtq.ctx, "IDs")
	if err = smtq.Select(servermemorytype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := smtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smtq *ServerMemoryTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, smtq.ctx, "Count")
	if err := smtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, smtq, querierCount[*ServerMemoryTypeQuery](), smtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) CountX(ctx context.Context) int {
	count, err := smtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smtq *ServerMemoryTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, smtq.ctx, "Exist")
	switch _, err := smtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (smtq *ServerMemoryTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := smtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServerMemoryTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smtq *ServerMemoryTypeQuery) Clone() *ServerMemoryTypeQuery {
	if smtq == nil {
		return nil
	}
	return &ServerMemoryTypeQuery{
		config:     smtq.config,
		ctx:        smtq.ctx.Clone(),
		order:      append([]servermemorytype.OrderOption{}, smtq.order...),
		inters:     append([]Interceptor{}, smtq.inters...),
		predicates: append([]predicate.ServerMemoryType{}, smtq.predicates...),
		withMemory: smtq.withMemory.Clone(),
		// clone intermediate query.
		sql:  smtq.sql.Clone(),
		path: smtq.path,
	}
}

// WithMemory tells the query-builder to eager-load the nodes that are connected to
// the "memory" edge. The optional arguments are used to configure the query builder of the edge.
func (smtq *ServerMemoryTypeQuery) WithMemory(opts ...func(*ServerMemoryQuery)) *ServerMemoryTypeQuery {
	query := (&ServerMemoryClient{config: smtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	smtq.withMemory = query
	return smtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServerMemoryType.Query().
//		GroupBy(servermemorytype.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (smtq *ServerMemoryTypeQuery) GroupBy(field string, fields ...string) *ServerMemoryTypeGroupBy {
	smtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ServerMemoryTypeGroupBy{build: smtq}
	grbuild.flds = &smtq.ctx.Fields
	grbuild.label = servermemorytype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ServerMemoryType.Query().
//		Select(servermemorytype.FieldCreatedAt).
//		Scan(ctx, &v)
func (smtq *ServerMemoryTypeQuery) Select(fields ...string) *ServerMemoryTypeSelect {
	smtq.ctx.Fields = append(smtq.ctx.Fields, fields...)
	sbuild := &ServerMemoryTypeSelect{ServerMemoryTypeQuery: smtq}
	sbuild.label = servermemorytype.Label
	sbuild.flds, sbuild.scan = &smtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ServerMemoryTypeSelect configured with the given aggregations.
func (smtq *ServerMemoryTypeQuery) Aggregate(fns ...AggregateFunc) *ServerMemoryTypeSelect {
	return smtq.Select().Aggregate(fns...)
}

func (smtq *ServerMemoryTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range smtq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, smtq); err != nil {
				return err
			}
		}
	}
	for _, f := range smtq.ctx.Fields {
		if !servermemorytype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if smtq.path != nil {
		prev, err := smtq.path(ctx)
		if err != nil {
			return err
		}
		smtq.sql = prev
	}
	return nil
}

func (smtq *ServerMemoryTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServerMemoryType, error) {
	var (
		nodes       = []*ServerMemoryType{}
		_spec       = smtq.querySpec()
		loadedTypes = [1]bool{
			smtq.withMemory != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ServerMemoryType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ServerMemoryType{config: smtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(smtq.modifiers) > 0 {
		_spec.Modifiers = smtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := smtq.withMemory; query != nil {
		if err := smtq.loadMemory(ctx, query, nodes,
			func(n *ServerMemoryType) { n.Edges.Memory = []*ServerMemory{} },
			func(n *ServerMemoryType, e *ServerMemory) { n.Edges.Memory = append(n.Edges.Memory, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range smtq.withNamedMemory {
		if err := smtq.loadMemory(ctx, query, nodes,
			func(n *ServerMemoryType) { n.appendNamedMemory(name) },
			func(n *ServerMemoryType, e *ServerMemory) { n.appendNamedMemory(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range smtq.loadTotal {
		if err := smtq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (smtq *ServerMemoryTypeQuery) loadMemory(ctx context.Context, query *ServerMemoryQuery, nodes []*ServerMemoryType, init func(*ServerMemoryType), assign func(*ServerMemoryType, *ServerMemory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID]*ServerMemoryType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(servermemory.FieldServerMemoryTypeID)
	}
	query.Where(predicate.ServerMemory(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(servermemorytype.MemoryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ServerMemoryTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "server_memory_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (smtq *ServerMemoryTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smtq.querySpec()
	if len(smtq.modifiers) > 0 {
		_spec.Modifiers = smtq.modifiers
	}
	_spec.Node.Columns = smtq.ctx.Fields
	if len(smtq.ctx.Fields) > 0 {
		_spec.Unique = smtq.ctx.Unique != nil && *smtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, smtq.driver, _spec)
}

func (smtq *ServerMemoryTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(servermemorytype.Table, servermemorytype.Columns, sqlgraph.NewFieldSpec(servermemorytype.FieldID, field.TypeString))
	_spec.From = smtq.sql
	if unique := smtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if smtq.path != nil {
		_spec.Unique = true
	}
	if fields := smtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servermemorytype.FieldID)
		for i := range fields {
			if fields[i] != servermemorytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smtq *ServerMemoryTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smtq.driver.Dialect())
	t1 := builder.Table(servermemorytype.Table)
	columns := smtq.ctx.Fields
	if len(columns) == 0 {
		columns = servermemorytype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smtq.sql != nil {
		selector = smtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smtq.ctx.Unique != nil && *smtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range smtq.predicates {
		p(selector)
	}
	for _, p := range smtq.order {
		p(selector)
	}
	if offset := smtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedMemory tells the query-builder to eager-load the nodes that are connected to the "memory"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (smtq *ServerMemoryTypeQuery) WithNamedMemory(name string, opts ...func(*ServerMemoryQuery)) *ServerMemoryTypeQuery {
	query := (&ServerMemoryClient{config: smtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if smtq.withNamedMemory == nil {
		smtq.withNamedMemory = make(map[string]*ServerMemoryQuery)
	}
	smtq.withNamedMemory[name] = query
	return smtq
}

// ServerMemoryTypeGroupBy is the group-by builder for ServerMemoryType entities.
type ServerMemoryTypeGroupBy struct {
	selector
	build *ServerMemoryTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smtgb *ServerMemoryTypeGroupBy) Aggregate(fns ...AggregateFunc) *ServerMemoryTypeGroupBy {
	smtgb.fns = append(smtgb.fns, fns...)
	return smtgb
}

// Scan applies the selector query and scans the result into the given value.
func (smtgb *ServerMemoryTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smtgb.build.ctx, "GroupBy")
	if err := smtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerMemoryTypeQuery, *ServerMemoryTypeGroupBy](ctx, smtgb.build, smtgb, smtgb.build.inters, v)
}

func (smtgb *ServerMemoryTypeGroupBy) sqlScan(ctx context.Context, root *ServerMemoryTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(smtgb.fns))
	for _, fn := range smtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*smtgb.flds)+len(smtgb.fns))
		for _, f := range *smtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*smtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ServerMemoryTypeSelect is the builder for selecting fields of ServerMemoryType entities.
type ServerMemoryTypeSelect struct {
	*ServerMemoryTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (smts *ServerMemoryTypeSelect) Aggregate(fns ...AggregateFunc) *ServerMemoryTypeSelect {
	smts.fns = append(smts.fns, fns...)
	return smts
}

// Scan applies the selector query and scans the result into the given value.
func (smts *ServerMemoryTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smts.ctx, "Select")
	if err := smts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerMemoryTypeQuery, *ServerMemoryTypeSelect](ctx, smts.ServerMemoryTypeQuery, smts, smts.inters, v)
}

func (smts *ServerMemoryTypeSelect) sqlScan(ctx context.Context, root *ServerMemoryTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(smts.fns))
	for _, fn := range smts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*smts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
