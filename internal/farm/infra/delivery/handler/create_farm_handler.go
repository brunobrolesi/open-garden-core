package handler

import (
	"net/http"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	createFarmHandler struct {
		CreateFarmUseCase usecase.CreateFarmUseCase
	}

	createFarmBodyRequest struct {
		Name    string `json:"name" validate:"required,min=3,max=150"`
		Address string `json:"address" validate:"required,min=3,max=255"`
	}
)

func NewCreateFarmHandler(u usecase.CreateFarmUseCase) Handler {
	return createFarmHandler{
		CreateFarmUseCase: u,
	}
}

func (h createFarmHandler) Handle(c *gin.Context) {
	var body createFarmBodyRequest

	userId, err := shared.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

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

	farm := usecase.CreateFarmInputDto{
		Name:    body.Name,
		Address: body.Address,
		Owner:   userId,
	}

	result, err := h.CreateFarmUseCase.Exec(c, farm)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": result,
	})
}
