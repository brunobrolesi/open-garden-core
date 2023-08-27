package factory

import (
	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/facade"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/delivery/handler"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/repository"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/service"
)

func GetFarmSensorsFactory() handler.Handler {
	postgresConn := db.GetPostreSQLInstance()
	farmFacade := facade.NewFarmFacade()
	farmService := service.NewFarmService(farmFacade)
	farmSensorRepository := repository.NewPostgresFarmSensorRepository(postgresConn)
	usecase := usecase.NewGetFarmSensorsUseCase(farmSensorRepository, farmService)
	handler := handler.NewGetFarmSensorsHandler(usecase)

	return handler
}
