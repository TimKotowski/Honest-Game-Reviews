package games_handler

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetSingleGameHandler(t *testing.T) {
	// Make sure server is running in a different terminal.
	// Seed the database with a game, so you have a single game, with an ID of 1.
	res, err := http.Get("http://localhost:8080" + "/games/1")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetAllGames(t *testing.T) {
	res, err := http.Get("http://localhost:8080" + "/games")
	if err != nil {
		t.Errorf("expecit nil, recived %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, recived %d", http.StatusOK, res.StatusCode)
	}
	fmt.Println("res", res)
}

func TestQueryGamesByCompany(t *testing.T) {
	res, err := http.Get("http://localhost:8080" + "/games/company?" + "publisher=blizzard&publisher=epic%20games&publisher=sega")
	if err != nil {
		t.Errorf("Expect nil, recived %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expeced %d, recived %d", http.StatusOK, res.StatusCode)
	}

}
