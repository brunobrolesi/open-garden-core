package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/usecase"
	mocks_usecase "github.com/brunobrolesi/open-garden-core/internal/user/business/usecase/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/user/infra/delivery/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthenticateUserHandler(t *testing.T) {
	type TestSuite struct {
		Sut                             *gin.Engine
		AuthenticateUserUserUseCaseMock *mocks_usecase.AuthenticateUserUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		authenticateUserUseCaseMock := mocks_usecase.NewAuthenticateUserUseCase(t)
		sut := handler.NewAuthenticateUserHandler(authenticateUserUseCaseMock)

		r := gin.Default()
		r.POST("/login", sut.Handle)

		return TestSuite{
			Sut:                             r,
			AuthenticateUserUserUseCaseMock: authenticateUserUseCaseMock,
		}
	}

	makeValidBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"email": "valid@mail.com",
			"password": "valid_password"
		}
	`))
	}

	makeUnprocessableBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"email": "valid@mail.com",
			"password": "valid_password",
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
					"email": "  ",
					"password": "valid_password"
				}
			`)),
				ExpectedResponseBody: `{"error":"email can't be empty"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"email": "invalidmail.com",
					"password": "valid_password"
				}
			`)),
				ExpectedResponseBody: `{"error":"email must be a valid format"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"email": "valid@mail.com",
					"password": "    "
				}
			`)),
				ExpectedResponseBody: `{"error":"password can't be empty"}`,
			},
		}
	}

	t.Run("Should return an error if body is invalid", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", makeUnprocessableBody())
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
			req, _ := http.NewRequest(http.MethodPost, "/login", test.InvalidBody)
			testSuite.Sut.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, test.ExpectedResponseBody, rr.Body.String())
		}
	})
	t.Run("Should call AuthenticateUserUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		credentials := usecase.AuthenticateUserInputDto{
			Email:    "valid@mail.com",
			Password: "valid_password",
		}
		testSuite.AuthenticateUserUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token(""), errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.AuthenticateUserUserUseCaseMock.AssertCalled(t, "Exec", credentials, mock.Anything)
	})
	t.Run("Should return 401 and correct message if AuthenticateUserUseCase returns an ErrAuthentication", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AuthenticateUserUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token(""), model.ErrAuthentication)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"invalid email or password"}`
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 500 and correct message if AuthenticateUserUseCase returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AuthenticateUserUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token(""), errors.New("any_error"))

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 200 and token on success", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.AuthenticateUserUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token("any_token"), nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/login", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"token":"any_token"}`
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
