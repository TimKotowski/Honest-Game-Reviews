package access_token

import (
	"Honest-Game-Reviews/src/domain/users"
	"Honest-Game-Reviews/src/logger"
	"Honest-Game-Reviews/src/services/users_service"
	"Honest-Game-Reviews/src/utils/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwt_secret = "jwt_secret"
)

var (
	JWTSecret = os.Getenv(jwt_secret)
)

type TokenDetails struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

type JWTAccessTokenData struct {
	JWTData string `json:"data"`
}

// Generate an access token
func CreateToken(userPassword string, userID int64) (string, *errors.RestErrors) {
	issued := time.Now()
	expires := issued.Add(time.Minute * 5)
	// Create the claims.
	claims := &TokenDetails{
		userID,
		jwt.StandardClaims{
			IssuedAt:  issued.Unix(),
			ExpiresAt: expires.Unix(),
		},
	}
	// Create and sign the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(GetJWTSigningKey(JWTSecret, userPassword))
	if err != nil {
		logger.Error("error when trying to get", err)
		return "", errors.NewBadRequestError("unable to get the signed token")
	}
	return signedToken, nil

}

func ExtractToken(r *http.Request) (string, *errors.RestErrors) {
	authHeader := strings.Split(r.Header.Get("Authorization"), " ")
	if authHeader[0] != "Bearer" && authHeader[0] != "Basic" {
		return "", errors.NewUnauthorizedError("unathorized user")
	}
	// return the extracted token
	return authHeader[1], nil
}

func GetUserFromJWT(headerToken string) (*users.User, *errors.RestErrors) {
	// verify user and get back a salt
	signedString, err := GetUserSigningKey(headerToken)
	if err != nil {
		return nil, err
	}

	// Parse the token.
	token, parseErr := jwt.ParseWithClaims(headerToken, &TokenDetails{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}

		return signedString, nil
	})
	if parseErr != nil {
		logger.Error("wrong method format for jwt", parseErr)
		return nil, errors.NewUnauthorizedError("wrong method")
	}

	claims, ok := token.Claims.(*TokenDetails)
	if !ok || !token.Valid {
		return nil, errors.NewUnauthorizedError("not a valid token or no Claims available")
	}

	// get memeber id from claims
	user, userErr := users_service.NewUsersService.GetUserByID(claims.UserID)
	if userErr != nil {
		return nil, err
	}

	return user, nil
}

func GetUserSigningKey(headerToken string) ([]byte, *errors.RestErrors) {
	// split token
	tokenParts := strings.Split(headerToken, " ")

	if len(tokenParts) != 3 {
		return nil, errors.NewUnauthorizedError("unable to get split the token into 3 parts for the full token request")
	}
	// Decode the token
	claimBytes, err := jwt.DecodeSegment(tokenParts[1])
	if err != nil {
		logger.Error("unable to decode token", err)
		return nil, errors.NewUnauthorizedError("unable to decode token")
	}

	// Decode the claim bytes into TokenClaims struct that match the json tag
	var claims TokenDetails
	if err := json.Unmarshal(claimBytes, &claims); err != nil {
		logger.Error("unable to decode claims into an object", err)
		return nil, errors.NewBadRequestError("unable to decode claims to an object")
	}

	fmt.Println("claims", claims)
	//make a route to get user by id
	user, userErr := users_service.NewUsersService.GetUserByID(claims.UserID)
	if err != nil {
		return nil, userErr
	}
	return GetJWTSigningKey(JWTSecret, user.Password), nil
}

// GetJWTSigningKey returns the JWT signing key.
// It is constructed using the member's hashed password and the application JWT secret.
// creating a salt with the jwt secret and concating the password to make the salt
func GetJWTSigningKey(jwtSecret, password string) []byte {
	return []byte(jwtSecret + password)
}
