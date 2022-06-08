package service

import (
	mockService "github.com/Thing-repository/backend-server/internal/service/mock"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	res := m.Run()
	os.Exit(res)
}

func TestSignIn(t *testing.T) {
	type generateTokenMockBehavior func(s *mockService.Mocktoken, userId int, token string)
	type getUserByEmailMockBehavior func(s *mockService.Mockdb, email string, userDB *core.UserDB)
	type validateHashMockBehavior func(s *mockService.Mockhash, hash string, password string)

	testImageUrl := "test_image"

	testTable := []struct {
		name                       string
		inputAuthData              core.UserSignInData
		outputResponse             *core.SignInResponse
		outputError                error
		token                      string
		userDb                     core.UserDB
		generateTokenMockBehavior  generateTokenMockBehavior
		getUserByEmailMockBehavior getUserByEmailMockBehavior
		validateHashMockBehavior   validateHashMockBehavior
	}{
		{
			name: "Success",
			inputAuthData: core.UserSignInData{
				UserPassword: "test_password",
			},
			outputResponse: new(core.SignInResponse),
			outputError:    nil,
			token:          "bar",
			userDb: core.UserDB{
				User: core.User{
					UserChange: core.UserChange{
						UserBaseData: core.UserBaseData{
							FirstName: "test_name",
							LastName:  "test_last_name",
							Email:     "foo@example.com",
						},
					},
					Id:           1,
					ImageURL:     &testImageUrl,
					CompanyId:    nil,
					DepartmentId: nil,
				},
				PasswordHash: "test_password_hash",
			},
			generateTokenMockBehavior: func(s *mockService.Mocktoken, userId int, token string) {
				s.EXPECT().GenerateToken(userId).Return(token, nil)
			},
			getUserByEmailMockBehavior: func(s *mockService.Mockdb, email string, userDb *core.UserDB) {
				s.EXPECT().GetUserByEmail(email).Return(userDb, nil)
			},
			validateHashMockBehavior: func(s *mockService.Mockhash, hash string, password string) {
				s.EXPECT().ValidateHash(hash, password).Return(nil)
			},
		},
		{
			name: "User not found",
			inputAuthData: core.UserSignInData{
				UserMail:     "foo@example.com",
				UserPassword: "test_password",
			},
			outputError: moduleErrors.ErrorServiceUserNotFound,
			getUserByEmailMockBehavior: func(s *mockService.Mockdb, email string, userDb *core.UserDB) {
				s.EXPECT().GetUserByEmail(email).Return(nil, moduleErrors.ErrorDatabaseUserNotFound)
			},
		},
		{
			name: "Invalid password",
			inputAuthData: core.UserSignInData{
				UserPassword: "test_password",
			},
			outputError: moduleErrors.ErrorServiceInvalidPassword,
			userDb: core.UserDB{
				User: core.User{
					UserChange: core.UserChange{
						UserBaseData: core.UserBaseData{
							FirstName: "test_name",
							LastName:  "test_last_name",
							Email:     "foo@example.com",
						},
					},
					Id:           1,
					ImageURL:     &testImageUrl,
					CompanyId:    nil,
					DepartmentId: nil,
				},
				PasswordHash: "test_password_hash",
			},
			getUserByEmailMockBehavior: func(s *mockService.Mockdb, email string, userDb *core.UserDB) {
				s.EXPECT().GetUserByEmail(email).Return(userDb, nil)
			},
			validateHashMockBehavior: func(s *mockService.Mockhash, hash string, password string) {
				s.EXPECT().ValidateHash(hash, password).Return(moduleErrors.ErrorHashValidationPassword)
			},
		},
		{
			name: "Database error",
			inputAuthData: core.UserSignInData{
				UserPassword: "test_password",
			},
			outputError: moduleErrors.ErrorDataBaseInternal,
			getUserByEmailMockBehavior: func(s *mockService.Mockdb, email string, userDb *core.UserDB) {
				s.EXPECT().GetUserByEmail(email).Return(nil, moduleErrors.ErrorDataBaseInternal)
			},
		},
	}

	for _, testCase := range testTable {
		if testCase.inputAuthData.UserMail == "" {
			testCase.inputAuthData.UserMail = testCase.userDb.User.Email
		}
		emptyUser := core.User{}
		if testCase.outputResponse != nil && testCase.userDb.User != emptyUser {
			testCase.outputResponse.User = testCase.userDb.User
		}
		if testCase.outputResponse != nil && testCase.outputResponse.Token == "" {
			testCase.outputResponse.Token = testCase.token
		}

		t.Run(testCase.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			token := mockService.NewMocktoken(c)
			hash := mockService.NewMockhash(c)
			db := mockService.NewMockdb(c)

			if testCase.generateTokenMockBehavior != nil {
				testCase.generateTokenMockBehavior(token, testCase.userDb.User.Id, testCase.token)
			}
			if testCase.validateHashMockBehavior != nil {
				testCase.validateHashMockBehavior(hash, testCase.userDb.PasswordHash, testCase.inputAuthData.UserPassword)
			}
			if testCase.getUserByEmailMockBehavior != nil {
				testCase.getUserByEmailMockBehavior(db, testCase.inputAuthData.UserMail, &testCase.userDb)
			}
			auth := NewAuth(token, db, hash)

			// Test handler
			resp, err := auth.SignIn(&testCase.inputAuthData)

			// Assert
			assert.Equal(t, testCase.outputError, err)
			assert.Equal(t, testCase.outputResponse, resp)
		})
	}
}

func TestSignUp(t *testing.T) {
	type generateTokenMockBehavior func(s *mockService.Mocktoken, userId int, token string)
	type addUserMockBehavior func(s *mockService.Mockdb, userAddDb *core.AddUserDB, userDb *core.UserDB)
	type generateHashMockBehavior func(s *mockService.Mockhash, password string, hash string)

	testTable := []struct {
		name                      string
		inputSingUpData           core.UserSignUpData
		outputResponse            *core.SignInResponse
		outputError               error
		token                     string
		userDb                    core.UserDB
		generateTokenMockBehavior generateTokenMockBehavior
		addUserMockBehavior       addUserMockBehavior
		generateHashMockBehavior  generateHashMockBehavior
	}{
		{
			name: "Success",
			inputSingUpData: core.UserSignUpData{
				UserBaseData: core.UserBaseData{
					FirstName: "Test_name",
					LastName:  "Test_last_name",
					Email:     "foo@example.com",
				},
				Password: "TestPassword123456",
			},
			outputResponse: new(core.SignInResponse),
			outputError:    nil,
			token:          "bar",
			userDb: core.UserDB{
				User: core.User{
					Id: 1,
				},
				PasswordHash: "test_password_hash",
			},
			generateTokenMockBehavior: func(s *mockService.Mocktoken, userId int, token string) {
				s.EXPECT().GenerateToken(userId).Return(token, nil)
			},
			addUserMockBehavior: func(s *mockService.Mockdb, userAddDb *core.AddUserDB, userDb *core.UserDB) {
				s.EXPECT().AddUser(userAddDb).Return(userDb, nil)
			},
			generateHashMockBehavior: func(s *mockService.Mockhash, hash string, password string) {
				s.EXPECT().GenerateHash(password).Return(hash, nil)
			},
		},
		{
			name: "User already has",
			inputSingUpData: core.UserSignUpData{
				UserBaseData: core.UserBaseData{
					FirstName: "Test_name",
					LastName:  "Test_last_name",
					Email:     "foo@example.com",
				},
				Password: "TestPassword123456",
			},
			userDb: core.UserDB{
				PasswordHash: "test_password_hash",
			},
			outputError: moduleErrors.ErrorServiceUserAlreadyHas,
			addUserMockBehavior: func(s *mockService.Mockdb, userAddDb *core.AddUserDB, userDb *core.UserDB) {
				s.EXPECT().AddUser(userAddDb).Return(nil, moduleErrors.ErrorDataBaseUserAlreadyHas)
			},
			generateHashMockBehavior: func(s *mockService.Mockhash, hash string, password string) {
				s.EXPECT().GenerateHash(password).Return(hash, nil)
			},
		},
		{
			name: "Database error",
			inputSingUpData: core.UserSignUpData{
				UserBaseData: core.UserBaseData{
					FirstName: "Test_name",
					LastName:  "Test_last_name",
					Email:     "foo@example.com",
				},
				Password: "TestPassword123456",
			},
			userDb: core.UserDB{
				PasswordHash: "test_password_hash",
			},
			outputError: moduleErrors.ErrorDataBaseInternal,
			addUserMockBehavior: func(s *mockService.Mockdb, userAddDb *core.AddUserDB, userDb *core.UserDB) {
				s.EXPECT().AddUser(userAddDb).Return(nil, moduleErrors.ErrorDataBaseInternal)
			},
			generateHashMockBehavior: func(s *mockService.Mockhash, hash string, password string) {
				s.EXPECT().GenerateHash(password).Return(hash, nil)
			},
		},
	}

	for _, testCase := range testTable {
		testCase.userDb.User.Email = testCase.inputSingUpData.Email
		testCase.userDb.User.FirstName = testCase.inputSingUpData.FirstName
		testCase.userDb.User.LastName = testCase.inputSingUpData.LastName

		emptyUser := core.User{}
		if testCase.outputResponse != nil && testCase.userDb.User != emptyUser {
			testCase.outputResponse.User = testCase.userDb.User
		}
		if testCase.outputResponse != nil && testCase.outputResponse.Token == "" {
			testCase.outputResponse.Token = testCase.token
		}

		t.Run(testCase.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			token := mockService.NewMocktoken(c)
			hash := mockService.NewMockhash(c)
			db := mockService.NewMockdb(c)

			if testCase.generateTokenMockBehavior != nil {
				testCase.generateTokenMockBehavior(token, testCase.userDb.User.Id, testCase.token)
			}
			if testCase.generateHashMockBehavior != nil {
				testCase.generateHashMockBehavior(hash, testCase.userDb.PasswordHash, testCase.inputSingUpData.Password)
			}
			if testCase.addUserMockBehavior != nil {
				testCase.addUserMockBehavior(db, &core.AddUserDB{
					UserBaseData: testCase.inputSingUpData.UserBaseData,
					PasswordHash: testCase.userDb.PasswordHash,
				}, &testCase.userDb)
			}
			auth := NewAuth(token, db, hash)

			// Test handler
			resp, err := auth.SignUp(&testCase.inputSingUpData)

			// Assert
			assert.Equal(t, testCase.outputError, err)
			assert.Equal(t, testCase.outputResponse, resp)
		})
	}
}
