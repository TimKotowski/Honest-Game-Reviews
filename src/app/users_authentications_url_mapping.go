package app

import (
	"Honest-Game-Reviews/src/http/rest/users_authentications_handler"

	"github.com/go-chi/chi"
)

func usersAuthenticationsURLMapping(router *chi.Mux) {
	router.Post("/api/v1/login", users_authentications_handler.NewusersAuthenticationsHandler().UserLogin)
}
