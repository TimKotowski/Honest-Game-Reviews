package auth_handler

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/services/tokens_service"
	"Honest-Game-Reviews/src/services/users_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"encoding/json"
	"net/http"
)

type AuthHandlerInterface interface {
	UserLogin(w http.ResponseWriter, r *http.Request)
}

type authHandler struct{}

func NewAuthHandler() AuthHandlerInterface {
	return &authHandler{}
}

func (handler *authHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var userCredentials users.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid json body", err)
		return
	}

	// get that specific user once they login
	user, userErr := users_service.NewUsersService.UserLogin(userCredentials)
	if userErr != nil {
		json_utils.JsonErrorResponse(w, userErr)
		return
	}
	// once we get more of the users info from the database, call the service and get the signedToken
	signedToken, jwtErr := tokens_service.TokensService.CreateToken(user.Password, user.ID)
	if jwtErr != nil {
		json_utils.JsonErrorResponse(w, jwtErr)
		return
	}
	json_utils.JsonResponse(w, http.StatusOK, signedToken)
}
