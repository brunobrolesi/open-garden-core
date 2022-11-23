package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
)

type (
	CreateFarmInputDto struct {
		Name    string
		Owner   int
		Address string
	}

	CreateFarmUseCase interface {
		Exec(farm CreateFarmInputDto, ctx context.Context) (model.Farm, error)
	}

	createFarm struct {
		FarmRepository gateway.FarmRepository
	}
)

func NewCreateFarmUseCase(farmRepository gateway.FarmRepository) CreateFarmUseCase {
	return createFarm{
		FarmRepository: farmRepository,
	}
}

func (c createFarm) Exec(farm CreateFarmInputDto, ctx context.Context) (model.Farm, error) {
	f := model.Farm{
		Name:    farm.Name,
		Owner:   farm.Owner,
		Address: farm.Address,
		Active:  true,
	}
	newFarm, err := c.FarmRepository.CreateFarm(f, ctx)

	if err != nil {
		return model.Farm{}, err
	}

	return newFarm, nil
}
