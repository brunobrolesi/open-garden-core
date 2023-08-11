package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	getSensorMeasurementsHandler struct {
		getSensorMeasurementsUseCase usecase.GetSensorPeriodMeasurementsUseCase
	}
)

func NewGetSensorMeasurementsHandler(u usecase.GetSensorPeriodMeasurementsUseCase) Handler {
	return getSensorMeasurementsHandler{
		getSensorMeasurementsUseCase: u,
	}
}

func (h getSensorMeasurementsHandler) Handle(c *gin.Context) {
	userId, err := shared.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	from, err := time.Parse(shared.DateParamLayout, c.Query("from"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	to, err := time.Parse(shared.DateParamLayout, c.Query("to"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

	i := usecase.GetSensorPeriodMeasurementsInputDto{
		SensorId: id,
		UserID:   userId,
		From:     from,
		To:       to,
	}
	result, err := h.getSensorMeasurementsUseCase.Exec(c, i)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No measurements found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
