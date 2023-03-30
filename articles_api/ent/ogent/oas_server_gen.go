// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateArticle implements createArticle operation.
	//
	// Creates a new Article and persists it to storage.
	//
	// POST /articles
	CreateArticle(ctx context.Context, req *CreateArticleReq) (CreateArticleRes, error)
	// DeleteArticle implements deleteArticle operation.
	//
	// Deletes the Article with the requested ID.
	//
	// DELETE /articles/{id}
	DeleteArticle(ctx context.Context, params DeleteArticleParams) (DeleteArticleRes, error)
	// ListArticle implements listArticle operation.
	//
	// List Articles.
	//
	// GET /articles
	ListArticle(ctx context.Context, params ListArticleParams) (ListArticleRes, error)
	// ReadArticle implements readArticle operation.
	//
	// Finds the Article with the requested ID and returns it.
	//
	// GET /articles/{id}
	ReadArticle(ctx context.Context, params ReadArticleParams) (ReadArticleRes, error)
	// UpdateArticle implements updateArticle operation.
	//
	// Updates a Article and persists changes to storage.
	//
	// PATCH /articles/{id}
	UpdateArticle(ctx context.Context, req *UpdateArticleReq, params UpdateArticleParams) (UpdateArticleRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
