package shared

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (int, error) {
	stringUserId := c.Request.Header.Get("X-User-Id")
	userId, err := strconv.Atoi(stringUserId)

	if err != nil {
		return 0, err
	}
	if userId <= 0 {
		return 0, ErrApiBadRequest
	}

	return userId, nil
}
