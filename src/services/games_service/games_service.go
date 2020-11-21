package games_service

import (
	"Honest-Game-Reviews/src/domain/games"
	"Honest-Game-Reviews/src/utils/errors"
)

var (
	GameService *gamesService = &gamesService{}
)

type gamesService struct{}

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
