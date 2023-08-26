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

// Type holds the schema definition for the ServerNetworkCard entity.
type ServerNetworkCard struct {
	ent.Schema
}

// Mixin of the ServerNetworkCard type
func (ServerNetworkCard) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerNetworkCard.
func (ServerNetworkCard) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server network card type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerNetworkCardPrefix) }),
		field.Text("serial").
			NotEmpty().
			Comment("The serial number for the server network card.").
			Annotations(
				entgql.OrderField("SERIAL"),
			),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server network card.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
		field.Text("network_card_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server network card.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("NETWORK_CARD_TYPE"),
			),
	}
}

// Edges of the ServerNetworkCard
func (ServerNetworkCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("network_card_type", ServerNetworkCardType.Type).
			Unique().
			Required().
			Immutable().
			Field("network_card_type_id").
			Annotations(),
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.From("network_port", ServerNetworkPort.Type).
			Ref("network_card").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the ServerNetworkCard
func (ServerNetworkCard) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerNetworkCard
func (ServerNetworkCard) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server network card type. ServerNetworkCard describes the available network card types for a server."),
		entgql.Type("ServerNetworkCard"),
		prefixIDDirective(ServerNetworkCardPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server network card type."),
			entgql.MutationUpdate().Description("Input information to update a server network card type."),
		),
	}
}
