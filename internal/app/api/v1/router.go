package v1

import (
	"github.com/go-chi/chi/v5"
)

// configureRouter Объявляем список доступных роутов
func (a *v1handler) ConfigureRouter() {
	a.api.Router.
		Route("/api/v1", func(r chi.Router) {
			r.Get("/route", a.GetRouteList)
			r.Get("/message", a.GetMessageList)
			r.Get("/request", a.GetRequestList)
			r.Post("/request", a.SaveRequest)
			r.Post("/org", a.SaveOrg)
			r.Post("/person", a.SavePerson)
			r.Post("/message", a.SaveMessage)
			r.Post("/user", a.SaveUser)
			r.Post("/push", a.Push)
		})
}
