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

func TestNewGetFarmSensorUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                      usecase.GetFarmSensorUseCase
		FarmServiceMock          *mocks_gateway.FarmService
		FarmSensorRepositoryMock *mocks_gateway.FarmSensorRepository
	}
	makeTestSuite := func() TestSuite {
		farmSensorRepositoryMock := mocks_gateway.NewFarmSensorRepository(t)
		farmServiceMock := mocks_gateway.NewFarmService(t)
		sut := usecase.NewGetFarmSensorUseCase(farmSensorRepositoryMock, farmServiceMock)

		return TestSuite{
			Sut:                      sut,
			FarmServiceMock:          farmServiceMock,
			FarmSensorRepositoryMock: farmSensorRepositoryMock,
		}
	}
	makeGetFarmSensorInputDto := func() usecase.GetFarmSensorInputDto {
		return usecase.GetFarmSensorInputDto{
			UserId:       1,
			FarmId:       2,
			FarmSensorId: 3,
		}
	}

	t.Run("Should call GetFarmByFarmIdAndUserId from FarmService with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorInputDto()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))
		testSuite.Sut.Exec(ctx, input)
		testSuite.FarmServiceMock.AssertCalled(t, "GetFarmByIdAndUserId", ctx, input.FarmId, input.UserId)
	})
	t.Run("Should return an error if GetFarmByFarmIdAndUserId from FarmService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("get_farm_error"))
		result, err := testSuite.Sut.Exec(ctx, makeGetFarmSensorInputDto())
		assert.Empty(t, result)
		assert.EqualError(t, err, "get_farm_error")
	})
	t.Run("Should return empty value if GetFarmByFarmIdAndUserId from FarmService returns an empty value", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, nil)
		result, err := testSuite.Sut.Exec(ctx, makeGetFarmSensorInputDto())
		assert.Equal(t, model.FarmSensor{}, result)
		assert.Nil(t, err)
	})
	t.Run("Should call GetFarmSensorById from FarmSensorRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorInputDto()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{Id: input.UserId}, nil)
		testSuite.FarmSensorRepositoryMock.On("GetFarmSensorById", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("any_error"))
		testSuite.Sut.Exec(ctx, input)
		testSuite.FarmSensorRepositoryMock.AssertCalled(t, "GetFarmSensorById", ctx, input.FarmSensorId)
	})
	t.Run("Should return an error if GetFarmSensorById from FarmSensorRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorInputDto()
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{Id: input.UserId}, nil)
		testSuite.FarmSensorRepositoryMock.On("GetFarmSensorById", mock.Anything, mock.Anything).Return(model.FarmSensor{}, errors.New("get_farm_sensor_error"))
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Equal(t, model.FarmSensor{}, result)
		assert.EqualError(t, err, "get_farm_sensor_error")
	})

	t.Run("Should return farm sensor on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetFarmSensorInputDto()
		farmSensor := model.FarmSensor{Id: input.FarmSensorId, FarmId: input.FarmId, Name: "any_name", SensorModel: 2, Description: "any_description", Active: true}
		testSuite.FarmServiceMock.On("GetFarmByIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{Id: input.UserId}, nil)
		testSuite.FarmSensorRepositoryMock.On("GetFarmSensorById", mock.Anything, mock.Anything).Return(farmSensor, nil)
		result, err := testSuite.Sut.Exec(ctx, input)

		assert.Equal(t, farmSensor, result)
		assert.Nil(t, err)
	})
}
