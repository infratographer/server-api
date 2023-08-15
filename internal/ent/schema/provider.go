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

// Provider holds the schema definition for the Provider entity.
type Provider struct {
	ent.Schema
}

// Mixin of the Provider type
func (Provider) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the Provider.
func (Provider) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server provider.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerProviderPrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the server provider.").
			Annotations(
				entgql.OrderField("NAME"),
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
	}
}

// Edges of the Provider
func (Provider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("servers", Server.Type).
			Ref("provider").
			Annotations(
				entgql.RelayConnection(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
	}
}

// Indexes of the Provider
func (Provider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

// Annotations for the Provider
func (Provider) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server provider. Server providers are responsible for provisioning and managing servers"),
		entgql.Type("ServerProvider"),
		prefixIDDirective(ServerProviderPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server provider."),
			entgql.MutationUpdate().Description("Input information to update a server provider."),
		),
	}
}
