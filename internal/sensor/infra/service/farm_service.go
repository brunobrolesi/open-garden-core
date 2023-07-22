package service

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/facade"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type farmService struct {
	farmFacade facade.FarmFacade
}

func NewFarmService(farmFacade facade.FarmFacade) gateway.FarmService {
	return &farmService{
		farmFacade: farmFacade,
	}
}

func (f *farmService) GetFarmByIdAndUserId(ctx context.Context, farmID int, userID int) (model.Farm, error) {
	farmFacadeOutput, err := f.farmFacade.GetUserFarm(ctx, userID, farmID)
	if err != nil {
		return model.Farm{}, err
	}

	return model.Farm{
		Id:      farmFacadeOutput.Id,
		Name:    farmFacadeOutput.Name,
		Owner:   farmFacadeOutput.Owner,
		Address: farmFacadeOutput.Address,
		Active:  farmFacadeOutput.Active,
	}, nil

}
