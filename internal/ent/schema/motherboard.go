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

// Type holds the schema definition for the ServerMotherboard entity.
type ServerMotherboard struct {
	ent.Schema
}

// Mixin of the ServerMotherboard type
func (ServerMotherboard) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerChassis.
func (ServerMotherboard) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server motherboard.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerMotherboardPrefix) }),
		field.Text("serial").
			NotEmpty().
			Comment("The serial of the server motherboard"),
		field.Text("server_motherboard_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server motherboard type of this server motherboard.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER_MOTHERBOARD_TYPE"),
			),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server motherboard.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
	}
}

// Edges of the ServerMotherboard
func (ServerMotherboard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.To("server_motherboard_type", ServerMotherboardType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_motherboard_type_id").
			Annotations(),
	}
}

// Indexes of the ServerMotherboard
func (ServerMotherboard) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerMotherboardType
func (ServerMotherboard) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server motherboard. ServerMotherboard describes the available motherboard for a server."),
		entgql.Type("ServerMotherboard"),
		prefixIDDirective(ServerMotherboardPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server motherboard."),
			entgql.MutationUpdate().Description("Input information to update a server motherboard."),
		),
	}
}
