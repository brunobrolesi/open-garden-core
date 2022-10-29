package model

import "errors"

var (
	ErrEmailInUse     = errors.New("email in use")
	ErrInternalServer = errors.New("internal server error")
)
