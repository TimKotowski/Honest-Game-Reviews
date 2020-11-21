package games

import "database/sql"

type Game struct {
	ID         int64          `json:"id"`
	Title      string         `json:"game_title"`
	Released   string         `json:"released"`
	Image      string         `json:"game_image"`
	Company    string         `json:"company"`
	Rating     int            `json:"rating"`
	Metacritic int            `json:"metacritic"`
	Platforms  sql.NullString `json:"platforms"`
	Genres     sql.NullString `json:"genres"`
}

// type Platforms struct {
// 	PC          string `json:"pc"`
// 	Xbox360     string `json:"xbox_360"`
// 	XboxOne     string `json:"xbox_one"`
// 	XboxSeriesX string `json:"xbox_series_x"`
// 	XboxSeriesS string `json:"xbox_series_s"`
// 	PS4         string `json:"ps4"`
// 	PS5         string `json:"ps5"`
// }

// type Genre struct {
// 	Action             string `json:"action"`
// 	Horror             string `json:"horror"`
// 	MMO                string `json:"mmo"`
// 	MMORPG             string `json:"mmorpg"`
// 	SinglePlayer       string `json:"single_player"`
// 	FirstPersonShooter string `json:"first_person_shooter"`
// 	Strategy           string `json:"strategy"`
// 	MOBA               string `json:"moba"`
// 	Simulator          string `json:"simulator"`
// 	Survival           string `json:"survival"`
// 	RolePlay           string `json:"role_play"`
// 	ActionRolePlayGame string `json:"ARPG"`
// 	Indie              string `json:"indie"`
// 	Sports             string `json:"sports"`
// 	Racing             string `json:"racing"`
// 	VirtualReality     string `json:"VR"`
// 	Adventure          string `json:"adventure"`
// 	JRPG               string `json:"JRPG"`
// 	Fighting           string `json:"fighting"`
// 	Stealth            string `json:"stealth"`
// 	BattleRoyale       string `json:"battle_royale"`
// 	TowerDefense       string `json:"tower_defense"`
// 	Trivia             string `json:"trivia"`
// }

type Games []Game
