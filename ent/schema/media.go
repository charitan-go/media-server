package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/charitan-go/media-server/internal/media/dto"
	"github.com/google/uuid"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("readable_id", uuid.UUID{}).
			Default(uuid.New).
			StructTag(`json:"readableId"`),
		field.Enum("media_type").
			GoType(dto.MediaTypeEnum("")).
			StructTag(`json:"mediaType"`),
		field.String("project_id").
			StructTag(`json:"projectId"`),
		// field.String("name").
		// 	NotEmpty(),
		// field.String("description").
		// 	Default("A charity media"),
		// field.Float("goal").
		// 	Positive(),
		// field.Time("start_date").
		// 	StructTag(`json:"startDate"`),
		// field.Time("end_date").
		// 	StructTag(`json:"endDate"`),
		// field.Enum("category").
		// 	GoType(dto.CategoryEnum("")),
		// field.String("countryCode").
		// 	NotEmpty(),
		// field.Enum("status").
		// 	GoType(dto.StatusEnum("")),
		// field.String("owner_charity_readable_id").
		// 	StructTag(`json:"ownerCharityReadableId"`),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return nil
}
