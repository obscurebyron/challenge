// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/byronka/articles_challenge/ent/article"
	"github.com/byronka/articles_challenge/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[1].Descriptor()
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = articleDescTitle.Validators[0].(func(string) error)
	// articleDescContent is the schema descriptor for content field.
	articleDescContent := articleFields[2].Descriptor()
	// article.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	article.ContentValidator = articleDescContent.Validators[0].(func(string) error)
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleFields[3].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleFields[4].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
}
