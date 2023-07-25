package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	addFarmSensorHandler struct {
		addFarmSensorUseCase usecase.AddFarmSensorUseCase
	}

	addFarmSensorBodyRequest struct {
		Name        string `json:"name" validate:"required,min=3,max=50"`
		SensorModel int    `json:"sensor_model" validate:"required,numeric,min=1"`
		Description string `json:"description" validate:"max=150"`
	}
)

func NewAddFarmSensorHandler(u usecase.AddFarmSensorUseCase) Handler {
	return addFarmSensorHandler{
		addFarmSensorUseCase: u,
	}
}

func (h addFarmSensorHandler) Handle(c *gin.Context) {
	var body addFarmSensorBodyRequest

	farmId, err := strconv.Atoi(c.Param("farm_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
		})
		return
	}

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

	farmSensor := usecase.AddFarmSensorInputDto{
		FarmId:      farmId,
		Name:        body.Name,
		SensorModel: body.SensorModel,
		Description: body.Description,
		UserId:      userId,
	}

	result, err := h.addFarmSensorUseCase.Exec(c, farmSensor)

	if errors.Is(err, model.ErrInvalidFarm) || errors.Is(err, model.ErrInvalidSensor) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": shared.ErrApiBadRequest.Error(),
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
		"data": result,
	})
}
