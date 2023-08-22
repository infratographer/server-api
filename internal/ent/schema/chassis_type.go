package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Type holds the schema definition for the ServerChassisType entity.
type ServerChassisType struct {
	ent.Schema
}

// Mixin of the ServerChassisType type
func (ServerChassisType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerChassisType.
func (ServerChassisType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server chassis type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerChassisTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server chassis type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The mode of the server chassis type."),
		field.Text("height").
			NotEmpty().
			Comment("The height of the server chassis type."),
		field.Bool("is_full_depth").
			Comment("Whether the server chassis type is full depth."),
		field.Text("parent_chassis_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the parent of this chassis type.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("PARENT_CHASSIS_TYPE"),
			),
	}
}

// Edges of the ServerChassisType
func (ServerChassisType) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the ServerChassisType
func (ServerChassisType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent_chassis_type_id"),
	}
}

// Annotations for the ServerChassisType
func (ServerChassisType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server chassis type. ServerChassisType describes the available chassis types for a server."),
		entgql.Type("ServerChassisType"),
		prefixIDDirective(ServerChassisTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server chassis type."),
			entgql.MutationUpdate().Description("Input information to update a server chassis type."),
		),
	}
}
