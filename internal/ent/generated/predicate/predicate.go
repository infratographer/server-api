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

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Provider is the predicate function for provider builders.
type Provider func(*sql.Selector)

// Server is the predicate function for server builders.
type Server func(*sql.Selector)

// ServerChassis is the predicate function for serverchassis builders.
type ServerChassis func(*sql.Selector)

// ServerChassisType is the predicate function for serverchassistype builders.
type ServerChassisType func(*sql.Selector)

// ServerComponent is the predicate function for servercomponent builders.
type ServerComponent func(*sql.Selector)

// ServerComponentType is the predicate function for servercomponenttype builders.
type ServerComponentType func(*sql.Selector)

// ServerType is the predicate function for servertype builders.
type ServerType func(*sql.Selector)
