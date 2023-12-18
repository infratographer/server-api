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
	"go.infratographer.com/server-api/internal/ent/generated/serverharddrive"
	"go.infratographer.com/server-api/internal/ent/generated/serverharddrivetype"
	"go.infratographer.com/x/gidx"
)

// Representation of a server hard drive. ServerHardDrive describes the available hard drives for a server.
type ServerHardDrive struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server hard drive type.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The serial for the server hard drive.
	Serial string `json:"serial,omitempty"`
	// The ID for the server of this server hard drive.
	ServerID gidx.PrefixedID `json:"server_id,omitempty"`
	// The ID for the server hard drive type of this server hard drive.
	ServerHardDriveTypeID gidx.PrefixedID `json:"server_hard_drive_type_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServerHardDriveQuery when eager-loading is set.
	Edges        ServerHardDriveEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ServerHardDriveEdges holds the relations/edges for other nodes in the graph.
type ServerHardDriveEdges struct {
	// Server holds the value of the server edge.
	Server *Server `json:"server,omitempty"`
	// ServerHardDriveType holds the value of the server_hard_drive_type edge.
	ServerHardDriveType *ServerHardDriveType `json:"server_hard_drive_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// ServerOrErr returns the Server value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerHardDriveEdges) ServerOrErr() (*Server, error) {
	if e.loadedTypes[0] {
		if e.Server == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: server.Label}
		}
		return e.Server, nil
	}
	return nil, &NotLoadedError{edge: "server"}
}

// ServerHardDriveTypeOrErr returns the ServerHardDriveType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServerHardDriveEdges) ServerHardDriveTypeOrErr() (*ServerHardDriveType, error) {
	if e.loadedTypes[1] {
		if e.ServerHardDriveType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: serverharddrivetype.Label}
		}
		return e.ServerHardDriveType, nil
	}
	return nil, &NotLoadedError{edge: "server_hard_drive_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServerHardDrive) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case serverharddrive.FieldID, serverharddrive.FieldServerID, serverharddrive.FieldServerHardDriveTypeID:
			values[i] = new(gidx.PrefixedID)
		case serverharddrive.FieldSerial:
			values[i] = new(sql.NullString)
		case serverharddrive.FieldCreatedAt, serverharddrive.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServerHardDrive fields.
func (shd *ServerHardDrive) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case serverharddrive.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				shd.ID = *value
			}
		case serverharddrive.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				shd.CreatedAt = value.Time
			}
		case serverharddrive.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				shd.UpdatedAt = value.Time
			}
		case serverharddrive.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				shd.Serial = value.String
			}
		case serverharddrive.FieldServerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_id", values[i])
			} else if value != nil {
				shd.ServerID = *value
			}
		case serverharddrive.FieldServerHardDriveTypeID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_hard_drive_type_id", values[i])
			} else if value != nil {
				shd.ServerHardDriveTypeID = *value
			}
		default:
			shd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServerHardDrive.
// This includes values selected through modifiers, order, etc.
func (shd *ServerHardDrive) Value(name string) (ent.Value, error) {
	return shd.selectValues.Get(name)
}

// QueryServer queries the "server" edge of the ServerHardDrive entity.
func (shd *ServerHardDrive) QueryServer() *ServerQuery {
	return NewServerHardDriveClient(shd.config).QueryServer(shd)
}

// QueryServerHardDriveType queries the "server_hard_drive_type" edge of the ServerHardDrive entity.
func (shd *ServerHardDrive) QueryServerHardDriveType() *ServerHardDriveTypeQuery {
	return NewServerHardDriveClient(shd.config).QueryServerHardDriveType(shd)
}

// Update returns a builder for updating this ServerHardDrive.
// Note that you need to call ServerHardDrive.Unwrap() before calling this method if this ServerHardDrive
// was returned from a transaction, and the transaction was committed or rolled back.
func (shd *ServerHardDrive) Update() *ServerHardDriveUpdateOne {
	return NewServerHardDriveClient(shd.config).UpdateOne(shd)
}

// Unwrap unwraps the ServerHardDrive entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (shd *ServerHardDrive) Unwrap() *ServerHardDrive {
	_tx, ok := shd.config.driver.(*txDriver)
	if !ok {
		panic("generated: ServerHardDrive is not a transactional entity")
	}
	shd.config.driver = _tx.drv
	return shd
}

// String implements the fmt.Stringer.
func (shd *ServerHardDrive) String() string {
	var builder strings.Builder
	builder.WriteString("ServerHardDrive(")
	builder.WriteString(fmt.Sprintf("id=%v, ", shd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(shd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(shd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(shd.Serial)
	builder.WriteString(", ")
	builder.WriteString("server_id=")
	builder.WriteString(fmt.Sprintf("%v", shd.ServerID))
	builder.WriteString(", ")
	builder.WriteString("server_hard_drive_type_id=")
	builder.WriteString(fmt.Sprintf("%v", shd.ServerHardDriveTypeID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (shd ServerHardDrive) IsEntity() {}

// ServerHardDrives is a parsable slice of ServerHardDrive.
type ServerHardDrives []*ServerHardDrive
