package v1

import (
	"net/http"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// SaveUser godoc
// @Summary Save User
// @Tags User
// @Description Save User
// @Produce  json
// @Success 200 {object} model.StatusResponse
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /user [post]
func (a *v1handler) SaveUser(w http.ResponseWriter, r *http.Request) {
	req := model.SaveUserRequest{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, a.api.ErrInvalidRequest(r, err))
		return
	}

	u := model.Device{
		DeviceID: req.DeviceID,
		UserID:   1,
	}
	err := a.api.Storage.SaveDevice(r.Context(), u)
	if err != nil {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}

	res := model.StatusResponse{
		Status: http.StatusText(http.StatusOK),
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}
