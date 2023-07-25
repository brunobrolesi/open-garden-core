package handler

import (
	"errors"
	"net/http"

	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/usecase"
	"github.com/gin-gonic/gin"
)

type (
	authenticateUserHandler struct {
		authenticateUserUseCase usecase.AuthenticateUserUseCase
	}

	authenticateUserBodyRequest struct {
		Email    string `json:"email" validate:"required,email,max=150"`
		Password string `json:"password" validate:"required,min=6"`
	}
)

func NewAuthenticateUserHandler(u usecase.AuthenticateUserUseCase) Handler {
	return authenticateUserHandler{
		authenticateUserUseCase: u,
	}
}

func (h authenticateUserHandler) Handle(c *gin.Context) {
	var body authenticateUserBodyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	validator := shared.GetValidator()
	if err := validator.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	credentials := usecase.AuthenticateUserInputDto{
		Email:    body.Email,
		Password: body.Password,
	}

	token, err := h.authenticateUserUseCase.Exec(c, credentials)

	if errors.Is(err, model.ErrAuthentication) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
