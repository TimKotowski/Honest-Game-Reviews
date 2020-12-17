package tokens_service

import (
	"Honest-Game-Reviews/src/domain/access_token"
	"Honest-Game-Reviews/src/utils/errors"
)

var (
	TokensService TokensServiceInterface = &tokensService{}
)

type TokensServiceInterface interface {
	CreateToken(string, int64) (string, *errors.RestErrors)
}

type tokensService struct{}

func (s *tokensService) CreateToken(userPassword string, userID int64) (string, *errors.RestErrors) {
	signedToken, err := access_token.CreateToken(userPassword, userID)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
