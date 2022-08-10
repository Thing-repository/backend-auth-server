package service

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/sirupsen/logrus"
)

type userDBUser interface {
	PathUser(ctx context.Context, user *core.UserDB, userId int) error
	GetUser(ctx context.Context, userId int) (*core.UserDB, error)
	GetUsersFilter(ctx context.Context, filter string, limit int, offset int) ([]core.User, error)
}

type departmentDBUser interface {
	GetDepartment(ctx context.Context, departmentId int) (*core.Department, error)
}

type credentialsDBUser interface {
	CreateCredential(ctx context.Context, credentials *core.AddCredentials) (int, error)
}

type transactionDBUser interface {
	InjectTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	RollbackTxDefer(ctx context.Context)
}

type User struct {
	userDB        userDBUser
	departmentDB  departmentDBUser
	credentialsDB credentialsDBUser
	transactionDB transactionDBUser
}

func NewUser(userDB userDBUser, departmentDB departmentDBUser, credentialsDB credentialsDBUser, transactionDB transactionDBUser) *User {
	return &User{
		userDB:        userDB,
		departmentDB:  departmentDB,
		credentialsDB: credentialsDB,
		transactionDB: transactionDB,
	}
}

func (U *User) FindUsersForInvite(ctx context.Context, filter string, limit int, offset int) ([]core.User, error) {
	logBase := logrus.Fields{
		"module":   "user",
		"function": "FindUsersForInvite",
		"filter":   filter,
		"limit":    limit,
		"offset":   offset,
		"context":  *core.LogContext(ctx),
	}

	users, err := U.userDB.GetUsersFilter(ctx, filter, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get users from database")
		return nil, err
	}

	return users, nil
}

func (U *User) AddUserToCompany(ctx context.Context, userId int, departmentId int) error {
	logBase := logrus.Fields{
		"module":       "user",
		"function":     "AddUserToCompany",
		"userId":       userId,
		"departmentId": departmentId,
		"context":      *core.LogContext(ctx),
	}
	ctx, err := U.transactionDB.InjectTx(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error create transaction")
		return err
	}

	defer U.transactionDB.RollbackTxDefer(ctx)

	credentials, err := core.ContextGetUserCredentials(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user credentials from context")
		return moduleErrors.ErrorServiceInvalidContext
	}

	departmentData, err := U.departmentDB.GetDepartment(ctx, departmentId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"departmentId": departmentId,
			"error":        err.Error(),
		}).Error("error get department data from db")
		return err
	}

	// TODO: add check rights for service admin, load and check permission form context
	isCompanyAdmin := core.CheckCredential(credentials, core.CredentialTypeCompanyAdmin, *departmentData.CompanyId)
	isDepartmentAdmin := core.CheckCredential(credentials, core.CredentialTypeDepartmentAdmin, departmentId)

	if !isCompanyAdmin || !isDepartmentAdmin {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"companyId":    departmentData.CompanyId,
			"departmentId": departmentId,
		}).Error("user has not credentials")
		return moduleErrors.ErrorServiceBadPermissions
	}

	userData, err := U.userDB.GetUser(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"userId": userId,
			"error":  err.Error(),
		}).Error("error get user data from db")
		return err
	}
	if userData.CompanyId != nil || userData.DepartmentId != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"userId": userId,
		}).Error("user already has company")
		return moduleErrors.ErrorServiceUserAlreadyHasCompany
	}

	newUserData := core.UserDB{
		User: core.User{
			CompanyId:    departmentData.CompanyId,
			DepartmentId: &departmentId,
		},
	}
	err = U.userDB.PathUser(ctx, &newUserData, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":        logBase,
			"userId":      userId,
			"newUserData": newUserData,
			"error":       err.Error(),
		}).Error("error update user data in db")
		return err
	}

	_, err = U.credentialsDB.CreateCredential(ctx, newCredential(departmentId, userId, core.CredentialTypeDepartmentUser))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"departmentId": departmentId,
			"userId":       userId,
			"error":        err.Error(),
		}).Error("error add department admin to database")
		return err
	}

	_, err = U.credentialsDB.CreateCredential(ctx, newCredential(*departmentData.CompanyId, userId, core.CredentialTypeCompanyUser))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": departmentData.CompanyId,
			"userId":    userId,
			"error":     err.Error(),
		}).Error("error add department admin to database")
		return err
	}

	if err = U.transactionDB.CommitTx(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error commit transaction")
		return err
	}

	return nil
}
