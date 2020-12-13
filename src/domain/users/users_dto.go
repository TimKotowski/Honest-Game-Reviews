package users

import (
	"Honest-Game-Reviews/src/logger"
	"Honest-Game-Reviews/src/utils/errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"user_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
	IsAdmin     bool   `json:"admin_privledge"`
	Status      string `json:"status"`
}

type UserLoginRequest struct {
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) GetHash() error {
	// GenerateFromPassword returns the bcrypt hash of the password at the given cost.
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(hash)
	if err != nil {
		logger.Error("unable to hash password in gethash function", err)
	}
	return nil
}

func (user *User) Validate() *errors.RestErrors {
	// remove white space
	user.Username = strings.TrimSpace(strings.ToLower(user.Username))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Username == "" {
		return errors.NewBadRequestError("username empty")
	}
	if user.Email == "" {
		return errors.NewBadRequestError("email is empty")
	}
	if user.Password == "" {
		return errors.NewNotFoundError("invalid password")
	}
	return nil
}
