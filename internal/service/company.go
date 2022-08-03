package service

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/sirupsen/logrus"
)

const departmentHeadName = "Head"

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type userDBTransaction interface {
	PathUser(ctx context.Context, user *core.UserDB) error
}

type companyDBCompany interface {
	AddCompany(ctx context.Context, companyBase *core.CompanyBase) (*core.Company, error)
	GetCompany(ctx context.Context, companyId int) (*core.Company, error)
	UpdateCompany(ctx context.Context, company core.Company) error
	DeleteCompany(ctx context.Context, companyId int) error
}

type departmentDBCompany interface {
	AddDepartment(ctx context.Context, departmentBase *core.DepartmentBase) (*core.Department, error)
}

type credentialsDBCompany interface {
	CreateCredential(ctx context.Context, credentials core.Credentials) (int, error)
}

type transactionDBCompany interface {
	InjectTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	RollbackTxDefer(ctx context.Context)
}

type Company struct {
	userDB        userDBTransaction
	companyDB     companyDBCompany
	departmentDB  departmentDBCompany
	credentialsDB credentialsDBCompany
	transactionDB transactionDBCompany
}

func NewCompany(userDB userDBTransaction, companyDB companyDBCompany,
	departmentDB departmentDBCompany, credentialsDB credentialsDBCompany,
	transactionDB transactionDBCompany) *Company {
	return &Company{userDB: userDB, companyDB: companyDB,
		departmentDB: departmentDB, credentialsDB: credentialsDB,
		transactionDB: transactionDB}
}

func (C *Company) AddCompany(companyAdd *core.CompanyBase, user *core.User) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "AddCompany",
	}

	ctx := context.TODO()

	ctx, err := C.transactionDB.InjectTx(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error create transaction")
		return nil, err
	}

	defer C.transactionDB.RollbackTxDefer(ctx)

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

	companyAdmin := core.Credentials{
		CredentialType: core.CredentialTypeCompanyAdmin,
		UserId:         user.Id,
		ObjectId:       companyData.Id,
	}

	_, err = C.credentialsDB.CreateCredential(ctx, companyAdmin)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"userId":    user.Id,
			"companyId": companyData.Id,
			"error":     err.Error(),
		}).Error("error add company admin to database")
		return nil, err
	}

	newUserDb := &core.UserDB{
		User: core.User{
			Id:           user.Id,
			CompanyId:    &companyData.Id,
			DepartmentId: &departmentData.Id,
		},
	}

	err = C.userDB.PathUser(ctx, newUserDb)
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

func (C *Company) GetCompany(companyId int) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "GetCompany",
	}

	ctx := context.TODO()

	companyData, err := C.companyDB.GetCompany(ctx, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("error add company to database")
		return nil, err
	}

	return companyData, nil
}

func (C *Company) UpdateCompany(companyBase core.CompanyBase, companyId int) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "UpdateCompany",
	}

	ctx := context.TODO()

	company := core.Company{
		CompanyBase: companyBase,
		Id:          companyId,
	}

	err := C.companyDB.UpdateCompany(ctx, company)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":        logBase,
			"companyId":   companyId,
			"companyBase": companyBase,
			"error":       err.Error(),
		}).Error("error update company in database")
		return nil, err
	}

	return &company, nil
}

func (C *Company) DeleteCompany(companyId int) error {
	logBase := logrus.Fields{
		"module":   "service",
		"function": "UpdateCompany",
	}

	ctx := context.TODO()

	err := C.companyDB.DeleteCompany(ctx, companyId)
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
