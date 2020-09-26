package access_token

import "github.com/wgarcia4190/bookstore_oauth_api/internal/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(string, int64) *errors.RestErr
}
