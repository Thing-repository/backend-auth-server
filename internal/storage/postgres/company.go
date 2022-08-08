package postgres

import (
	"context"
	"fmt"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"strings"
)

type dbDriver interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
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
		"file":     "company.go",
		"function": "AddCompany",
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
		CompanyUpdate: core.CompanyUpdate{
			CompanyBase: *companyBase,
		},
		Id: companyId,
	}, nil
}

func (C *CompanyDB) GetCompany(ctx context.Context, companyId int) (*core.Company, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "company.go",
		"function": "GetCompany",
	}

	db := C.db
	tx, ok := C.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				SELECT
				    id,
					company_name, 
					address,
					image_url
				FROM
				    companies
				WHERE 
					id = $1`

	row := db.QueryRow(ctx, query, companyId)

	var companyData core.Company

	if err := row.Scan(&companyData.Id, &companyData.CompanyName, &companyData.Address, &companyData.ImageURL); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":      logBase,
					"companyId": companyId,
					"massage":   pgErr.Message,
					"where":     pgErr.Where,
					"detail":    pgErr.Detail,
					"code":      pgErr.Code,
					"query":     query,
				}).Error("error get company from postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":      logBase,
				"query":     query,
				"companyId": companyId,
				"error":     err,
			}).Error("error get company from postgres")
			return nil, err
		}
	}

	return &companyData, nil
}

func (C *CompanyDB) UpdateCompany(ctx context.Context, company core.CompanyUpdate, companyId int) error {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "company.go",
		"function": "UpdateCompany",
	}

	db := C.db
	tx, ok := C.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if company.CompanyName != nil {
		setValues = append(setValues, fmt.Sprintf("company_name = $%d", argId))
		args = append(args, *company.CompanyName)
		argId++
	}
	if company.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address = $%d", argId))
		args = append(args, *company.Address)
		argId++
	}
	if argId == 1 {
		return moduleErrors.ErrorDataBaseHasNotDataToChange
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`
				UPDATE 
					companies 
				SET 
					%s 
				WHERE 
					id = $%d
`, setQuery, argId)

	args = append(args, companyId)

	cmdTag, err := db.Exec(ctx, query, args...)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":    logBase,
					"massage": pgErr.Message,
					"where":   pgErr.Where,
					"detail":  pgErr.Detail,
					"code":    pgErr.Code,
					"query":   logQuery(query),
					"args":    args,
					"cmdTag":  cmdTag,
				}).Error("error update company to postgres")
				return err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":   logBase,
				"query":  logQuery(query),
				"error":  err,
				"args":   args,
				"cmdTag": cmdTag,
			}).Error("error update company to postgres")
			return err
		}
	}

	return nil
}

func (C *CompanyDB) DeleteCompany(ctx context.Context, companyId int) error {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "company.go",
		"function": "UpdateCompany",
	}

	db := C.db
	tx, ok := C.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				DELETE
				FROM
				    companies
				WHERE 
					id = $1`

	cmdTag, err := db.Exec(ctx, query, companyId)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":      logBase,
					"massage":   pgErr.Message,
					"where":     pgErr.Where,
					"detail":    pgErr.Detail,
					"code":      pgErr.Code,
					"query":     logQuery(query),
					"companyId": companyId,
					"cmdTag":    cmdTag,
				}).Error("error delete company from postgres")
				return err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":      logBase,
				"query":     logQuery(query),
				"error":     err,
				"companyId": companyId,
				"cmdTag":    cmdTag,
			}).Error("error delete company from postgres")
			return err
		}
	}

	return nil
}
