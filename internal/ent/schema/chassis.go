package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Type holds the schema definition for the ServerChassis entity.
type ServerChassis struct {
	ent.Schema
}

// Mixin of the ServerChassis type
func (ServerChassis) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerChassis.
func (ServerChassis) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server chassis.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerChassisTypePrefix) }),
		field.Text("server_chassis_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server chassis type of this server chassis.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER_CHASSIS_TYPE"),
			),
		field.Text("parent_chassis_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the parent of this chassis.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("PARENT_CHASSIS"),
			),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server chassis.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
		field.Text("serial").
			NotEmpty().
			Comment("The serial number of the server chassis."),
	}
}

// Edges of the ServerChassis
func (ServerChassis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.To("server_chassis_type", ServerChassisType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_chassis_type_id").
			Annotations(),
	}
}

// Indexes of the ServerChassis
func (ServerChassis) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_chassis_type_id"),
		index.Fields("server_id"),
		index.Fields("parent_chassis_id"),
	}
}

// Annotations for the ServerChassis
func (ServerChassis) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server chassis. ServerChassis describes the available chassis types for a server."),
		entgql.Type("ServerChassis"),
		prefixIDDirective(ServerChassisPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server chassis."),
			entgql.MutationUpdate().Description("Input information to update a server chassis."),
		),
	}
}
