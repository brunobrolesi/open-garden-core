package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type FarmSensorRepository interface {
	CreateFarmSensor(context.Context, model.FarmSensor) (model.FarmSensor, error)
}