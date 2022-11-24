package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	getUserFarmHandler struct {
		GetUserFarmUseCase usecase.GetUserFarmUseCase
	}
)

func NewGetUserFarmHandler(u usecase.GetUserFarmUseCase) Handler {
	return getUserFarmHandler{
		GetUserFarmUseCase: u,
	}
}

func (h getUserFarmHandler) Handle(c *gin.Context) {
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

	i := usecase.GetUserFarmInputDto{
		UserId: userId,
		FarmId: id,
	}
	result, err := h.GetUserFarmUseCase.Exec(i, c)

	if err != nil {
		fmt.Println("ERROR", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	if result.IdIsEmpty() {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
