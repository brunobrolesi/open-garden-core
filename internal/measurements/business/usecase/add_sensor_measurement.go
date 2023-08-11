package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/gateway"
)

type AddSensorMeasurementInputDto struct {
	SensorId int
	Value    float64
}

type AddSensorMeasurementUseCase interface {
	Exec(ctx context.Context, input AddSensorMeasurementInputDto) error
}

type addSensorMeasurement struct {
	sensorMeasurementRepository gateway.SensorMeasurementRepository
}

func NewAddSensorMeasurementUseCase(sensorMeasurementRepository gateway.SensorMeasurementRepository) AddSensorMeasurementUseCase {
	return &addSensorMeasurement{
		sensorMeasurementRepository: sensorMeasurementRepository,
	}
}

func (g *addSensorMeasurement) Exec(ctx context.Context, input AddSensorMeasurementInputDto) error {
	return g.sensorMeasurementRepository.AddSensorMeasurement(ctx, input.SensorId, input.Value)
}
