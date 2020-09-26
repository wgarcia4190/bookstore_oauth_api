package db

import (
	"github.com/gocql/gocql"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/domain/access_token"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/utils/errors"
)

const (
	queryGetAccessToken = `
		SELECT
			access_token,
			user_id,
			client_id,
			expires
		FROM
			access_tokens
		WHERE
			access_token = ?;
	`
	queryCreateAccessToken = `
		INSERT INTO access_tokens (
			access_token,
			user_id,
			client_id,
			expires
		) VALUES (
			?, ?, ?, ?
		);
	`
	queryUpdateExpires = `
		UPDATE
			access_tokens
		SET
			expires = ?
		WHERE
			access_token = ?;
	`
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(string, int64) *errors.RestErr
}

type dbRepository struct {
	session *gocql.Session
}

func NewRepository(session *gocql.Session) DBRepository {
	return &dbRepository{
		session: session,
	}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	var at access_token.AccessToken
	if err := r.session.Query(queryGetAccessToken, id).Scan(
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	); err != nil {
		return nil, errors.ParseError(err)
	}

	return &at, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := r.session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at string, expires int64) *errors.RestErr {
	if err := r.session.Query(queryUpdateExpires,
		expires,
		at,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
