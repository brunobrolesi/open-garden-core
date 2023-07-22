package dto

type GetUserFarmOutputDto struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   int    `json:"owner"`
	Address string `json:"address"`
	Active  bool   `json:"active"`
}
