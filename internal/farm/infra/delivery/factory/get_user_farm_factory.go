package factory

import (
	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/handler"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/repository"
)

func GetUserFarmFactory() handler.Handler {
	postgresConn := db.GetInstance()
	farmRepository := repository.NewPostgresFarmRepository(postgresConn)
	usecase := usecase.NewGetUserFarmUseCase(farmRepository)
	handler := handler.NewGetUserFarmHandler(usecase)

	return handler
}