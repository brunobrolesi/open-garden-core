package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
)

type FarmRepository interface {
	CreateFarm(context.Context, model.Farm) (model.Farm, error)
	GetFarmsByUserId(context.Context, int) (model.Farms, error)
	GetFarmByIdAndUserId(ctx context.Context, farmId int, userId int) (model.Farm, error)
}
