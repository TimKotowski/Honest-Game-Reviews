package games_handler

import (
	"Honest-Game-Reviews/src/services/games_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
)

var (
	GamesHandler *gamesHandler = &gamesHandler{}
)

type gamesHandler struct{}


func (handler *gamesHandler) GetGame(w http.ResponseWriter, r *http.Request) {
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
	games, gamesErr := games_service.GameService.GetAllGames()
	if gamesErr != nil {
		json_utils.JsonErrorResponse(w, gamesErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, games)
}




func (handler *gamesHandler) QueryGamesByCompany(w http.ResponseWriter, r *http.Request) {
	//Extract the value of the publisher parameter from the query string
	query_params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	publishers, ok := query_params["publisher"]
	if !ok {
		w.Write([]byte("Publisher paramter not set"))
		return
	}
		queriedGame, gameErr := games_service.GameService.QueryGame(publishers)
		if gameErr != nil {
			json_utils.JsonErrorResponse(w, gameErr)
			return
		}
		json_utils.JsonResponse(w, http.StatusOK, queriedGame)

}
