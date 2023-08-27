package handler

import (
	"net/http"
	"strconv"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	getFarmSensorHandler struct {
		getFarmSensorUseCase usecase.GetFarmSensorUseCase
	}
)

func NewGetFarmSensorHandler(u usecase.GetFarmSensorUseCase) Handler {
	return getFarmSensorHandler{
		getFarmSensorUseCase: u,
	}
}

func (h getFarmSensorHandler) Handle(c *gin.Context) {
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

	farmSensorId, err := strconv.Atoi(c.Param("farm_sensor_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	i := usecase.GetFarmSensorInputDto{
		UserId:       userId,
		FarmId:       farmId,
		FarmSensorId: farmSensorId,
	}
	result, err := h.getFarmSensorUseCase.Exec(c, i)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	if result.IsEmpty() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no sensor farm found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
