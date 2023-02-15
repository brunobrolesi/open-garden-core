package handler_test

import (
	"bytes"
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

func TestAddFarmSensorHandler_Handle(t *testing.T) {
	type TestSuite struct {
		Sut                      *gin.Engine
		AddFarmSensorUseCaseMock *mocks_usecase.AddFarmSensorUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		addFarmSensorUseCaseMock := mocks_usecase.NewAddFarmSensorUseCase(t)
		sut := handler.NewAddFarmSensorHandler(addFarmSensorUseCaseMock)

		r := gin.Default()
		r.POST("/farm/:farm_id/sensor", sut.Handle)

		return TestSuite{
			Sut:                      r,
			AddFarmSensorUseCaseMock: addFarmSensorUseCaseMock,
		}
	}

	makeValidBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"name": "valid_name",
			"sensor_model": 1,
			"description": "valid_description"
		}
	`))
	}

	makeUnprocessableBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"name": "valid_name",
			"sensor_model": 1,
	`))
	}

	type InvalidBodyTestCases struct {
		InvalidBody          *bytes.Buffer
		ExpectedResponseBody string
	}

	makeInvalidBodyTestCases := func() []InvalidBodyTestCases {
		return []InvalidBodyTestCases{
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"name": "a",
					"sensor_model": 1,
					"description": "valid_description"
				}
			`)),
				ExpectedResponseBody: `{"error":"Key: 'addFarmSensorBodyRequest.Name' Error:Field validation for 'Name' failed on the 'min' tag"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"name": "valid_name",
					"sensor_model": 0,
					"description": "valid_description"
				}
			`)),
				ExpectedResponseBody: `{"error":"Key: 'addFarmSensorBodyRequest.SensorModel' Error:Field validation for 'SensorModel' failed on the 'required' tag"}`,
			},
		}
	}

	t.Run("Should return an error if can't get farm id param", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/a/sensor", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return an error if can't get user id header", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/1/sensor", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return an error if body is invalid", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/1/sensor", makeUnprocessableBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"unexpected EOF"}`
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return an error if body field is invalid", func(t *testing.T) {
		testCases := makeInvalidBodyTestCases()

		for _, test := range testCases {
			testSuite := makeTestSuite()

			rr := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/farm/1/sensor", test.InvalidBody)
			req.Header.Set("X-User-Id", "1")
			testSuite.Sut.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, test.ExpectedResponseBody, rr.Body.String())
		}
	})
	t.Run("Should call AddFarmSensorUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		farmSensor := usecase.AddFarmSensorInputDto{
			Name:        "valid_name",
			SensorModel: 1,
			Description: "valid_description",
			FarmId:      2,
			UserId:      3,
		}
		testSuite.AddFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/2/sensor", makeValidBody())
		req.Header.Set("X-User-Id", "3")
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.AddFarmSensorUseCaseMock.AssertCalled(t, "Exec", mock.Anything, farmSensor)
	})
	t.Run("Should return 400 and correct message if AddFarmSensorUseCase returns ErrInvalidFarm", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AddFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, model.ErrInvalidFarm)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/1/sensor", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 400 and correct message if AddFarmSensorUseCase returns ErrInvalidSensor", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AddFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, model.ErrInvalidSensor)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/1/sensor", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 500 and correct message if AddFarmSensorUseCase returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AddFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/1/sensor", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 201 and created farm sensor on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		farmSensor := model.FarmSensor{
			Id:          1,
			Name:        "valid_name",
			SensorModel: 1,
			FarmId:      2,
			Description: "valid_description",
			Active:      true,
		}
		testSuite.AddFarmSensorUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(farmSensor, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm/2/sensor", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"data":{"id":1,"name":"valid_name","sensor_model":1,"farm_id":2,"description":"valid_description","active":true}}`
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
