package router

import (
	"github.com/go-chi/chi"
	"github.com/h-hiwatashi/go-practice-sample/handler"
	"github.com/h-hiwatashi/go-practice-sample/setting"
)

func Get(dbSetting setting.DB) *chi.Mux {
	r := chi.NewRouter()

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(dbSetting)

	// httpルーティング
	r.Route("/user", func(r chi.Router) {
		r.Post("/add", userHandler.Add)
		r.Get("/detail", userHandler.Get)
	})
	return r
}