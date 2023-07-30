package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserFarmsHandler(t *testing.T) {
	type TestSuite struct {
		Sut                     *gin.Engine
		GetUserFarmsUseCaseMock *mocks_usecase.GetUserFarmsUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		getUserFarmsUseCaseMock := mocks_usecase.NewGetUserFarmsUseCase(t)
		sut := handler.NewGetUserFarmsHandler(getUserFarmsUseCaseMock)

		r := gin.Default()
		r.GET("/farm", sut.Handle)

		return TestSuite{
			Sut:                     r,
			GetUserFarmsUseCaseMock: getUserFarmsUseCaseMock,
		}
	}

	t.Run("Should return an error if can't get user id header", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm", nil)
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should call GetUserFarmsUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.GetUserFarmsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farms{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.GetUserFarmsUseCaseMock.AssertCalled(t, "Exec", mock.Anything, 1)
	})
	t.Run("Should return 500 and correct message if GetUserFarmsUseCase returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.GetUserFarmsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farms{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 404 if no farms was found", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.GetUserFarmsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farms{}, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, `{"message":"no farms found for this user"}`, rr.Body.String())
	})
	t.Run("Should return 200 and array of farms on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		farms := model.Farms{
			model.Farm{Id: 1, Name: "farm_1", Address: "address_1", Owner: 1, Active: true},
			model.Farm{Id: 2, Name: "farm_2", Address: "address_2", Owner: 1, Active: true},
		}
		testSuite.GetUserFarmsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(farms, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"data":[{"id":1,"name":"farm_1","owner":1,"address":"address_1","active":true},{"id":2,"name":"farm_2","owner":1,"address":"address_2","active":true}]}`
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
