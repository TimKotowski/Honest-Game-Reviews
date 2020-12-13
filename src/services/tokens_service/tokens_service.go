package tokens_service

import (
	"Honest-Game-Reviews/src/domain/access_token"
)



var (
	TokensService TokensServiceInterface = &tokensService{}
)

type TokensServiceInterface interface {
	CreateToken(string, int64) (*access_token.JWTAccessTokenData, error)
}

type tokensService struct{}


func (s *tokensService) CreateToken(userPassword string, userID int64) (*access_token.JWTAccessTokenData, error) {
	signedToken, err := access_token.CreateToken(userPassword, userID)
	if err != nil {
		return nil, err
	}
	// store the jwt signedToken in a struct to send off to the handler ass jwtData
	jwtData := access_token.JWTAccessTokenData{
		JWTData: signedToken,
	}
	return &jwtData, nil
}
