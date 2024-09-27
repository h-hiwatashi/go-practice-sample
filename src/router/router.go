package router

import (
	"github.com/h-hiwatashi/go-practice-sample/handler"
	"github.com/go-chi/chi"
)

func Get() *chi.Mux {
	r := chi.NewRouter()

	userHandler := handler.NewUserHandler()
	r.Route("/user", func(r chi.Router) {
		r.Post("/add", userHandler.Add)
		r.Get("/detail", userHandler.Get)
	})
	return r
}