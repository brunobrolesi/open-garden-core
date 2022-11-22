package shared

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	POSTGRESQL_UNIQUE_VIOLATION_CODE = "23505"
)

func IsPostgreSqlError(err error, expectedCode string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == expectedCode {
			return true
		}
	}

	return false
}
