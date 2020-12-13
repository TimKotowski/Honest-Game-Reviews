package app

import (
	"Honest-Game-Reviews/src/http/rest/games_handler"

	"github.com/go-chi/chi"
)

func gamesURLMapping(router *chi.Mux) {
	router.Get("/api/v1/games", games_handler.NewGamesHandler().GetAllGames)
	router.Get("/api/v1/games/{game_id}", games_handler.NewGamesHandler().GetGame)
	router.Get("/api/v1/games/company", games_handler.NewGamesHandler().QueryGamesByCompany)
	router.Get("/api/v1/games/platforms", games_handler.NewGamesHandler().QueryGameByPlatforms)
	router.Get("/api/v1/games/metacriticScore", games_handler.NewGamesHandler().QueryGameByMetacriticScore)
}
