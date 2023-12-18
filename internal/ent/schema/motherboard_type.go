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

// Type holds the schema definition for the ServerMotherboardType entity.
type ServerMotherboardType struct {
	ent.Schema
}

// Mixin of the ServerMotherboardType type
func (ServerMotherboardType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerChassisType.
func (ServerMotherboardType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server motherboard type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerMotherboardTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server motherboard type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The mode of the server chassis type."),
	}
}

// Edges of the ServerMotherboardType
func (ServerMotherboardType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("motherboard", ServerMotherboard.Type).
			Ref("server_motherboard_type").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the ServerMotherboardType
func (ServerMotherboardType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerMotherboardType
func (ServerMotherboardType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server motherboard type. ServerMotherboardType describes the available motherboard types for a server."),
		entgql.Type("ServerMotherboardType"),
		prefixIDDirective(ServerMotherboardTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server motherboard type."),
			entgql.MutationUpdate().Description("Input information to update a server motherboard type."),
		),
	}
}
