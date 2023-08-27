package model

import "reflect"

type FarmSensor struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SensorModel int    `json:"sensor_model"`
	FarmId      int    `json:"farm_id"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type FarmSensors []FarmSensor

func (f FarmSensor) IsEmpty() bool {
	return reflect.DeepEqual(f, FarmSensor{})
}
