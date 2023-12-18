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

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

// Mixin to use for Server type
func (Server) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
		// softdelete.Mixin{},
	}
}

// Fields of the Instance.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerPrefix) }).
			Unique().
			Immutable().
			Comment("The ID for the server.").
			Annotations(
				entgql.OrderField("ID"),
			),
		field.Text("name").
			NotEmpty().
			Comment("The name of the server.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("description").
			Optional().
			Comment("The description of the server.").
			Annotations(
				entgql.OrderField("DESCRIPTION"),
			),
		field.String("owner_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the owner for this server.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("OWNER"),
			),
		field.String("location_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the location of this server.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
		field.String("provider_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the server provider for this server.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
		field.String("server_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			NotEmpty().
			Comment("The ID for the server type of this server.").
			Annotations(
				entgql.Type("ID"),
				entgql.Skip(^entgql.SkipMutationCreateInput),
			),
	}
}

// Edges of the Instance.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("provider", Provider.Type).
			Unique().
			Required().
			Immutable().
			Field("provider_id").
			Comment("The server provider for the server.").
			Annotations(
				entgql.MapsTo("serverProvider"),
			),
		edge.To("server_type", ServerType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_type_id").
			Comment("The server type for the server.").
			Annotations(
				entgql.MapsTo("serverType"),
			),
		edge.From("components", ServerComponent.Type).
			Ref("server").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the LoadBalancer
func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider_id"),
		index.Fields("location_id"),
		index.Fields("owner_id"),
		index.Fields("server_type_id"),
	}
}

// Annotations for the LoadBalancer
func (Server) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server."),
		prefixIDDirective(ServerPrefix),
		entgql.Implements("IPAddressable"),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server."),
			entgql.MutationUpdate().Description("Input information to update a server."),
		),
	}
}
