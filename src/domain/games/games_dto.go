package games

import "database/sql"

type Game struct {
	ID         int64          `json:"id"`
	Title      string         `json:"game_title"`
	Released   string         `json:"released"`
	Image      string         `json:"game_image"`
	Company    string         `json:"company"`
	Rating     float64        `json:"rating"`
	Metacritic int            `json:"metacritic"`
	Platforms  sql.NullString `json:"platforms"`
	Genres     sql.NullString `json:"genres"`
}

type Games []Game
