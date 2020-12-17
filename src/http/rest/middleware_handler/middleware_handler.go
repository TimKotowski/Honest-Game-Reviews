package middleware_handler

import (
	"Honest-Game-Reviews/src/domain/access_token"
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/utils/errors"
	"Honest-Game-Reviews/src/utils/json_utils"
	"context"
	"fmt"
	"net/http"
)

// key is the key type used by this package for the request context.
type key int

// AuthKey is the key used for storing and retrieving the member data from the
// request context.
var AuthKey key = 1

func GetUserFromRequest(r *http.Request) (*users.User, *errors.RestErrors) {
	user, ok := r.Context().Value(AuthKey).(*users.User)
	if !ok {
		return nil, errors.NewUnauthorizedError("could not type assert authenicated member")
	}
	fmt.Println("r context in get use from request", user)
	return user, nil
}

func AuthenticateEndpoint(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, tokenStrErr := access_token.ExtractToken(r)
		if tokenStrErr != nil {
			json_utils.JsonErrorResponse(w, tokenStrErr)
		}

		user, userErr := access_token.GetUserFromJWT(tokenString)
		if userErr != nil {
			json_utils.JsonErrorResponse(w, userErr)
			return
		}
		ctx := context.WithValue(r.Context(), AuthKey, user)
		fmt.Println("ctx", ctx.Value(AuthKey))
		h(w, r.WithContext(ctx))
	}
}
