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
	"go.infratographer.com/server-api/internal/ent/generated/serverchassis"
	"go.infratographer.com/x/gidx"
)

// Representation of a server chassis. ServerChassis describes the available chassis types for a server.
type ServerChassis struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server chassis.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The ID for the server chassis type of this server chassis.
	ServerChassisTypeID gidx.PrefixedID `json:"server_chassis_type_id,omitempty"`
	// The ID for the parent of this chassis.
	ParentChassisID gidx.PrefixedID `json:"parent_chassis_id,omitempty"`
	// The ID for the server of this server chassis.
	ServerID gidx.PrefixedID `json:"server_id,omitempty"`
	// The serial number of the server chassis.
	Serial       string `json:"serial,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServerChassis) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case serverchassis.FieldID, serverchassis.FieldServerChassisTypeID, serverchassis.FieldParentChassisID, serverchassis.FieldServerID:
			values[i] = new(gidx.PrefixedID)
		case serverchassis.FieldSerial:
			values[i] = new(sql.NullString)
		case serverchassis.FieldCreatedAt, serverchassis.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServerChassis fields.
func (sc *ServerChassis) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case serverchassis.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sc.ID = *value
			}
		case serverchassis.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case serverchassis.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case serverchassis.FieldServerChassisTypeID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_chassis_type_id", values[i])
			} else if value != nil {
				sc.ServerChassisTypeID = *value
			}
		case serverchassis.FieldParentChassisID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_chassis_id", values[i])
			} else if value != nil {
				sc.ParentChassisID = *value
			}
		case serverchassis.FieldServerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field server_id", values[i])
			} else if value != nil {
				sc.ServerID = *value
			}
		case serverchassis.FieldSerial:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial", values[i])
			} else if value.Valid {
				sc.Serial = value.String
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServerChassis.
// This includes values selected through modifiers, order, etc.
func (sc *ServerChassis) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// Update returns a builder for updating this ServerChassis.
// Note that you need to call ServerChassis.Unwrap() before calling this method if this ServerChassis
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *ServerChassis) Update() *ServerChassisUpdateOne {
	return NewServerChassisClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the ServerChassis entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *ServerChassis) Unwrap() *ServerChassis {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("generated: ServerChassis is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *ServerChassis) String() string {
	var builder strings.Builder
	builder.WriteString("ServerChassis(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("server_chassis_type_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.ServerChassisTypeID))
	builder.WriteString(", ")
	builder.WriteString("parent_chassis_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.ParentChassisID))
	builder.WriteString(", ")
	builder.WriteString("server_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.ServerID))
	builder.WriteString(", ")
	builder.WriteString("serial=")
	builder.WriteString(sc.Serial)
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (sc ServerChassis) IsEntity() {}

// ServerChasses is a parsable slice of ServerChassis.
type ServerChasses []*ServerChassis
