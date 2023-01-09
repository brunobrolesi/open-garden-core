package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	createFarmHandler struct {
		CreateFarmUseCase usecase.CreateFarmUseCase
	}

	createFarmBodyRequest struct {
		Name    string `json:"name"`
		Address string `json:"address"`
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

	if err := body.Validate(); err != nil {
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

func (b *createFarmBodyRequest) Validate() error {
	if strings.TrimSpace(b.Name) == "" {
		return errors.New("name can't be empty")
	}

	if strings.TrimSpace(b.Address) == "" {
		return errors.New("address can't be empty")
	}

	return nil
}
