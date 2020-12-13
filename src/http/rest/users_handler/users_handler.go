package users_handler

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/services/users_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"encoding/json"
	"net/http"
)


type UsersHandler interface {
	CreateUser(http.ResponseWriter, *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type usersHandler struct {}

func NewUsersHandler() UsersHandler {
	return &usersHandler{}
}

func (handler *usersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// create a user and genreate a hashed password
	// get a body of the request
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid json body", err)
		return
	}

	createdUser, userErr := users_service.NewUsersService.CreateUser(user)
	if userErr != nil {
		json_utils.JsonErrorResponse(w, userErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, createdUser)
}


func (handler *usersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var user users.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid json body", err)
		return
	}
	// get that user by email service
	foundUser, userErr := users_service.NewUsersService.GetUser(user)
	if userErr != nil {
		json_utils.JsonErrorResponse(w, userErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, foundUser)

}

