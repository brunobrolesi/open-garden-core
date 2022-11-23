package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
)

type FarmRepository interface {
	CreateFarm(model.Farm, context.Context) (model.Farm, error)
	GetFarmsByUserId(int, context.Context) (model.Farms, error)
}
