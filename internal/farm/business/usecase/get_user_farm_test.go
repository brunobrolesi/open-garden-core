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

func TestGetUserFarmUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                usecase.GetUserFarmUseCase
		FarmRepositoryMock *mocks_gateway.FarmRepository
	}
	makeTestSuite := func() TestSuite {
		farmRepositoryMock := mocks_gateway.NewFarmRepository(t)
		sut := usecase.NewGetUserFarmUseCase(farmRepositoryMock)
		return TestSuite{
			Sut:                sut,
			FarmRepositoryMock: farmRepositoryMock,
		}
	}

	makeGetUserFarmInputDto := func() usecase.GetUserFarmInputDto {
		return usecase.GetUserFarmInputDto{
			UserId: 1,
			FarmId: 2,
		}
	}

	t.Run("Should call GetFarmByFarmIdAndUserId from FarmRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		input := makeGetUserFarmInputDto()
		testSuite.FarmRepositoryMock.On("GetFarmByFarmIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))
		testSuite.Sut.Exec(input, ctx)
		testSuite.FarmRepositoryMock.AssertCalled(t, "GetFarmByFarmIdAndUserId", input.FarmId, input.UserId, ctx)
	})
	t.Run("Should return an error if GetFarmByFarmIdAndUserId from FarmRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("GetFarmByFarmIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("get_farm_error"))
		result, err := testSuite.Sut.Exec(makeGetUserFarmInputDto(), ctx)
		assert.Empty(t, result)
		assert.Error(t, err, "get_farm_error")
	})
	t.Run("Should return a farm on success", func(t *testing.T) {
		expected := model.Farm{Id: 1, Name: "farm", Address: "address", Owner: 1, Active: true}
		testSuite := makeTestSuite()
		ctx := context.Background()
		testSuite.FarmRepositoryMock.On("GetFarmByFarmIdAndUserId", mock.Anything, mock.Anything, mock.Anything).Return(expected, nil)
		result, err := testSuite.Sut.Exec(makeGetUserFarmInputDto(), ctx)
		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})
}
