package users_service

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/utils/date_utils"
	"Honest-Game-Reviews/src/utils/errors"
)



var (
	NewUsersService UsersServiceInterface = &usersService{}
	statusActive string = "active"
)

type UsersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErrors)
	GetUser(users.UserLoginRequest) (*users.User, *errors.RestErrors)
}

type usersService struct {}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErrors) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// hash password
	user.GetHash() // &userBody.GetHash() implicitly dose it
	user.Status = statusActive
	// set dateTime and format it right to be in mysql DB
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.SaveUser(); err != nil {
		return nil, err
	}
	return &user, nil
}



func (s *usersService) GetUser(UserLoginRequest users.UserLoginRequest) (*users.User, *errors.RestErrors) {
	return UserLoginRequest.GetUser()
}
