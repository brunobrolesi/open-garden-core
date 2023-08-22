package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type (
	AddFarmSensorInputDto struct {
		Name        string
		SensorModel int
		Description string
		FarmId      int
		UserId      int
	}

	AddFarmSensorUseCase interface {
		Exec(ctx context.Context, farm AddFarmSensorInputDto) (model.FarmSensor, error)
	}

	addFarmSensor struct {
		sensorRepository     gateway.SensorRepository
		farmService          gateway.FarmService
		farmSensorRepository gateway.FarmSensorRepository
	}
)

func NewAddFarmSensorUseCase(sensorRepository gateway.SensorRepository, farmService gateway.FarmService, farmSensorRepository gateway.FarmSensorRepository) AddFarmSensorUseCase {
	return addFarmSensor{
		sensorRepository:     sensorRepository,
		farmService:          farmService,
		farmSensorRepository: farmSensorRepository,
	}
}

func (a addFarmSensor) Exec(ctx context.Context, input AddFarmSensorInputDto) (model.FarmSensor, error) {
	s, err := a.sensorRepository.GetSensorById(ctx, input.SensorModel)

	if err != nil {
		return model.FarmSensor{}, err
	}

	if s.IsEmpty() {
		return model.FarmSensor{}, model.ErrInvalidSensor
	}

	f, err := a.farmService.GetFarmByIdAndUserId(ctx, input.FarmId, input.UserId)

	if err != nil {
		return model.FarmSensor{}, err
	}

	if f.IsEmpty() {
		return model.FarmSensor{}, model.ErrInvalidFarm
	}

	fs := model.FarmSensor{
		Name:        input.Name,
		SensorModel: s.Id,
		FarmId:      f.Id,
		Description: input.Description,
		Active:      true,
	}

	newFarmSensor, err := a.farmSensorRepository.CreateFarmSensor(ctx, fs)

	if err != nil {
		return model.FarmSensor{}, err
	}

	return newFarmSensor, nil
}
