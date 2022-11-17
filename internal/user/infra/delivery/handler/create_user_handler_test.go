package handler_test

import (
	"bytes"
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

func TestCreateUserHandler(t *testing.T) {
	type TestSuite struct {
		Sut                   *gin.Engine
		CreateUserUseCaseMock *mocks_usecase.CreateUserUseCase
	}

	makeTestSuite := func() TestSuite {
		gin.SetMode(gin.TestMode)

		createUserUseCaseMock := mocks_usecase.NewCreateUserUseCase(t)
		sut := handler.NewCreateUserHandler(createUserUseCaseMock)

		r := gin.Default()
		r.POST("/register", sut.Handle)

		return TestSuite{
			Sut:                   r,
			CreateUserUseCaseMock: createUserUseCaseMock,
		}
	}

	makeValidBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"company_name": "valid_company_name",
			"email": "valid@mail.com",
			"password": "valid_password",
			"confirm_password": "valid_password"
		}
	`))
	}

	makeUnprocessableBody := func() *bytes.Buffer {
		return bytes.NewBuffer([]byte(`
		{
			"company_name": "valid_company_name",
			"email": "valid@mail.com",
			"password": "valid_password",
			"confirm_password": "valid_password"
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
					"company_name": "   ",
					"email": "valid@mail.com",
					"password": "valid_password",
					"confirm_password": "valid_password"
				}
			`)),
				ExpectedResponseBody: `{"error":"company_name can't be empty"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"company_name": "valid_company_name",
					"email": "  ",
					"password": "valid_password",
					"confirm_password": "valid_password"
				}
			`)),
				ExpectedResponseBody: `{"error":"email can't be empty"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"company_name": "valid_company_name",
					"email": "invalidmail.com",
					"password": "valid_password",
					"confirm_password": "valid_password"
				}
			`)),
				ExpectedResponseBody: `{"error":"email must be a valid format"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"company_name": "valid_company_name",
					"email": "valid@mail.com",
					"password": "    ",
					"confirm_password": "valid_password"
				}
			`)),
				ExpectedResponseBody: `{"error":"password can't be empty"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"company_name": "valid_company_name",
					"email": "valid@mail.com",
					"password": "valid_password",
					"confirm_password": "  "
				}
			`)),
				ExpectedResponseBody: `{"error":"confirm_password can't be empty"}`,
			},
			{
				InvalidBody: bytes.NewBuffer([]byte(`
				{
					"company_name": "valid_company_name",
					"email": "valid@mail.com",
					"password": "valid_password",
					"confirm_password": "no_match"
				}
			`)),
				ExpectedResponseBody: `{"error":"passwords must be equal"}`,
			},
		}
	}

	t.Run("Should return an error if body is invalid", func(t *testing.T) {
		testSuite := makeTestSuite()

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", makeUnprocessableBody())
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
			req, _ := http.NewRequest(http.MethodPost, "/register", test.InvalidBody)
			testSuite.Sut.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, test.ExpectedResponseBody, rr.Body.String())
		}
	})
	t.Run("Should call CreateUserUseCase with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()

		user := usecase.CreateUserInputDto{
			CompanyName: "valid_company_name",
			Email:       "valid@mail.com",
			Password:    "valid_password",
		}
		testSuite.CreateUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token(""), model.ErrEmailInUse)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		testSuite.CreateUserUseCaseMock.AssertCalled(t, "Exec", user, mock.Anything)
	})
	t.Run("Should return 400 and correct message if CreateUserUseCase returns an ErrEmailInUse", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.CreateUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token(""), model.ErrEmailInUse)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"email in use"}`
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 500 and correct message if CreateUserUseCase returns an no treated error", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.CreateUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token(""), model.ErrInternalServer)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"error":"internal server error"}`
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
	t.Run("Should return 201 and token on success", func(t *testing.T) {
		testSuite := makeTestSuite()

		testSuite.CreateUserUseCaseMock.On("Exec", mock.Anything, mock.Anything).Return(model.Token("any_token"), nil)

		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/register", makeValidBody())
		testSuite.Sut.ServeHTTP(rr, req)

		expected := `{"token":"any_token"}`
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, expected, rr.Body.String())
	})
}
