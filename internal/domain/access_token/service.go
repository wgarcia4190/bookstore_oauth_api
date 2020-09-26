package access_token

import (
	"strings"

	"github.com/wgarcia4190/bookstore_oauth_api/internal/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(*AccessToken, int64) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId, err := validateAccessToken(accessTokenId)

	if err != nil {
		return nil, err
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at *AccessToken, expires int64) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	err := s.repository.UpdateExpirationTime(at.AccessToken, expires)
	if err == nil {
		at.Expires = expires
	}

	return err
}

func validateAccessToken(accessToken string) (string, *errors.RestErr) {
	accessToken = strings.TrimSpace(accessToken)

	if len(accessToken) == 0 {
		return "", errors.NewBadRequestError("invalid access token id")
	}

	return accessToken, nil
}
