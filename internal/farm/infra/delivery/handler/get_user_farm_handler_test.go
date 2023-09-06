package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserFarmHandler(t *testing.T) {
	type TestSuite struct {
		Sut                    *gin.Engine
		GetUserFarmUseCaseMock *mocks_usecase.GetUserFarmUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		getUserFarmUseCaseMock := mocks_usecase.NewGetUserFarmUseCase(t)
		sut := handler.NewGetUserFarmHandler(getUserFarmUseCaseMock)

		r := gin.Default()
		r.GET("/farm/:farm_id", sut.Handle)

		return TestSuite{
			Sut:                    r,
			GetUserFarmUseCaseMock: getUserFarmUseCaseMock,
		}
	}

	t.Run("Should return an error if can't get user id header", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/1", nil)
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return an error if can't get farm id in params", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/invalid_id", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should call GetUserFarmUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.GetUserFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := usecase.GetUserFarmInputDto{
			UserId: 1,
			FarmId: 2,
		}
		testSuite.GetUserFarmUseCaseMock.AssertCalled(t, "Exec", mock.Anything, expected)
	})
	t.Run("Should return 500 and correct message if GetUserFarmUseCaseMock returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.GetUserFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 404 if no farm was found", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.GetUserFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farm{}, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, `{"message":"no farm found for this user"}`, rr.Body.String())
	})
	t.Run("Should return 200 and farm on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		farm := model.Farm{Id: 2, Name: "farm_2", Address: "address_2", Owner: 1, Active: true}
		testSuite.GetUserFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(farm, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"data":{"id":2,"name":"farm_2","owner":1,"address":"address_2","active":true}}`
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
