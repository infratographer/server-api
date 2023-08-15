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

// Mixin of the Server type
func (Server) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerPrefix) }),
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
			Comment("The ID for the owner of this server.").
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
			Comment("The ID for the server type of this server.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER_TYPE"),
			),
	}
}

// Edges of the Server.
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
			Annotations(),
		edge.From("components", ServerComponent.Type).
			Ref("server").
			Annotations(
				entgql.RelayConnection(),
			),
		edge.From("attributes", ServerAttribute.Type).
			Ref("server").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the Server.
func (Server) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
		index.Fields("location_id"),
		index.Fields("provider_id"),
		index.Fields("server_type_id"),
	}
}

// Annotations of the Server
func (Server) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents a server on the graph."),
		entgql.RelayConnection(),
		prefixIDDirective(ServerPrefix),
		entgql.Implements("IPAddressable"),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new server."),
			entgql.MutationUpdate().Description("Update an existing server."),
		),
	}
}
