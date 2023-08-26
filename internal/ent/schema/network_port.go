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

// Type holds the schema definition for the ServerNetworkPort entity.
type ServerNetworkPort struct {
	ent.Schema
}

// Mixin of the ServerNetworkPort type
func (ServerNetworkPort) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerNetworkPort.
func (ServerNetworkPort) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server network card type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerNetworkPortPrefix) }),
		field.Text("mac_address").
			NotEmpty().
			Comment("The mac address for the server network port.").
			Annotations(
				entgql.OrderField("MAC_ADDRESS"),
			),
		field.Text("network_card_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server network card of this server network port.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("NETWORK_CARD"),
			),
	}
}

// Edges of the ServerNetworkPort
func (ServerNetworkPort) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("network_card", ServerNetworkCard.Type).
			Unique().
			Required().
			Immutable().
			Field("network_card_id").
			Annotations(),
	}
}

// Indexes of the ServerNetworkPort
func (ServerNetworkPort) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerNetworkPort
func (ServerNetworkPort) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server network card type. ServerNetworkPort describes the available network card types for a server."),
		entgql.Type("ServerNetworkPort"),
		prefixIDDirective(ServerNetworkPortPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server network card type."),
			entgql.MutationUpdate().Description("Input information to update a server network card type."),
		),
	}
}
