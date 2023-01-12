package model

type (
	SensorType string
)

const (
	Temperature  SensorType = "temperature"
	Conductivity SensorType = "conductivity"
	Ph           SensorType = "pH"
	WaterLevel   SensorType = "water-level"
)

type Sensor struct {
	Id   int
	Name string
	Type SensorType
	Unit string
}
