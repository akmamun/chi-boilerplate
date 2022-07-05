package routers

import (
	"chi-boilerplate/controllers"
	"chi-boilerplate/pkg/database"
	"github.com/go-chi/chi/v5"
)

func ExamplesRoutes(router *chi.Mux) {
	ctrl := controllers.BaseController{DB: database.GetDB()}
	router.Group(func(r chi.Router) {
		router.Group(func(r chi.Router) {
			r.Post("/test/", ctrl.CreateExample)
			r.Get("/test/", ctrl.GetData)
		})
	})
}
