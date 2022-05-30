package service

import (
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type token interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(userId int) error
}

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type hash interface {
	GenerateHash(password string) (string, error)
	ValidateHash(password string, hash string) error
}

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type db interface {
	GetUserByEmail(email string) (*core.UserDB, error)
	GetUser(userId int) (*core.UserDB, error)
	AddUser(user *core.AddUserDB) (*core.UserDB, error)
}

type Auth struct {
	token token
	db    db
	hash  hash
}

func NewAuth(token token, db db, hash hash) *Auth {
	return &Auth{
		token: token,
		db:    db,
		hash:  hash,
	}
}

func (a *Auth) SignIn(authData *core.UserSignInData) (*core.SignInResponse, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "signIn",
	}

	// get user data
	userData, err := a.db.GetUserByEmail(authData.UserMail)
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
	err = a.hash.ValidateHash(authData.UserPassword, userData.PasswordHash)
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
	token, err := a.token.GenerateToken(userData.Id)
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

func (a *Auth) SignUp(authData *core.UserSignUpData) (*core.SignInResponse, error) {
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

	// add user to database
	userData, err := a.db.AddUser(&userDb)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("error generate password hash")
		switch err {
		case moduleErrors.ErrorDataBaseUserAlreadyHas:
			return nil, moduleErrors.ErrorServiceUserAlreadyHas
		default:
			return nil, err
		}
	}

	token, err := a.token.GenerateToken(userData.Id)
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
