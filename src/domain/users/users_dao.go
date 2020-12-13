package users

import (
	"Honest-Game-Reviews/src/datasource/mysql/database"
	"Honest-Game-Reviews/src/logger"
	"Honest-Game-Reviews/src/utils/errors"
)

var (
	queryCreateUser = "INSERT INTO users (user_name, email, password, date_created, isAdmin, status) VALUES (?, ?, ?, ?, ?, ?)"
	queryGetUser = "SELECT * FROM users WHERE email=?"
)

func (user *User) SaveUser() *errors.RestErrors {
	stmt, err := database.DatabaseClient.Client.Prepare(queryCreateUser)
	if err != nil {
		logger.Error("error in preparing sql statment", err)
		return errors.NewInternalServerError("database error")
	}

	result, err := stmt.Exec(user.Username, user.Email, user.Password, user.DateCreated, user.IsAdmin, user.Status)
	if err != nil {
		logger.Error("error when saving user in database", err)
		return errors.NewInternalServerError("database error")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get user id from query result", err)
		return errors.NewNotFoundError("database error")
	}
	user.ID = userID
	return nil
}


func (UserLoginRequest *UserLoginRequest) GetUser() (*User, *errors.RestErrors) {
	stmt, stmtErr := database.DatabaseClient.Client.Prepare(queryGetUser)
	if stmtErr != nil {
		logger.Error("error in preparing sql statment", stmtErr)
		return nil, errors.NewInternalServerError("database error")
	}

	 row := stmt.QueryRow(UserLoginRequest.Email)
	 var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.DateCreated, &user.IsAdmin, &user.Status)
	if err != nil {
		logger.Error("error in query statment", err)
		return nil, errors.NewInternalServerError("database error")
	}
	return &user, nil
}
