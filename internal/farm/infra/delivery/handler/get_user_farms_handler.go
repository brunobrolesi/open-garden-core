package handler

import (
	"net/http"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	getUserFarmsHandler struct {
		getUserFarmsUseCase usecase.GetUserFarmsUseCase
	}
)

func NewGetUserFarmsHandler(u usecase.GetUserFarmsUseCase) Handler {
	return getUserFarmsHandler{
		getUserFarmsUseCase: u,
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

	result, err := h.getUserFarmsUseCase.Exec(c, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": shared.ErrApiInternalServer.Error(),
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
