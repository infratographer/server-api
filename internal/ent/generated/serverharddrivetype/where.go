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

package serverharddrivetype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/server-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldUpdatedAt, v))
}

// Vendor applies equality check predicate on the "vendor" field. It's identical to VendorEQ.
func Vendor(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldVendor, v))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldModel, v))
}

// Speed applies equality check predicate on the "speed" field. It's identical to SpeedEQ.
func Speed(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldSpeed, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldType, v))
}

// Capacity applies equality check predicate on the "capacity" field. It's identical to CapacityEQ.
func Capacity(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldCapacity, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldUpdatedAt, v))
}

// VendorEQ applies the EQ predicate on the "vendor" field.
func VendorEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldVendor, v))
}

// VendorNEQ applies the NEQ predicate on the "vendor" field.
func VendorNEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldVendor, v))
}

// VendorIn applies the In predicate on the "vendor" field.
func VendorIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldVendor, vs...))
}

// VendorNotIn applies the NotIn predicate on the "vendor" field.
func VendorNotIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldVendor, vs...))
}

// VendorGT applies the GT predicate on the "vendor" field.
func VendorGT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldVendor, v))
}

// VendorGTE applies the GTE predicate on the "vendor" field.
func VendorGTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldVendor, v))
}

// VendorLT applies the LT predicate on the "vendor" field.
func VendorLT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldVendor, v))
}

// VendorLTE applies the LTE predicate on the "vendor" field.
func VendorLTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldVendor, v))
}

// VendorContains applies the Contains predicate on the "vendor" field.
func VendorContains(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContains(FieldVendor, v))
}

// VendorHasPrefix applies the HasPrefix predicate on the "vendor" field.
func VendorHasPrefix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasPrefix(FieldVendor, v))
}

// VendorHasSuffix applies the HasSuffix predicate on the "vendor" field.
func VendorHasSuffix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasSuffix(FieldVendor, v))
}

// VendorEqualFold applies the EqualFold predicate on the "vendor" field.
func VendorEqualFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEqualFold(FieldVendor, v))
}

// VendorContainsFold applies the ContainsFold predicate on the "vendor" field.
func VendorContainsFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContainsFold(FieldVendor, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasSuffix(FieldModel, v))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContainsFold(FieldModel, v))
}

// SpeedEQ applies the EQ predicate on the "speed" field.
func SpeedEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldSpeed, v))
}

// SpeedNEQ applies the NEQ predicate on the "speed" field.
func SpeedNEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldSpeed, v))
}

// SpeedIn applies the In predicate on the "speed" field.
func SpeedIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldSpeed, vs...))
}

// SpeedNotIn applies the NotIn predicate on the "speed" field.
func SpeedNotIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldSpeed, vs...))
}

// SpeedGT applies the GT predicate on the "speed" field.
func SpeedGT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldSpeed, v))
}

// SpeedGTE applies the GTE predicate on the "speed" field.
func SpeedGTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldSpeed, v))
}

// SpeedLT applies the LT predicate on the "speed" field.
func SpeedLT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldSpeed, v))
}

// SpeedLTE applies the LTE predicate on the "speed" field.
func SpeedLTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldSpeed, v))
}

// SpeedContains applies the Contains predicate on the "speed" field.
func SpeedContains(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContains(FieldSpeed, v))
}

// SpeedHasPrefix applies the HasPrefix predicate on the "speed" field.
func SpeedHasPrefix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasPrefix(FieldSpeed, v))
}

// SpeedHasSuffix applies the HasSuffix predicate on the "speed" field.
func SpeedHasSuffix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasSuffix(FieldSpeed, v))
}

// SpeedEqualFold applies the EqualFold predicate on the "speed" field.
func SpeedEqualFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEqualFold(FieldSpeed, v))
}

// SpeedContainsFold applies the ContainsFold predicate on the "speed" field.
func SpeedContainsFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContainsFold(FieldSpeed, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContainsFold(FieldType, v))
}

// CapacityEQ applies the EQ predicate on the "capacity" field.
func CapacityEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEQ(FieldCapacity, v))
}

// CapacityNEQ applies the NEQ predicate on the "capacity" field.
func CapacityNEQ(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNEQ(FieldCapacity, v))
}

// CapacityIn applies the In predicate on the "capacity" field.
func CapacityIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldIn(FieldCapacity, vs...))
}

// CapacityNotIn applies the NotIn predicate on the "capacity" field.
func CapacityNotIn(vs ...string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldNotIn(FieldCapacity, vs...))
}

// CapacityGT applies the GT predicate on the "capacity" field.
func CapacityGT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGT(FieldCapacity, v))
}

// CapacityGTE applies the GTE predicate on the "capacity" field.
func CapacityGTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldGTE(FieldCapacity, v))
}

// CapacityLT applies the LT predicate on the "capacity" field.
func CapacityLT(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLT(FieldCapacity, v))
}

// CapacityLTE applies the LTE predicate on the "capacity" field.
func CapacityLTE(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldLTE(FieldCapacity, v))
}

// CapacityContains applies the Contains predicate on the "capacity" field.
func CapacityContains(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContains(FieldCapacity, v))
}

// CapacityHasPrefix applies the HasPrefix predicate on the "capacity" field.
func CapacityHasPrefix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasPrefix(FieldCapacity, v))
}

// CapacityHasSuffix applies the HasSuffix predicate on the "capacity" field.
func CapacityHasSuffix(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldHasSuffix(FieldCapacity, v))
}

// CapacityEqualFold applies the EqualFold predicate on the "capacity" field.
func CapacityEqualFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldEqualFold(FieldCapacity, v))
}

// CapacityContainsFold applies the ContainsFold predicate on the "capacity" field.
func CapacityContainsFold(v string) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(sql.FieldContainsFold(FieldCapacity, v))
}

// HasHardDrive applies the HasEdge predicate on the "hard_drive" edge.
func HasHardDrive() predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, HardDriveTable, HardDriveColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHardDriveWith applies the HasEdge predicate on the "hard_drive" edge with a given conditions (other predicates).
func HasHardDriveWith(preds ...predicate.ServerHardDrive) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(func(s *sql.Selector) {
		step := newHardDriveStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ServerHardDriveType) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ServerHardDriveType) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ServerHardDriveType) predicate.ServerHardDriveType {
	return predicate.ServerHardDriveType(func(s *sql.Selector) {
		p(s.Not())
	})
}
