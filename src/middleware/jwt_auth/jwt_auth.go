package jwt_auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	jwt_secret = "jwt_secret"
)

var (
	JWTSecret = os.Getenv(jwt_secret)
)

func NewJWT(userPassword string, userID int64) (string, error) {
	issued := time.Now()
	expires := issued.Add(time.Minute * 5)

	// Create the claims.
	claims := &TokenClaims{
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
		return "", err
	}
	return signedToken, nil
}

// GetJWTSigningKey returns the JWT signing key.
// It is constructed using the member's hashed password and the application
// JWT secret.
func GetJWTSigningKey(jwtSecret, password string) []byte {
	return []byte(jwtSecret + password)
}
