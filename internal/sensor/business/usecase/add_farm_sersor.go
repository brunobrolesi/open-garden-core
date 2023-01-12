package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type (
	AddFarmSensorInputDto struct {
		SensorId    int
		Description string
		FarmId      int
		UserId      int
	}

	AddFarmSensorUseCase interface {
		Exec(ctx context.Context, farm AddFarmSensorInputDto) (model.FarmSensor, error)
	}

	addFarmSensor struct {
		SensorRepository     gateway.SensorRepository
		FarmRepository       gateway.FarmRepository
		FarmSensorRepository gateway.FarmSensorRepository
	}
)

func NewAddFarmSensorUseCase(sensorRepository gateway.SensorRepository, farmRepository gateway.FarmRepository, farmSensorRepository gateway.FarmSensorRepository) AddFarmSensorUseCase {
	return addFarmSensor{
		SensorRepository:     sensorRepository,
		FarmRepository:       farmRepository,
		FarmSensorRepository: farmSensorRepository,
	}
}

func (a addFarmSensor) Exec(ctx context.Context, input AddFarmSensorInputDto) (model.FarmSensor, error) {
	s, err := a.SensorRepository.GetSensorById(ctx, input.SensorId)

	if err != nil {
		return model.FarmSensor{}, err
	}

	f, err := a.FarmRepository.GetFarmByIdAndUserId(ctx, input.FarmId, input.UserId)

	if err != nil {
		return model.FarmSensor{}, err
	}

	fs := model.FarmSensor{
		SensorId:    s.Id,
		FarmId:      f.Id,
		Description: input.Description,
		Active:      true,
	}

	newFarmSensor, err := a.FarmSensorRepository.CreateFarmSensor(ctx, fs)

	if err != nil {
		return model.FarmSensor{}, err
	}

	return newFarmSensor, nil
}