package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StructTag(`json:"oid,omitempty"`),
		field.Text("slug").
			NotEmpty(),
		field.Text("title").
			NotEmpty(),
		field.Text("excerpt").
			NotEmpty(),
		field.Text("coverImage").
			NotEmpty(),
		field.Text("date").
			NotEmpty(),
		field.Text("author_name").
			NotEmpty(),
		field.Text("author_picture_url").
			NotEmpty(),
		field.Text("open_graph_image_url").
			NotEmpty(),
		field.Text("content").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
