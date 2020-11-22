package app

import (
	"Honest-Game-Reviews/src/http/rest/games_handler"

	"github.com/go-chi/chi"
)

func gamesURLMapping(router *chi.Mux) {
	router.Get("/games", games_handler.GamesHandler.GetAllGames)
	router.Get("/games/{game_id}", games_handler.GamesHandler.GetGame)
	router.Get("/games/company", games_handler.GamesHandler.QueryGamesByCompany)
}
