package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type (
	GetFarmSensorInputDto struct {
		UserId       int
		FarmId       int
		FarmSensorId int
	}

	GetFarmSensorUseCase interface {
		Exec(ctx context.Context, input GetFarmSensorInputDto) (model.FarmSensor, error)
	}

	getFarmSensor struct {
		farmService          gateway.FarmService
		farmSensorRepository gateway.FarmSensorRepository
	}
)

func NewGetFarmSensorUseCase(farmSensorRepository gateway.FarmSensorRepository, farmService gateway.FarmService) GetFarmSensorUseCase {
	return getFarmSensor{
		farmSensorRepository: farmSensorRepository,
		farmService:          farmService,
	}
}

func (g getFarmSensor) Exec(ctx context.Context, input GetFarmSensorInputDto) (model.FarmSensor, error) {
	farm, err := g.farmService.GetFarmByIdAndUserId(ctx, input.FarmId, input.UserId)
	if err != nil {
		return model.FarmSensor{}, err
	}
	if farm.IsEmpty() {
		return model.FarmSensor{}, nil
	}

	farmSensor, err := g.farmSensorRepository.GetFarmSensorById(ctx, input.FarmSensorId)

	if err != nil {
		return model.FarmSensor{}, err
	}

	return farmSensor, nil
}
