package postgres

import (
	"context"
	"fmt"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
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

func NewCredentialsDB(dbDriver dbDriverRightsDB, transactionDB transactionDBRightsDB) *RightsDB {
	return &RightsDB{dbDriver: dbDriver, transactionDB: transactionDB}
}

func (R *RightsDB) CreateCredential(ctx context.Context,
	credentials *core.AddCredentials) (int, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "CreateCredential",
	}

	var table string

	if slices.Index(core.CompanyCredential, credentials.CredentialType) != -1 {
		table = "company_credentials"
	} else if slices.Index(core.DepartmentCredential, credentials.CredentialType) != -1 {
		table = "department_credentials"
	} else {
		logrus.WithFields(logrus.Fields{
			"base":           logBase,
			"credentials":    credentials,
			"CredentialType": credentials.CredentialType,
		}).Error("error add credential to postgres")
		return 0, moduleErrors.ErrorDataBaseInvalidCredentialsType
	}

	db := R.dbDriver
	tx, ok := R.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
			INSERT INTO %s
				(credential_type, user_id, object_id)
			VALUES 
				($1, $2, $3)
			RETURNING 
				id`

	query = fmt.Sprintf(query, table)

	row := db.QueryRow(ctx, query, credentials.CredentialType, credentials.UserId, credentials.ObjectId)

	var credentialId int

	if err := row.Scan(&credentialId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":        logBase,
					"credentials": credentials,
					"massage":     pgErr.Message,
					"where":       pgErr.Where,
					"detail":      pgErr.Detail,
					"code":        pgErr.Code,
					"query":       query,
				}).Error("error add credential to postgres")
				return 0, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":        logBase,
				"credentials": credentials,
				"query":       query,
				"error":       err,
			}).Error("error add credential to postgres")
			return 0, err
		}
	}

	return credentialId, nil

}

func (R *RightsDB) getCredentials(ctx context.Context,
	userId int, table string) ([]core.Credentials, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "getCredentials",
	}

	db := R.dbDriver
	tx, ok := R.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
			SELECT 
				credential_type, 
				object_id 
			FROM 
				%s 
			WHERE 
				user_id = $1`

	query = fmt.Sprintf(query, table)

	rows, err := db.Query(ctx, query, userId)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			default:
				logrus.WithFields(logrus.Fields{
					"base":    logBase,
					"userId":  userId,
					"table":   table,
					"massage": pgErr.Message,
					"where":   pgErr.Where,
					"detail":  pgErr.Detail,
					"code":    pgErr.Code,
					"query":   logQuery(query),
				}).Error("error getting user credentials from postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":   logBase,
				"userId": userId,
				"table":  table,
				"query":  logQuery(query),
				"error":  err,
			}).Error("error getting user credentials from postgres")
			return nil, err
		}
	}

	var ret []core.Credentials

	for rows.Next() {
		var credential core.Credentials
		err = rows.Scan(&credential.CredentialType, &credential.ObjectId)
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
						"userId":  userId,
					}).Error("error scan rows")
					return nil, err
				}
			} else {
				logrus.WithFields(logrus.Fields{
					"base":   logBase,
					"query":  logQuery(query),
					"error":  err,
					"userId": userId,
				}).Error("error scan rows")
				return nil, err
			}
		}
		ret = append(ret, credential)
	}
	return ret, nil
}

func (R *RightsDB) GetUserCredential(ctx context.Context,
	userId int) ([]core.Credentials, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "getCredentials",
	}

	companyCredentials, err := R.getCredentials(ctx, userId, "company_credentials")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"userId": userId,
			"error":  err,
		}).Error("error getting user credentials from company_credentials postgres")
		return nil, err
	}

	departmentCredentials, err := R.getCredentials(ctx, userId, "department_credentials")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"userId": userId,
			"error":  err,
		}).Error("error getting user credentials from department_credentials postgres")
		return nil, err
	}

	return append(companyCredentials, departmentCredentials...), nil
}
