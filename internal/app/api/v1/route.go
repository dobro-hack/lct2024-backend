package v1

import (
	"database/sql"
	"net/http"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// GetRouteList godoc
// @Summary Get list of Route
// @Tags Route
// @Description Route in park
// @Produce  json
// @Success 200 {object} model.RouteList
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /route [get]
func (a *v1handler) GetRouteList(w http.ResponseWriter, r *http.Request) {
	list, cnt, err := a.api.Storage.ListRoute(r.Context())
	if err != nil && err != sql.ErrNoRows {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}
	res := model.RouteList{
		Items:           list,
		TotalItemsCount: cnt,
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}
