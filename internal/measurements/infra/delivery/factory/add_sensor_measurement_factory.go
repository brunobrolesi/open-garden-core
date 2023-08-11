package factory

import (
	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/infra/delivery/handler"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/infra/repository"
)

func AddSensorMeasurementFactory() handler.Handler {
	timescaleConn := db.GetTimescaleInstance()
	sensorMeasurementRepository := repository.NewTimeScaleSensorMeasurementRepository(timescaleConn)
	usecase := usecase.NewAddSensorMeasurementUseCase(sensorMeasurementRepository)
	handler := handler.NewAddSensorMeasurementHandler(usecase)

	return handler
}
