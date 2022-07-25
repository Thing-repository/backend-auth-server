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
				department_id,
				vacation_time_start, 
				vacation_time_end 
			FROM 
				users 
			WHERE 
				email = $1`

	row := db.QueryRow(ctx, query, email)

	var userData core.UserDB

	err := row.Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email,
		&userData.ImageURL, &userData.PasswordHash, &userData.CompanyId, &userData.DepartmentId,
		&userData.VacationTimeStart, &userData.VacationTimeEnd)
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
				department_id,
				vacation_time_start, 
				vacation_time_end 
			FROM 
				users 
			WHERE 
				id = $1`

	row := db.QueryRow(ctx, query, userId)

	var userData core.UserDB

	err := row.Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email,
		&userData.ImageURL, &userData.PasswordHash, &userData.CompanyId, &userData.DepartmentId,
		&userData.VacationTimeStart, &userData.VacationTimeEnd)
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
			UserChange: core.UserChange{
				UserBaseData: user.UserBaseData,
			},
			Id: &userId,
		},
		PasswordHash: &user.PasswordHash,
	}

	return &ret, nil
}

func (U *UserDB) PathUser(ctx context.Context, user *core.UserDB) error {
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

	if user.User.UserChange.UserBaseData.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name = $%d", argId))
		args = append(args, *user.User.UserChange.UserBaseData.FirstName)
		argId++
	}
	if user.User.UserChange.UserBaseData.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name = $%d", argId))
		args = append(args, *user.User.UserChange.UserBaseData.LastName)
		argId++
	}
	if user.User.UserChange.UserBaseData.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email = $%d", argId))
		args = append(args, *user.User.UserChange.UserBaseData.Email)
		argId++
	}
	if user.User.UserChange.VacationTimeStart != nil {
		setValues = append(setValues, fmt.Sprintf("vacation_time_start = $%d", argId))
		args = append(args, *user.User.UserChange.VacationTimeStart)
		argId++
	}
	if user.User.UserChange.VacationTimeEnd != nil {
		setValues = append(setValues, fmt.Sprintf("vacation_time_end = $%d", argId))
		args = append(args, *user.User.UserChange.VacationTimeEnd)
		argId++
	}
	if user.User.EmailIsValidated != nil {
		setValues = append(setValues, fmt.Sprintf("email_is_validated = $%d", argId))
		args = append(args, *user.User.EmailIsValidated)
		argId++
	}
	if user.User.ImageURL != nil {
		setValues = append(setValues, fmt.Sprintf("image_url = $%d", argId))
		args = append(args, *user.User.ImageURL)
		argId++
	}
	if user.User.CompanyId != nil {
		setValues = append(setValues, fmt.Sprintf("company_id = $%d", argId))
		args = append(args, *user.User.CompanyId)
		argId++
	}
	if user.User.DepartmentId != nil {
		setValues = append(setValues, fmt.Sprintf("department_id = $%d", argId))
		args = append(args, *user.User.DepartmentId)
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

	args = append(args, *user.User.Id)

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
