package factory

import (
	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/handler"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/repository"
)

func GetUserFarmsFactory() handler.Handler {
	postgresConn := db.GetPostreSQLInstance()
	farmRepository := repository.NewPostgresFarmRepository(postgresConn)
	usecase := usecase.NewGetUserFarmsUseCase(farmRepository)
	handler := handler.NewGetUserFarmsHandler(usecase)

	return handler
}
