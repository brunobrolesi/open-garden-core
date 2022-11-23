package usecase_test

import (
	"context"
	"errors"
	"testing"

	mocks_gateway "github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateFarmUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                usecase.CreateFarmUseCase
		FarmRepositoryMock *mocks_gateway.FarmRepository
	}
	makeTestSuite := func() TestSuite {
		farmRepositoryMock := mocks_gateway.NewFarmRepository(t)
		sut := usecase.NewCreateFarmUseCase(farmRepositoryMock)
		return TestSuite{
			Sut:                sut,
			FarmRepositoryMock: farmRepositoryMock,
		}
	}
	makeCreateFarmInputDto := func() usecase.CreateFarmInputDto {
		return usecase.CreateFarmInputDto{
			Name:    "valid_name",
			Address: "valid_address",
			Owner:   1,
		}
	}

	t.Run("Should call CreateFarm from FarmRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		farm := makeCreateFarmInputDto()
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("CreateFarm", mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))
		testSuite.Sut.Exec(farm, ctx)
		expectedFarmCall := model.Farm{
			Name:    farm.Name,
			Address: farm.Address,
			Owner:   farm.Owner,
			Active:  true,
		}
		testSuite.FarmRepositoryMock.AssertCalled(t, "CreateFarm", expectedFarmCall, ctx)
	})
	t.Run("Should return an error if CreateFarm from FarmRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		farm := makeCreateFarmInputDto()
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("CreateFarm", mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("create_farm_error"))
		result, err := testSuite.Sut.Exec(farm, ctx)
		assert.Empty(t, result)
		assert.Error(t, err, "create_farm_error")
	})
	t.Run("Should return a farm on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		farm := makeCreateFarmInputDto()
		ctx := context.Background()
		expected := model.Farm{
			Id:      1,
			Name:    farm.Name,
			Address: farm.Address,
			Owner:   1,
			Active:  true,
		}
		testSuite.FarmRepositoryMock.On("CreateFarm", mock.Anything, mock.Anything).Return(expected, nil)
		result, err := testSuite.Sut.Exec(farm, ctx)
		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
}
