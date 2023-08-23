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

// Type holds the schema definition for the ServerHardDrive entity.
type ServerHardDrive struct {
	ent.Schema
}

// Mixin of the ServerHardDrive type
func (ServerHardDrive) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerHardDrive.
func (ServerHardDrive) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server hard drive type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerHardDrivePrefix) }),
		field.Text("serial").
			NotEmpty().
			Comment("The serial for the server hard drive.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("server_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server of this server hard drive.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER"),
			),
		field.Text("server_hard_drive_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the server hard drive type of this server hard drive.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("SERVER_HARD_DRIVE_TYPE"),
			),
	}
}

// Edges of the ServerHardDrive
func (ServerHardDrive) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("server", Server.Type).
			Unique().
			Required().
			Immutable().
			Field("server_id").
			Annotations(),
		edge.To("hard_drive_type", ServerHardDriveType.Type).
			Unique().
			Required().
			Immutable().
			Field("server_hard_drive_type_id").
			Annotations(),
	}
}

// Indexes of the ServerHardDrive
func (ServerHardDrive) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("server_id"),
		index.Fields("server_hard_drive_type_id"),
	}
}

// Annotations for the ServerHardDrive
func (ServerHardDrive) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server hard drive. ServerHardDrive describes the available hard drives for a server."),
		entgql.Type("ServerHardDrive"),
		prefixIDDirective(ServerHardDrivePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server hard drive."),
			entgql.MutationUpdate().Description("Input information to update a server hard drive."),
		),
	}
}
