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

// Type holds the schema definition for the Type entity.
type ServerType struct {
	ent.Schema
}

// Mixin of the ServerType type
func (ServerType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerType.
func (ServerType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerTypePrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the server type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("owner_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the owner of this server type.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("OWNER"),
			),
		// TODO: Add description field
	}
}

// Edges of the ServerType
func (ServerType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("servers", Server.Type).
			Ref("server_type").
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
	}
}

// Indexes of the ServerType
func (ServerType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

// Annotations for the ServerType
func (ServerType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server type. Server types describe the type of a server."),
		entgql.Type("ServerType"),
		prefixIDDirective(ServerTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server type."),
			entgql.MutationUpdate().Description("Input information to update a server type."),
		),
	}
}
