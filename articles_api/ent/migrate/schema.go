// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "slug", Type: field.TypeString, Size: 2147483647},
		{Name: "title", Type: field.TypeString, Size: 2147483647},
		{Name: "excerpt", Type: field.TypeString, Size: 2147483647},
		{Name: "cover_image", Type: field.TypeString, Size: 2147483647},
		{Name: "date", Type: field.TypeString, Size: 2147483647},
		{Name: "author_name", Type: field.TypeString, Size: 2147483647},
		{Name: "author_picture_url", Type: field.TypeString, Size: 2147483647},
		{Name: "open_graph_image_url", Type: field.TypeString, Size: 2147483647},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
	}
)

func init() {
}
