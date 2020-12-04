package app

import (
	"Honest-Game-Reviews/src/http/rest/games_handler"

	"github.com/go-chi/chi"
)

func gamesURLMapping(router *chi.Mux) {
	// GET methods for games
	router.Get("/api/v1/games", games_handler.GamesHandler.GetAllGames)
	router.Get("/api/v1/games/{game_id}", games_handler.GamesHandler.GetGame)
	router.Get("/api/v1/games/company", games_handler.GamesHandler.QueryGamesByCompany)
	router.Get("/api/v1/games/platforms", games_handler.GamesHandler.QueryGameByPlatforms)
	router.Get("/api/v1/games/metacriticScore", games_handler.GamesHandler.QueryGameByMetacriticScore)
}
