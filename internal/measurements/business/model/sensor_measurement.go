package model

import "time"

type SensorMeasurement struct {
	SensorID int       `json:"sensor_id"`
	Time     time.Time `json:"time"`
	Value    float64   `json:"value"`
}

type SensorMeasurements []SensorMeasurement
