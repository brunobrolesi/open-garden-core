package handler

import (
	"net/http"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	addSensorMeasurementHandler struct {
		addSensorMeasurementUseCase usecase.AddSensorMeasurementUseCase
	}

	addSensorMeasurementBodyRequest struct {
		SensorID int     `json:"sensor_id" validate:"required,numeric"`
		Value    float64 `json:"value" validate:"required,numeric"`
	}
)

func NewAddSensorMeasurementHandler(u usecase.AddSensorMeasurementUseCase) Handler {
	return addSensorMeasurementHandler{
		addSensorMeasurementUseCase: u,
	}
}

func (h addSensorMeasurementHandler) Handle(c *gin.Context) {
	var body addSensorMeasurementBodyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	i := usecase.AddSensorMeasurementInputDto{
		SensorId: body.SensorID,
		Value:    body.Value,
	}

	err := h.addSensorMeasurementUseCase.Exec(c, i)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
