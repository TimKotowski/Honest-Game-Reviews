package app

import (
	"Honest-Game-Reviews/src/http/rest/users_handler"

	"github.com/go-chi/chi"
)

func usersURLMapping(router *chi.Mux) {
	router.Post("/api/v1/create", users_handler.NewUsersHandler().CreateUserAccount)
	router.Get("/api/v1/user/{user_id}", users_handler.NewUsersHandler().GetUserByID)
}
