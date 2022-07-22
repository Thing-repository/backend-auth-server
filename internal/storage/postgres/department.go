package postgres

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type dbDriverDepartmentDB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type transactionDBDepartmentDB interface {
	ExtractTx(ctx context.Context) (pgx.Tx, bool)
}

type DepartmentDB struct {
	dbDriver      dbDriverDepartmentDB
	transactionDB transactionDBDepartmentDB
}

func NewDepartmentDB(dbDriver dbDriverDepartmentDB, transactionDB transactionDBDepartmentDB) *DepartmentDB {
	return &DepartmentDB{dbDriver: dbDriver, transactionDB: transactionDB}
}

func (D *DepartmentDB) AddDepartment(ctx context.Context, departmentBase *core.DepartmentBase) (*core.Department, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "AddDepartment",
	}

	db := D.dbDriver
	tx, ok := D.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				INSERT INTO departments
				    (department_name) 
				VALUES 
				    ($1) 
				RETURNING 
					id`

	row := db.QueryRow(ctx, query, departmentBase)

	var departmentId int

	if err := row.Scan(&departmentId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":        logBase,
					"companyBase": departmentBase,
					"massage":     pgErr.Message,
					"where":       pgErr.Where,
					"detail":      pgErr.Detail,
					"code":        pgErr.Code,
					"query":       query,
				}).Error("error add department to postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":  logBase,
				"query": query,
				"error": err,
			}).Error("error add department to postgres")
			return nil, err
		}
	}
	return &core.Department{
		DepartmentBase: *departmentBase,
		Id:             departmentId,
	}, nil
}
