package model

import "errors"

var (
	ErrInvalidSensor = errors.New("invalid sensor")
	ErrInvalidFarm   = errors.New("invalid farm")
)
