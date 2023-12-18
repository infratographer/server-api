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
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkcard"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkport"
	"go.infratographer.com/x/gidx"
)

// ServerNetworkPortQuery is the builder for querying ServerNetworkPort entities.
type ServerNetworkPortQuery struct {
	config
	ctx             *QueryContext
	order           []servernetworkport.OrderOption
	inters          []Interceptor
	predicates      []predicate.ServerNetworkPort
	withNetworkCard *ServerNetworkCardQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*ServerNetworkPort) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServerNetworkPortQuery builder.
func (snpq *ServerNetworkPortQuery) Where(ps ...predicate.ServerNetworkPort) *ServerNetworkPortQuery {
	snpq.predicates = append(snpq.predicates, ps...)
	return snpq
}

// Limit the number of records to be returned by this query.
func (snpq *ServerNetworkPortQuery) Limit(limit int) *ServerNetworkPortQuery {
	snpq.ctx.Limit = &limit
	return snpq
}

// Offset to start from.
func (snpq *ServerNetworkPortQuery) Offset(offset int) *ServerNetworkPortQuery {
	snpq.ctx.Offset = &offset
	return snpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (snpq *ServerNetworkPortQuery) Unique(unique bool) *ServerNetworkPortQuery {
	snpq.ctx.Unique = &unique
	return snpq
}

// Order specifies how the records should be ordered.
func (snpq *ServerNetworkPortQuery) Order(o ...servernetworkport.OrderOption) *ServerNetworkPortQuery {
	snpq.order = append(snpq.order, o...)
	return snpq
}

// QueryNetworkCard chains the current query on the "network_card" edge.
func (snpq *ServerNetworkPortQuery) QueryNetworkCard() *ServerNetworkCardQuery {
	query := (&ServerNetworkCardClient{config: snpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := snpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := snpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servernetworkport.Table, servernetworkport.FieldID, selector),
			sqlgraph.To(servernetworkcard.Table, servernetworkcard.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, servernetworkport.NetworkCardTable, servernetworkport.NetworkCardColumn),
		)
		fromU = sqlgraph.SetNeighbors(snpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServerNetworkPort entity from the query.
// Returns a *NotFoundError when no ServerNetworkPort was found.
func (snpq *ServerNetworkPortQuery) First(ctx context.Context) (*ServerNetworkPort, error) {
	nodes, err := snpq.Limit(1).All(setContextOp(ctx, snpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{servernetworkport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) FirstX(ctx context.Context) *ServerNetworkPort {
	node, err := snpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServerNetworkPort ID from the query.
// Returns a *NotFoundError when no ServerNetworkPort ID was found.
func (snpq *ServerNetworkPortQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = snpq.Limit(1).IDs(setContextOp(ctx, snpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{servernetworkport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := snpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServerNetworkPort entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServerNetworkPort entity is found.
// Returns a *NotFoundError when no ServerNetworkPort entities are found.
func (snpq *ServerNetworkPortQuery) Only(ctx context.Context) (*ServerNetworkPort, error) {
	nodes, err := snpq.Limit(2).All(setContextOp(ctx, snpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{servernetworkport.Label}
	default:
		return nil, &NotSingularError{servernetworkport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) OnlyX(ctx context.Context) *ServerNetworkPort {
	node, err := snpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServerNetworkPort ID in the query.
// Returns a *NotSingularError when more than one ServerNetworkPort ID is found.
// Returns a *NotFoundError when no entities are found.
func (snpq *ServerNetworkPortQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = snpq.Limit(2).IDs(setContextOp(ctx, snpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{servernetworkport.Label}
	default:
		err = &NotSingularError{servernetworkport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := snpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServerNetworkPorts.
func (snpq *ServerNetworkPortQuery) All(ctx context.Context) ([]*ServerNetworkPort, error) {
	ctx = setContextOp(ctx, snpq.ctx, "All")
	if err := snpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ServerNetworkPort, *ServerNetworkPortQuery]()
	return withInterceptors[[]*ServerNetworkPort](ctx, snpq, qr, snpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) AllX(ctx context.Context) []*ServerNetworkPort {
	nodes, err := snpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServerNetworkPort IDs.
func (snpq *ServerNetworkPortQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if snpq.ctx.Unique == nil && snpq.path != nil {
		snpq.Unique(true)
	}
	ctx = setContextOp(ctx, snpq.ctx, "IDs")
	if err = snpq.Select(servernetworkport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := snpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (snpq *ServerNetworkPortQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, snpq.ctx, "Count")
	if err := snpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, snpq, querierCount[*ServerNetworkPortQuery](), snpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) CountX(ctx context.Context) int {
	count, err := snpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (snpq *ServerNetworkPortQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, snpq.ctx, "Exist")
	switch _, err := snpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (snpq *ServerNetworkPortQuery) ExistX(ctx context.Context) bool {
	exist, err := snpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServerNetworkPortQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (snpq *ServerNetworkPortQuery) Clone() *ServerNetworkPortQuery {
	if snpq == nil {
		return nil
	}
	return &ServerNetworkPortQuery{
		config:          snpq.config,
		ctx:             snpq.ctx.Clone(),
		order:           append([]servernetworkport.OrderOption{}, snpq.order...),
		inters:          append([]Interceptor{}, snpq.inters...),
		predicates:      append([]predicate.ServerNetworkPort{}, snpq.predicates...),
		withNetworkCard: snpq.withNetworkCard.Clone(),
		// clone intermediate query.
		sql:  snpq.sql.Clone(),
		path: snpq.path,
	}
}

// WithNetworkCard tells the query-builder to eager-load the nodes that are connected to
// the "network_card" edge. The optional arguments are used to configure the query builder of the edge.
func (snpq *ServerNetworkPortQuery) WithNetworkCard(opts ...func(*ServerNetworkCardQuery)) *ServerNetworkPortQuery {
	query := (&ServerNetworkCardClient{config: snpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	snpq.withNetworkCard = query
	return snpq
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
//	client.ServerNetworkPort.Query().
//		GroupBy(servernetworkport.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (snpq *ServerNetworkPortQuery) GroupBy(field string, fields ...string) *ServerNetworkPortGroupBy {
	snpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ServerNetworkPortGroupBy{build: snpq}
	grbuild.flds = &snpq.ctx.Fields
	grbuild.label = servernetworkport.Label
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
//	client.ServerNetworkPort.Query().
//		Select(servernetworkport.FieldCreatedAt).
//		Scan(ctx, &v)
func (snpq *ServerNetworkPortQuery) Select(fields ...string) *ServerNetworkPortSelect {
	snpq.ctx.Fields = append(snpq.ctx.Fields, fields...)
	sbuild := &ServerNetworkPortSelect{ServerNetworkPortQuery: snpq}
	sbuild.label = servernetworkport.Label
	sbuild.flds, sbuild.scan = &snpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ServerNetworkPortSelect configured with the given aggregations.
func (snpq *ServerNetworkPortQuery) Aggregate(fns ...AggregateFunc) *ServerNetworkPortSelect {
	return snpq.Select().Aggregate(fns...)
}

func (snpq *ServerNetworkPortQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range snpq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, snpq); err != nil {
				return err
			}
		}
	}
	for _, f := range snpq.ctx.Fields {
		if !servernetworkport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if snpq.path != nil {
		prev, err := snpq.path(ctx)
		if err != nil {
			return err
		}
		snpq.sql = prev
	}
	return nil
}

func (snpq *ServerNetworkPortQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServerNetworkPort, error) {
	var (
		nodes       = []*ServerNetworkPort{}
		_spec       = snpq.querySpec()
		loadedTypes = [1]bool{
			snpq.withNetworkCard != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ServerNetworkPort).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ServerNetworkPort{config: snpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(snpq.modifiers) > 0 {
		_spec.Modifiers = snpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, snpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := snpq.withNetworkCard; query != nil {
		if err := snpq.loadNetworkCard(ctx, query, nodes, nil,
			func(n *ServerNetworkPort, e *ServerNetworkCard) { n.Edges.NetworkCard = e }); err != nil {
			return nil, err
		}
	}
	for i := range snpq.loadTotal {
		if err := snpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (snpq *ServerNetworkPortQuery) loadNetworkCard(ctx context.Context, query *ServerNetworkCardQuery, nodes []*ServerNetworkPort, init func(*ServerNetworkPort), assign func(*ServerNetworkPort, *ServerNetworkCard)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*ServerNetworkPort)
	for i := range nodes {
		fk := nodes[i].ServerNetworkCardID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(servernetworkcard.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "server_network_card_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (snpq *ServerNetworkPortQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := snpq.querySpec()
	if len(snpq.modifiers) > 0 {
		_spec.Modifiers = snpq.modifiers
	}
	_spec.Node.Columns = snpq.ctx.Fields
	if len(snpq.ctx.Fields) > 0 {
		_spec.Unique = snpq.ctx.Unique != nil && *snpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, snpq.driver, _spec)
}

func (snpq *ServerNetworkPortQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(servernetworkport.Table, servernetworkport.Columns, sqlgraph.NewFieldSpec(servernetworkport.FieldID, field.TypeString))
	_spec.From = snpq.sql
	if unique := snpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if snpq.path != nil {
		_spec.Unique = true
	}
	if fields := snpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servernetworkport.FieldID)
		for i := range fields {
			if fields[i] != servernetworkport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if snpq.withNetworkCard != nil {
			_spec.Node.AddColumnOnce(servernetworkport.FieldServerNetworkCardID)
		}
	}
	if ps := snpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := snpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := snpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := snpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (snpq *ServerNetworkPortQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(snpq.driver.Dialect())
	t1 := builder.Table(servernetworkport.Table)
	columns := snpq.ctx.Fields
	if len(columns) == 0 {
		columns = servernetworkport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if snpq.sql != nil {
		selector = snpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if snpq.ctx.Unique != nil && *snpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range snpq.predicates {
		p(selector)
	}
	for _, p := range snpq.order {
		p(selector)
	}
	if offset := snpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := snpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServerNetworkPortGroupBy is the group-by builder for ServerNetworkPort entities.
type ServerNetworkPortGroupBy struct {
	selector
	build *ServerNetworkPortQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (snpgb *ServerNetworkPortGroupBy) Aggregate(fns ...AggregateFunc) *ServerNetworkPortGroupBy {
	snpgb.fns = append(snpgb.fns, fns...)
	return snpgb
}

// Scan applies the selector query and scans the result into the given value.
func (snpgb *ServerNetworkPortGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, snpgb.build.ctx, "GroupBy")
	if err := snpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerNetworkPortQuery, *ServerNetworkPortGroupBy](ctx, snpgb.build, snpgb, snpgb.build.inters, v)
}

func (snpgb *ServerNetworkPortGroupBy) sqlScan(ctx context.Context, root *ServerNetworkPortQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(snpgb.fns))
	for _, fn := range snpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*snpgb.flds)+len(snpgb.fns))
		for _, f := range *snpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*snpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := snpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ServerNetworkPortSelect is the builder for selecting fields of ServerNetworkPort entities.
type ServerNetworkPortSelect struct {
	*ServerNetworkPortQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (snps *ServerNetworkPortSelect) Aggregate(fns ...AggregateFunc) *ServerNetworkPortSelect {
	snps.fns = append(snps.fns, fns...)
	return snps
}

// Scan applies the selector query and scans the result into the given value.
func (snps *ServerNetworkPortSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, snps.ctx, "Select")
	if err := snps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerNetworkPortQuery, *ServerNetworkPortSelect](ctx, snps.ServerNetworkPortQuery, snps, snps.inters, v)
}

func (snps *ServerNetworkPortSelect) sqlScan(ctx context.Context, root *ServerNetworkPortQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(snps.fns))
	for _, fn := range snps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*snps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := snps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
