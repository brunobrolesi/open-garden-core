package factory

import (
	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/delivery/handler"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/repository"
)

func AddFarmSensorFactory() handler.Handler {
	postgresConn := db.GetInstance()
	farmRepository := repository.NewPostgresFarmRepository(postgresConn)
	sensorRepository := repository.NewPostgresSensorRepository(postgresConn)
	farmSensorRepository := repository.NewPostgresFarmSensorRepository(postgresConn)
	usecase := usecase.NewAddFarmSensorUseCase(sensorRepository, farmRepository, farmSensorRepository)
	handler := handler.NewAddFarmSensorHandler(usecase)

	return handler
}
