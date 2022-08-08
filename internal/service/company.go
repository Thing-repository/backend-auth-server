package service

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/sirupsen/logrus"
)

const departmentHeadName = "Head"

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type userDBCompany interface {
	PathUser(ctx context.Context, user *core.UserDB, userId int) error
	GetUser(ctx context.Context, userId int) (*core.UserDB, error)
}

type companyDBCompany interface {
	AddCompany(ctx context.Context, companyBase *core.CompanyBase) (*core.Company, error)
	GetCompany(ctx context.Context, companyId int) (*core.Company, error)
	UpdateCompany(ctx context.Context, company core.CompanyUpdate, companyId int) error
	DeleteCompany(ctx context.Context, companyId int) error
}

type departmentDBCompany interface {
	AddDepartment(ctx context.Context, departmentBase *core.DepartmentBase) (*core.Department, error)
}

type credentialsDBCompany interface {
	CreateCredential(ctx context.Context, credentials *core.AddCredentials) (int, error)
}

type transactionDBCompany interface {
	InjectTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	RollbackTxDefer(ctx context.Context)
}

type Company struct {
	userDB        userDBCompany
	companyDB     companyDBCompany
	departmentDB  departmentDBCompany
	credentialsDB credentialsDBCompany
	transactionDB transactionDBCompany
}

func NewCompany(userDB userDBCompany, companyDB companyDBCompany,
	departmentDB departmentDBCompany, credentialsDB credentialsDBCompany,
	transactionDB transactionDBCompany) *Company {
	return &Company{userDB: userDB, companyDB: companyDB,
		departmentDB: departmentDB, credentialsDB: credentialsDB,
		transactionDB: transactionDB}
}

func (C *Company) AddCompany(ctx context.Context, companyAdd *core.CompanyBase) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "AddCompany",
		"context":  *core.LogContext(ctx),
	}

	userId, err := core.ContextGetUserId(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user id from context")
		return nil, moduleErrors.ErrorServiceInvalidContext
	}

	ctx, err = C.transactionDB.InjectTx(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error create transaction")
		return nil, err
	}

	defer C.transactionDB.RollbackTxDefer(ctx)

	userData, err := C.userDB.GetUser(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user data")
		return nil, moduleErrors.ErrorServiceGetUserData
	}

	if userData.CompanyId != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": userData,
		}).Error(moduleErrors.ErrorServiceUserAlreadyHasCompany.Error())

		return nil, moduleErrors.ErrorServiceUserAlreadyHasCompany
	}

	companyData, err := C.companyDB.AddCompany(ctx, companyAdd)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":       logBase,
			"companyAdd": companyAdd,
			"error":      err.Error(),
		}).Error("error add company to database")
		return nil, err
	}

	departmentAdd := &core.DepartmentBase{
		DepartmentName: departmentHeadName,
		CompanyId:      companyData.Id,
	}

	departmentData, err := C.departmentDB.AddDepartment(ctx, departmentAdd)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":          logBase,
			"departmentAdd": departmentAdd,
			"error":         err.Error(),
		}).Error("error add department to database")
		return nil, err
	}

	_, err = C.credentialsDB.CreateCredential(ctx, newCredential(companyData.Id, userId, core.CredentialTypeCompanyAdmin))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyData.Id,
			"userId":    userId,
			"error":     err.Error(),
		}).Error("error add company admin to database")
		return nil, err
	}

	_, err = C.credentialsDB.CreateCredential(ctx, newCredential(companyData.Id, userId, core.CredentialTypeCompanyUser))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyData.Id,
			"userId":    userId,
			"error":     err.Error(),
		}).Error("error add company user to database")
		return nil, err
	}

	_, err = C.credentialsDB.CreateCredential(ctx, newCredential(departmentData.Id, userId, core.CredentialTypeDepartmentAdmin))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyData.Id,
			"userId":    userId,
			"error":     err.Error(),
		}).Error("error add department admin to database")
		return nil, err
	}

	_, err = C.credentialsDB.CreateCredential(ctx, newCredential(departmentData.Id, userId, core.CredentialTypeDepartmentUser))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyData.Id,
			"userId":    userId,
			"error":     err.Error(),
		}).Error("error add department admin to database")
		return nil, err
	}

	newUserDb := &core.UserDB{
		User: core.User{
			CompanyId:    &companyData.Id,
			DepartmentId: &departmentData.Id,
		},
	}

	err = C.userDB.PathUser(ctx, newUserDb, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"newUserDb": newUserDb,
			"error":     err.Error(),
		}).Error("error change user data in database")
		return nil, err
	}

	if err = C.transactionDB.CommitTx(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error commit transaction")
		return nil, err
	}

	return companyData, nil
}

func (C *Company) GetCompany(ctx context.Context, companyId int) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "GetCompany",
		"context":  *core.LogContext(ctx),
	}

	credentials, err := core.ContextGetUserCredentials(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user credentials from context")
		return nil, moduleErrors.ErrorServiceInvalidContext
	}

	// TODO: add check rights for service admin, load and check permission form context
	if !core.CheckCredential(credentials, core.CredentialTypeCompanyUser, companyId) {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
		}).Error("user has not credentials")
		return nil, moduleErrors.ErrorServiceBadPermissions
	}

	companyData, err := C.companyDB.GetCompany(ctx, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("error get company from database")
		return nil, err
	}

	return companyData, nil
}

func (C *Company) UpdateCompany(ctx context.Context, companyBase core.CompanyBase, companyId int) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "UpdateCompany",
		"context":  *core.LogContext(ctx),
	}

	credentials, err := core.ContextGetUserCredentials(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user credentials from context")
		return nil, moduleErrors.ErrorServiceInvalidContext
	}

	if !core.CheckCredential(credentials, core.CredentialTypeCompanyAdmin, companyId) {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
		}).Error("user has not credentials")
		return nil, moduleErrors.ErrorServiceBadPermissions
	}

	company := core.CompanyUpdate{
		CompanyBase: companyBase,
	}

	err = C.companyDB.UpdateCompany(ctx, company, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":        logBase,
			"companyId":   companyId,
			"companyBase": companyBase,
			"error":       err.Error(),
		}).Error("error update company in database")
		return nil, err
	}

	return C.GetCompany(ctx, companyId)
}

func (C *Company) DeleteCompany(ctx context.Context, companyId int) error {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "UpdateCompany",
	}

	credentials, err := core.ContextGetUserCredentials(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user credentials from context")
		return moduleErrors.ErrorServiceInvalidContext
	}

	if !core.CheckCredential(credentials, core.CredentialTypeCompanyAdmin, companyId) {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
		}).Error("user has not credentials")
		return moduleErrors.ErrorServiceBadPermissions
	}

	err = C.companyDB.DeleteCompany(ctx, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("error delete company from database")
		return err
	}

	return nil
}
