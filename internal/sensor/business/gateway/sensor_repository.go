package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

type SensorRepository interface {
	GetSensorById(ctx context.Context, id int) (model.Sensor, error)
}
