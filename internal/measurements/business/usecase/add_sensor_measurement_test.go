package usecase_test

import (
	"context"
	"errors"
	"testing"

	mocks_gateway "github.com/brunobrolesi/open-garden-core/internal/measurements/business/gateway/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddSensorMeasurement(t *testing.T) {
	type TestSuite struct {
		Sut                             usecase.AddSensorMeasurementUseCase
		SensorMeasurementRepositoryMock *mocks_gateway.SensorMeasurementRepository
	}
	makeTestSuite := func() TestSuite {
		sensorMeasurementRepository := mocks_gateway.NewSensorMeasurementRepository(t)
		sut := usecase.NewAddSensorMeasurementUseCase(sensorMeasurementRepository)
		return TestSuite{
			Sut:                             sut,
			SensorMeasurementRepositoryMock: sensorMeasurementRepository,
		}
	}
	makeInput := func() usecase.AddSensorMeasurementInputDto {
		return usecase.AddSensorMeasurementInputDto{
			SensorId: 1,
			Value:    12.7,
		}
	}
	t.Run("should call add sensor measurement repository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		input := makeInput()
		ctx := context.Background()
		testSuite.SensorMeasurementRepositoryMock.On("AddSensorMeasurement", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		_ = testSuite.Sut.Exec(ctx, input)

		testSuite.SensorMeasurementRepositoryMock.AssertCalled(t, "AddSensorMeasurement", ctx, input.SensorId, input.Value)
	})

	t.Run("should return error if add sensor measurement repository returns error", func(t *testing.T) {
		testSuite := makeTestSuite()
		input := makeInput()
		testSuite.SensorMeasurementRepositoryMock.On("AddSensorMeasurement", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("add measurement error")).Once()
		err := testSuite.Sut.Exec(context.Background(), input)

		assert.EqualError(t, err, "add measurement error")
	})

	t.Run("should return nil on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		input := makeInput()
		testSuite.SensorMeasurementRepositoryMock.On("AddSensorMeasurement", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		err := testSuite.Sut.Exec(context.Background(), input)

		assert.Nil(t, err)
	})
}
