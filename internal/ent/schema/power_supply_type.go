package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// Type holds the schema definition for the ServerPowerSupplyType entity.
type ServerPowerSupplyType struct {
	ent.Schema
}

// Mixin of the ServerPowerSupplyType type
func (ServerPowerSupplyType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerPowerSupplyType.
func (ServerPowerSupplyType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server power supply type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerPowerSupplyTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server power supply type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The mode of the server power supply type."),
		field.Text("watts").
			NotEmpty().
			Comment("The watts of the server power supply type."),
	}
}

// Edges of the ServerPowerSupplyType
func (ServerPowerSupplyType) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the ServerPowerSupplyType
func (ServerPowerSupplyType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerPowerSupplyType
func (ServerPowerSupplyType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server power supply type. ServerPowerSupplyType describes the available power supply types for a server."),
		entgql.Type("ServerPowerSupplyType"),
		prefixIDDirective(ServerPowerSupplyTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server power supply type."),
			entgql.MutationUpdate().Description("Input information to update a server power supply type."),
		),
	}
}
