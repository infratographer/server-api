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
	"go.infratographer.com/x/gidx"
)

// CreateServerProviderInput represents a mutation input for creating serverproviders.
type CreateServerProviderInput struct {
	Name               string
	ResourceProviderID gidx.PrefixedID
}

// Mutate applies the CreateServerProviderInput on the ProviderMutation builder.
func (i *CreateServerProviderInput) Mutate(m *ProviderMutation) {
	m.SetName(i.Name)
	m.SetResourceProviderID(i.ResourceProviderID)
}

// SetInput applies the change-set in the CreateServerProviderInput on the ProviderCreate builder.
func (c *ProviderCreate) SetInput(i CreateServerProviderInput) *ProviderCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerProviderInput represents a mutation input for updating serverproviders.
type UpdateServerProviderInput struct {
	Name *string
}

// Mutate applies the UpdateServerProviderInput on the ProviderMutation builder.
func (i *UpdateServerProviderInput) Mutate(m *ProviderMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateServerProviderInput on the ProviderUpdate builder.
func (c *ProviderUpdate) SetInput(i UpdateServerProviderInput) *ProviderUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerProviderInput on the ProviderUpdateOne builder.
func (c *ProviderUpdateOne) SetInput(i UpdateServerProviderInput) *ProviderUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerInput represents a mutation input for creating servers.
type CreateServerInput struct {
	Name         string
	Description  *string
	OwnerID      gidx.PrefixedID
	LocationID   gidx.PrefixedID
	ProviderID   gidx.PrefixedID
	ServerTypeID gidx.PrefixedID
	ComponentIDs []gidx.PrefixedID
}

// Mutate applies the CreateServerInput on the ServerMutation builder.
func (i *CreateServerInput) Mutate(m *ServerMutation) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetOwnerID(i.OwnerID)
	m.SetLocationID(i.LocationID)
	m.SetProviderID(i.ProviderID)
	m.SetServerTypeID(i.ServerTypeID)
	if v := i.ComponentIDs; len(v) > 0 {
		m.AddComponentIDs(v...)
	}
}

// SetInput applies the change-set in the CreateServerInput on the ServerCreate builder.
func (c *ServerCreate) SetInput(i CreateServerInput) *ServerCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerInput represents a mutation input for updating servers.
type UpdateServerInput struct {
	Name               *string
	ClearDescription   bool
	Description        *string
	ClearComponents    bool
	AddComponentIDs    []gidx.PrefixedID
	RemoveComponentIDs []gidx.PrefixedID
}

// Mutate applies the UpdateServerInput on the ServerMutation builder.
func (i *UpdateServerInput) Mutate(m *ServerMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if i.ClearComponents {
		m.ClearComponents()
	}
	if v := i.AddComponentIDs; len(v) > 0 {
		m.AddComponentIDs(v...)
	}
	if v := i.RemoveComponentIDs; len(v) > 0 {
		m.RemoveComponentIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateServerInput on the ServerUpdate builder.
func (c *ServerUpdate) SetInput(i UpdateServerInput) *ServerUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerInput on the ServerUpdateOne builder.
func (c *ServerUpdateOne) SetInput(i UpdateServerInput) *ServerUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerChassisInput represents a mutation input for creating serverchasses.
type CreateServerChassisInput struct {
	ParentChassisID     gidx.PrefixedID
	Serial              string
	ServerID            gidx.PrefixedID
	ServerChassisTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerChassisInput on the ServerChassisMutation builder.
func (i *CreateServerChassisInput) Mutate(m *ServerChassisMutation) {
	m.SetParentChassisID(i.ParentChassisID)
	m.SetSerial(i.Serial)
	m.SetServerID(i.ServerID)
	m.SetServerChassisTypeID(i.ServerChassisTypeID)
}

// SetInput applies the change-set in the CreateServerChassisInput on the ServerChassisCreate builder.
func (c *ServerChassisCreate) SetInput(i CreateServerChassisInput) *ServerChassisCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerChassisInput represents a mutation input for updating serverchasses.
type UpdateServerChassisInput struct {
	Serial *string
}

// Mutate applies the UpdateServerChassisInput on the ServerChassisMutation builder.
func (i *UpdateServerChassisInput) Mutate(m *ServerChassisMutation) {
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerChassisInput on the ServerChassisUpdate builder.
func (c *ServerChassisUpdate) SetInput(i UpdateServerChassisInput) *ServerChassisUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerChassisInput on the ServerChassisUpdateOne builder.
func (c *ServerChassisUpdateOne) SetInput(i UpdateServerChassisInput) *ServerChassisUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerChassisTypeInput represents a mutation input for creating serverchassistypes.
type CreateServerChassisTypeInput struct {
	Vendor              string
	Model               string
	Height              string
	IsFullDepth         bool
	ParentChassisTypeID gidx.PrefixedID
}

// Mutate applies the CreateServerChassisTypeInput on the ServerChassisTypeMutation builder.
func (i *CreateServerChassisTypeInput) Mutate(m *ServerChassisTypeMutation) {
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetHeight(i.Height)
	m.SetIsFullDepth(i.IsFullDepth)
	m.SetParentChassisTypeID(i.ParentChassisTypeID)
}

// SetInput applies the change-set in the CreateServerChassisTypeInput on the ServerChassisTypeCreate builder.
func (c *ServerChassisTypeCreate) SetInput(i CreateServerChassisTypeInput) *ServerChassisTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerChassisTypeInput represents a mutation input for updating serverchassistypes.
type UpdateServerChassisTypeInput struct {
	Vendor      *string
	Model       *string
	Height      *string
	IsFullDepth *bool
}

// Mutate applies the UpdateServerChassisTypeInput on the ServerChassisTypeMutation builder.
func (i *UpdateServerChassisTypeInput) Mutate(m *ServerChassisTypeMutation) {
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Height; v != nil {
		m.SetHeight(*v)
	}
	if v := i.IsFullDepth; v != nil {
		m.SetIsFullDepth(*v)
	}
}

// SetInput applies the change-set in the UpdateServerChassisTypeInput on the ServerChassisTypeUpdate builder.
func (c *ServerChassisTypeUpdate) SetInput(i UpdateServerChassisTypeInput) *ServerChassisTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerChassisTypeInput on the ServerChassisTypeUpdateOne builder.
func (c *ServerChassisTypeUpdateOne) SetInput(i UpdateServerChassisTypeInput) *ServerChassisTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerComponentInput represents a mutation input for creating servercomponents.
type CreateServerComponentInput struct {
	Name            string
	Vendor          string
	Model           string
	Serial          string
	ComponentTypeID gidx.PrefixedID
	ServerID        gidx.PrefixedID
}

// Mutate applies the CreateServerComponentInput on the ServerComponentMutation builder.
func (i *CreateServerComponentInput) Mutate(m *ServerComponentMutation) {
	m.SetName(i.Name)
	m.SetVendor(i.Vendor)
	m.SetModel(i.Model)
	m.SetSerial(i.Serial)
	m.SetComponentTypeID(i.ComponentTypeID)
	m.SetServerID(i.ServerID)
}

// SetInput applies the change-set in the CreateServerComponentInput on the ServerComponentCreate builder.
func (c *ServerComponentCreate) SetInput(i CreateServerComponentInput) *ServerComponentCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerComponentInput represents a mutation input for updating servercomponents.
type UpdateServerComponentInput struct {
	Name   *string
	Vendor *string
	Model  *string
	Serial *string
}

// Mutate applies the UpdateServerComponentInput on the ServerComponentMutation builder.
func (i *UpdateServerComponentInput) Mutate(m *ServerComponentMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Vendor; v != nil {
		m.SetVendor(*v)
	}
	if v := i.Model; v != nil {
		m.SetModel(*v)
	}
	if v := i.Serial; v != nil {
		m.SetSerial(*v)
	}
}

// SetInput applies the change-set in the UpdateServerComponentInput on the ServerComponentUpdate builder.
func (c *ServerComponentUpdate) SetInput(i UpdateServerComponentInput) *ServerComponentUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerComponentInput on the ServerComponentUpdateOne builder.
func (c *ServerComponentUpdateOne) SetInput(i UpdateServerComponentInput) *ServerComponentUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerComponentTypeInput represents a mutation input for creating servercomponenttypes.
type CreateServerComponentTypeInput struct {
	Name string
}

// Mutate applies the CreateServerComponentTypeInput on the ServerComponentTypeMutation builder.
func (i *CreateServerComponentTypeInput) Mutate(m *ServerComponentTypeMutation) {
	m.SetName(i.Name)
}

// SetInput applies the change-set in the CreateServerComponentTypeInput on the ServerComponentTypeCreate builder.
func (c *ServerComponentTypeCreate) SetInput(i CreateServerComponentTypeInput) *ServerComponentTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerComponentTypeInput represents a mutation input for updating servercomponenttypes.
type UpdateServerComponentTypeInput struct {
	Name *string
}

// Mutate applies the UpdateServerComponentTypeInput on the ServerComponentTypeMutation builder.
func (i *UpdateServerComponentTypeInput) Mutate(m *ServerComponentTypeMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateServerComponentTypeInput on the ServerComponentTypeUpdate builder.
func (c *ServerComponentTypeUpdate) SetInput(i UpdateServerComponentTypeInput) *ServerComponentTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerComponentTypeInput on the ServerComponentTypeUpdateOne builder.
func (c *ServerComponentTypeUpdateOne) SetInput(i UpdateServerComponentTypeInput) *ServerComponentTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateServerTypeInput represents a mutation input for creating servertypes.
type CreateServerTypeInput struct {
	Name    string
	OwnerID gidx.PrefixedID
}

// Mutate applies the CreateServerTypeInput on the ServerTypeMutation builder.
func (i *CreateServerTypeInput) Mutate(m *ServerTypeMutation) {
	m.SetName(i.Name)
	m.SetOwnerID(i.OwnerID)
}

// SetInput applies the change-set in the CreateServerTypeInput on the ServerTypeCreate builder.
func (c *ServerTypeCreate) SetInput(i CreateServerTypeInput) *ServerTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateServerTypeInput represents a mutation input for updating servertypes.
type UpdateServerTypeInput struct {
	Name *string
}

// Mutate applies the UpdateServerTypeInput on the ServerTypeMutation builder.
func (i *UpdateServerTypeInput) Mutate(m *ServerTypeMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
}

// SetInput applies the change-set in the UpdateServerTypeInput on the ServerTypeUpdate builder.
func (c *ServerTypeUpdate) SetInput(i UpdateServerTypeInput) *ServerTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateServerTypeInput on the ServerTypeUpdateOne builder.
func (c *ServerTypeUpdateOne) SetInput(i UpdateServerTypeInput) *ServerTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
