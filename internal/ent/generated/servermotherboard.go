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
	"go.infratographer.com/server-api/internal/ent/generated/servermotherboard"
	"go.infratographer.com/server-api/internal/ent/generated/servermotherboardtype"
	"go.infratographer.com/x/gidx"
)

// Representation of a server motherboard. ServerMotherboard describes the available motherboard for a server.
type ServerMotherboard struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server motherboard.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The serial of the server motherboard
	Serial string `json:"serial,omitempty"`
	// The ID for the server motherboard type of this server motherboard.
	ServerMotherboardTypeID gidx.PrefixedID `json:"server_motherboard_type_id,omitempty"`
	// The ID for the server of this server motherboard.
	ServerID gidx.PrefixedID `json:"server_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServerMotherboardQuery when eager-loading is set.
	Edges        ServerMotherboardEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ServerMotherboardEdges holds the relations/edges for other nodes in the graph.
type ServerMotherboardEdges struct {
	// Server holds the value of the server edge.
	Server *Server `json:"server,omitempty"`
	// ServerMotherboardType holds the value of the server_motherboard_type edge.
	ServerMotherboardType *ServerMotherboardType `json:"server_motherboard_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// ServerOrErr returns the Server value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerMotherboardEdges) ServerOrErr() (*Server, error) {
	if e.loadedTypes[0] {
		if e.Server == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: server.Label}
		}
		return e.Server, nil
	}
	return nil, &NotLoadedError{edge: "server"}
}

// ServerMotherboardTypeOrErr returns the ServerMotherboardType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerMotherboardEdges) ServerMotherboardTypeOrErr() (*ServerMotherboardType, error) {
	if e.loadedTypes[1] {
		if e.ServerMotherboardType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: servermotherboardtype.Label}
		}
		return e.ServerMotherboardType, nil
	}
	return nil, &NotLoadedError{edge: "server_motherboard_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServerMotherboard) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case servermotherboard.FieldID, servermotherboard.FieldServerMotherboardTypeID, servermotherboard.FieldServerID:
			values[i] = new(gidx.PrefixedID)
		case servermotherboard.FieldSerial:
			values[i] = new(sql.NullString)
		case servermotherboard.FieldCreatedAt, servermotherboard.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServerMotherboard fields.
func (sm *ServerMotherboard) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case servermotherboard.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sm.ID = *value
			}
		case servermotherboard.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = value.Time
			}
		case servermotherboard.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = value.Time
			}
		case servermotherboard.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				sm.Serial = value.String
			}
		case servermotherboard.FieldServerMotherboardTypeID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_motherboard_type_id", values[i])
			} else if value != nil {
				sm.ServerMotherboardTypeID = *value
			}
		case servermotherboard.FieldServerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_id", values[i])
			} else if value != nil {
				sm.ServerID = *value
			}
		default:
			sm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServerMotherboard.
// This includes values selected through modifiers, order, etc.
func (sm *ServerMotherboard) Value(name string) (ent.Value, error) {
	return sm.selectValues.Get(name)
}

// QueryServer queries the "server" edge of the ServerMotherboard entity.
func (sm *ServerMotherboard) QueryServer() *ServerQuery {
	return NewServerMotherboardClient(sm.config).QueryServer(sm)
}

// QueryServerMotherboardType queries the "server_motherboard_type" edge of the ServerMotherboard entity.
func (sm *ServerMotherboard) QueryServerMotherboardType() *ServerMotherboardTypeQuery {
	return NewServerMotherboardClient(sm.config).QueryServerMotherboardType(sm)
}

// Update returns a builder for updating this ServerMotherboard.
// Note that you need to call ServerMotherboard.Unwrap() before calling this method if this ServerMotherboard
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *ServerMotherboard) Update() *ServerMotherboardUpdateOne {
	return NewServerMotherboardClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the ServerMotherboard entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *ServerMotherboard) Unwrap() *ServerMotherboard {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("generated: ServerMotherboard is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *ServerMotherboard) String() string {
	var builder strings.Builder
	builder.WriteString("ServerMotherboard(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(sm.Serial)
	builder.WriteString(", ")
	builder.WriteString("server_motherboard_type_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.ServerMotherboardTypeID))
	builder.WriteString(", ")
	builder.WriteString("server_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.ServerID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (sm ServerMotherboard) IsEntity() {}

// ServerMotherboards is a parsable slice of ServerMotherboard.
type ServerMotherboards []*ServerMotherboard
