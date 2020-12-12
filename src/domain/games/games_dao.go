package games

import (
	"Honest-Game-Reviews/src/datasource/mysql/database"
	"Honest-Game-Reviews/src/logger"
	"Honest-Game-Reviews/src/utils/errors"
	"fmt"
	"strings"
)

const (
	queryGetGame     = "SELECT id, title, released, image, company, rating, metacritic, platforms, genres FROM games WHERE id=?;"
	queryGetAllGames = "SELECT id, title, released, image, company, rating, metacritic, platforms, genres FROM games;"
)

func (games *Game) GetGame() *errors.RestErrors {
	stmt, err := database.DatabaseClient.Client.Prepare(queryGetGame)
	if err != nil {
		logger.Error("error when trying to prepare get user statment", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(games.ID)
	if getErr := row.Scan(&games.ID, &games.Title, &games.Released, &games.Image, &games.Company, &games.Rating, &games.Metacritic, &games.Platforms, &games.Genres); err != nil {
		logger.Error("error in row scan", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func GetAllGames() ([]Game, *errors.RestErrors) {
	stmt, err := database.DatabaseClient.Client.Prepare(queryGetAllGames)
	if err != nil {
		logger.Error("error when trying to prepare get user statment", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error wen getting rows from stmt", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	listOfGames := []Game{}

	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Released, &game.Image, &game.Company, &game.Rating, &game.Metacritic, &game.Platforms, &game.Genres); err != nil {
			logger.Error("error when scanning rows", err)
			return nil, errors.NewInternalServerError("database error")
		}
		listOfGames = append(listOfGames, game)
	}

	if len(listOfGames) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("there were no games in the database %v", listOfGames))
	}
	return listOfGames, nil
}

func QueryGamesByCompany(publishers []string) ([]Game, *errors.RestErrors) {
	var args []interface{}
	for _, publisher := range publishers {
		args = append(args, publisher)
	}
	sql := "SELECT * FROM games WHERE company IN (?" + strings.Repeat(",?", len(args)-1) + ")"

	stmt, err := database.DatabaseClient.Client.Prepare(sql)
	if err != nil {
		logger.Error("error when trying to prepare get user statment", err)
		return nil, errors.NewInternalServerError("database error")
	}
	rows, err := stmt.Query(args...)
	defer stmt.Close()
	if err != nil {
		logger.Error("error when getting rows from stmt", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	queryListOfGames := []Game{}
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Released, &game.Image, &game.Company, &game.Rating, &game.Metacritic, &game.Platforms, &game.Genres); err != nil {
			logger.Error("error when scanning rows", err)
			return nil, errors.NewInternalServerError("database error")
		}
		queryListOfGames = append(queryListOfGames, game)
	}

	if len(queryListOfGames) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("there were no companies found in the database based on those query results %v", queryListOfGames))
	}
	return queryListOfGames, nil
}


func QueryGamesByPlatform(platforms []string) ([]Game, *errors.RestErrors) {
	var args []interface{}

	for _, platform := range platforms {
		args = append(args, platform)
	}
	sql := "SELECT * FROM games WHERE platforms IN (?" + strings.Repeat(",?", len(args)-1) + ")"
	stmt, err := database.DatabaseClient.Client.Prepare(sql)
	if err != nil {
		logger.Error("error when trying to prepare get user statment", err)
		return nil, errors.NewInternalServerError("database error")
	}
	rows, err := stmt.Query(args...)
	defer stmt.Close()
	if err != nil {
		logger.Error("error when getting rows from stmt", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	queryListOfGames := []Game{}
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Released, &game.Image, &game.Company, &game.Rating, &game.Metacritic, &game.Platforms, &game.Genres); err != nil {
			logger.Error("error when scanning rows", err)
			return nil, errors.NewInternalServerError("database error")
		}
		queryListOfGames = append(queryListOfGames, game)
	}

	if len(queryListOfGames) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("there were no companies found in the database based on those query results %v", queryListOfGames))
	}
	return queryListOfGames, nil
}

func QueryGamesByMetacriticScore(metacriticScore []string) ([]Game, *errors.RestErrors) {
	var args []interface{}
	for _, metaScore := range metacriticScore {
		args = append(args, metaScore)
	}
	sql := 	"SELECT * FROM games WHERE metacritic IN (?" + strings.Repeat(",?", len(args)-1) + ")"
	stmt, err := database.DatabaseClient.Client.Prepare(sql)

	if err != nil {
		logger.Error("error when trying to prepare get user statment", err)
		return nil, errors.NewInternalServerError("database error")
	}
	rows, err := stmt.Query(args...)
	defer stmt.Close()
	if err != nil {
		logger.Error("error when getting rows from stmt", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	queryListOfGames := []Game{}
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Released, &game.Image, &game.Company, &game.Rating, &game.Metacritic, &game.Platforms, &game.Genres); err != nil {
			logger.Error("error when scanning rows", err)
			return nil, errors.NewInternalServerError("database error")
		}
		queryListOfGames = append(queryListOfGames, game)
	}
	if len(queryListOfGames) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("there were no companies found in the database based on those query results %v", queryListOfGames))
	}
	return queryListOfGames, nil
}
