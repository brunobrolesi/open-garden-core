package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	mocks_gateway "github.com/brunobrolesi/open-garden-core/internal/measurements/business/gateway/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPeriodMeasurements(t *testing.T) {
	type TestSuite struct {
		Sut                             usecase.GetSensorPeriodMeasurementsUseCase
		SensorMeasurementRepositoryMock *mocks_gateway.SensorMeasurementRepository
	}
	makeTestSuite := func() TestSuite {
		sensorMeasurementRepository := mocks_gateway.NewSensorMeasurementRepository(t)
		sut := usecase.NewGetSensorPeriodMeasurementsUseCase(sensorMeasurementRepository)
		return TestSuite{
			Sut:                             sut,
			SensorMeasurementRepositoryMock: sensorMeasurementRepository,
		}
	}
	makeInput := func() usecase.GetSensorPeriodMeasurementsInputDto {
		return usecase.GetSensorPeriodMeasurementsInputDto{
			SensorId: 1,
			UserID:   2,
			From:     time.Now().Add(-time.Hour),
			To:       time.Now(),
		}
	}
	t.Run("should call get sensor period measurements repository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		input := makeInput()
		ctx := context.Background()
		testSuite.SensorMeasurementRepositoryMock.On("GetSensorPeriodMeasurements", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(model.SensorMeasurements{}, nil).Once()
		_, _ = testSuite.Sut.Exec(ctx, input)

		testSuite.SensorMeasurementRepositoryMock.AssertCalled(t, "GetSensorPeriodMeasurements", ctx, input.SensorId, input.UserID, input.From, input.To)
	})

	t.Run("should return error if get sensor period measurements repository returns error", func(t *testing.T) {
		testSuite := makeTestSuite()
		input := makeInput()
		testSuite.SensorMeasurementRepositoryMock.On("GetSensorPeriodMeasurements", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(model.SensorMeasurements{}, errors.New("get measurements error")).Once()
		_, err := testSuite.Sut.Exec(context.Background(), input)

		assert.EqualError(t, err, "get measurements error")
	})

	t.Run("should return sensor measurements on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		input := makeInput()
		expected := model.SensorMeasurements{{SensorID: 1, Value: 10.5, Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)}}
		testSuite.SensorMeasurementRepositoryMock.On("GetSensorPeriodMeasurements", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(expected, nil).Once()

		result, err := testSuite.Sut.Exec(context.Background(), input)

		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
}
