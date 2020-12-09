package games_service

import (
	"Honest-Game-Reviews/src/domain/games"
	"Honest-Game-Reviews/src/utils/errors"
	"fmt"
	"strings"
)


type GameServiceInterface interface {
	GetGame(int64) (*games.Game, *errors.RestErrors)
	GetAllGames() (games.Games, *errors.RestErrors)
	QueryGamesByCompany([]string) (games.Games, *errors.RestErrors)
	QueryGamesByPlatform([]string) (games.Games, *errors.RestErrors)
	QueryGamesByMetacriticScore([]string) (games.Games, *errors.RestErrors)
}

type gamesService struct{}

// call to get a new instande of NewGameService to hiding implementation for the rest of the world
func NewGameService() GameServiceInterface {
	return &gamesService{}
}


func (s *gamesService) GetGame(gameID int64) (*games.Game, *errors.RestErrors) {
	game := &games.Game{ID: gameID}
	if err := game.GetGame(); err != nil {
		return nil, err
	}
	return game, nil
}

func (s *gamesService) GetAllGames() (games.Games, *errors.RestErrors) {
	return games.GetAllGames()
}

func (s *gamesService) QueryGamesByCompany(publishers []string) (games.Games, *errors.RestErrors) {
	for _, publisher := range publishers {
		if strings.Contains(publisher, "Reset List") {
			return games.GetAllGames()
		}
	}
	return games.QueryGamesByCompany(publishers)
}


func (s *gamesService) QueryGamesByPlatform(platforms []string) (games.Games, *errors.RestErrors) {
	for _, platform := range platforms {
		if strings.Contains(platform, "Reset List") {
			fmt.Println("hitting")
			return games.GetAllGames()
		}
	}
	return games.QueryGamesByPlatform(platforms)
}

func (s *gamesService) QueryGamesByMetacriticScore(metacriticScore []string) (games.Games, *errors.RestErrors)  {
	for _, score := range metacriticScore {
		if strings.Contains(score,"Reset List") {
			return games.GetAllGames()
		}
	}
	return games.QueryGamesByMetacriticScore(metacriticScore)
}
