package postgres

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type dbDriverRightsDB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type transactionDBRightsDB interface {
	ExtractTx(ctx context.Context) (pgx.Tx, bool)
}

type RightsDB struct {
	dbDriver      dbDriverRightsDB
	transactionDB transactionDBRightsDB
}

func NewRightsDB(dbDriver dbDriverRightsDB, transactionDB transactionDBRightsDB) *RightsDB {
	return &RightsDB{dbDriver: dbDriver, transactionDB: transactionDB}
}

func (R *RightsDB) AddCompanyAdmin(ctx context.Context, userId int, companyId int) (*core.CompanyManager, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "AddCompanyAdmin",
	}

	db := R.dbDriver
	tx, ok := R.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
			INSERT INTO companies_admins	
				(user_id, company_id)
			VALUES 
				($1, $2)
			RETURNING 
				id`

	row := db.QueryRow(ctx, query, userId, companyId)

	var companyAdminId int

	if err := row.Scan(&companyAdminId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":      logBase,
					"userId":    userId,
					"companyId": companyId,
					"massage":   pgErr.Message,
					"where":     pgErr.Where,
					"detail":    pgErr.Detail,
					"code":      pgErr.Code,
					"query":     query,
				}).Error("error add company admin to postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":      logBase,
				"userId":    userId,
				"companyId": companyId,
				"query":     query,
				"error":     err,
			}).Error("error add company admin to postgres")
			return nil, err
		}
	}

	return &core.CompanyManager{
		Id:        companyAdminId,
		UserId:    userId,
		CompanyId: companyId,
	}, nil

}

func (R *RightsDB) AddDepartmentAdmin(ctx context.Context, userId int, departmentId int) (*core.DepartmentManager, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "AddDepartmentAdmin",
	}

	db := R.dbDriver
	tx, ok := R.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
			INSERT INTO departments_admins	
				(user_id, department_id)
			VALUES 
				($1, $2)
			RETURNING 
				id`

	row := db.QueryRow(ctx, query, userId, departmentId)

	var companyAdminId int

	if err := row.Scan(&companyAdminId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":      logBase,
					"userId":    userId,
					"companyId": departmentId,
					"massage":   pgErr.Message,
					"where":     pgErr.Where,
					"detail":    pgErr.Detail,
					"code":      pgErr.Code,
					"query":     query,
				}).Error("error add department admin to postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":      logBase,
				"userId":    userId,
				"companyId": departmentId,
				"query":     query,
				"error":     err,
			}).Error("error add department admin to postgres")
			return nil, err
		}
	}

	return &core.DepartmentManager{
		Id:           companyAdminId,
		UserId:       userId,
		DepartmentId: departmentId,
	}, nil

}
