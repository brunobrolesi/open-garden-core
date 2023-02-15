package model

import "reflect"

type Farm struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   int    `json:"owner"`
	Address string `json:"address"`
	Active  bool   `json:"active"`
}

func (f Farm) IsEmpty() bool {
	return reflect.DeepEqual(f, Farm{})
}
