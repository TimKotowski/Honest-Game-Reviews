package users

import (
	"Honest-Game-Reviews/src/datasource/mysql/database"
	"Honest-Game-Reviews/src/logger"
	"Honest-Game-Reviews/src/utils/errors"
)

var (
	queryCreateUser = "INSERT INTO users (user_name, email, password, date_created, isAdmin, status) VALUES (?, ?, ?, ?, ?, ?)"
	queryGetUser    = "SELECT * FROM users WHERE email=?"
	queryGetUserByID = "SELECT * FROM users WHERE id=?"
)

func (user *User) SaveUser() *errors.RestErrors {
	stmt, err := database.DatabaseClient.Client.Prepare(queryCreateUser)
	if err != nil {
		logger.Error("error in preparing sql statment", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Email, user.Password, user.DateCreated, user.IsAdmin, user.Status)
	if err != nil {
		logger.Error("error when saving user in database", err)
		return errors.NewInternalServerError("email or username already in use, please use a different username or check to see if the email is right")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get user id from query result", err)
		return errors.NewNotFoundError("database error")
	}
	user.ID = userID
	return nil
}

func (userLoginRequest *UserLoginRequest) GetUserByEmail() (*User, *errors.RestErrors) {
	stmt, stmtErr := database.DatabaseClient.Client.Prepare(queryGetUser)
	if stmtErr != nil {
		logger.Error("error in preparing sql statment", stmtErr)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(userLoginRequest.Email)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.DateCreated, &user.IsAdmin, &user.Status)
	if err != nil {
		logger.Error("error when scanning user row", err)
		return nil, errors.NewInternalServerError("database error")
	}

	if user.ID == 0 {
		logger.Info("unable to fetch the user from the database")
		return nil, errors.NewBadRequestError("no user found")
	}
	return &user, nil
}

func (user *User) GetUserByID() *errors.RestErrors {
	stmt, stmtErr := database.DatabaseClient.Client.Prepare(queryGetUserByID)
	if stmtErr != nil {
		logger.Error("error when preparing sql statement", stmtErr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	row := stmt.QueryRow(user.ID)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.DateCreated, &user.IsAdmin, &user.Status)
	if err != nil {
		logger.Error("erorr when trying to scan user row", err)
		return errors.NewInternalServerError("database error")
	}

	if user.Username == "" {
		logger.Info("unable to get the user from the database")
		return errors.NewBadRequestError("no user found")
	}
	return nil
}
