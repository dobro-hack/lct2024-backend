package v1

import (
	"database/sql"
	"net/http"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// GetRequestList godoc
// @Summary Get Request List
// @Tags Request
// @Description Request List
// @Produce  json
// @Success 200 {object} model.RequestList
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /request [get]
func (a *v1handler) GetRequestList(w http.ResponseWriter, r *http.Request) {
	list, cnt, err := a.api.Storage.ListRequest(r.Context())
	if err != nil && err != sql.ErrNoRows {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}
	res := model.RequestList{
		Items:           list,
		TotalItemsCount: cnt,
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}

// SaveRequest godoc
// @Summary Save Request
// @Tags Request
// @Description Save Request
// @Produce  json
// @Success 200 {object} model.StatusResponse
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /request [post]
func (a *v1handler) SaveRequest(w http.ResponseWriter, r *http.Request) {
	req := model.Request{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, a.api.ErrInvalidRequest(r, err))
		return
	}

	err := a.api.Storage.SaveRequest(r.Context(), req)
	if err != nil && err != sql.ErrNoRows {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}
	res := model.StatusResponse{
		Status: http.StatusText(http.StatusOK),
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}
