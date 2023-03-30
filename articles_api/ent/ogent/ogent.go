// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/byronka/articles_challenge/ent"
	"github.com/byronka/articles_challenge/ent/article"
	"github.com/go-faster/jx"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateArticle handles POST /articles requests.
func (h *OgentHandler) CreateArticle(ctx context.Context, req *CreateArticleReq) (CreateArticleRes, error) {
	b := h.client.Article.Create()
	// Add all fields.
	b.SetTitle(req.Title)
	b.SetContent(req.Content)
	b.SetCreatedAt(req.CreatedAt)
	b.SetUpdatedAt(req.UpdatedAt)
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Article.Query().Where(article.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewArticleCreate(e), nil
}

// ReadArticle handles GET /articles/{id} requests.
func (h *OgentHandler) ReadArticle(ctx context.Context, params ReadArticleParams) (ReadArticleRes, error) {
	q := h.client.Article.Query().Where(article.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewArticleRead(e), nil
}

// UpdateArticle handles PATCH /articles/{id} requests.
func (h *OgentHandler) UpdateArticle(ctx context.Context, req *UpdateArticleReq, params UpdateArticleParams) (UpdateArticleRes, error) {
	b := h.client.Article.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Title.Get(); ok {
		b.SetTitle(v)
	}
	if v, ok := req.Content.Get(); ok {
		b.SetContent(v)
	}
	if v, ok := req.UpdatedAt.Get(); ok {
		b.SetUpdatedAt(v)
	}
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Article.Query().Where(article.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewArticleUpdate(e), nil
}

// DeleteArticle handles DELETE /articles/{id} requests.
func (h *OgentHandler) DeleteArticle(ctx context.Context, params DeleteArticleParams) (DeleteArticleRes, error) {
	err := h.client.Article.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteArticleNoContent), nil

}

// ListArticle handles GET /articles requests.
func (h *OgentHandler) ListArticle(ctx context.Context, params ListArticleParams) (ListArticleRes, error) {
	q := h.client.Article.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewArticleLists(es)
	return (*ListArticleOKApplicationJSON)(&r), nil
}