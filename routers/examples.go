package routers

import (
	"chi-boilerplate/controllers"
	"github.com/go-chi/chi/v5"
)

func ExamplesRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Post("/test/", controllers.CreateExample)
	})
}
