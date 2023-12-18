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

// Type holds the schema definition for the ServerComponent entity.
type ServerComponent struct {
	ent.Schema
}

// Mixin of the ServerComponent type
func (ServerComponent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerComponent.
func (ServerComponent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server component.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerComponentPrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the server component.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor of the server component.").
			Annotations(
				entgql.OrderField("VENDOR"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The model of the server component.").
			Annotations(
				entgql.OrderField("MODEL"),
			),
		field.Text("serial").
			NotEmpty().
			Comment("The serial number of the server component.").
			Annotations(
				entgql.OrderField("SERIAL"),
			),
		field.String("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the server of this server component.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
		field.String("component_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the server component type of this server component.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
	}
}

// Edges of the ServerComponent
func (ServerComponent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("component_type", ServerComponentType.Type).
			Unique().
			Required().
			Immutable().
			Field("component_type_id").
			Comment("The server component type for the server component.").
			Annotations(
				entgql.MapsTo("serverComponentType"),
			),
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
	}
}

// Indexes of the ServerComponent
func (ServerComponent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_id"),
		index.Fields("component_type_id"),
	}
}

// Annotations for the ServerComponent
func (ServerComponent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server component. ServerComponent describes the components for a server."),
		entgql.Type("ServerComponent"),
		prefixIDDirective(ServerComponentPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server component."),
			entgql.MutationUpdate().Description("Input information to update a server component."),
		),
	}
}
