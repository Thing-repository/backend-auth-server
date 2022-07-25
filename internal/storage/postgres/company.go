package postgres

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type dbDriver interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type transactionDB interface {
	ExtractTx(ctx context.Context) (pgx.Tx, bool)
}

type CompanyDB struct {
	db            dbDriver
	transactionDB transactionDB
}

func NewCompanyDB(db dbDriver, transactionDB transactionDB) *CompanyDB {
	return &CompanyDB{
		db:            db,
		transactionDB: transactionDB,
	}
}

func (C *CompanyDB) AddCompany(ctx context.Context, companyBase *core.CompanyBase) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "AddUser",
	}

	db := C.db
	tx, ok := C.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				INSERT INTO companies
				    (company_name, address) 
				VALUES 
				    ($1, $2) 
				RETURNING 
					id`

	row := db.QueryRow(ctx, query, companyBase.CompanyName, companyBase.Address)

	var companyId int

	if err := row.Scan(&companyId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":        logBase,
					"companyBase": companyBase,
					"massage":     pgErr.Message,
					"where":       pgErr.Where,
					"detail":      pgErr.Detail,
					"code":        pgErr.Code,
					"query":       query,
				}).Error("error add company to postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":  logBase,
				"query": query,
				"error": err,
			}).Error("error add company to postgres")
			return nil, err
		}
	}

	return &core.Company{
		CompanyBase: *companyBase,
		Id:          companyId,
	}, nil
}
