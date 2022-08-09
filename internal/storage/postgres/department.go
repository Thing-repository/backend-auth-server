package postgres

import "C"
import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type dbDepartmentDB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type transactionDBDepartmentDB interface {
	ExtractTx(ctx context.Context) (pgx.Tx, bool)
}

type DepartmentDB struct {
	db            dbDepartmentDB
	transactionDB transactionDBDepartmentDB
}

func NewDepartmentDB(dbDriver dbDepartmentDB, transactionDB transactionDBDepartmentDB) *DepartmentDB {
	return &DepartmentDB{db: dbDriver, transactionDB: transactionDB}
}

func (D *DepartmentDB) AddDepartment(ctx context.Context, departmentBase *core.DepartmentBase) (*core.Department, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "AddDepartment",
	}

	db := D.db
	tx, ok := D.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				INSERT INTO departments
				    (department_name, company_id) 
				VALUES 
				    ($1, $2) 
				RETURNING 
					id`

	row := db.QueryRow(ctx, query, departmentBase.DepartmentName, departmentBase.CompanyId)

	var departmentId int

	if err := row.Scan(&departmentId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":           logBase,
					"departmentBase": departmentBase,
					"massage":        pgErr.Message,
					"where":          pgErr.Where,
					"detail":         pgErr.Detail,
					"code":           pgErr.Code,
					"query":          query,
				}).Error("error add department to postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":           logBase,
				"query":          query,
				"error":          err,
				"departmentBase": departmentBase,
			}).Error("error add department to postgres")
			return nil, err
		}
	}
	return &core.Department{
		DepartmentBase: *departmentBase,
		Id:             &departmentId,
	}, nil
}

func (D *DepartmentDB) GetDepartment(ctx context.Context, departmentId int) (*core.Department, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "department.go",
		"function": "GetDepartment",
	}

	db := D.db
	tx, ok := D.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				SELECT
				    id,
					department_name, 
					company_id,
					image_url
				FROM
				    departments
				WHERE 
					id = $1`

	row := db.QueryRow(ctx, query, departmentId)

	var departmentData core.Department

	if err := row.Scan(&departmentData.Id, &departmentData.DepartmentName, &departmentData.CompanyId, &departmentData.ImageURL); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":         logBase,
					"departmentId": departmentId,
					"massage":      pgErr.Message,
					"where":        pgErr.Where,
					"detail":       pgErr.Detail,
					"code":         pgErr.Code,
					"query":        query,
				}).Error("error get department from postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":         logBase,
				"query":        query,
				"departmentId": departmentId,
				"error":        err,
			}).Error("error get department from postgres")
			return nil, err
		}
	}

	return &departmentData, nil
}
