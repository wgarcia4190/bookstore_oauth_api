package access_token_test

import (
	"testing"
	"time"

	"github.com/wgarcia4190/bookstore_oauth_api/internal/domain/access_token"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, access_token.ExpirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := access_token.GetNewAccessToken(123)

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined token id")
	assert.EqualValues(t, 0, at.UserId, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := access_token.AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
