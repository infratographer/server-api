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
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/server-api/internal/ent/generated/server"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponent"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponenttype"
	"go.infratographer.com/x/gidx"
)

// Representation of a server component. ServerComponent describes the components for a server.
type ServerComponent struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server component.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the server component.
	Name string `json:"name,omitempty"`
	// The name of the vendor of the server component.
	Vendor string `json:"vendor,omitempty"`
	// The model of the server component.
	Model string `json:"model,omitempty"`
	// The serial number of the server component.
	Serial string `json:"serial,omitempty"`
	// The ID for the server of this server component.
	ServerID gidx.PrefixedID `json:"server_id,omitempty"`
	// The ID for the server component type of this server component.
	ComponentTypeID gidx.PrefixedID `json:"component_type_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServerComponentQuery when eager-loading is set.
	Edges        ServerComponentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ServerComponentEdges holds the relations/edges for other nodes in the graph.
type ServerComponentEdges struct {
	// The server component type for the server component.
	ComponentType *ServerComponentType `json:"component_type,omitempty"`
	// Server holds the value of the server edge.
	Server *Server `json:"server,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// ComponentTypeOrErr returns the ComponentType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerComponentEdges) ComponentTypeOrErr() (*ServerComponentType, error) {
	if e.loadedTypes[0] {
		if e.ComponentType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: servercomponenttype.Label}
		}
		return e.ComponentType, nil
	}
	return nil, &NotLoadedError{edge: "component_type"}
}

// ServerOrErr returns the Server value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerComponentEdges) ServerOrErr() (*Server, error) {
	if e.loadedTypes[1] {
		if e.Server == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: server.Label}
		}
		return e.Server, nil
	}
	return nil, &NotLoadedError{edge: "server"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServerComponent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case servercomponent.FieldID, servercomponent.FieldServerID, servercomponent.FieldComponentTypeID:
			values[i] = new(gidx.PrefixedID)
		case servercomponent.FieldName, servercomponent.FieldVendor, servercomponent.FieldModel, servercomponent.FieldSerial:
			values[i] = new(sql.NullString)
		case servercomponent.FieldCreatedAt, servercomponent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServerComponent fields.
func (sc *ServerComponent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case servercomponent.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sc.ID = *value
			}
		case servercomponent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case servercomponent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case servercomponent.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sc.Name = value.String
			}
		case servercomponent.FieldVendor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vendor", values[i])
			} else if value.Valid {
				sc.Vendor = value.String
			}
		case servercomponent.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				sc.Model = value.String
			}
		case servercomponent.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				sc.Serial = value.String
			}
		case servercomponent.FieldServerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_id", values[i])
			} else if value != nil {
				sc.ServerID = *value
			}
		case servercomponent.FieldComponentTypeID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field component_type_id", values[i])
			} else if value != nil {
				sc.ComponentTypeID = *value
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServerComponent.
// This includes values selected through modifiers, order, etc.
func (sc *ServerComponent) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// QueryComponentType queries the "component_type" edge of the ServerComponent entity.
func (sc *ServerComponent) QueryComponentType() *ServerComponentTypeQuery {
	return NewServerComponentClient(sc.config).QueryComponentType(sc)
}

// QueryServer queries the "server" edge of the ServerComponent entity.
func (sc *ServerComponent) QueryServer() *ServerQuery {
	return NewServerComponentClient(sc.config).QueryServer(sc)
}

// Update returns a builder for updating this ServerComponent.
// Note that you need to call ServerComponent.Unwrap() before calling this method if this ServerComponent
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *ServerComponent) Update() *ServerComponentUpdateOne {
	return NewServerComponentClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the ServerComponent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *ServerComponent) Unwrap() *ServerComponent {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("generated: ServerComponent is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *ServerComponent) String() string {
	var builder strings.Builder
	builder.WriteString("ServerComponent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sc.Name)
	builder.WriteString(", ")
	builder.WriteString("vendor=")
	builder.WriteString(sc.Vendor)
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(sc.Model)
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(sc.Serial)
	builder.WriteString(", ")
	builder.WriteString("server_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.ServerID))
	builder.WriteString(", ")
	builder.WriteString("component_type_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.ComponentTypeID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (sc ServerComponent) IsEntity() {}

// ServerComponents is a parsable slice of ServerComponent.
type ServerComponents []*ServerComponent
