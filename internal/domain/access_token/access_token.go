package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/wgarcia4190/bookstore_oauth_api/internal/utils/crypto_utils"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/utils/errors"
)

const (
	ExpirationTime       = 24
	GrantTypePassword    = "password"
	GrandTypeCredentials = "credentials"
)

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case GrantTypePassword:
		break
	case GrandTypeCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	return nil
}

type AccessToken struct {
	AccessToken string `json:"token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(ExpirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Validate() *errors.RestErr {
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

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
