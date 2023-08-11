package factory

import (
	"log"
	"os"
	"strconv"

	"github.com/brunobrolesi/open-garden-core/db"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/user/infra/delivery/handler"
	"github.com/brunobrolesi/open-garden-core/internal/user/infra/repository"
	"github.com/brunobrolesi/open-garden-core/internal/user/infra/service"
)

func AuthenticateUserFactory() handler.Handler {
	hashCost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		log.Fatal("fail get hash cost")
	}
	hashService := service.NewBcryptHashService(hashCost)
	postgresConn := db.GetPostreSQLInstance()
	userRepository := repository.NewPostgresUserRepository(postgresConn)
	tokenService := service.NewJwtTokenService(os.Getenv("JWT_SECRET"))
	usecase := usecase.NewAuthenticateUserUseCase(hashService, userRepository, tokenService)
	handler := handler.NewAuthenticateUserHandler(usecase)

	return handler
}
