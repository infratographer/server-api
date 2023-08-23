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
	"go.infratographer.com/server-api/internal/ent/generated/servercputype"
	"go.infratographer.com/x/gidx"
)

// Representation of a server cpu type. ServerCPUType describes the available cpu types for a server.
type ServerCPUType struct {
	config `json:"-"`
	// ID of the ent.
	// The ID of the server cpu type.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the vendor for the server cpu type.
	Vendor string `json:"vendor,omitempty"`
	// The mode of the server cpu type.
	Model string `json:"model,omitempty"`
	// The clock speed of the server cpu type.
	ClockSpeed string `json:"clock_speed,omitempty"`
	// The number of cores for the server cpu type.
	CoreCount int `json:"core_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServerCPUTypeQuery when eager-loading is set.
	Edges        ServerCPUTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ServerCPUTypeEdges holds the relations/edges for other nodes in the graph.
type ServerCPUTypeEdges struct {
	// CPU holds the value of the cpu edge.
	CPU []*ServerCPU `json:"cpu,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedCPU map[string][]*ServerCPU
}

// CPUOrErr returns the CPU value or an error if the edge
// was not loaded in eager-loading.
func (e ServerCPUTypeEdges) CPUOrErr() ([]*ServerCPU, error) {
	if e.loadedTypes[0] {
		return e.CPU, nil
	}
	return nil, &NotLoadedError{edge: "cpu"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ServerCPUType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case servercputype.FieldID:
			values[i] = new(gidx.PrefixedID)
		case servercputype.FieldCoreCount:
			values[i] = new(sql.NullInt64)
		case servercputype.FieldVendor, servercputype.FieldModel, servercputype.FieldClockSpeed:
			values[i] = new(sql.NullString)
		case servercputype.FieldCreatedAt, servercputype.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ServerCPUType fields.
func (sct *ServerCPUType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case servercputype.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				sct.ID = *value
			}
		case servercputype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sct.CreatedAt = value.Time
			}
		case servercputype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sct.UpdatedAt = value.Time
			}
		case servercputype.FieldVendor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vendor", values[i])
			} else if value.Valid {
				sct.Vendor = value.String
			}
		case servercputype.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				sct.Model = value.String
			}
		case servercputype.FieldClockSpeed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clock_speed", values[i])
			} else if value.Valid {
				sct.ClockSpeed = value.String
			}
		case servercputype.FieldCoreCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field core_count", values[i])
			} else if value.Valid {
				sct.CoreCount = int(value.Int64)
			}
		default:
			sct.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ServerCPUType.
// This includes values selected through modifiers, order, etc.
func (sct *ServerCPUType) Value(name string) (ent.Value, error) {
	return sct.selectValues.Get(name)
}

// QueryCPU queries the "cpu" edge of the ServerCPUType entity.
func (sct *ServerCPUType) QueryCPU() *ServerCPUQuery {
	return NewServerCPUTypeClient(sct.config).QueryCPU(sct)
}

// Update returns a builder for updating this ServerCPUType.
// Note that you need to call ServerCPUType.Unwrap() before calling this method if this ServerCPUType
// was returned from a transaction, and the transaction was committed or rolled back.
func (sct *ServerCPUType) Update() *ServerCPUTypeUpdateOne {
	return NewServerCPUTypeClient(sct.config).UpdateOne(sct)
}

// Unwrap unwraps the ServerCPUType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sct *ServerCPUType) Unwrap() *ServerCPUType {
	_tx, ok := sct.config.driver.(*txDriver)
	if !ok {
		panic("generated: ServerCPUType is not a transactional entity")
	}
	sct.config.driver = _tx.drv
	return sct
}

// String implements the fmt.Stringer.
func (sct *ServerCPUType) String() string {
	var builder strings.Builder
	builder.WriteString("ServerCPUType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sct.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sct.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sct.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("vendor=")
	builder.WriteString(sct.Vendor)
	builder.WriteString(", ")
	builder.WriteString("model=")
	builder.WriteString(sct.Model)
	builder.WriteString(", ")
	builder.WriteString("clock_speed=")
	builder.WriteString(sct.ClockSpeed)
	builder.WriteString(", ")
	builder.WriteString("core_count=")
	builder.WriteString(fmt.Sprintf("%v", sct.CoreCount))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (sct ServerCPUType) IsEntity() {}

// NamedCPU returns the CPU named value or an error if the edge was not
// loaded in eager-loading with this name.
func (sct *ServerCPUType) NamedCPU(name string) ([]*ServerCPU, error) {
	if sct.Edges.namedCPU == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := sct.Edges.namedCPU[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (sct *ServerCPUType) appendNamedCPU(name string, edges ...*ServerCPU) {
	if sct.Edges.namedCPU == nil {
		sct.Edges.namedCPU = make(map[string][]*ServerCPU)
	}
	if len(edges) == 0 {
		sct.Edges.namedCPU[name] = []*ServerCPU{}
	} else {
		sct.Edges.namedCPU[name] = append(sct.Edges.namedCPU[name], edges...)
	}
}

// ServerCPUTypes is a parsable slice of ServerCPUType.
type ServerCPUTypes []*ServerCPUType
