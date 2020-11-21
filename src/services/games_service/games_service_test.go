package games_service

import (
	"Honest-Game-Reviews/src/datasource/mysql/database"
	"fmt"
	"net/http"
	"testing"
)

func TestGamesService(t *testing.T) {
	database.NewDatabase()
	var gamedID int64 = 1
	game, err := GameService.GetGame(gamedID)
	if err != nil {
		if err.Status != http.StatusOK {
			t.Errorf("Expected nil, received %d", err.Status)
		}
	}
	fmt.Println("game", game)
}

func TestAllGamesService(t *testing.T) {
	database.NewDatabase()
	game, err := GameService.GetAllGames()
	if err != nil {
		if err.Status != http.StatusOK {
			t.Errorf("Expected nil, received %d", err.Status)
		}
	}
	fmt.Println("game", game)
}
