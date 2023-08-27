package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFarmSensorHandler(t *testing.T) {
	type TestSuite struct {
		Sut                      *gin.Engine
		GetFarmSensorUseCaseMock *mocks_usecase.GetFarmSensorUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		getFarmSensorUseCaseMock := mocks_usecase.NewGetFarmSensorUseCase(t)
		sut := handler.NewGetFarmSensorHandler(getFarmSensorUseCaseMock)

		r := gin.Default()
		r.GET("/farm/:farm_id/sensor/:farm_sensor_id", sut.Handle)

		return TestSuite{
			Sut:                      r,
			GetFarmSensorUseCaseMock: getFarmSensorUseCaseMock,
		}
	}

	t.Run("Should return an error if can't get user id header", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/1/sensor/2", nil)
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return an error if can't get farm id in params", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/invalid_id/sensor/2", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})

	t.Run("Should return an error if can't get farm sensor id in params", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/1/sensor/invalid_id", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should call GetUserFarmSensorUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.GetFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2/sensor/3", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := usecase.GetFarmSensorInputDto{
			UserId:       1,
			FarmId:       2,
			FarmSensorId: 3,
		}
		testSuite.GetFarmSensorUseCaseMock.AssertCalled(t, "Exec", mock.Anything, expected)
	})
	t.Run("Should return 500 and correct message if GetUserFarmSensorUseCase returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.GetFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2/sensor/3", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 404 if no farm was found", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.GetFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2/sensor/3", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, `{"message":"no sensor farm found"}`, rr.Body.String())
	})
	t.Run("Should return 200 and farm sensor on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		farmSensor := model.FarmSensor{Id: 3, Name: "any_name", FarmId: 2, Active: true, SensorModel: 3, Description: "any_description"}
		testSuite.GetFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(farmSensor, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/farm/2/sensor/3", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"data":{"id":3,"name":"any_name","sensor_model":3,"farm_id":2,"description":"any_description","active":true}}`
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
