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

func TestGetFarmSensorsUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                      usecase.GetFarmSensorsUseCase
		FarmServiceMock          *mocks_gateway.FarmService
		FarmSensorRepositoryMock *mocks_gateway.FarmSensorRepository
	}
	makeTestSuite := func() TestSuite {
		farmSensorRepositoryMock := mocks_gateway.NewFarmSensorRepository(t)
		farmServiceMock := mocks_gateway.NewFarmService(t)
		sut := usecase.NewGetFarmSensorsUseCase(farmSensorRepositoryMock, farmServiceMock)

		return TestSuite{
			Sut:                      sut,
			FarmServiceMock:          farmServiceMock,
			FarmSensorRepositoryMock: farmSensorRepositoryMock,
		}
	}
	makeGetFarmSensorsInputDto := func() usecase.GetFarmSensorsInputDto {
		return usecase.GetFarmSensorsInputDto{
			UserId: 1,
			FarmId: 2,
		}
	}

	t.Run("Should call GetFarmByFarmIdAndUserId from FarmService with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorsInputDto()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))
		testSuite.Sut.Exec(ctx, input)
		testSuite.FarmServiceMock.AssertCalled(t, "GetFarmByIdAndUserId", ctx, input.FarmId, input.UserId)
	})
	t.Run("Should return an error if GetFarmByFarmIdAndUserId from FarmService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("get_farm_error"))
		result, err := testSuite.Sut.Exec(ctx, makeGetFarmSensorsInputDto())
		assert.Empty(t, result)
		assert.EqualError(t, err, "get_farm_error")
	})
	t.Run("Should return empty value if GetFarmByFarmIdAndUserId from FarmService returns an empty value", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, nil)
		result, err := testSuite.Sut.Exec(ctx, makeGetFarmSensorsInputDto())
		assert.Equal(t, model.FarmSensors{}, result)
		assert.Nil(t, err)
	})
	t.Run("Should call GetFarmSensorsByFarmId from FarmSensorRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorsInputDto()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{Id: input.UserId}, nil)
		testSuite.FarmSensorRepositoryMock.On("GetFarmSensorsByFarmId", mock.Anything, mock.Anything).Return(model.FarmSensors{}, errors.New("any_error"))
		testSuite.Sut.Exec(ctx, input)
		testSuite.FarmSensorRepositoryMock.AssertCalled(t, "GetFarmSensorsByFarmId", ctx, input.FarmId)
	})
	t.Run("Should return an error if GetFarmSensorsByFarmId from FarmSensorRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorsInputDto()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{Id: input.UserId}, nil)
		testSuite.FarmSensorRepositoryMock.On("GetFarmSensorsByFarmId", mock.Anything, mock.Anything).Return(model.FarmSensors{}, errors.New("get_farm_sensor_error"))
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Equal(t, model.FarmSensors{}, result)
		assert.EqualError(t, err, "get_farm_sensor_error")
	})

	t.Run("Should return farm sensors on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorsInputDto()
		farmSensors := model.FarmSensors{{Id: 1, FarmId: input.FarmId, Name: "any_name", SensorModel: 2, Description: "any_description", Active: true}, {Id: 2, FarmId: input.FarmId, Name: "any_name_2", SensorModel: 3, Description: "any_description_2", Active: true}}
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{Id: input.UserId}, nil)
		testSuite.FarmSensorRepositoryMock.On("GetFarmSensorsByFarmId", mock.Anything, mock.Anything).Return(farmSensors, nil)
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Equal(t, farmSensors, result)
		assert.Nil(t, err)
	})
}
