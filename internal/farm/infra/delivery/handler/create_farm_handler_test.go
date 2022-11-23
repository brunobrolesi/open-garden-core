package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateFarmHandler(t *testing.T) {
	type TestSuite struct {
		Sut                   *gin.Engine
		CreateFarmUseCaseMock *mocks_usecase.CreateFarmUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		createFarmUseCaseMock := mocks_usecase.NewCreateFarmUseCase(t)
		sut := handler.NewCreateFarmHandler(createFarmUseCaseMock)

		r := gin.Default()
		r.POST("/farm", sut.Handle)

		return TestSuite{
			Sut:                   r,
			CreateFarmUseCaseMock: createFarmUseCaseMock,
		}
	}

	makeValidBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"name": "valid_name",
			"address": "valid_address"
		}
	`))
	}

	makeUnprocessableBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"name": "valid_name",
			"address": "valid_address"
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
					"name": "  ",
					"address": "valid_address"
				}
			`)),
				ExpectedResponseBody: `{"error":"name can't be empty"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"name": "valid_name",
					"address": "   "
				}
			`)),
				ExpectedResponseBody: `{"error":"address can't be empty"}`,
			},
		}
	}

	t.Run("Should return an error if can't get user id header", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"bad request"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return an error if body is invalid", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm", makeUnprocessableBody())
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
			req, _ := http.NewRequest(http.MethodPost, "/farm", test.InvalidBody)
			req.Header.Set("X-User-Id", "1")
			testSuite.Sut.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, test.ExpectedResponseBody, rr.Body.String())
		}
	})
	t.Run("Should call CreateFarmUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		farm := usecase.CreateFarmInputDto{
			Name:    "valid_name",
			Address: "valid_address",
			Owner:   1,
		}
		testSuite.CreateFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.CreateFarmUseCaseMock.AssertCalled(t, "Exec", farm, mock.Anything)
	})
	t.Run("Should return 500 and correct message if CreateFarmUseCase returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.CreateFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Farm{}, errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 201 and created farm on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		farm := model.Farm{
			Id:      1,
			Name:    "valid_name",
			Address: "valid_address",
			Owner:   1,
			Active:  true,
		}
		testSuite.CreateFarmUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(farm, nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/farm", makeValidBody())
		req.Header.Set("X-User-Id", "1")
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"data":{"id":1,"name":"valid_name","owner":1,"address":"valid_address","active":true}}`
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
