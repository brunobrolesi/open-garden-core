package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type FarmService interface {
	GetFarmByIdAndUserId(ctx context.Context, farmID int, userID int) (model.Farm, error)
}
