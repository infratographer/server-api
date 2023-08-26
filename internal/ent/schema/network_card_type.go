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

// Type holds the schema definition for the ServerNetworkCardType entity.
type ServerNetworkCardType struct {
	ent.Schema
}

// Mixin of the ServerNetworkCardType type
func (ServerNetworkCardType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerNetworkCardType.
func (ServerNetworkCardType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server network card type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerNetworkCardTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server network card type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The model of the server network card type."),
		field.Int("port_count").
			Positive().
			Comment("The number of ports on the server network card type."),
	}
}

// Edges of the ServerNetworkCardType
func (ServerNetworkCardType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("network_card", ServerNetworkCard.Type).
			Ref("network_card_type").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the ServerNetworkCardType
func (ServerNetworkCardType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerNetworkCardType
func (ServerNetworkCardType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server network card type. ServerNetworkCardType describes the available network card types for a server."),
		entgql.Type("ServerNetworkCardType"),
		prefixIDDirective(ServerNetworkCardTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server network card type."),
			entgql.MutationUpdate().Description("Input information to update a server network card type."),
		),
	}
}
