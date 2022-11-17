package model

import "errors"

var (
	ErrEmailInUse     = errors.New("email in use")
	ErrInternalServer = errors.New("internal server error")
	ErrAuthentication = errors.New("invalid email or password")
)
