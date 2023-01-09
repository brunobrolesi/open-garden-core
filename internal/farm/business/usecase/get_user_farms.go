package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
)

type (
	GetUserFarmsUseCase interface {
		Exec(ctx context.Context, userId int) (model.Farms, error)
	}

	getUserFarms struct {
		FarmRepository gateway.FarmRepository
	}
)

func NewGetUserFarmsUseCase(farmRepository gateway.FarmRepository) GetUserFarmsUseCase {
	return getUserFarms{
		FarmRepository: farmRepository,
	}
}

func (g getUserFarms) Exec(ctx context.Context, userId int) (model.Farms, error) {
	farms, err := g.FarmRepository.GetFarmsByUserId(ctx, userId)

	if err != nil {
		return model.Farms{}, err
	}

	return farms, nil
}
