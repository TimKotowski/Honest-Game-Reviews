package users_service

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/utils/date_utils"
	"Honest-Game-Reviews/src/utils/errors"
)



var (
	NewUsersService UsersServiceInterface = &usersService{}
	statusActive = "active"
)

type UsersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErrors)
}

type usersService struct {}


func (s *usersService) CreateUser(userBody users.User) (*users.User, *errors.RestErrors) {
	// hash password
	userBody.GetHash()
	userBody.Status = statusActive
	userBody.DateCreated = date_utils.GetNowDBFormat()
	if err := userBody.SaveUser(); err != nil {
		return nil, err
	}
	return &userBody, nil
}
