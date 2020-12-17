package users_handler

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/services/users_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type UsersHandler interface {
	CreateUserAccount(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)
}

type usersHandler struct{}

func NewUsersHandler() UsersHandler {
	return &usersHandler{}
}

func (handler *usersHandler) CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	// create a user and genreate a hashed password
	// get a body of the request
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid json body", err)
		return
	}

	_, userErr := users_service.NewUsersService.CreateUserAccount(user)
	if userErr != nil {
		json_utils.JsonErrorResponse(w, userErr)
		return
	}

	createdUserMessage := make(map[string]string)
	createdUserMessage["success"] = "user created"
	json_utils.JsonResponse(w, http.StatusOK, createdUserMessage)
}

func (handler *usersHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid url paramter", err)
		return
	}

	user, userErr := users_service.NewUsersService.GetUserByID(userID)
	if userErr != nil {
		json_utils.JsonErrorResponse(w, userErr)
	}
	json_utils.JsonResponse(w, http.StatusOK, user)
}
