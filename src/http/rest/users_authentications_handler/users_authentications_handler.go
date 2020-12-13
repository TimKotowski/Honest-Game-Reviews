package users_authentications_handler

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/middleware/jwt_auth"
	"Honest-Game-Reviews/src/services/users_service"
	"Honest-Game-Reviews/src/utils/json_utils"
	"encoding/json"
	"net/http"
)

type UsersAuthenticationsHandlerInterface interface {
	UserLogin(w http.ResponseWriter, r *http.Request)
}

type usersAuthenticationsHandler struct{}

func NewusersAuthenticationsHandler() UsersAuthenticationsHandlerInterface {
	return &usersAuthenticationsHandler{}
}

func (handler *usersAuthenticationsHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var user users.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid json body", err)
		return
	}
	// get that specific user once they login
	specificUser, userErr := users_service.NewUsersService.GetUser(user)
	if userErr != nil {
		json_utils.JsonErrorResponse(w, userErr)
		return
	}

	// once we get more of the users info from the database assign a JET to that user
	jwtString, jwtErr := jwt_auth.NewJWT(specificUser.Password, specificUser.ID)
	if jwtErr != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid json body", jwtErr)
		return
	}
	// return the JWT string back?
	json_utils.JsonResponse(w, http.StatusOK, jwtString)
}
