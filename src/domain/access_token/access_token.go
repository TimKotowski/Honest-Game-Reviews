package access_token

import (
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
// TODO: Improve the error handler of the token so the client dosnt get back sensitive into
func CreateToken(userPassword string, userID int64) (string, error) {
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
		return "", err
	}
	return signedToken, nil
}


func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// GetJWTSigningKey returns the JWT signing key.
// It is constructed using the member's hashed password and the application
// JWT secret.
func GetJWTSigningKey(jwtSecret, password string) []byte {
	return []byte(jwtSecret + password)
}


