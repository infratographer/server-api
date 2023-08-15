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

// ServerAttribute holds the schema definition for the Server entity.
type ServerAttribute struct {
	ent.Schema
}

// Mixin of the ServerAttribute type
func (ServerAttribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerAttribute.
func (ServerAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server attribute.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerAttributePrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the server attribute.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("value").
			NotEmpty().
			Comment("The value of the server attribute.").
			Annotations(
				entgql.OrderField("VALUE"),
			),
		field.String("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the server of this attribute.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
	}
}

// Edges of the Provider
func (ServerAttribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
	}
}

// Indexes of the Provider
func (ServerAttribute) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the Provider
func (ServerAttribute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server attribute. Server attributes describe various characteristics of a server"),
		entgql.Type("ServerAttribute"),
		prefixIDDirective(ServerAttributePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server attribute."),
			entgql.MutationUpdate().Description("Input information to update a server attribute."),
		),
	}
}
