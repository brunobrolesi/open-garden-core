package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
)

type (
	GetUserFarmInputDto struct {
		UserId int
		FarmId int
	}

	GetUserFarmUseCase interface {
		Exec(input GetUserFarmInputDto, ctx context.Context) (model.Farm, error)
	}

	getUserFarm struct {
		FarmRepository gateway.FarmRepository
	}
)

func NewGetUserFarmUseCase(farmRepository gateway.FarmRepository) GetUserFarmUseCase {
	return getUserFarm{
		FarmRepository: farmRepository,
	}
}

func (g getUserFarm) Exec(input GetUserFarmInputDto, ctx context.Context) (model.Farm, error) {
	farm, err := g.FarmRepository.GetFarmByFarmIdAndUserId(input.FarmId, input.UserId, ctx)

	if err != nil {
		return model.Farm{}, err
	}

	return farm, nil
}
