package rest

import (
	"encoding/json"
	"time"

	"github.com/wgarcia4190/go-rest/gorest"

	"github.com/wgarcia4190/bookstore_oauth_api/internal/domain/users"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/utils/errors"
)

var (
	usersRestClient = gorest.NewBuilder().
		SetBaseUrl("http://localhost:8081").
		SetConnectionTimeout(100 * time.Millisecond).
		Build()
)

type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response, err := usersRestClient.Post("/users/login", request)

	if err != nil || response == nil {
		return nil, errors.NewInternalServerError("invalid client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Body, &restErr)

		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response", errors.NewError("json parsing error"))
	}
	return &user, nil
}
