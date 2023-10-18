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

// Type holds the schema definition for the ServerHardDriveType entity.
type ServerHardDriveType struct {
	ent.Schema
}

// Mixin of the ServerHardDriveType type
func (ServerHardDriveType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the ServerHardDriveType.
func (ServerHardDriveType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the server hard drive type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(ServerHardDriveTypePrefix) }),
		field.Text("vendor").
			NotEmpty().
			Comment("The name of the vendor for the server hard drive type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("model").
			NotEmpty().
			Comment("The mode of the server chassis type."),
		field.Text("speed").
			NotEmpty().
			Comment("The speed of the server hard drive type."),
		field.Enum("type").
			Values("ssd", "hdd").
			Comment("The type of the server hard drive type."),
		// TODO: should capacity go on the hard drive itself?
		field.Text("capacity").
			NotEmpty().
			Comment("The capacity of the server hard drive type."),
	}
}

// Edges of the ServerHardDriveType
func (ServerHardDriveType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hard_drive", ServerHardDrive.Type).
			Ref("hard_drive_type").
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the ServerHardDriveType
func (ServerHardDriveType) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations for the ServerHardDriveType
func (ServerHardDriveType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Representation of a server hard drive type. ServerHardDriveType describes the available hard drive types for a server."),
		entgql.Type("ServerHardDriveType"),
		prefixIDDirective(ServerHardDriveTypePrefix),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Input information to create a server hard drive type."),
			entgql.MutationUpdate().Description("Input information to update a server hard drive type."),
		),
	}
}
