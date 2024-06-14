package v1

import (
	"database/sql"
	"net/http"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// SaveOrg godoc
// @Summary Save Org
// @Tags Org
// @Description Save Org
// @Produce  json
// @Success 200 {object} model.StatusResponse
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /org [post]
func (a *v1handler) SaveOrg(w http.ResponseWriter, r *http.Request) {
	req := model.Org{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, a.api.ErrInvalidRequest(r, err))
		return
	}

	err := a.api.Storage.SaveOrg(r.Context(), req)
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
