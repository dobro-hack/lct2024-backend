package v1

import (
	"database/sql"
	"log"
	"net/http"

	"firebase.google.com/go/messaging"
	"github.com/dobro-hack/lct2024-backend/internal/app/model"
	"github.com/go-chi/render"
)

// Push godoc
// @Summary Push
// @Tags Push
// @Description Push
// @Produce  json
// @Success 200 {object} model.StatusResponse
// @Failure 404 {object} api.ErrResponse
// @Failure 500 {object} api.ErrResponse
// @Router /push [post]
func (a *v1handler) Push(w http.ResponseWriter, r *http.Request) {
	req := model.Push{}
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, a.api.ErrInvalidRequest(r, err))
		return
	}

	list, err := a.api.Storage.ListDevice(r.Context())
	if err != nil && err != sql.ErrNoRows {
		render.Render(w, r, a.api.ErrInternalServerError(r, err))
		return
	}

	for _, device := range list {
		apsData := make(map[string]interface{})
		apsData["text"] = "here"
		pushMessage := &messaging.Message{
			Token: device.DeviceID,
			Android: &messaging.AndroidConfig{
				Priority: "high",
				Data: map[string]string{
					"apns-priority": "10",
				},
			},
			APNS: &messaging.APNSConfig{
				Headers: map[string]string{
					"apns-priority": "10",
				},
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						Alert: &messaging.ApsAlert{
							Title: req.Title,
							Body:  req.Message,
						},
						CustomData: apsData,
					},
				},
			},
		}
		pushMessage.Android.Notification = &messaging.AndroidNotification{
			Title: req.Title,
			Body:  req.Message,
			Tag:   "some",
		}
		_, err = a.api.Push.Send(r.Context(), pushMessage)
		if err != nil {
			log.Print(err)
		}
	}

	res := model.StatusResponse{
		Status: http.StatusText(http.StatusOK),
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &res)
}
