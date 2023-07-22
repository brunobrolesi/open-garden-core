package facade

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/facade/dto"
)

type FarmFacade interface {
	GetUserFarm(ctx context.Context, userID int, farmID int) (dto.GetUserFarmOutputDto, error)
}

type farmFacade struct {
	GetUserFarmUseCase usecase.GetUserFarmUseCase
}

func newFarmFacade(getUserFarmUseCase usecase.GetUserFarmUseCase) FarmFacade {
	return &farmFacade{
		GetUserFarmUseCase: getUserFarmUseCase,
	}
}

func (f *farmFacade) GetUserFarm(ctx context.Context, userID int, farmID int) (dto.GetUserFarmOutputDto, error) {
	usecaseInput := usecase.GetUserFarmInputDto{
		UserId: userID,
		FarmId: farmID,
	}
	farm, err := f.GetUserFarmUseCase.Exec(ctx, usecaseInput)
	if err != nil {
		return dto.GetUserFarmOutputDto{}, err
	}

	return dto.GetUserFarmOutputDto{
		Id:      farm.Id,
		Name:    farm.Name,
		Owner:   farm.Owner,
		Address: farm.Address,
		Active:  farm.Active,
	}, nil
}
