package users_service

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/logger"
	"Honest-Game-Reviews/src/utils/date_utils"
	"Honest-Game-Reviews/src/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	NewUsersService UsersServiceInterface = &usersService{}
	statusActive    string                = "active"
)

type UsersServiceInterface interface {
	CreateUserAccount(users.User) (*users.User, *errors.RestErrors)
	UserLogin(users.UserLoginRequest) (*users.User, *errors.RestErrors)
	GetUserByID(int64) (*users.User, *errors.RestErrors)
}

type usersService struct{}

func (s *usersService) CreateUserAccount(user users.User) (*users.User, *errors.RestErrors) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// hash password
	user.GetHash()
	user.Status = statusActive
	// set dateTime and format it right to be in mysql DB
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.SaveUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

// route is only meant for getting a user by email once they login so we can assign a token once verified
func (s *usersService) UserLogin(UserLoginRequest users.UserLoginRequest) (*users.User, *errors.RestErrors) {
	if err := UserLoginRequest.LoginRequestValidation(); err != nil {
		return nil, err
	}
	user, err := UserLoginRequest.GetUserByEmail()
	if err != nil {
		return nil, err
	}
	userPass := []byte(UserLoginRequest.Password)
	userFromDataBase := []byte(user.Password)

	// CompareHashAndPassword compares a bcrypt hashed password from the user in the databse with its possible match when the user trying to log in
	// if both passwords are a match we have the right user to move forward
	passErr := bcrypt.CompareHashAndPassword(userFromDataBase, userPass)
	if passErr != nil {
		logger.Error("error in comparing passwords", passErr)
		return nil, errors.NewBadRequestError("error when comparing passwords")
	}
	return user, nil
}

func (s *usersService) GetUserByID(userId int64) (*users.User, *errors.RestErrors) {
	user := users.User{
		ID: userId,
	}
	if err := user.GetUserByID(); err != nil {
		return nil, err
	}
	return &user, nil
}
