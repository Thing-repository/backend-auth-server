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

type dbDriverUserDB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type transactionDBUserDB interface {
	ExtractTx(ctx context.Context) (pgx.Tx, bool)
}

type UserDB struct {
	dbDriver      dbDriverUserDB
	transactionDB transactionDBUserDB
}

func NewUser(dbDriver dbDriverUserDB, transactionDB transactionDBUserDB) *UserDB {
	return &UserDB{
		dbDriver:      dbDriver,
		transactionDB: transactionDB,
	}
}

func (U *UserDB) GetUserByEmail(ctx context.Context, email string) (*core.UserDB, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "GetUserByEmail",
	}

	db := U.dbDriver
	tx, ok := U.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `SELECT 
				id, 
				first_name, 
				last_name, 
				email, 
				image_url, 
				password_hash, 
				company_id, 
				department_id
			FROM 
				users 
			WHERE 
				email = $1`

	row := db.QueryRow(ctx, query, email)

	var userData core.UserDB

	err := row.Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email,
		&userData.ImageURL, &userData.PasswordHash, &userData.CompanyId, &userData.DepartmentId)
	if err != nil {
		switch err.Error() {
		case "no rows in result set":
			logrus.WithFields(logrus.Fields{
				"base":  logBase,
				"email": email,
				"query": query,
				"error": err,
			}).Error("user not found")
			return nil, moduleErrors.ErrorDatabaseUserNotFound
		default:
			logrus.WithFields(logrus.Fields{
				"base":  logBase,
				"query": query,
				"error": err,
			}).Error("error get user by email")
			return nil, err
		}
	}
	return &userData, nil
}

func (U *UserDB) GetUser(ctx context.Context, userId int) (*core.UserDB, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "GetUserByEmail",
	}

	db := U.dbDriver
	tx, ok := U.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
		SELECT 
				id, 
				first_name, 
				last_name, 
				email, 
				image_url, 
				password_hash, 
				company_id, 
				department_id
			FROM 
				users 
			WHERE 
				id = $1`

	row := db.QueryRow(ctx, query, userId)

	var userData core.UserDB

	err := row.Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email,
		&userData.ImageURL, &userData.PasswordHash, &userData.CompanyId, &userData.DepartmentId)
	if err != nil {
		switch err.Error() {
		case "no rows in result set":
			logrus.WithFields(logrus.Fields{
				"base":   logBase,
				"userId": userId,
				"query":  query,
				"error":  err,
			}).Error("user not found")
			return nil, moduleErrors.ErrorDatabaseUserNotFound
		default:
			logrus.WithFields(logrus.Fields{
				"base":   logBase,
				"userId": userId,
				"query":  query,
				"error":  err,
			}).Error("error get user by email")
			return nil, err
		}
	}
	return &userData, nil
}

func (U *UserDB) AddUser(ctx context.Context, user *core.AddUserDB) (*core.UserDB, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "AddUser",
	}

	db := U.dbDriver
	tx, ok := U.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
				INSERT INTO users
				    (first_name, last_name, email, password_hash) 
				VALUES 
				    ($1, $2, $3, $4) 
				RETURNING 
					id`

	row := db.QueryRow(ctx, query, user.FirstName, user.LastName, user.Email, user.PasswordHash)

	var userId int
	if err := row.Scan(&userId); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				logrus.WithFields(logrus.Fields{
					"base":    logBase,
					"email":   user.Email,
					"massage": pgErr.Message,
					"where":   pgErr.Where,
					"detail":  pgErr.Detail,
					"code":    pgErr.Code,
					"query":   query,
				}).Error("user already has")
				return nil, moduleErrors.ErrorDataBaseUserAlreadyHas
			default:
				logrus.WithFields(logrus.Fields{
					"base":    logBase,
					"email":   user.Email,
					"massage": pgErr.Message,
					"where":   pgErr.Where,
					"detail":  pgErr.Detail,
					"code":    pgErr.Code,
					"query":   query,
				}).Error("error add user to postgres")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":  logBase,
				"query": query,
				"error": err,
			}).Error("error add user to postgres")
			return nil, err
		}
	}

	ret := core.UserDB{
		User: core.User{
			UserBaseData: user.UserBaseData,
			Id:           userId,
		},
		PasswordHash: &user.PasswordHash,
	}

	return &ret, nil
}

func (U *UserDB) PathUser(ctx context.Context, user *core.UserDB, userId int) error {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "PathUser",
	}

	db := U.dbDriver
	tx, ok := U.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name = $%d", argId))
		args = append(args, *user.FirstName)
		argId++
	}
	if user.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name = $%d", argId))
		args = append(args, *user.LastName)
		argId++
	}
	if user.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email = $%d", argId))
		args = append(args, *user.Email)
		argId++
	}
	if user.EmailIsValidated != nil {
		setValues = append(setValues, fmt.Sprintf("email_is_validated = $%d", argId))
		args = append(args, *user.EmailIsValidated)
		argId++
	}
	if user.ImageURL != nil {
		setValues = append(setValues, fmt.Sprintf("image_url = $%d", argId))
		args = append(args, *user.ImageURL)
		argId++
	}
	if user.CompanyId != nil {
		setValues = append(setValues, fmt.Sprintf("company_id = $%d", argId))
		args = append(args, *user.CompanyId)
		argId++
	}
	if user.DepartmentId != nil {
		setValues = append(setValues, fmt.Sprintf("department_id = $%d", argId))
		args = append(args, *user.DepartmentId)
		argId++
	}
	if user.PasswordHash != nil {
		setValues = append(setValues, fmt.Sprintf("password_hash = $%d", argId))
		args = append(args, *user.PasswordHash)
		argId++
	}
	if user.EmailValidationToken != nil {
		setValues = append(setValues, fmt.Sprintf("email_validation_token = $%d", argId))
		args = append(args, *user.EmailValidationToken)
		argId++
	}

	if argId == 1 {
		return moduleErrors.ErrorDataBaseHasNotDataToChange
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`
				UPDATE 
					users 
				SET 
					%s 
				WHERE 
					id = $%d
`, setQuery, argId)

	args = append(args, userId)

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
				}).Error("error path user to postgres")
				return err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":   logBase,
				"query":  logQuery(query),
				"error":  err,
				"args":   args,
				"cmdTag": cmdTag,
			}).Error("error path user to postgres")
			return err
		}
	}

	return nil
}

func (U *UserDB) UserInCompany(ctx context.Context, companyId int) ([]int, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "UserInCompany",
	}

	db := U.dbDriver
	tx, ok := U.transactionDB.ExtractTx(ctx)
	if ok {
		db = tx
	}

	query := `
		SELECT 
			id
		FROM 
			users 
		WHERE 
			company_id = $1`

	var ret []int
	rows, err := db.Query(ctx, query, companyId)

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
				}).Error("error find user in company")
				return nil, err
			}
		} else {
			logrus.WithFields(logrus.Fields{
				"base":      logBase,
				"query":     logQuery(query),
				"error":     err,
				"companyId": companyId,
			}).Error("error find user in company")
			return nil, err
		}
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
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
					}).Error("error scan rows")
					return nil, err
				}
			} else {
				logrus.WithFields(logrus.Fields{
					"base":      logBase,
					"query":     logQuery(query),
					"error":     err,
					"companyId": companyId,
				}).Error("error scan rows")
				return nil, err
			}
		}
		ret = append(ret, id)
	}

	return ret, nil
}
