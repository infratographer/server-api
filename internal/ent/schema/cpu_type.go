package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Type holds the schema definition for the ServerCPUType entity.
type ServerCPUType struct {
	ent.Schema
}

// Mixin of the ServerCPUType type
func (ServerCPUType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerCPUType.
func (ServerCPUType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server cpu type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerCPUTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server cpu type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The mode of the server cpu type."),
		field.Text("clock_speed").
			NotEmpty().
			Comment("The clock speed of the server cpu type."),
		field.Int("core_count").
			Positive().
			Comment("The number of cores for the server cpu type."),
	}
}

// Edges of the ServerCPUType
func (ServerCPUType) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the ServerCPUType
func (ServerCPUType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerCPUType
func (ServerCPUType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server cpu type. ServerCPUType describes the available cpu types for a server."),
		entgql.Type("ServerCPUType"),
		prefixIDDirective(ServerCPUTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server cpu type."),
			entgql.MutationUpdate().Description("Input information to update a server cpu type."),
		),
	}
}
