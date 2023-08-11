package gateway

import (
	"context"
	"time"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/model"
)

type SensorMeasurementRepository interface {
	GetSensorPeriodMeasurements(ctx context.Context, sensorID int, userID int, from time.Time, to time.Time) (model.SensorMeasurements, error)
	AddSensorMeasurement(ctx context.Context, sensorID int, value float64) error
}
