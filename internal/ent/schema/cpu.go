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

// Type holds the schema definition for the ServerCPU entity.
type ServerCPU struct {
	ent.Schema
}

// Mixin of the ServerCPU type
func (ServerCPU) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerCPU.
func (ServerCPU) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server cpu.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerCPUPrefix) }),
		field.Text("server_cpu_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server cpu type of this server cpu.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER_CPU_TYPE"),
			),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server cpu.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
		field.Text("serial").
			NotEmpty().
			Comment("The serial number of the server cpu."),
	}
}

// Edges of the ServerCPU
func (ServerCPU) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.To("server_cpu_type", ServerCPUType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_cpu_type_id").
			Annotations(),
	}
}

// Indexes of the ServerCPU
func (ServerCPU) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_cpu_type_id"),
		index.Fields("server_id"),
	}
}

// Annotations for the ServerCPU
func (ServerCPU) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server cpu. ServerCPU describes the available cpu for a server."),
		entgql.Type("ServerCPU"),
		prefixIDDirective(ServerCPUPrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server cpu."),
			entgql.MutationUpdate().Description("Input information to update a server cpu."),
		),
	}
}
