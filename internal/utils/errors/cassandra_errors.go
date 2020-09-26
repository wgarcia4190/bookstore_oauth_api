package errors

import "github.com/gocql/gocql"

func ParseError(err error) *RestErr {
	switch true {
	case err == gocql.ErrNotFound:
		return NewNotFoundError("no records found")
	default:
		return NewInternalServerError("internal database error")
	}
}
