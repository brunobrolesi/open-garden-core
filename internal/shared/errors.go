package shared

import "errors"

var (
	ErrApiInternalServer = errors.New("internal server error")
	ErrApiUnauthorized   = errors.New("unauthorized")
	ErrApiBadRequest     = errors.New("bad request")
)
