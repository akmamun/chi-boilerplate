package routers

import (
	"chi-boilerplate/pkg/helpers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.SuccessResponse(w, "alive ok")
	})
	//Add All route
	ExamplesRoutes(router)
}
