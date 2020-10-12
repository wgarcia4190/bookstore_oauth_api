package rest

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wgarcia4190/go-rest/gorest_mock"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting mock")
	gorest_mock.MockupServer.Start()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	gorest_mock.MockupServer.DeleteMocks()
	gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
		Method:      http.MethodPost,
		Url:         "https://api.bookstore.com/users/login",
		RequestBody: `{"email":"email@gmail.com","password":"password"}`,
		Error:       errors.New("invalid client response when trying to login user"),
	})

	repository := NewRepository()

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid client response when trying to login user", err.Message)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	gorest_mock.MockupServer.DeleteMocks()
	gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
		Method:             http.MethodPost,
		Url:                "https://api.bookstore.com/users/login",
		RequestBody:        `{"email":"email@gmail.com","password":"password"}`,
		ResponseStatusCode: http.StatusNotFound,
		ResponseBody:       `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	repository := NewRepository()

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	fmt.Println(err.Message)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	gorest_mock.MockupServer.DeleteMocks()
	gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
		Method:             http.MethodPost,
		Url:                "https://api.bookstore.com/users/login",
		RequestBody:        `{"email":"email@gmail.com","password":"password"}`,
		ResponseStatusCode: http.StatusNotFound,
		ResponseBody:       `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repository := NewRepository()

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	gorest_mock.MockupServer.DeleteMocks()
	gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
		Method:             http.MethodPost,
		Url:                "https://api.bookstore.com/users/login",
		RequestBody:        `{"email":"email@gmail.com","password":"password"}`,
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"id": "1", "first_name": "John", "last_name": "Doe", "email": "johndoe@gmail.com"}`,
	})

	repository := NewRepository()

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	gorest_mock.MockupServer.DeleteMocks()
	gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
		Method:             http.MethodPost,
		Url:                "https://api.bookstore.com/users/login",
		RequestBody:        `{"email":"email@gmail.com","password":"password"}`,
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"id": 1, "first_name": "John", "last_name": "Doe", "email": "johndoe@gmail.com"}`,
	})

	repository := NewRepository()

	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, user.ID, 1)
	assert.EqualValues(t, user.FirstName, "John")
	assert.EqualValues(t, user.LastName, "Doe")
	assert.EqualValues(t, user.Email, "johndoe@gmail.com")
	fmt.Println(user)
}
