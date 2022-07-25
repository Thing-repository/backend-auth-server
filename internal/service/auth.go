package service

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type token interface {
	GenerateToken(userId int) (string, error)
}

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type hash interface {
	GenerateHash(password string) (string, error)
	ValidateHash(hash string, password string) error
}

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type userDB interface {
	GetUserByEmail(ctx context.Context, email string) (*core.UserDB, error)
	GetUser(ctx context.Context, userId int) (*core.UserDB, error)
	AddUser(ctx context.Context, user *core.AddUserDB) (*core.UserDB, error)
}

type AuthService struct {
	token token
	db    userDB
	hash  hash
}

func NewAuth(token token, db userDB, hash hash) *AuthService {
	return &AuthService{
		token: token,
		db:    db,
		hash:  hash,
	}
}

func (a *AuthService) SignIn(authData *core.UserSignInData) (*core.SignInResponse, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "signIn",
	}

	ctx := context.TODO()

	// get user data
	userData, err := a.db.GetUserByEmail(ctx, authData.UserMail)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"email": authData.UserMail,
			"error": err,
		}).Error("error get user data")
		switch err {
		case moduleErrors.ErrorDatabaseUserNotFound:
			return nil, moduleErrors.ErrorServiceUserNotFound

		default:
			return nil, err
		}
	}

	// validation password
	err = a.hash.ValidateHash(*userData.PasswordHash, authData.UserPassword)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"email": authData.UserMail,
		}).Error("error validation password")
		switch err {
		case moduleErrors.ErrorHashValidationPassword:
			return nil, moduleErrors.ErrorServiceInvalidPassword
		default:
			return nil, err
		}
	}

	//generate token
	token, err := a.token.GenerateToken(*userData.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"email":  authData.UserMail,
			"userId": userData.Id,
		}).Error("error generate token")
		switch err {
		default:
			return nil, err
		}
	}

	return &core.SignInResponse{
		User:  userData.User,
		Token: token,
	}, nil
}

func (a *AuthService) SignUp(authData *core.UserSignUpData) (*core.SignInResponse, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "signUp",
	}
	// create struct for add user to database
	userDb := core.AddUserDB{
		UserBaseData: authData.UserBaseData,
	}

	// generate hash for user password
	hash, err := a.hash.GenerateHash(authData.Password)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("error generate password hash")
		switch err {
		default:
			return nil, err
		}
	}

	// add hash to data
	userDb.PasswordHash = hash

	ctx := context.TODO()

	// add user to database
	userData, err := a.db.AddUser(ctx, &userDb)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("error add user to database")
		switch err {
		case moduleErrors.ErrorDataBaseUserAlreadyHas:
			return nil, moduleErrors.ErrorServiceUserAlreadyHas
		default:
			return nil, err
		}
	}

	token, err := a.token.GenerateToken(*userData.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"userId": userData.Id,
			"error":  err,
		}).Error("error generate token")
		switch err {
		default:
			return nil, err
		}
	}

	responce := core.SignInResponse{
		User:  userData.User,
		Token: token,
	}

	return &responce, nil
}
