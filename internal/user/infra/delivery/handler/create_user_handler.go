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
	createUserHandler struct {
		CreateUserUseCase usecase.CreateUserUseCase
	}

	createUserBodyRequest struct {
		CompanyName     string `json:"company_name" validate:"required,min=3,max=150"`
		Email           string `json:"email" validate:"required,email,max=150"`
		Password        string `json:"password" validate:"required,min=6"`
		ConfirmPassword string `json:"confirm_password" validate:"required,min=6,eqfield=Password"`
	}
)

func NewCreateUserHandler(u usecase.CreateUserUseCase) Handler {
	return createUserHandler{
		CreateUserUseCase: u,
	}
}

func (h createUserHandler) Handle(c *gin.Context) {
	var body createUserBodyRequest

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

	user := usecase.CreateUserInputDto{
		CompanyName: body.CompanyName,
		Email:       body.Email,
		Password:    body.Password,
	}

	token, err := h.CreateUserUseCase.Exec(c, user)

	if errors.Is(err, model.ErrEmailInUse) {
		c.JSON(http.StatusBadRequest, gin.H{
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

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
