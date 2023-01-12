package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type FarmRepository interface {
	GetFarmByIdAndUserId(ctx context.Context, id int, userId int) (model.Farm, error)
}
