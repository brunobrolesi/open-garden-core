package handler

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/usecase"
	"github.com/gin-gonic/gin"
)

type (
	authenticateUserHandler struct {
		AuthenticateUserUseCase usecase.AuthenticateUserUseCase
	}

	authenticateUserBodyRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewAuthenticateUserHandler(u usecase.AuthenticateUserUseCase) Handler {
	return authenticateUserHandler{
		AuthenticateUserUseCase: u,
	}
}

func (h authenticateUserHandler) Handle(c *gin.Context) {
	var body authenticateUserBodyRequest

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

	credentials := usecase.Credentials{
		Email:    body.Email,
		Password: body.Password,
	}

	token, err := h.AuthenticateUserUseCase.AuthenticateUser(credentials, c)

	if errors.Is(err, model.ErrAuthentication) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": model.ErrInternalServer.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (b *authenticateUserBodyRequest) Validate() error {
	if strings.TrimSpace(b.Email) == "" {
		return errors.New("email can't be empty")
	}

	emailRegex := regexp.MustCompile("^[\\w!#$%&’*+/=?`{|}~^-]+(?:\\.[\\w!#$%&’*+/=?`{|}~^-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,6}$")

	if !emailRegex.MatchString(b.Email) {
		return errors.New("email must be a valid format")
	}

	if strings.TrimSpace(b.Password) == "" {
		return errors.New("password can't be empty")
	}

	return nil
}
