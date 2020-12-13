package jwt_auth

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}
