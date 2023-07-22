package facade

import (
	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/repository"
)

func NewFarmFacade() FarmFacade {
	postgresConn := db.GetInstance()
	farmRepository := repository.NewPostgresFarmRepository(postgresConn)
	uc := usecase.NewGetUserFarmUseCase(farmRepository)
	return newFarmFacade(uc)
}
