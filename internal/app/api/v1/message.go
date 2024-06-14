package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// GetMessageList godoc
// @Summary Get Message List
// @Tags Message
// @Description Message List
// @Produce  json
// @Success 200 {object} model.MessageList
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /message [get]
func (a *v1handler) GetMessageList(w http.ResponseWriter, r *http.Request) {
	list, cnt, err := a.api.Storage.ListMessage(r.Context())
	if err != nil && err != sql.ErrNoRows {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}
	res := model.MessageList{
		Items:           list,
		TotalItemsCount: cnt,
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}

// SaveMessage godoc
// @Summary Save Message
// @Tags Message
// @Description Save Message
// @Produce  json
// @Success 200 {object} model.SaveMessageResponse
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /message [post]
func (a *v1handler) SaveMessage(w http.ResponseWriter, r *http.Request) {
	req := model.SaveMessageRequest{}
	req.Lat = r.PostFormValue("lat")
	if req.Lat == "" {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("empty lat")))
		return
	}
	req.Lon = r.PostFormValue("lon")
	if req.Lon == "" {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("empty lon")))
		return
	}
	req.Type = r.PostFormValue("type")
	if req.Type == "" {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("empty type")))
		return
	}
	req.Message = r.PostFormValue("message")
	if req.Message == "" {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("empty message")))
		return
	}
	req.Phone = r.PostFormValue("phone")
	if req.Phone == "" {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("empty phone")))
		return
	}

	// get file
	f, handler, err := r.FormFile("file")
	if err != nil {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("something went wrong on getting file")))
		return
	}
	defer f.Close()

	// create folders
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	req.FileUrl = path + "/" + handler.Filename

	// open and copy files
	file, err := os.OpenFile(req.FileUrl, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("something went wrong on getting file")))
		return
	}
	defer file.Close()

	_, err = io.Copy(file, f)
	if err != nil {
		render.Render(w, r, a.api.ErrInternalServerError(r, errors.New("something went wrong on getting file")))
		return
	}

	j, _ := json.Marshal(struct {
		Lon string
		Lat string
	}{
		Lon: req.Lon,
		Lat: req.Lat,
	})
	m := model.Message{
		UserID:   1,
		SentAt:   time.Now(),
		Status:   "pending",
		Type:     req.Type,
		Message:  req.Message,
		Location: j,
		FileUrl:  req.FileUrl,
		Phone:    req.Phone,
	}
	err = a.api.Storage.SaveMessage(r.Context(), m)
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
