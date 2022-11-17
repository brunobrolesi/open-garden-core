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
	createUserHandler struct {
		CreateUserUseCase usecase.CreateUserUseCase
	}

	createUserBodyRequest struct {
		CompanyName     string `json:"company_name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
)

func NewCreateUserHandler(u usecase.CreateUserUseCase) Handler {
	return createUserHandler{
		CreateUserUseCase: u,
	}
}

func (h createUserHandler) Handle(c *gin.Context) {
	var body createUserBodyRequest

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

	user := usecase.CreateUserInputDto{
		CompanyName: body.CompanyName,
		Email:       body.Email,
		Password:    body.Password,
	}

	token, err := h.CreateUserUseCase.Exec(user, c)

	if errors.Is(err, model.ErrEmailInUse) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func (b *createUserBodyRequest) Validate() error {
	if strings.TrimSpace(b.CompanyName) == "" {
		return errors.New("company_name can't be empty")
	}

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

	if strings.TrimSpace(b.ConfirmPassword) == "" {
		return errors.New("confirm_password can't be empty")
	}

	if strings.Compare(b.Password, b.ConfirmPassword) != 0 {
		return errors.New("passwords must be equal")
	}

	return nil
}
