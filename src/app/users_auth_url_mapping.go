package app

import (
	"Honest-Game-Reviews/src/http/rest/auth_handler"

	"github.com/go-chi/chi"
)

func usersAuthURLMapping(router *chi.Mux) {
	router.Post("/api/v1/login", auth_handler.NewAuthHandler().UserLogin)
}
