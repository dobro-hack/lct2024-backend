package v1

import (
	"github.com/go-chi/chi/v5"
)

// configureRouter Объявляем список доступных роутов
func (a *v1handler) ConfigureRouter() {
	a.api.Router.
		Route("/api/v1", func(r chi.Router) {
			r.Get("/route", a.GetRouteList)
		})
}
