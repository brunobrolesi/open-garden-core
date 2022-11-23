package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
)

type (
	GetUserFarmsUseCase interface {
		Exec(userId int, ctx context.Context) (model.Farms, error)
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

func (g getUserFarms) Exec(userId int, ctx context.Context) (model.Farms, error) {
	newFarm, err := g.FarmRepository.GetFarmsByUserId(userId, ctx)

	if err != nil {
		return model.Farms{}, err
	}

	return newFarm, nil
}
