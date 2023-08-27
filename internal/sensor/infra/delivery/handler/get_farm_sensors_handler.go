package handler

import (
	"net/http"
	"strconv"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	getUserFarmsHandler struct {
		getFarmSensorsUseCase usecase.GetFarmSensorsUseCase
	}
)

func NewGetFarmSensorsHandler(u usecase.GetFarmSensorsUseCase) Handler {
	return getUserFarmsHandler{
		getFarmSensorsUseCase: u,
	}
}

func (h getUserFarmsHandler) Handle(c *gin.Context) {
	userId, err := shared.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	farmId, err := strconv.Atoi(c.Param("farm_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	input := usecase.GetFarmSensorsInputDto{
		UserId: userId,
		FarmId: farmId,
	}

	result, err := h.getFarmSensorsUseCase.Exec(c, input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no sensors found for this farm",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
