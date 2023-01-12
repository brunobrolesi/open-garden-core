package usecase_test

import (
	"context"
	"errors"
	"testing"

	mocks_gateway "github.com/brunobrolesi/open-garden-core/internal/sensor/business/gateway/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFarmSensor(t *testing.T) {
	type TestSuite struct {
		Sut                      usecase.AddFarmSensorUseCase
		SensorRepositoryMock     *mocks_gateway.SensorRepository
		FarmRepositoryMock       *mocks_gateway.FarmRepository
		FarmSensorRepositoryMock *mocks_gateway.FarmSensorRepository
	}
	makeTestSuite := func() TestSuite {
		sensorRepositoryMock := mocks_gateway.NewSensorRepository(t)
		farmRepositoryMock := mocks_gateway.NewFarmRepository(t)
		farmSensorRepositoryMock := mocks_gateway.NewFarmSensorRepository(t)

		sut := usecase.NewAddFarmSensorUseCase(sensorRepositoryMock, farmRepositoryMock, farmSensorRepositoryMock)
		return TestSuite{
			Sut:                      sut,
			SensorRepositoryMock:     sensorRepositoryMock,
			FarmRepositoryMock:       farmRepositoryMock,
			FarmSensorRepositoryMock: farmSensorRepositoryMock,
		}
	}
	makeAddFarmSensorInputDto := func() usecase.AddFarmSensorInputDto {
		return usecase.AddFarmSensorInputDto{
			SensorId:    1,
			Description: "valid_description",
			FarmId:      2,
			UserId:      3,
		}
	}
	makeSensor := func() model.Sensor {
		return model.Sensor{
			Id:   1,
			Name: "temp_sensor",
			Type: model.Temperature,
			Unit: "°C",
		}
	}
	makeFarm := func() model.Farm {
		return model.Farm{
			Id:      2,
			Name:    "any_farm",
			Owner:   3,
			Address: "any_address",
			Active:  true,
		}
	}
	t.Run("Should call GetSensorById from SensorRepository with correct id", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(model.Sensor{}, errors.New("any_error"))

		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		testSuite.Sut.Exec(ctx, input)

		testSuite.SensorRepositoryMock.AssertCalled(t, "GetSensorById", ctx, input.SensorId)
	})
	t.Run("Should return an error if GetSensorById from SensorRepository return an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(model.Sensor{}, errors.New("get_sensor_error"))

		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Empty(t, result)
		assert.Error(t, err, "get_sensor_error")
	})
	t.Run("Should call GetFarmByIdAndUserId from FarmRepository with correct id", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(makeSensor(), nil)
		testSuite.FarmRepositoryMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))

		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		testSuite.Sut.Exec(ctx, input)

		testSuite.FarmRepositoryMock.AssertCalled(t, "GetFarmByIdAndUserId", ctx, input.FarmId, input.UserId)
	})
	t.Run("Should return an error if GetFarmByIdAndUserId from FarmRepository return an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(makeSensor(), nil)
		testSuite.FarmRepositoryMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("get_farm_error"))

		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Empty(t, result)
		assert.Error(t, err, "get_farm_error")
	})
	t.Run("Should call CreateFarmSensor from FarmSensorRepository with correct id", func(t *testing.T) {
		testSuite := makeTestSuite()
		sensor := makeSensor()
		farm := makeFarm()
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(sensor, nil)
		testSuite.FarmRepositoryMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(farm, nil)
		testSuite.FarmSensorRepositoryMock.On("CreateFarmSensor", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("any_error"))

		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		testSuite.Sut.Exec(ctx, input)

		expectedInput := model.FarmSensor{
			SensorId:    sensor.Id,
			FarmId:      farm.Id,
			Description: input.Description,
			Active:      true,
		}

		testSuite.FarmSensorRepositoryMock.AssertCalled(t, "CreateFarmSensor", ctx, expectedInput)
	})
	t.Run("Should return an error if CreateFarmSensor from FarmSensorRepository return an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(makeSensor(), nil)
		testSuite.FarmRepositoryMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(makeFarm(), nil)
		testSuite.FarmSensorRepositoryMock.On("CreateFarmSensor", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("create_farm_sensor_error"))

		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Empty(t, result)
		assert.Error(t, err, "create_farm_sensor_error")
	})
	t.Run("Should return a Farm Sensor on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeAddFarmSensorInputDto()
		sensor := makeSensor()
		farm := makeFarm()
		farmSensor := model.FarmSensor{
			Id:          4,
			SensorId:    sensor.Id,
			FarmId:      farm.Id,
			Description: input.Description,
			Active:      true,
		}
		testSuite.SensorRepositoryMock.On("GetSensorById", mock.Anything, mock.Anything).Return(makeSensor(), nil)
		testSuite.FarmRepositoryMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(makeFarm(), nil)
		testSuite.FarmSensorRepositoryMock.On("CreateFarmSensor", mock.Anything, mock.Anything).Return(farmSensor, nil)

		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Equal(t, farmSensor, result)
		assert.Nil(t, err)
	})
}
