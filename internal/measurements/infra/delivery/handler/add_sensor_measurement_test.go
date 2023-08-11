package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddSensorMeasurementHandler(t *testing.T) {
	type TestSuite struct {
		Sut                             *gin.Engine
		AddSensorMeasurementUseCaseMock *mocks_usecase.AddSensorMeasurementUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		addSensorMeasurementUseCaseMock := mocks_usecase.NewAddSensorMeasurementUseCase(t)
		sut := handler.NewAddSensorMeasurementHandler(addSensorMeasurementUseCaseMock)

		r := gin.Default()
		r.POST("/sensor", sut.Handle)

		return TestSuite{
			Sut:                             r,
			AddSensorMeasurementUseCaseMock: addSensorMeasurementUseCaseMock,
		}
	}

	makeValidBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"sensor_id": 1,
			"value": 10.7
		}
	`))
	}

	makeUnprocessableBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"sensor_id": 1,
			"value": 10.7
	`))
	}

	t.Run("Should return an error if body is invalid", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/sensor", makeUnprocessableBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"unexpected EOF"}`
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should call AddSensorMeasurementUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		measurement := usecase.AddSensorMeasurementInputDto{
			SensorId: 1,
			Value:    10.7,
		}
		testSuite.AddSensorMeasurementUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/sensor", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.AddSensorMeasurementUseCaseMock.AssertCalled(t, "Exec", mock.Anything, measurement)
	})
	t.Run("Should return 500 and correct message if AddSensorMeasurementUseCase returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AddSensorMeasurementUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/sensor", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 201 and on success", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AddSensorMeasurementUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/sensor", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Empty(t, rr.Body.String())
	})
}
