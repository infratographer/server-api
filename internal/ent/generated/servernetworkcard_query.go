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
	"go.infratographer.com/server-api/internal/ent/generated/server"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkcard"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkcardtype"
	"go.infratographer.com/server-api/internal/ent/generated/servernetworkport"
	"go.infratographer.com/x/gidx"
)

// ServerNetworkCardQuery is the builder for querying ServerNetworkCard entities.
type ServerNetworkCardQuery struct {
	config
	ctx                  *QueryContext
	order                []servernetworkcard.OrderOption
	inters               []Interceptor
	predicates           []predicate.ServerNetworkCard
	withNetworkCardType  *ServerNetworkCardTypeQuery
	withServer           *ServerQuery
	withNetworkPort      *ServerNetworkPortQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*ServerNetworkCard) error
	withNamedNetworkPort map[string]*ServerNetworkPortQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServerNetworkCardQuery builder.
func (sncq *ServerNetworkCardQuery) Where(ps ...predicate.ServerNetworkCard) *ServerNetworkCardQuery {
	sncq.predicates = append(sncq.predicates, ps...)
	return sncq
}

// Limit the number of records to be returned by this query.
func (sncq *ServerNetworkCardQuery) Limit(limit int) *ServerNetworkCardQuery {
	sncq.ctx.Limit = &limit
	return sncq
}

// Offset to start from.
func (sncq *ServerNetworkCardQuery) Offset(offset int) *ServerNetworkCardQuery {
	sncq.ctx.Offset = &offset
	return sncq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sncq *ServerNetworkCardQuery) Unique(unique bool) *ServerNetworkCardQuery {
	sncq.ctx.Unique = &unique
	return sncq
}

// Order specifies how the records should be ordered.
func (sncq *ServerNetworkCardQuery) Order(o ...servernetworkcard.OrderOption) *ServerNetworkCardQuery {
	sncq.order = append(sncq.order, o...)
	return sncq
}

// QueryNetworkCardType chains the current query on the "network_card_type" edge.
func (sncq *ServerNetworkCardQuery) QueryNetworkCardType() *ServerNetworkCardTypeQuery {
	query := (&ServerNetworkCardTypeClient{config: sncq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servernetworkcard.Table, servernetworkcard.FieldID, selector),
			sqlgraph.To(servernetworkcardtype.Table, servernetworkcardtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, servernetworkcard.NetworkCardTypeTable, servernetworkcard.NetworkCardTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(sncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryServer chains the current query on the "server" edge.
func (sncq *ServerNetworkCardQuery) QueryServer() *ServerQuery {
	query := (&ServerClient{config: sncq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servernetworkcard.Table, servernetworkcard.FieldID, selector),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, servernetworkcard.ServerTable, servernetworkcard.ServerColumn),
		)
		fromU = sqlgraph.SetNeighbors(sncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNetworkPort chains the current query on the "network_port" edge.
func (sncq *ServerNetworkCardQuery) QueryNetworkPort() *ServerNetworkPortQuery {
	query := (&ServerNetworkPortClient{config: sncq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servernetworkcard.Table, servernetworkcard.FieldID, selector),
			sqlgraph.To(servernetworkport.Table, servernetworkport.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, servernetworkcard.NetworkPortTable, servernetworkcard.NetworkPortColumn),
		)
		fromU = sqlgraph.SetNeighbors(sncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServerNetworkCard entity from the query.
// Returns a *NotFoundError when no ServerNetworkCard was found.
func (sncq *ServerNetworkCardQuery) First(ctx context.Context) (*ServerNetworkCard, error) {
	nodes, err := sncq.Limit(1).All(setContextOp(ctx, sncq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{servernetworkcard.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) FirstX(ctx context.Context) *ServerNetworkCard {
	node, err := sncq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServerNetworkCard ID from the query.
// Returns a *NotFoundError when no ServerNetworkCard ID was found.
func (sncq *ServerNetworkCardQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = sncq.Limit(1).IDs(setContextOp(ctx, sncq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{servernetworkcard.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := sncq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServerNetworkCard entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServerNetworkCard entity is found.
// Returns a *NotFoundError when no ServerNetworkCard entities are found.
func (sncq *ServerNetworkCardQuery) Only(ctx context.Context) (*ServerNetworkCard, error) {
	nodes, err := sncq.Limit(2).All(setContextOp(ctx, sncq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{servernetworkcard.Label}
	default:
		return nil, &NotSingularError{servernetworkcard.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) OnlyX(ctx context.Context) *ServerNetworkCard {
	node, err := sncq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServerNetworkCard ID in the query.
// Returns a *NotSingularError when more than one ServerNetworkCard ID is found.
// Returns a *NotFoundError when no entities are found.
func (sncq *ServerNetworkCardQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = sncq.Limit(2).IDs(setContextOp(ctx, sncq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{servernetworkcard.Label}
	default:
		err = &NotSingularError{servernetworkcard.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := sncq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServerNetworkCards.
func (sncq *ServerNetworkCardQuery) All(ctx context.Context) ([]*ServerNetworkCard, error) {
	ctx = setContextOp(ctx, sncq.ctx, "All")
	if err := sncq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ServerNetworkCard, *ServerNetworkCardQuery]()
	return withInterceptors[[]*ServerNetworkCard](ctx, sncq, qr, sncq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) AllX(ctx context.Context) []*ServerNetworkCard {
	nodes, err := sncq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServerNetworkCard IDs.
func (sncq *ServerNetworkCardQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if sncq.ctx.Unique == nil && sncq.path != nil {
		sncq.Unique(true)
	}
	ctx = setContextOp(ctx, sncq.ctx, "IDs")
	if err = sncq.Select(servernetworkcard.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := sncq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sncq *ServerNetworkCardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sncq.ctx, "Count")
	if err := sncq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sncq, querierCount[*ServerNetworkCardQuery](), sncq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) CountX(ctx context.Context) int {
	count, err := sncq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sncq *ServerNetworkCardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sncq.ctx, "Exist")
	switch _, err := sncq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sncq *ServerNetworkCardQuery) ExistX(ctx context.Context) bool {
	exist, err := sncq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServerNetworkCardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sncq *ServerNetworkCardQuery) Clone() *ServerNetworkCardQuery {
	if sncq == nil {
		return nil
	}
	return &ServerNetworkCardQuery{
		config:              sncq.config,
		ctx:                 sncq.ctx.Clone(),
		order:               append([]servernetworkcard.OrderOption{}, sncq.order...),
		inters:              append([]Interceptor{}, sncq.inters...),
		predicates:          append([]predicate.ServerNetworkCard{}, sncq.predicates...),
		withNetworkCardType: sncq.withNetworkCardType.Clone(),
		withServer:          sncq.withServer.Clone(),
		withNetworkPort:     sncq.withNetworkPort.Clone(),
		// clone intermediate query.
		sql:  sncq.sql.Clone(),
		path: sncq.path,
	}
}

// WithNetworkCardType tells the query-builder to eager-load the nodes that are connected to
// the "network_card_type" edge. The optional arguments are used to configure the query builder of the edge.
func (sncq *ServerNetworkCardQuery) WithNetworkCardType(opts ...func(*ServerNetworkCardTypeQuery)) *ServerNetworkCardQuery {
	query := (&ServerNetworkCardTypeClient{config: sncq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sncq.withNetworkCardType = query
	return sncq
}

// WithServer tells the query-builder to eager-load the nodes that are connected to
// the "server" edge. The optional arguments are used to configure the query builder of the edge.
func (sncq *ServerNetworkCardQuery) WithServer(opts ...func(*ServerQuery)) *ServerNetworkCardQuery {
	query := (&ServerClient{config: sncq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sncq.withServer = query
	return sncq
}

// WithNetworkPort tells the query-builder to eager-load the nodes that are connected to
// the "network_port" edge. The optional arguments are used to configure the query builder of the edge.
func (sncq *ServerNetworkCardQuery) WithNetworkPort(opts ...func(*ServerNetworkPortQuery)) *ServerNetworkCardQuery {
	query := (&ServerNetworkPortClient{config: sncq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sncq.withNetworkPort = query
	return sncq
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
//	client.ServerNetworkCard.Query().
//		GroupBy(servernetworkcard.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (sncq *ServerNetworkCardQuery) GroupBy(field string, fields ...string) *ServerNetworkCardGroupBy {
	sncq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ServerNetworkCardGroupBy{build: sncq}
	grbuild.flds = &sncq.ctx.Fields
	grbuild.label = servernetworkcard.Label
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
//	client.ServerNetworkCard.Query().
//		Select(servernetworkcard.FieldCreatedAt).
//		Scan(ctx, &v)
func (sncq *ServerNetworkCardQuery) Select(fields ...string) *ServerNetworkCardSelect {
	sncq.ctx.Fields = append(sncq.ctx.Fields, fields...)
	sbuild := &ServerNetworkCardSelect{ServerNetworkCardQuery: sncq}
	sbuild.label = servernetworkcard.Label
	sbuild.flds, sbuild.scan = &sncq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ServerNetworkCardSelect configured with the given aggregations.
func (sncq *ServerNetworkCardQuery) Aggregate(fns ...AggregateFunc) *ServerNetworkCardSelect {
	return sncq.Select().Aggregate(fns...)
}

func (sncq *ServerNetworkCardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sncq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sncq); err != nil {
				return err
			}
		}
	}
	for _, f := range sncq.ctx.Fields {
		if !servernetworkcard.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if sncq.path != nil {
		prev, err := sncq.path(ctx)
		if err != nil {
			return err
		}
		sncq.sql = prev
	}
	return nil
}

func (sncq *ServerNetworkCardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServerNetworkCard, error) {
	var (
		nodes       = []*ServerNetworkCard{}
		_spec       = sncq.querySpec()
		loadedTypes = [3]bool{
			sncq.withNetworkCardType != nil,
			sncq.withServer != nil,
			sncq.withNetworkPort != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ServerNetworkCard).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ServerNetworkCard{config: sncq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sncq.modifiers) > 0 {
		_spec.Modifiers = sncq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sncq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sncq.withNetworkCardType; query != nil {
		if err := sncq.loadNetworkCardType(ctx, query, nodes, nil,
			func(n *ServerNetworkCard, e *ServerNetworkCardType) { n.Edges.NetworkCardType = e }); err != nil {
			return nil, err
		}
	}
	if query := sncq.withServer; query != nil {
		if err := sncq.loadServer(ctx, query, nodes, nil,
			func(n *ServerNetworkCard, e *Server) { n.Edges.Server = e }); err != nil {
			return nil, err
		}
	}
	if query := sncq.withNetworkPort; query != nil {
		if err := sncq.loadNetworkPort(ctx, query, nodes,
			func(n *ServerNetworkCard) { n.Edges.NetworkPort = []*ServerNetworkPort{} },
			func(n *ServerNetworkCard, e *ServerNetworkPort) { n.Edges.NetworkPort = append(n.Edges.NetworkPort, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range sncq.withNamedNetworkPort {
		if err := sncq.loadNetworkPort(ctx, query, nodes,
			func(n *ServerNetworkCard) { n.appendNamedNetworkPort(name) },
			func(n *ServerNetworkCard, e *ServerNetworkPort) { n.appendNamedNetworkPort(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range sncq.loadTotal {
		if err := sncq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sncq *ServerNetworkCardQuery) loadNetworkCardType(ctx context.Context, query *ServerNetworkCardTypeQuery, nodes []*ServerNetworkCard, init func(*ServerNetworkCard), assign func(*ServerNetworkCard, *ServerNetworkCardType)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*ServerNetworkCard)
	for i := range nodes {
		fk := nodes[i].ServerNetworkCardTypeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(servernetworkcardtype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "server_network_card_type_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sncq *ServerNetworkCardQuery) loadServer(ctx context.Context, query *ServerQuery, nodes []*ServerNetworkCard, init func(*ServerNetworkCard), assign func(*ServerNetworkCard, *Server)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*ServerNetworkCard)
	for i := range nodes {
		fk := nodes[i].ServerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(server.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "server_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sncq *ServerNetworkCardQuery) loadNetworkPort(ctx context.Context, query *ServerNetworkPortQuery, nodes []*ServerNetworkCard, init func(*ServerNetworkCard), assign func(*ServerNetworkCard, *ServerNetworkPort)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID]*ServerNetworkCard)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(servernetworkport.FieldNetworkCardID)
	}
	query.Where(predicate.ServerNetworkPort(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(servernetworkcard.NetworkPortColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NetworkCardID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "network_card_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sncq *ServerNetworkCardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sncq.querySpec()
	if len(sncq.modifiers) > 0 {
		_spec.Modifiers = sncq.modifiers
	}
	_spec.Node.Columns = sncq.ctx.Fields
	if len(sncq.ctx.Fields) > 0 {
		_spec.Unique = sncq.ctx.Unique != nil && *sncq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sncq.driver, _spec)
}

func (sncq *ServerNetworkCardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(servernetworkcard.Table, servernetworkcard.Columns, sqlgraph.NewFieldSpec(servernetworkcard.FieldID, field.TypeString))
	_spec.From = sncq.sql
	if unique := sncq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sncq.path != nil {
		_spec.Unique = true
	}
	if fields := sncq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servernetworkcard.FieldID)
		for i := range fields {
			if fields[i] != servernetworkcard.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sncq.withNetworkCardType != nil {
			_spec.Node.AddColumnOnce(servernetworkcard.FieldServerNetworkCardTypeID)
		}
		if sncq.withServer != nil {
			_spec.Node.AddColumnOnce(servernetworkcard.FieldServerID)
		}
	}
	if ps := sncq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sncq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sncq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sncq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sncq *ServerNetworkCardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sncq.driver.Dialect())
	t1 := builder.Table(servernetworkcard.Table)
	columns := sncq.ctx.Fields
	if len(columns) == 0 {
		columns = servernetworkcard.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sncq.sql != nil {
		selector = sncq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sncq.ctx.Unique != nil && *sncq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sncq.predicates {
		p(selector)
	}
	for _, p := range sncq.order {
		p(selector)
	}
	if offset := sncq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sncq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedNetworkPort tells the query-builder to eager-load the nodes that are connected to the "network_port"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (sncq *ServerNetworkCardQuery) WithNamedNetworkPort(name string, opts ...func(*ServerNetworkPortQuery)) *ServerNetworkCardQuery {
	query := (&ServerNetworkPortClient{config: sncq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if sncq.withNamedNetworkPort == nil {
		sncq.withNamedNetworkPort = make(map[string]*ServerNetworkPortQuery)
	}
	sncq.withNamedNetworkPort[name] = query
	return sncq
}

// ServerNetworkCardGroupBy is the group-by builder for ServerNetworkCard entities.
type ServerNetworkCardGroupBy struct {
	selector
	build *ServerNetworkCardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sncgb *ServerNetworkCardGroupBy) Aggregate(fns ...AggregateFunc) *ServerNetworkCardGroupBy {
	sncgb.fns = append(sncgb.fns, fns...)
	return sncgb
}

// Scan applies the selector query and scans the result into the given value.
func (sncgb *ServerNetworkCardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sncgb.build.ctx, "GroupBy")
	if err := sncgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerNetworkCardQuery, *ServerNetworkCardGroupBy](ctx, sncgb.build, sncgb, sncgb.build.inters, v)
}

func (sncgb *ServerNetworkCardGroupBy) sqlScan(ctx context.Context, root *ServerNetworkCardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sncgb.fns))
	for _, fn := range sncgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sncgb.flds)+len(sncgb.fns))
		for _, f := range *sncgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sncgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sncgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ServerNetworkCardSelect is the builder for selecting fields of ServerNetworkCard entities.
type ServerNetworkCardSelect struct {
	*ServerNetworkCardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sncs *ServerNetworkCardSelect) Aggregate(fns ...AggregateFunc) *ServerNetworkCardSelect {
	sncs.fns = append(sncs.fns, fns...)
	return sncs
}

// Scan applies the selector query and scans the result into the given value.
func (sncs *ServerNetworkCardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sncs.ctx, "Select")
	if err := sncs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerNetworkCardQuery, *ServerNetworkCardSelect](ctx, sncs.ServerNetworkCardQuery, sncs, sncs.inters, v)
}

func (sncs *ServerNetworkCardSelect) sqlScan(ctx context.Context, root *ServerNetworkCardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sncs.fns))
	for _, fn := range sncs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sncs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sncs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
