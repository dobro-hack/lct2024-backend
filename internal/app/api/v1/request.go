package v1

import (
	"database/sql"
	"net/http"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// SaveRequest godoc
// @Summary Save Request
// @Tags Request
// @Description Save Request
// @Produce  json
// @Success 200 {object} model.SaveRequestResponse
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /request [post]
func (a *v1handler) SaveRequest(w http.ResponseWriter, r *http.Request) {
	req := model.SaveRequest{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, a.api.ErrInvalidRequest(r, err))
		return
	}

	err := a.api.Storage.SaveRequest(r.Context(), req)
	if err != nil && err != sql.ErrNoRows {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}
	res := model.SaveRequestResponse{
		Status: http.StatusText(http.StatusOK),
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}
