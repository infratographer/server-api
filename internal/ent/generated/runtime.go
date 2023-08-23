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
	"time"

	"go.infratographer.com/server-api/internal/ent/generated/provider"
	"go.infratographer.com/server-api/internal/ent/generated/server"
	"go.infratographer.com/server-api/internal/ent/generated/serverchassis"
	"go.infratographer.com/server-api/internal/ent/generated/serverchassistype"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponent"
	"go.infratographer.com/server-api/internal/ent/generated/servercomponenttype"
	"go.infratographer.com/server-api/internal/ent/generated/servercpu"
	"go.infratographer.com/server-api/internal/ent/generated/servercputype"
	"go.infratographer.com/server-api/internal/ent/generated/serverharddrive"
	"go.infratographer.com/server-api/internal/ent/generated/serverharddrivetype"
	"go.infratographer.com/server-api/internal/ent/generated/servermemory"
	"go.infratographer.com/server-api/internal/ent/generated/servermemorytype"
	"go.infratographer.com/server-api/internal/ent/generated/servermotherboard"
	"go.infratographer.com/server-api/internal/ent/generated/servermotherboardtype"
	"go.infratographer.com/server-api/internal/ent/generated/servertype"
	"go.infratographer.com/server-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	providerMixin := schema.Provider{}.Mixin()
	providerMixinFields0 := providerMixin[0].Fields()
	_ = providerMixinFields0
	providerFields := schema.Provider{}.Fields()
	_ = providerFields
	// providerDescCreatedAt is the schema descriptor for created_at field.
	providerDescCreatedAt := providerMixinFields0[0].Descriptor()
	// provider.DefaultCreatedAt holds the default value on creation for the created_at field.
	provider.DefaultCreatedAt = providerDescCreatedAt.Default.(func() time.Time)
	// providerDescUpdatedAt is the schema descriptor for updated_at field.
	providerDescUpdatedAt := providerMixinFields0[1].Descriptor()
	// provider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	provider.DefaultUpdatedAt = providerDescUpdatedAt.Default.(func() time.Time)
	// provider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	provider.UpdateDefaultUpdatedAt = providerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// providerDescName is the schema descriptor for name field.
	providerDescName := providerFields[1].Descriptor()
	// provider.NameValidator is a validator for the "name" field. It is called by the builders before save.
	provider.NameValidator = providerDescName.Validators[0].(func(string) error)
	// providerDescID is the schema descriptor for id field.
	providerDescID := providerFields[0].Descriptor()
	// provider.DefaultID holds the default value on creation for the id field.
	provider.DefaultID = providerDescID.Default.(func() gidx.PrefixedID)
	serverMixin := schema.Server{}.Mixin()
	serverMixinFields0 := serverMixin[0].Fields()
	_ = serverMixinFields0
	serverFields := schema.Server{}.Fields()
	_ = serverFields
	// serverDescCreatedAt is the schema descriptor for created_at field.
	serverDescCreatedAt := serverMixinFields0[0].Descriptor()
	// server.DefaultCreatedAt holds the default value on creation for the created_at field.
	server.DefaultCreatedAt = serverDescCreatedAt.Default.(func() time.Time)
	// serverDescUpdatedAt is the schema descriptor for updated_at field.
	serverDescUpdatedAt := serverMixinFields0[1].Descriptor()
	// server.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	server.DefaultUpdatedAt = serverDescUpdatedAt.Default.(func() time.Time)
	// server.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	server.UpdateDefaultUpdatedAt = serverDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverDescName is the schema descriptor for name field.
	serverDescName := serverFields[1].Descriptor()
	// server.NameValidator is a validator for the "name" field. It is called by the builders before save.
	server.NameValidator = serverDescName.Validators[0].(func(string) error)
	// serverDescLocationID is the schema descriptor for location_id field.
	serverDescLocationID := serverFields[4].Descriptor()
	// server.LocationIDValidator is a validator for the "location_id" field. It is called by the builders before save.
	server.LocationIDValidator = serverDescLocationID.Validators[0].(func(string) error)
	// serverDescProviderID is the schema descriptor for provider_id field.
	serverDescProviderID := serverFields[5].Descriptor()
	// server.ProviderIDValidator is a validator for the "provider_id" field. It is called by the builders before save.
	server.ProviderIDValidator = serverDescProviderID.Validators[0].(func(string) error)
	// serverDescID is the schema descriptor for id field.
	serverDescID := serverFields[0].Descriptor()
	// server.DefaultID holds the default value on creation for the id field.
	server.DefaultID = serverDescID.Default.(func() gidx.PrefixedID)
	servercpuMixin := schema.ServerCPU{}.Mixin()
	servercpuMixinFields0 := servercpuMixin[0].Fields()
	_ = servercpuMixinFields0
	servercpuFields := schema.ServerCPU{}.Fields()
	_ = servercpuFields
	// servercpuDescCreatedAt is the schema descriptor for created_at field.
	servercpuDescCreatedAt := servercpuMixinFields0[0].Descriptor()
	// servercpu.DefaultCreatedAt holds the default value on creation for the created_at field.
	servercpu.DefaultCreatedAt = servercpuDescCreatedAt.Default.(func() time.Time)
	// servercpuDescUpdatedAt is the schema descriptor for updated_at field.
	servercpuDescUpdatedAt := servercpuMixinFields0[1].Descriptor()
	// servercpu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servercpu.DefaultUpdatedAt = servercpuDescUpdatedAt.Default.(func() time.Time)
	// servercpu.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servercpu.UpdateDefaultUpdatedAt = servercpuDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servercpuDescSerial is the schema descriptor for serial field.
	servercpuDescSerial := servercpuFields[3].Descriptor()
	// servercpu.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	servercpu.SerialValidator = servercpuDescSerial.Validators[0].(func(string) error)
	// servercpuDescID is the schema descriptor for id field.
	servercpuDescID := servercpuFields[0].Descriptor()
	// servercpu.DefaultID holds the default value on creation for the id field.
	servercpu.DefaultID = servercpuDescID.Default.(func() gidx.PrefixedID)
	servercputypeMixin := schema.ServerCPUType{}.Mixin()
	servercputypeMixinFields0 := servercputypeMixin[0].Fields()
	_ = servercputypeMixinFields0
	servercputypeFields := schema.ServerCPUType{}.Fields()
	_ = servercputypeFields
	// servercputypeDescCreatedAt is the schema descriptor for created_at field.
	servercputypeDescCreatedAt := servercputypeMixinFields0[0].Descriptor()
	// servercputype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servercputype.DefaultCreatedAt = servercputypeDescCreatedAt.Default.(func() time.Time)
	// servercputypeDescUpdatedAt is the schema descriptor for updated_at field.
	servercputypeDescUpdatedAt := servercputypeMixinFields0[1].Descriptor()
	// servercputype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servercputype.DefaultUpdatedAt = servercputypeDescUpdatedAt.Default.(func() time.Time)
	// servercputype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servercputype.UpdateDefaultUpdatedAt = servercputypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servercputypeDescVendor is the schema descriptor for vendor field.
	servercputypeDescVendor := servercputypeFields[1].Descriptor()
	// servercputype.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	servercputype.VendorValidator = servercputypeDescVendor.Validators[0].(func(string) error)
	// servercputypeDescModel is the schema descriptor for model field.
	servercputypeDescModel := servercputypeFields[2].Descriptor()
	// servercputype.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	servercputype.ModelValidator = servercputypeDescModel.Validators[0].(func(string) error)
	// servercputypeDescClockSpeed is the schema descriptor for clock_speed field.
	servercputypeDescClockSpeed := servercputypeFields[3].Descriptor()
	// servercputype.ClockSpeedValidator is a validator for the "clock_speed" field. It is called by the builders before save.
	servercputype.ClockSpeedValidator = servercputypeDescClockSpeed.Validators[0].(func(string) error)
	// servercputypeDescCoreCount is the schema descriptor for core_count field.
	servercputypeDescCoreCount := servercputypeFields[4].Descriptor()
	// servercputype.CoreCountValidator is a validator for the "core_count" field. It is called by the builders before save.
	servercputype.CoreCountValidator = servercputypeDescCoreCount.Validators[0].(func(int) error)
	// servercputypeDescID is the schema descriptor for id field.
	servercputypeDescID := servercputypeFields[0].Descriptor()
	// servercputype.DefaultID holds the default value on creation for the id field.
	servercputype.DefaultID = servercputypeDescID.Default.(func() gidx.PrefixedID)
	serverchassisMixin := schema.ServerChassis{}.Mixin()
	serverchassisMixinFields0 := serverchassisMixin[0].Fields()
	_ = serverchassisMixinFields0
	serverchassisFields := schema.ServerChassis{}.Fields()
	_ = serverchassisFields
	// serverchassisDescCreatedAt is the schema descriptor for created_at field.
	serverchassisDescCreatedAt := serverchassisMixinFields0[0].Descriptor()
	// serverchassis.DefaultCreatedAt holds the default value on creation for the created_at field.
	serverchassis.DefaultCreatedAt = serverchassisDescCreatedAt.Default.(func() time.Time)
	// serverchassisDescUpdatedAt is the schema descriptor for updated_at field.
	serverchassisDescUpdatedAt := serverchassisMixinFields0[1].Descriptor()
	// serverchassis.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serverchassis.DefaultUpdatedAt = serverchassisDescUpdatedAt.Default.(func() time.Time)
	// serverchassis.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	serverchassis.UpdateDefaultUpdatedAt = serverchassisDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverchassisDescSerial is the schema descriptor for serial field.
	serverchassisDescSerial := serverchassisFields[4].Descriptor()
	// serverchassis.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	serverchassis.SerialValidator = serverchassisDescSerial.Validators[0].(func(string) error)
	// serverchassisDescID is the schema descriptor for id field.
	serverchassisDescID := serverchassisFields[0].Descriptor()
	// serverchassis.DefaultID holds the default value on creation for the id field.
	serverchassis.DefaultID = serverchassisDescID.Default.(func() gidx.PrefixedID)
	serverchassistypeMixin := schema.ServerChassisType{}.Mixin()
	serverchassistypeMixinFields0 := serverchassistypeMixin[0].Fields()
	_ = serverchassistypeMixinFields0
	serverchassistypeFields := schema.ServerChassisType{}.Fields()
	_ = serverchassistypeFields
	// serverchassistypeDescCreatedAt is the schema descriptor for created_at field.
	serverchassistypeDescCreatedAt := serverchassistypeMixinFields0[0].Descriptor()
	// serverchassistype.DefaultCreatedAt holds the default value on creation for the created_at field.
	serverchassistype.DefaultCreatedAt = serverchassistypeDescCreatedAt.Default.(func() time.Time)
	// serverchassistypeDescUpdatedAt is the schema descriptor for updated_at field.
	serverchassistypeDescUpdatedAt := serverchassistypeMixinFields0[1].Descriptor()
	// serverchassistype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serverchassistype.DefaultUpdatedAt = serverchassistypeDescUpdatedAt.Default.(func() time.Time)
	// serverchassistype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	serverchassistype.UpdateDefaultUpdatedAt = serverchassistypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverchassistypeDescVendor is the schema descriptor for vendor field.
	serverchassistypeDescVendor := serverchassistypeFields[1].Descriptor()
	// serverchassistype.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	serverchassistype.VendorValidator = serverchassistypeDescVendor.Validators[0].(func(string) error)
	// serverchassistypeDescModel is the schema descriptor for model field.
	serverchassistypeDescModel := serverchassistypeFields[2].Descriptor()
	// serverchassistype.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	serverchassistype.ModelValidator = serverchassistypeDescModel.Validators[0].(func(string) error)
	// serverchassistypeDescHeight is the schema descriptor for height field.
	serverchassistypeDescHeight := serverchassistypeFields[3].Descriptor()
	// serverchassistype.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	serverchassistype.HeightValidator = serverchassistypeDescHeight.Validators[0].(func(string) error)
	// serverchassistypeDescID is the schema descriptor for id field.
	serverchassistypeDescID := serverchassistypeFields[0].Descriptor()
	// serverchassistype.DefaultID holds the default value on creation for the id field.
	serverchassistype.DefaultID = serverchassistypeDescID.Default.(func() gidx.PrefixedID)
	servercomponentMixin := schema.ServerComponent{}.Mixin()
	servercomponentMixinFields0 := servercomponentMixin[0].Fields()
	_ = servercomponentMixinFields0
	servercomponentFields := schema.ServerComponent{}.Fields()
	_ = servercomponentFields
	// servercomponentDescCreatedAt is the schema descriptor for created_at field.
	servercomponentDescCreatedAt := servercomponentMixinFields0[0].Descriptor()
	// servercomponent.DefaultCreatedAt holds the default value on creation for the created_at field.
	servercomponent.DefaultCreatedAt = servercomponentDescCreatedAt.Default.(func() time.Time)
	// servercomponentDescUpdatedAt is the schema descriptor for updated_at field.
	servercomponentDescUpdatedAt := servercomponentMixinFields0[1].Descriptor()
	// servercomponent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servercomponent.DefaultUpdatedAt = servercomponentDescUpdatedAt.Default.(func() time.Time)
	// servercomponent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servercomponent.UpdateDefaultUpdatedAt = servercomponentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servercomponentDescName is the schema descriptor for name field.
	servercomponentDescName := servercomponentFields[1].Descriptor()
	// servercomponent.NameValidator is a validator for the "name" field. It is called by the builders before save.
	servercomponent.NameValidator = servercomponentDescName.Validators[0].(func(string) error)
	// servercomponentDescVendor is the schema descriptor for vendor field.
	servercomponentDescVendor := servercomponentFields[2].Descriptor()
	// servercomponent.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	servercomponent.VendorValidator = servercomponentDescVendor.Validators[0].(func(string) error)
	// servercomponentDescModel is the schema descriptor for model field.
	servercomponentDescModel := servercomponentFields[3].Descriptor()
	// servercomponent.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	servercomponent.ModelValidator = servercomponentDescModel.Validators[0].(func(string) error)
	// servercomponentDescSerial is the schema descriptor for serial field.
	servercomponentDescSerial := servercomponentFields[4].Descriptor()
	// servercomponent.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	servercomponent.SerialValidator = servercomponentDescSerial.Validators[0].(func(string) error)
	// servercomponentDescServerID is the schema descriptor for server_id field.
	servercomponentDescServerID := servercomponentFields[5].Descriptor()
	// servercomponent.ServerIDValidator is a validator for the "server_id" field. It is called by the builders before save.
	servercomponent.ServerIDValidator = servercomponentDescServerID.Validators[0].(func(string) error)
	// servercomponentDescComponentTypeID is the schema descriptor for component_type_id field.
	servercomponentDescComponentTypeID := servercomponentFields[6].Descriptor()
	// servercomponent.ComponentTypeIDValidator is a validator for the "component_type_id" field. It is called by the builders before save.
	servercomponent.ComponentTypeIDValidator = servercomponentDescComponentTypeID.Validators[0].(func(string) error)
	// servercomponentDescID is the schema descriptor for id field.
	servercomponentDescID := servercomponentFields[0].Descriptor()
	// servercomponent.DefaultID holds the default value on creation for the id field.
	servercomponent.DefaultID = servercomponentDescID.Default.(func() gidx.PrefixedID)
	servercomponenttypeMixin := schema.ServerComponentType{}.Mixin()
	servercomponenttypeMixinFields0 := servercomponenttypeMixin[0].Fields()
	_ = servercomponenttypeMixinFields0
	servercomponenttypeFields := schema.ServerComponentType{}.Fields()
	_ = servercomponenttypeFields
	// servercomponenttypeDescCreatedAt is the schema descriptor for created_at field.
	servercomponenttypeDescCreatedAt := servercomponenttypeMixinFields0[0].Descriptor()
	// servercomponenttype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servercomponenttype.DefaultCreatedAt = servercomponenttypeDescCreatedAt.Default.(func() time.Time)
	// servercomponenttypeDescUpdatedAt is the schema descriptor for updated_at field.
	servercomponenttypeDescUpdatedAt := servercomponenttypeMixinFields0[1].Descriptor()
	// servercomponenttype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servercomponenttype.DefaultUpdatedAt = servercomponenttypeDescUpdatedAt.Default.(func() time.Time)
	// servercomponenttype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servercomponenttype.UpdateDefaultUpdatedAt = servercomponenttypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servercomponenttypeDescName is the schema descriptor for name field.
	servercomponenttypeDescName := servercomponenttypeFields[1].Descriptor()
	// servercomponenttype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	servercomponenttype.NameValidator = servercomponenttypeDescName.Validators[0].(func(string) error)
	// servercomponenttypeDescID is the schema descriptor for id field.
	servercomponenttypeDescID := servercomponenttypeFields[0].Descriptor()
	// servercomponenttype.DefaultID holds the default value on creation for the id field.
	servercomponenttype.DefaultID = servercomponenttypeDescID.Default.(func() gidx.PrefixedID)
	serverharddriveMixin := schema.ServerHardDrive{}.Mixin()
	serverharddriveMixinFields0 := serverharddriveMixin[0].Fields()
	_ = serverharddriveMixinFields0
	serverharddriveFields := schema.ServerHardDrive{}.Fields()
	_ = serverharddriveFields
	// serverharddriveDescCreatedAt is the schema descriptor for created_at field.
	serverharddriveDescCreatedAt := serverharddriveMixinFields0[0].Descriptor()
	// serverharddrive.DefaultCreatedAt holds the default value on creation for the created_at field.
	serverharddrive.DefaultCreatedAt = serverharddriveDescCreatedAt.Default.(func() time.Time)
	// serverharddriveDescUpdatedAt is the schema descriptor for updated_at field.
	serverharddriveDescUpdatedAt := serverharddriveMixinFields0[1].Descriptor()
	// serverharddrive.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serverharddrive.DefaultUpdatedAt = serverharddriveDescUpdatedAt.Default.(func() time.Time)
	// serverharddrive.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	serverharddrive.UpdateDefaultUpdatedAt = serverharddriveDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverharddriveDescSerial is the schema descriptor for serial field.
	serverharddriveDescSerial := serverharddriveFields[1].Descriptor()
	// serverharddrive.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	serverharddrive.SerialValidator = serverharddriveDescSerial.Validators[0].(func(string) error)
	// serverharddriveDescID is the schema descriptor for id field.
	serverharddriveDescID := serverharddriveFields[0].Descriptor()
	// serverharddrive.DefaultID holds the default value on creation for the id field.
	serverharddrive.DefaultID = serverharddriveDescID.Default.(func() gidx.PrefixedID)
	serverharddrivetypeMixin := schema.ServerHardDriveType{}.Mixin()
	serverharddrivetypeMixinFields0 := serverharddrivetypeMixin[0].Fields()
	_ = serverharddrivetypeMixinFields0
	serverharddrivetypeFields := schema.ServerHardDriveType{}.Fields()
	_ = serverharddrivetypeFields
	// serverharddrivetypeDescCreatedAt is the schema descriptor for created_at field.
	serverharddrivetypeDescCreatedAt := serverharddrivetypeMixinFields0[0].Descriptor()
	// serverharddrivetype.DefaultCreatedAt holds the default value on creation for the created_at field.
	serverharddrivetype.DefaultCreatedAt = serverharddrivetypeDescCreatedAt.Default.(func() time.Time)
	// serverharddrivetypeDescUpdatedAt is the schema descriptor for updated_at field.
	serverharddrivetypeDescUpdatedAt := serverharddrivetypeMixinFields0[1].Descriptor()
	// serverharddrivetype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	serverharddrivetype.DefaultUpdatedAt = serverharddrivetypeDescUpdatedAt.Default.(func() time.Time)
	// serverharddrivetype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	serverharddrivetype.UpdateDefaultUpdatedAt = serverharddrivetypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// serverharddrivetypeDescVendor is the schema descriptor for vendor field.
	serverharddrivetypeDescVendor := serverharddrivetypeFields[1].Descriptor()
	// serverharddrivetype.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	serverharddrivetype.VendorValidator = serverharddrivetypeDescVendor.Validators[0].(func(string) error)
	// serverharddrivetypeDescModel is the schema descriptor for model field.
	serverharddrivetypeDescModel := serverharddrivetypeFields[2].Descriptor()
	// serverharddrivetype.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	serverharddrivetype.ModelValidator = serverharddrivetypeDescModel.Validators[0].(func(string) error)
	// serverharddrivetypeDescSpeed is the schema descriptor for speed field.
	serverharddrivetypeDescSpeed := serverharddrivetypeFields[3].Descriptor()
	// serverharddrivetype.SpeedValidator is a validator for the "speed" field. It is called by the builders before save.
	serverharddrivetype.SpeedValidator = serverharddrivetypeDescSpeed.Validators[0].(func(string) error)
	// serverharddrivetypeDescType is the schema descriptor for type field.
	serverharddrivetypeDescType := serverharddrivetypeFields[4].Descriptor()
	// serverharddrivetype.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	serverharddrivetype.TypeValidator = serverharddrivetypeDescType.Validators[0].(func(string) error)
	// serverharddrivetypeDescCapacity is the schema descriptor for capacity field.
	serverharddrivetypeDescCapacity := serverharddrivetypeFields[5].Descriptor()
	// serverharddrivetype.CapacityValidator is a validator for the "capacity" field. It is called by the builders before save.
	serverharddrivetype.CapacityValidator = serverharddrivetypeDescCapacity.Validators[0].(func(string) error)
	// serverharddrivetypeDescID is the schema descriptor for id field.
	serverharddrivetypeDescID := serverharddrivetypeFields[0].Descriptor()
	// serverharddrivetype.DefaultID holds the default value on creation for the id field.
	serverharddrivetype.DefaultID = serverharddrivetypeDescID.Default.(func() gidx.PrefixedID)
	servermemoryMixin := schema.ServerMemory{}.Mixin()
	servermemoryMixinFields0 := servermemoryMixin[0].Fields()
	_ = servermemoryMixinFields0
	servermemoryFields := schema.ServerMemory{}.Fields()
	_ = servermemoryFields
	// servermemoryDescCreatedAt is the schema descriptor for created_at field.
	servermemoryDescCreatedAt := servermemoryMixinFields0[0].Descriptor()
	// servermemory.DefaultCreatedAt holds the default value on creation for the created_at field.
	servermemory.DefaultCreatedAt = servermemoryDescCreatedAt.Default.(func() time.Time)
	// servermemoryDescUpdatedAt is the schema descriptor for updated_at field.
	servermemoryDescUpdatedAt := servermemoryMixinFields0[1].Descriptor()
	// servermemory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servermemory.DefaultUpdatedAt = servermemoryDescUpdatedAt.Default.(func() time.Time)
	// servermemory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servermemory.UpdateDefaultUpdatedAt = servermemoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servermemoryDescSerial is the schema descriptor for serial field.
	servermemoryDescSerial := servermemoryFields[1].Descriptor()
	// servermemory.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	servermemory.SerialValidator = servermemoryDescSerial.Validators[0].(func(string) error)
	// servermemoryDescID is the schema descriptor for id field.
	servermemoryDescID := servermemoryFields[0].Descriptor()
	// servermemory.DefaultID holds the default value on creation for the id field.
	servermemory.DefaultID = servermemoryDescID.Default.(func() gidx.PrefixedID)
	servermemorytypeMixin := schema.ServerMemoryType{}.Mixin()
	servermemorytypeMixinFields0 := servermemorytypeMixin[0].Fields()
	_ = servermemorytypeMixinFields0
	servermemorytypeFields := schema.ServerMemoryType{}.Fields()
	_ = servermemorytypeFields
	// servermemorytypeDescCreatedAt is the schema descriptor for created_at field.
	servermemorytypeDescCreatedAt := servermemorytypeMixinFields0[0].Descriptor()
	// servermemorytype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servermemorytype.DefaultCreatedAt = servermemorytypeDescCreatedAt.Default.(func() time.Time)
	// servermemorytypeDescUpdatedAt is the schema descriptor for updated_at field.
	servermemorytypeDescUpdatedAt := servermemorytypeMixinFields0[1].Descriptor()
	// servermemorytype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servermemorytype.DefaultUpdatedAt = servermemorytypeDescUpdatedAt.Default.(func() time.Time)
	// servermemorytype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servermemorytype.UpdateDefaultUpdatedAt = servermemorytypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servermemorytypeDescVendor is the schema descriptor for vendor field.
	servermemorytypeDescVendor := servermemorytypeFields[1].Descriptor()
	// servermemorytype.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	servermemorytype.VendorValidator = servermemorytypeDescVendor.Validators[0].(func(string) error)
	// servermemorytypeDescModel is the schema descriptor for model field.
	servermemorytypeDescModel := servermemorytypeFields[2].Descriptor()
	// servermemorytype.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	servermemorytype.ModelValidator = servermemorytypeDescModel.Validators[0].(func(string) error)
	// servermemorytypeDescSpeed is the schema descriptor for speed field.
	servermemorytypeDescSpeed := servermemorytypeFields[3].Descriptor()
	// servermemorytype.SpeedValidator is a validator for the "speed" field. It is called by the builders before save.
	servermemorytype.SpeedValidator = servermemorytypeDescSpeed.Validators[0].(func(string) error)
	// servermemorytypeDescSize is the schema descriptor for size field.
	servermemorytypeDescSize := servermemorytypeFields[4].Descriptor()
	// servermemorytype.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	servermemorytype.SizeValidator = servermemorytypeDescSize.Validators[0].(func(string) error)
	// servermemorytypeDescID is the schema descriptor for id field.
	servermemorytypeDescID := servermemorytypeFields[0].Descriptor()
	// servermemorytype.DefaultID holds the default value on creation for the id field.
	servermemorytype.DefaultID = servermemorytypeDescID.Default.(func() gidx.PrefixedID)
	servermotherboardMixin := schema.ServerMotherboard{}.Mixin()
	servermotherboardMixinFields0 := servermotherboardMixin[0].Fields()
	_ = servermotherboardMixinFields0
	servermotherboardFields := schema.ServerMotherboard{}.Fields()
	_ = servermotherboardFields
	// servermotherboardDescCreatedAt is the schema descriptor for created_at field.
	servermotherboardDescCreatedAt := servermotherboardMixinFields0[0].Descriptor()
	// servermotherboard.DefaultCreatedAt holds the default value on creation for the created_at field.
	servermotherboard.DefaultCreatedAt = servermotherboardDescCreatedAt.Default.(func() time.Time)
	// servermotherboardDescUpdatedAt is the schema descriptor for updated_at field.
	servermotherboardDescUpdatedAt := servermotherboardMixinFields0[1].Descriptor()
	// servermotherboard.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servermotherboard.DefaultUpdatedAt = servermotherboardDescUpdatedAt.Default.(func() time.Time)
	// servermotherboard.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servermotherboard.UpdateDefaultUpdatedAt = servermotherboardDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servermotherboardDescSerial is the schema descriptor for serial field.
	servermotherboardDescSerial := servermotherboardFields[1].Descriptor()
	// servermotherboard.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	servermotherboard.SerialValidator = servermotherboardDescSerial.Validators[0].(func(string) error)
	// servermotherboardDescID is the schema descriptor for id field.
	servermotherboardDescID := servermotherboardFields[0].Descriptor()
	// servermotherboard.DefaultID holds the default value on creation for the id field.
	servermotherboard.DefaultID = servermotherboardDescID.Default.(func() gidx.PrefixedID)
	servermotherboardtypeMixin := schema.ServerMotherboardType{}.Mixin()
	servermotherboardtypeMixinFields0 := servermotherboardtypeMixin[0].Fields()
	_ = servermotherboardtypeMixinFields0
	servermotherboardtypeFields := schema.ServerMotherboardType{}.Fields()
	_ = servermotherboardtypeFields
	// servermotherboardtypeDescCreatedAt is the schema descriptor for created_at field.
	servermotherboardtypeDescCreatedAt := servermotherboardtypeMixinFields0[0].Descriptor()
	// servermotherboardtype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servermotherboardtype.DefaultCreatedAt = servermotherboardtypeDescCreatedAt.Default.(func() time.Time)
	// servermotherboardtypeDescUpdatedAt is the schema descriptor for updated_at field.
	servermotherboardtypeDescUpdatedAt := servermotherboardtypeMixinFields0[1].Descriptor()
	// servermotherboardtype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servermotherboardtype.DefaultUpdatedAt = servermotherboardtypeDescUpdatedAt.Default.(func() time.Time)
	// servermotherboardtype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servermotherboardtype.UpdateDefaultUpdatedAt = servermotherboardtypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servermotherboardtypeDescVendor is the schema descriptor for vendor field.
	servermotherboardtypeDescVendor := servermotherboardtypeFields[1].Descriptor()
	// servermotherboardtype.VendorValidator is a validator for the "vendor" field. It is called by the builders before save.
	servermotherboardtype.VendorValidator = servermotherboardtypeDescVendor.Validators[0].(func(string) error)
	// servermotherboardtypeDescModel is the schema descriptor for model field.
	servermotherboardtypeDescModel := servermotherboardtypeFields[2].Descriptor()
	// servermotherboardtype.ModelValidator is a validator for the "model" field. It is called by the builders before save.
	servermotherboardtype.ModelValidator = servermotherboardtypeDescModel.Validators[0].(func(string) error)
	// servermotherboardtypeDescID is the schema descriptor for id field.
	servermotherboardtypeDescID := servermotherboardtypeFields[0].Descriptor()
	// servermotherboardtype.DefaultID holds the default value on creation for the id field.
	servermotherboardtype.DefaultID = servermotherboardtypeDescID.Default.(func() gidx.PrefixedID)
	servertypeMixin := schema.ServerType{}.Mixin()
	servertypeMixinFields0 := servertypeMixin[0].Fields()
	_ = servertypeMixinFields0
	servertypeFields := schema.ServerType{}.Fields()
	_ = servertypeFields
	// servertypeDescCreatedAt is the schema descriptor for created_at field.
	servertypeDescCreatedAt := servertypeMixinFields0[0].Descriptor()
	// servertype.DefaultCreatedAt holds the default value on creation for the created_at field.
	servertype.DefaultCreatedAt = servertypeDescCreatedAt.Default.(func() time.Time)
	// servertypeDescUpdatedAt is the schema descriptor for updated_at field.
	servertypeDescUpdatedAt := servertypeMixinFields0[1].Descriptor()
	// servertype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	servertype.DefaultUpdatedAt = servertypeDescUpdatedAt.Default.(func() time.Time)
	// servertype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	servertype.UpdateDefaultUpdatedAt = servertypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// servertypeDescName is the schema descriptor for name field.
	servertypeDescName := servertypeFields[1].Descriptor()
	// servertype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	servertype.NameValidator = servertypeDescName.Validators[0].(func(string) error)
	// servertypeDescID is the schema descriptor for id field.
	servertypeDescID := servertypeFields[0].Descriptor()
	// servertype.DefaultID holds the default value on creation for the id field.
	servertype.DefaultID = servertypeDescID.Default.(func() gidx.PrefixedID)
}
