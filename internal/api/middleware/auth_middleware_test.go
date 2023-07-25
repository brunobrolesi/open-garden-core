package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/api/middleware"
	mocks_middleware "github.com/brunobrolesi/open-garden-core/internal/api/middleware/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthMiddleware(t *testing.T) {
	type TestSuite struct {
		Sut              *gin.Engine
		TokenServiceMock *mocks_middleware.TokenService
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		tokenServiceMock := mocks_middleware.NewTokenService(t)
		sut := middleware.NewAuthMiddleware(tokenServiceMock)

		r := gin.Default()
		r.GET("/auth", sut.Handle, func(c *gin.Context) {
			userId, err := shared.GetUserId(c)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": shared.ErrApiBadRequest.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{"authorized": userId})
		})

		return TestSuite{
			Sut:              r,
			TokenServiceMock: tokenServiceMock,
		}
	}

	t.Run("Should return an error if can't get Authorization header", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"unauthorized"}`
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should call ValidateToken from TokenService with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.TokenServiceMock.On("ValidateToken", mock.Anything).Return(0, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		req.Header.Set("Authorization", "valid_token")
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.TokenServiceMock.AssertCalled(t, "ValidateToken", middleware.Token("valid_token"))
	})
	t.Run("Should return 401 and correct message if ValidateToken from TokenService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.TokenServiceMock.On("ValidateToken", mock.Anything).Return(0, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		req.Header.Set("Authorization", "valid_token")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"unauthorized"}`
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 200 if X-User-Id header is set", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.TokenServiceMock.On("ValidateToken", mock.Anything).Return(1, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		req.Header.Set("Authorization", "valid_token")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"authorized":1}`
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
