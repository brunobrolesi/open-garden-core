package model

import "errors"

var (
	ErrEmailInUse     = errors.New("email in use")
	ErrAuthentication = errors.New("invalid email or password")
)
