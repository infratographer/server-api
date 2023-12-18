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
	"go.infratographer.com/server-api/internal/ent/generated/servercomponenttype"
	"go.infratographer.com/x/gidx"
)

// Representation of a server component type. ServerComponentType describes the available component types for a server.
type ServerComponentType struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server component type.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the server component type.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServerComponentType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case servercomponenttype.FieldID:
			values[i] = new(gidx.PrefixedID)
		case servercomponenttype.FieldName:
			values[i] = new(sql.NullString)
		case servercomponenttype.FieldCreatedAt, servercomponenttype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServerComponentType fields.
func (sct *ServerComponentType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case servercomponenttype.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sct.ID = *value
			}
		case servercomponenttype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sct.CreatedAt = value.Time
			}
		case servercomponenttype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sct.UpdatedAt = value.Time
			}
		case servercomponenttype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sct.Name = value.String
			}
		default:
			sct.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServerComponentType.
// This includes values selected through modifiers, order, etc.
func (sct *ServerComponentType) Value(name string) (ent.Value, error) {
	return sct.selectValues.Get(name)
}

// Update returns a builder for updating this ServerComponentType.
// Note that you need to call ServerComponentType.Unwrap() before calling this method if this ServerComponentType
// was returned from a transaction, and the transaction was committed or rolled back.
func (sct *ServerComponentType) Update() *ServerComponentTypeUpdateOne {
	return NewServerComponentTypeClient(sct.config).UpdateOne(sct)
}

// Unwrap unwraps the ServerComponentType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sct *ServerComponentType) Unwrap() *ServerComponentType {
	_tx, ok := sct.config.driver.(*txDriver)
	if !ok {
		panic("generated: ServerComponentType is not a transactional entity")
	}
	sct.config.driver = _tx.drv
	return sct
}

// String implements the fmt.Stringer.
func (sct *ServerComponentType) String() string {
	var builder strings.Builder
	builder.WriteString("ServerComponentType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sct.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sct.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sct.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sct.Name)
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (sct ServerComponentType) IsEntity() {}

// ServerComponentTypes is a parsable slice of ServerComponentType.
type ServerComponentTypes []*ServerComponentType
