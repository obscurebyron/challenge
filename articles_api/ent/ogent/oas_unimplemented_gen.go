// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateArticle implements createArticle operation.
//
// Creates a new Article and persists it to storage.
//
// POST /articles
func (UnimplementedHandler) CreateArticle(ctx context.Context, req *CreateArticleReq) (r CreateArticleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteArticle implements deleteArticle operation.
//
// Deletes the Article with the requested ID.
//
// DELETE /articles/{id}
func (UnimplementedHandler) DeleteArticle(ctx context.Context, params DeleteArticleParams) (r DeleteArticleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListArticle implements listArticle operation.
//
// List Articles.
//
// GET /articles
func (UnimplementedHandler) ListArticle(ctx context.Context, params ListArticleParams) (r ListArticleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadArticle implements readArticle operation.
//
// Finds the Article with the requested ID and returns it.
//
// GET /articles/{id}
func (UnimplementedHandler) ReadArticle(ctx context.Context, params ReadArticleParams) (r ReadArticleRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateArticle implements updateArticle operation.
//
// Updates a Article and persists changes to storage.
//
// PATCH /articles/{id}
func (UnimplementedHandler) UpdateArticle(ctx context.Context, req *UpdateArticleReq, params UpdateArticleParams) (r UpdateArticleRes, _ error) {
	return r, ht.ErrNotImplemented
}
