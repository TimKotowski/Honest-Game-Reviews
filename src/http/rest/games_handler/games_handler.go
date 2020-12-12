package games_handler

import (
	"Honest-Game-Reviews/src/services/games_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
)

type GamesHandler interface {
	GetGame(w http.ResponseWriter, r *http.Request)
	GetAllGames(w http.ResponseWriter, r *http.Request)
	QueryGamesByCompany(w http.ResponseWriter, r *http.Request)
	QueryGameByPlatforms(w http.ResponseWriter, r *http.Request)
	QueryGameByMetacriticScore(w http.ResponseWriter, r *http.Request)
}

type gamesHandler struct{}

func NewGamesHandler() GamesHandler {
	return &gamesHandler{}
}

func (handler *gamesHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.ParseInt(chi.URLParam(r, "game_id"), 10, 64)
	if err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid parameter, expect to get a number got something else", err)
		return
	}

	game, gameErr := games_service.NewGameService().GetGame(gameID)
	if gameErr != nil {
		json_utils.JsonErrorResponse(w, gameErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, game)
}

func (handler *gamesHandler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	games, gamesErr := games_service.NewGameService().GetAllGames()
	if gamesErr != nil {
		json_utils.JsonErrorResponse(w, gamesErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, games)
}

func (handler *gamesHandler) QueryGamesByCompany(w http.ResponseWriter, r *http.Request) {
	//Extract the value of the publisher parameter from the query string
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "error when parsing query, may not be set", err)
		return
	}
	publishers, ok := queryParams["publisher"]
	if !ok {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "no value to get out of publisher query", err)
		return
	}
	queriedGame, gameErr := games_service.NewGameService().QueryGamesByCompany(publishers)
	if gameErr != nil {
		json_utils.JsonErrorResponse(w, gameErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, queriedGame)
}

func (handler *gamesHandler) QueryGameByPlatforms(w http.ResponseWriter, r *http.Request) {
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		if queryParams == nil {
			json_utils.ClientErrorResponse(w, http.StatusBadRequest, "error when parsing query, may not be set", err)
			return
		}
	}
	platforms, ok := queryParams["platforms"]
	if !ok {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "no values given", err)
		return
	}

	queriedPlatforms, platformErr := games_service.NewGameService().QueryGamesByPlatform(platforms)
	if platformErr != nil {
		json_utils.JsonErrorResponse(w, platformErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, queriedPlatforms)
}

func (handler *gamesHandler) QueryGameByMetacriticScore(w http.ResponseWriter, r *http.Request) {
	queryParams, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "error when parsing query, may be bad query parmeter or may not be set", err)
		return
	}
	queriedMetacrticScore, ok := queryParams["metacritic"]
	if !ok {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "no values given", err)
		return
	}
	metacrticScore, metaErr := games_service.NewGameService().QueryGamesByMetacriticScore(queriedMetacrticScore)
	if metaErr != nil {
		json_utils.JsonErrorResponse(w, metaErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, metacrticScore)
}
