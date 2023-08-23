package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Type holds the schema definition for the ServerMemoryType entity.
type ServerMemoryType struct {
	ent.Schema
}

// Mixin of the ServerMemoryType type
func (ServerMemoryType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerMemoryType.
func (ServerMemoryType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server memory type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerMemoryTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server chassis type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The mode of the server chassis type."),
		field.Text("speed").
			NotEmpty().
			Comment("The speed of the server memory type."),
		field.Text("size").
			NotEmpty().
			Comment("The size of the server memory type."),
	}
}

// Edges of the ServerMemoryType
func (ServerMemoryType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("memory", ServerMemory.Type).
			Ref("server_memory_type").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the ServerMemoryType
func (ServerMemoryType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerMemoryType
func (ServerMemoryType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server memory type. ServerMemoryType describes the available memory types for a server."),
		entgql.Type("ServerMemoryType"),
		prefixIDDirective(ServerMemoryTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server memory type."),
			entgql.MutationUpdate().Description("Input information to update a server memory type."),
		),
	}
}
