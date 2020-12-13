package app

import (
	"Honest-Game-Reviews/src/http/rest/users_handler"

	"github.com/go-chi/chi"
)

func usersURLMapping(router *chi.Mux) {
	router.Post("/api/v1/create", users_handler.NewUsersHandler().CreateUser)
	router.Post("/api/v1/user", users_handler.NewUsersHandler().GetUser)
}
