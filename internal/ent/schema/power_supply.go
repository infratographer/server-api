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

// Type holds the schema definition for the ServerPowerSupply entity.
type ServerPowerSupply struct {
	ent.Schema
}

// Mixin of the ServerPowerSupply type
func (ServerPowerSupply) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerPowerSupply.
func (ServerPowerSupply) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server power supply type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerPowerSupplyPrefix) }),
		field.Text("serial").
			NotEmpty().
			Comment("The serial of the server power supply."),
		field.Text("server_power_supply_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server power supply type.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("PARENT_CHASSIS"),
			),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server power supply.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
	}
}

// Edges of the ServerPowerSupply
func (ServerPowerSupply) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.To("server_power_supply_type", ServerPowerSupplyType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_power_supply_type_id").
			Annotations(),
	}
}

// Indexes of the ServerPowerSupply
func (ServerPowerSupply) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_power_supply_type_id"),
		index.Fields("server_id"),
	}
}

// Annotations for the ServerPowerSupply
func (ServerPowerSupply) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server power supply. ServerPowerSupply describes the available power supply options for a server."),
		entgql.Type("ServerPowerSupply"),
		prefixIDDirective(ServerPowerSupplyPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server power supply."),
			entgql.MutationUpdate().Description("Input information to update a server power supply."),
		),
	}
}
