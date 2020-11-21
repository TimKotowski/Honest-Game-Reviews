package games_handler

import (
	"Honest-Game-Reviews/src/services/games_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var (
	GamesHandler *gamesHandler = &gamesHandler{}
)

type gamesHandler struct{}


func (handler *gamesHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	gameID, err := strconv.ParseInt(chi.URLParam(r, "game_id"), 10, 64)
	if err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	game, gameErr := games_service.GameService.GetGame(gameID)
	if gameErr != nil {
		json_utils.JsonErrorResponse(w, gameErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, game)
}

func (handler *gamesHandler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	games, gamesErr := games_service.GameService.GetAllGames()
	if gamesErr != nil {
		json_utils.JsonErrorResponse(w, gamesErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, games)
}
