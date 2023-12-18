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

// Type holds the schema definition for the ServerMemory entity.
type ServerMemory struct {
	ent.Schema
}

// Mixin of the ServerMemory type
func (ServerMemory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerMemory.
func (ServerMemory) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server memory.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerMemoryPrefix) }),
		field.Text("serial").
			NotEmpty().
			Comment("The serial of the server memory."),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server chassis.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
		field.Text("server_memory_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server memory type of this server memory.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER_MEMORY_TYPE"),
			),
	}
}

// Edges of the ServerMemory
func (ServerMemory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.To("server_memory_type", ServerMemoryType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_memory_type_id").
			Annotations(),
	}
}

// Indexes of the ServerMemory
func (ServerMemory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_id"),
		index.Fields("server_memory_type_id"),
	}
}

// Annotations for the ServerMemory
func (ServerMemory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server memory. ServerMemory describes the available memory for a server."),
		entgql.Type("ServerMemory"),
		prefixIDDirective(ServerMemoryPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server memory."),
			entgql.MutationUpdate().Description("Input information to update a server memory."),
		),
	}
}
