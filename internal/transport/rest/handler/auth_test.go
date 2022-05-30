package handler

import (
	"bytes"
	mockhandler "github.com/Thing-repository/backend-server/internal/transport/rest/handler/mocks"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	res := m.Run()
	os.Exit(res)
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *mockhandler.MockAuth, user *core.UserSignInData, userData *core.SignInResponse)

	invalidDataMassage := `{"message":"invalid username or password"}`

	testTable := []struct {
		name                 string
		inputBody            string
		inputAuthData        core.UserSignInData
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		outputResponse       core.SignInResponse
	}{
		{
			name:      "Success",
			inputBody: `{"user_mail":"Test","user_password":"TestTest"}`,
			inputAuthData: core.UserSignInData{
				UserMail:     "Test",
				UserPassword: "TestTest",
			},
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignInData, userData *core.SignInResponse) {
				s.EXPECT().SignIn(user).Return(userData, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{` +
				`"id":0,` +
				`"first_name":"test_name",` +
				`"last_name":"test_last_name",` +
				`"email":"test_email",` +
				`"image_url":"test_image",` +
				`"company_id":1,` +
				`"department_id":2,` +
				`"is_company_admin":true,` +
				`"is_department_admin":true,` +
				`"vacation_time_start":1653484250,` +
				`"vacation_time_end":1653484250,` +
				`"token":"test_token"` +
				`}`,
			outputResponse: core.SignInResponse{
				User: core.User{
					Id:                0,
					FirstName:         "test_name",
					LastName:          "test_last_name",
					Email:             "test_email",
					ImageURL:          "test_image",
					CompanyId:         pointy.Int(1),
					DepartmentId:      pointy.Int(2),
					IsCompanyAdmin:    pointy.Bool(true),
					IsDepartmentAdmin: pointy.Bool(true),
					VacationTimeStart: pointy.Uint32(1653484250),
					VacationTimeEnd:   pointy.Uint32(1653484250),
				},
				Token: "test_token",
			},
		},
		{
			name:      "Short password",
			inputBody: `{"user_mail":"Test","user_password":"test"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignInData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: invalidDataMassage,
		},
		{
			name:      "Empty email",
			inputBody: `{"user_mail":"","user_password":"test"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignInData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: invalidDataMassage,
		},
		{
			name:      "Empty password",
			inputBody: `{"user_mail":"Test","user_password":""}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignInData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: invalidDataMassage,
		},
		{
			name:      "Empty input data",
			inputBody: `{"user_mail":"Test","user_password":""}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignInData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: invalidDataMassage,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mockhandler.NewMockAuth(c)
			testCase.mockBehavior(auth, &testCase.inputAuthData, &testCase.outputResponse)

			handler := NewHandler(auth)

			// Test handler
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse)

	testTable := []struct {
		name                 string
		inputBody            string
		inputAuthData        core.UserSignUpData
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
		outputResponse       core.SignInResponse
	}{
		{
			name:      "OK",
			inputBody: `{"first_name":"test_name","last_name":"test_last_name","email":"test@test.com","password":"TestTest24"}`,
			inputAuthData: core.UserSignUpData{
				FirstName: "test_name",
				LastName:  "test_last_name",
				Email:     "test@test.com",
				Password:  "TestTest24",
			},
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
				s.EXPECT().SignUp(user).Return(userData, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{` +
				`"id":0,` +
				`"first_name":"test_name",` +
				`"last_name":"test_last_name",` +
				`"email":"test@test.com",` +
				`"image_url":"test_image",` +
				`"token":"test_token"` +
				`}`,
			outputResponse: core.SignInResponse{
				User: core.User{
					Id:                0,
					FirstName:         "test_name",
					LastName:          "test_last_name",
					Email:             "test@test.com",
					ImageURL:          "test_image",
					CompanyId:         nil,
					DepartmentId:      nil,
					IsCompanyAdmin:    nil,
					IsDepartmentAdmin: nil,
					VacationTimeStart: nil,
					VacationTimeEnd:   nil,
				},
				Token: "test_token",
			},
		},
		{
			name:      "Short password",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test.com","password":"Test24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid password, too short"}`,
		},
		{
			name:      "No numbers in password",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test.com","password":"TestTestTest"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid password, no numbers"}`,
		},
		{
			name:      "No capital letters in password",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test.com","password":"test_test12342"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid password, no uppercase letters"}`,
		},
		{
			name:      "No capital letters in password and numbers",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test.com","password":"TEST_TEST_TEST_TEST345"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid password, no lowercase letters"}`,
		},
		{
			name:      "Invalid email no .",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email no @",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test.test","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email no @ and .",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test_test","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email has ;",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@t;est.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email has ,",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@t,est.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email has [",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@t[est.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email has ]",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@t]est.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email has ,",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@tes,t.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Invalid email has \\",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@tes\\t.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid email"}`,
		},
		{
			name:      "Empty first name",
			inputBody: `{"first_name":"","last_name":"TestTest","email":"test@test.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"}`,
		},
		{
			name:      "Hasn't first name",
			inputBody: `{"last_name":"TestTest","email":"test@test.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"}`,
		},
		{
			name:      "Empty last name",
			inputBody: `{"first_name":"Test","last_name":"","email":"test@test.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.LastName' Error:Field validation for 'LastName' failed on the 'required' tag"}`,
		},
		{
			name:      "No last name",
			inputBody: `{"first_name":"Test","email":"test@test.com","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.LastName' Error:Field validation for 'LastName' failed on the 'required' tag"}`,
		},
		{
			name:      "Empty password",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test.com","password":""}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name:      "No password",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"test@test.com"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name:      "Empty email",
			inputBody: `{"first_name":"Test","last_name":"TestTest","email":"","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			name:      "Empty email",
			inputBody: `{"first_name":"Test","last_name":"TestTest","password":"TestTest24"}`,
			mockBehavior: func(s *mockhandler.MockAuth, user *core.UserSignUpData, userData *core.SignInResponse) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Key: 'UserSignUpData.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mockhandler.NewMockAuth(c)
			testCase.mockBehavior(auth, &testCase.inputAuthData, &testCase.outputResponse)

			handler := NewHandler(auth)

			// Test handler
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
