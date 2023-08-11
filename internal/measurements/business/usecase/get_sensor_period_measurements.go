package usecase

import (
	"context"
	"time"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/model"
)

type GetSensorPeriodMeasurementsInputDto struct {
	SensorId int
	UserID   int
	From     time.Time
	To       time.Time
}

type GetSensorPeriodMeasurementsUseCase interface {
	Exec(ctx context.Context, input GetSensorPeriodMeasurementsInputDto) (model.SensorMeasurements, error)
}

type getSensorPeriodMeasurements struct {
	sensorMeasurementRepository gateway.SensorMeasurementRepository
}

func NewGetSensorPeriodMeasurementsUseCase(sensorMeasurementRepository gateway.SensorMeasurementRepository) GetSensorPeriodMeasurementsUseCase {
	return &getSensorPeriodMeasurements{
		sensorMeasurementRepository: sensorMeasurementRepository,
	}
}

func (g *getSensorPeriodMeasurements) Exec(ctx context.Context, input GetSensorPeriodMeasurementsInputDto) (model.SensorMeasurements, error) {
	return g.sensorMeasurementRepository.GetSensorPeriodMeasurements(ctx, input.SensorId, input.UserID, input.From, input.To)
}
