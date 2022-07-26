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

type companyDBTransaction interface {
	AddCompany(ctx context.Context, companyBase *core.CompanyBase) (*core.Company, error)
	GetCompany(ctx context.Context, companyId int) (*core.Company, error)
	UpdateCompany(ctx context.Context, company core.Company) error
	DeleteCompany(ctx context.Context, companyId int) error
}

type departmentDBTransaction interface {
	AddDepartment(ctx context.Context, departmentBase *core.DepartmentBase) (*core.Department, error)
}

type rightsDBTransaction interface {
	AddCompanyAdmin(ctx context.Context, userId int, companyId int) (*core.CompanyManager, error)
	AddDepartmentAdmin(ctx context.Context, userId int, departmentId int) (*core.DepartmentManager, error)
}

type transactionDBTransaction interface {
	InjectTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	RollbackTxDefer(ctx context.Context)
}

type Company struct {
	userDB        userDBTransaction
	companyDB     companyDBTransaction
	departmentDB  departmentDBTransaction
	rightsDB      rightsDBTransaction
	transactionDB transactionDBTransaction
}

func NewCompany(userDB userDBTransaction, companyDB companyDBTransaction,
	departmentDB departmentDBTransaction, rightsDB rightsDBTransaction,
	transactionDB transactionDBTransaction) *Company {
	return &Company{userDB: userDB, companyDB: companyDB,
		departmentDB: departmentDB, rightsDB: rightsDB,
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

	_, err = C.rightsDB.AddCompanyAdmin(ctx, user.Id, companyData.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"userId":    user.Id,
			"companyId": companyData.Id,
			"error":     err.Error(),
		}).Error("error add company admin to database")
		return nil, err
	}

	_, err = C.rightsDB.AddDepartmentAdmin(ctx, user.Id, departmentData.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"userId":    user.Id,
			"companyId": departmentData.Id,
			"error":     err.Error(),
		}).Error("error add department admin to database")
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
