package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSensorMeasurements(t *testing.T) {
	type TestSuite struct {
		Sut                                    *gin.Engine
		GetSensorPeriodMeasurementsUseCaseMock *mocks_usecase.GetSensorPeriodMeasurementsUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		getSensorPeriodMeasurementsUseCaseMock := mocks_usecase.NewGetSensorPeriodMeasurementsUseCase(t)
		sut := handler.NewGetSensorMeasurementsHandler(getSensorPeriodMeasurementsUseCaseMock)

		r := gin.Default()
		r.GET("/measurements/sensors/:id", sut.Handle)

		return TestSuite{
			Sut:                                    r,
			GetSensorPeriodMeasurementsUseCaseMock: getSensorPeriodMeasurementsUseCaseMock,
		}
	}

	t.Run("should return bad request if can't get user id", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1", nil)
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("should return bad request if can't get sensor id", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/invalid_id", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("should return bad request if can't get from param", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("should return bad request if can't get to param", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1?from=2023-01-02T15:04:05Z", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("should call use case with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.GetSensorPeriodMeasurementsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.SensorMeasurements{}, nil)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1?from=2023-01-02T15:04:05Z&to=2023-01-03T15:04:05Z", nil)
		req.Header.Set("X-User-Id", "2")
		testSuite.Sut.ServeHTTP(rr, req)

		from, _ := time.Parse(time.RFC3339, "2023-01-02T15:04:05Z")
		to, _ := time.Parse(time.RFC3339, "2023-01-03T15:04:05Z")
		expectedCall := usecase.GetSensorPeriodMeasurementsInputDto{
			SensorId: 1,
			UserID:   2,
			From:     from,
			To:       to,
		}
		testSuite.GetSensorPeriodMeasurementsUseCaseMock.AssertCalled(t, "Exec", mock.Anything, expectedCall)
	})
	t.Run("should return internal server error if use case returns error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.GetSensorPeriodMeasurementsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.SensorMeasurements{}, errors.New("use case error"))
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1?from=2023-01-02T15:04:05Z&to=2023-01-03T15:04:05Z", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("should return not found if no measurements are found", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.GetSensorPeriodMeasurementsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.SensorMeasurements{}, nil)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1?from=2023-01-02T15:04:05Z&to=2023-01-03T15:04:05Z", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"message":"No measurements found"}`
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("should return measurements on success", func(t *testing.T) {
		testSuite := makeTestSuite()

		timeValue, _ := time.Parse(time.RFC3339, "2023-01-02T15:04:05Z")
		testSuite.GetSensorPeriodMeasurementsUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.SensorMeasurements{{SensorID: 1, Time: timeValue, Value: 10.5}}, nil)
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/measurements/sensors/1?from=2023-01-02T15:04:05Z&to=2023-01-03T15:04:05Z", nil)
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"data":[{"sensor_id":1,"time":"2023-01-02T15:04:05Z","value":10.5}]}`
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
