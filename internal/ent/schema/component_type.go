package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Type holds the schema definition for the ServerComponentType entity.
type ServerComponentType struct {
	ent.Schema
}

// Mixin of the ServerComponentType type
func (ServerComponentType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerComponentType.
func (ServerComponentType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server component type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerComponentTypePrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the server component type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the ServerComponentType
func (ServerComponentType) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the ServerComponentType
func (ServerComponentType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerType
func (ServerComponentType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server component type. ServerComponentType describes the available component types for a server."),
		entgql.Type("ServerComponentType"),
		prefixIDDirective(ServerComponentTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server component type."),
			entgql.MutationUpdate().Description("Input information to update a server component type."),
		),
	}
}
