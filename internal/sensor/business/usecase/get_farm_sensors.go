package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type (
	GetFarmSensorsInputDto struct {
		UserId int
		FarmId int
	}

	GetFarmSensorsUseCase interface {
		Exec(ctx context.Context, input GetFarmSensorsInputDto) (model.FarmSensors, error)
	}

	getFarmSensors struct {
		farmService          gateway.FarmService
		farmSensorRepository gateway.FarmSensorRepository
	}
)

func NewGetFarmSensorsUseCase(farmSensorRepository gateway.FarmSensorRepository, farmService gateway.FarmService) GetFarmSensorsUseCase {
	return getFarmSensors{
		farmSensorRepository: farmSensorRepository,
		farmService:          farmService,
	}
}

func (g getFarmSensors) Exec(ctx context.Context, input GetFarmSensorsInputDto) (model.FarmSensors, error) {
	farm, err := g.farmService.GetFarmByIdAndUserId(ctx, input.FarmId, input.UserId)
	if err != nil {
		return model.FarmSensors{}, err
	}
	if farm.IsEmpty() {
		return model.FarmSensors{}, nil
	}

	farmSensors, err := g.farmSensorRepository.GetFarmSensorsByFarmId(ctx, input.FarmId)

	if err != nil {
		return model.FarmSensors{}, err
	}

	return farmSensors, nil
}
