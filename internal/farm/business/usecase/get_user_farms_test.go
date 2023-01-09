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

func TestGetUserFarmsUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                usecase.GetUserFarmsUseCase
		FarmRepositoryMock *mocks_gateway.FarmRepository
	}
	makeTestSuite := func() TestSuite {
		farmRepositoryMock := mocks_gateway.NewFarmRepository(t)
		sut := usecase.NewGetUserFarmsUseCase(farmRepositoryMock)
		return TestSuite{
			Sut:                sut,
			FarmRepositoryMock: farmRepositoryMock,
		}
	}

	t.Run("Should call GetFarmsByUserId from FarmRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		userId := 1
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("GetFarmsByUserId", mock.Anything, mock.Anything).Return(model.Farms{}, errors.New("any_error"))
		testSuite.Sut.Exec(ctx, userId)
		testSuite.FarmRepositoryMock.AssertCalled(t, "GetFarmsByUserId", ctx, userId)
	})
	t.Run("Should return an error if GetFarmsByUserId from FarmRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		userId := 1
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("GetFarmsByUserId", mock.Anything, mock.Anything).Return(model.Farms{}, errors.New("get_farms_error"))
		result, err := testSuite.Sut.Exec(ctx, userId)
		assert.Empty(t, result)
		assert.Error(t, err, "get_farms_error")
	})
	t.Run("Should return a farms on success", func(t *testing.T) {
		expected := model.Farms{
			model.Farm{Id: 1, Name: "farm_1", Address: "address_1", Owner: 1, Active: true},
			model.Farm{Id: 2, Name: "farm_2", Address: "address_2", Owner: 1, Active: true},
		}
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("GetFarmsByUserId", mock.Anything, mock.Anything).Return(expected, nil)
		result, err := testSuite.Sut.Exec(ctx, 1)
		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
}
