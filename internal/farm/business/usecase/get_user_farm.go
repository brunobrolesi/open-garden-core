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
		Exec(ctx context.Context, input GetUserFarmInputDto) (model.Farm, error)
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

func (g getUserFarm) Exec(ctx context.Context, input GetUserFarmInputDto) (model.Farm, error) {
	farm, err := g.FarmRepository.GetFarmByIdAndUserId(ctx, input.FarmId, input.UserId)

	if err != nil {
		return model.Farm{}, err
	}

	return farm, nil
}