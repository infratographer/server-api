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

	"github.com/99designs/gqlgen/graphql"
)

func (pr *Provider) Servers(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerOrder, where *ServerWhereInput,
) (*ServerConnection, error) {
	opts := []ServerPaginateOption{
		WithServerOrder(orderBy),
		WithServerFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := pr.Edges.totalCount[0][alias]
	if nodes, err := pr.NamedServers(alias); err == nil || hasTotalCount {
		pager, err := newServerPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerConnection{Edges: []*ServerEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return pr.QueryServers().Paginate(ctx, after, first, before, last, opts...)
}

func (s *Server) Provider(ctx context.Context) (*Provider, error) {
	result, err := s.Edges.ProviderOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryProvider().Only(ctx)
	}
	return result, err
}

func (s *Server) ServerType(ctx context.Context) (*ServerType, error) {
	result, err := s.Edges.ServerTypeOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryServerType().Only(ctx)
	}
	return result, err
}

func (s *Server) Components(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerComponentOrder, where *ServerComponentWhereInput,
) (*ServerComponentConnection, error) {
	opts := []ServerComponentPaginateOption{
		WithServerComponentOrder(orderBy),
		WithServerComponentFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := s.Edges.totalCount[2][alias]
	if nodes, err := s.NamedComponents(alias); err == nil || hasTotalCount {
		pager, err := newServerComponentPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerComponentConnection{Edges: []*ServerComponentEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return s.QueryComponents().Paginate(ctx, after, first, before, last, opts...)
}

func (sc *ServerCPU) Server(ctx context.Context) (*Server, error) {
	result, err := sc.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryServer().Only(ctx)
	}
	return result, err
}

func (sc *ServerCPU) ServerCPUType(ctx context.Context) (*ServerCPUType, error) {
	result, err := sc.Edges.ServerCPUTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryServerCPUType().Only(ctx)
	}
	return result, err
}

func (sct *ServerCPUType) CPU(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerCPUOrder, where *ServerCPUWhereInput,
) (*ServerCPUConnection, error) {
	opts := []ServerCPUPaginateOption{
		WithServerCPUOrder(orderBy),
		WithServerCPUFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := sct.Edges.totalCount[0][alias]
	if nodes, err := sct.NamedCPU(alias); err == nil || hasTotalCount {
		pager, err := newServerCPUPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerCPUConnection{Edges: []*ServerCPUEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return sct.QueryCPU().Paginate(ctx, after, first, before, last, opts...)
}

func (sc *ServerChassis) Server(ctx context.Context) (*Server, error) {
	result, err := sc.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryServer().Only(ctx)
	}
	return result, err
}

func (sc *ServerChassis) ServerChassisType(ctx context.Context) (*ServerChassisType, error) {
	result, err := sc.Edges.ServerChassisTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryServerChassisType().Only(ctx)
	}
	return result, err
}

func (sct *ServerChassisType) Chassis(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerChassisOrder, where *ServerChassisWhereInput,
) (*ServerChassisConnection, error) {
	opts := []ServerChassisPaginateOption{
		WithServerChassisOrder(orderBy),
		WithServerChassisFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := sct.Edges.totalCount[0][alias]
	if nodes, err := sct.NamedChassis(alias); err == nil || hasTotalCount {
		pager, err := newServerChassisPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerChassisConnection{Edges: []*ServerChassisEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return sct.QueryChassis().Paginate(ctx, after, first, before, last, opts...)
}

func (sc *ServerComponent) ComponentType(ctx context.Context) (*ServerComponentType, error) {
	result, err := sc.Edges.ComponentTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryComponentType().Only(ctx)
	}
	return result, err
}

func (sc *ServerComponent) Server(ctx context.Context) (*Server, error) {
	result, err := sc.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = sc.QueryServer().Only(ctx)
	}
	return result, err
}

func (shd *ServerHardDrive) Server(ctx context.Context) (*Server, error) {
	result, err := shd.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = shd.QueryServer().Only(ctx)
	}
	return result, err
}

func (shd *ServerHardDrive) ServerHardDriveType(ctx context.Context) (*ServerHardDriveType, error) {
	result, err := shd.Edges.ServerHardDriveTypeOrErr()
	if IsNotLoaded(err) {
		result, err = shd.QueryServerHardDriveType().Only(ctx)
	}
	return result, err
}

func (shdt *ServerHardDriveType) HardDrive(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerHardDriveOrder, where *ServerHardDriveWhereInput,
) (*ServerHardDriveConnection, error) {
	opts := []ServerHardDrivePaginateOption{
		WithServerHardDriveOrder(orderBy),
		WithServerHardDriveFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := shdt.Edges.totalCount[0][alias]
	if nodes, err := shdt.NamedHardDrive(alias); err == nil || hasTotalCount {
		pager, err := newServerHardDrivePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerHardDriveConnection{Edges: []*ServerHardDriveEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return shdt.QueryHardDrive().Paginate(ctx, after, first, before, last, opts...)
}

func (sm *ServerMemory) Server(ctx context.Context) (*Server, error) {
	result, err := sm.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = sm.QueryServer().Only(ctx)
	}
	return result, err
}

func (sm *ServerMemory) ServerMemoryType(ctx context.Context) (*ServerMemoryType, error) {
	result, err := sm.Edges.ServerMemoryTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sm.QueryServerMemoryType().Only(ctx)
	}
	return result, err
}

func (smt *ServerMemoryType) Memory(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerMemoryOrder, where *ServerMemoryWhereInput,
) (*ServerMemoryConnection, error) {
	opts := []ServerMemoryPaginateOption{
		WithServerMemoryOrder(orderBy),
		WithServerMemoryFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := smt.Edges.totalCount[0][alias]
	if nodes, err := smt.NamedMemory(alias); err == nil || hasTotalCount {
		pager, err := newServerMemoryPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerMemoryConnection{Edges: []*ServerMemoryEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return smt.QueryMemory().Paginate(ctx, after, first, before, last, opts...)
}

func (sm *ServerMotherboard) Server(ctx context.Context) (*Server, error) {
	result, err := sm.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = sm.QueryServer().Only(ctx)
	}
	return result, err
}

func (sm *ServerMotherboard) ServerMotherboardType(ctx context.Context) (*ServerMotherboardType, error) {
	result, err := sm.Edges.ServerMotherboardTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sm.QueryServerMotherboardType().Only(ctx)
	}
	return result, err
}

func (smt *ServerMotherboardType) Motherboard(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerMotherboardOrder, where *ServerMotherboardWhereInput,
) (*ServerMotherboardConnection, error) {
	opts := []ServerMotherboardPaginateOption{
		WithServerMotherboardOrder(orderBy),
		WithServerMotherboardFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := smt.Edges.totalCount[0][alias]
	if nodes, err := smt.NamedMotherboard(alias); err == nil || hasTotalCount {
		pager, err := newServerMotherboardPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerMotherboardConnection{Edges: []*ServerMotherboardEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return smt.QueryMotherboard().Paginate(ctx, after, first, before, last, opts...)
}

func (snc *ServerNetworkCard) NetworkCardType(ctx context.Context) (*ServerNetworkCardType, error) {
	result, err := snc.Edges.NetworkCardTypeOrErr()
	if IsNotLoaded(err) {
		result, err = snc.QueryNetworkCardType().Only(ctx)
	}
	return result, err
}

func (snc *ServerNetworkCard) Server(ctx context.Context) (*Server, error) {
	result, err := snc.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = snc.QueryServer().Only(ctx)
	}
	return result, err
}

func (snc *ServerNetworkCard) NetworkPort(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerNetworkPortOrder, where *ServerNetworkPortWhereInput,
) (*ServerNetworkPortConnection, error) {
	opts := []ServerNetworkPortPaginateOption{
		WithServerNetworkPortOrder(orderBy),
		WithServerNetworkPortFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := snc.Edges.totalCount[2][alias]
	if nodes, err := snc.NamedNetworkPort(alias); err == nil || hasTotalCount {
		pager, err := newServerNetworkPortPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerNetworkPortConnection{Edges: []*ServerNetworkPortEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return snc.QueryNetworkPort().Paginate(ctx, after, first, before, last, opts...)
}

func (snct *ServerNetworkCardType) NetworkCard(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerNetworkCardOrder, where *ServerNetworkCardWhereInput,
) (*ServerNetworkCardConnection, error) {
	opts := []ServerNetworkCardPaginateOption{
		WithServerNetworkCardOrder(orderBy),
		WithServerNetworkCardFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := snct.Edges.totalCount[0][alias]
	if nodes, err := snct.NamedNetworkCard(alias); err == nil || hasTotalCount {
		pager, err := newServerNetworkCardPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerNetworkCardConnection{Edges: []*ServerNetworkCardEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return snct.QueryNetworkCard().Paginate(ctx, after, first, before, last, opts...)
}

func (snp *ServerNetworkPort) NetworkCard(ctx context.Context) (*ServerNetworkCard, error) {
	result, err := snp.Edges.NetworkCardOrErr()
	if IsNotLoaded(err) {
		result, err = snp.QueryNetworkCard().Only(ctx)
	}
	return result, err
}

func (sps *ServerPowerSupply) Server(ctx context.Context) (*Server, error) {
	result, err := sps.Edges.ServerOrErr()
	if IsNotLoaded(err) {
		result, err = sps.QueryServer().Only(ctx)
	}
	return result, err
}

func (sps *ServerPowerSupply) ServerPowerSupplyType(ctx context.Context) (*ServerPowerSupplyType, error) {
	result, err := sps.Edges.ServerPowerSupplyTypeOrErr()
	if IsNotLoaded(err) {
		result, err = sps.QueryServerPowerSupplyType().Only(ctx)
	}
	return result, err
}

func (st *ServerType) Servers(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ServerOrder, where *ServerWhereInput,
) (*ServerConnection, error) {
	opts := []ServerPaginateOption{
		WithServerOrder(orderBy),
		WithServerFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := st.Edges.totalCount[0][alias]
	if nodes, err := st.NamedServers(alias); err == nil || hasTotalCount {
		pager, err := newServerPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ServerConnection{Edges: []*ServerEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return st.QueryServers().Paginate(ctx, after, first, before, last, opts...)
}
