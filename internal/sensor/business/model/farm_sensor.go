package model

type FarmSensor struct {
	Id          int    `json:"id"`
	SensorId    int    `json:"sensor_id"`
	FarmId      int    `json:"farm_id"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
