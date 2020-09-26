package access_token

import (
	"strings"
	"time"

	"github.com/wgarcia4190/bookstore_oauth_api/internal/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	switch true {
	case at.AccessToken == "":
		return errors.NewBadRequestError("invalid access token id")
	case at.UserId <= 0:
		return errors.NewBadRequestError("invalid user id")
	case at.ClientId <= 0:
		return errors.NewBadRequestError("invalid client id")
	case at.Expires <= 0:
		return errors.NewBadRequestError("invalid expiration time")
	case at.IsExpired():
		return errors.NewBadRequestError("token is expired")
	default:
		return nil
	}
}
