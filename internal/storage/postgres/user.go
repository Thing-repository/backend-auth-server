package postgres

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type User struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *User {
	return &User{db: db}
}

func (u User) GetUserByEmail(email string) (*core.UserDB, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "GetUserByEmail",
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
				is_company_admin, 
				is_department_admin, 
				vacation_time_start, 
				vacation_time_end 
			FROM 
				users 
			WHERE 
				email = $1`

	ctx := context.TODO()

	row := u.db.QueryRow(ctx, query, email)

	var userData core.UserDB

	err := row.Scan(&userData.Id, &userData.FirstName, &userData.LastName, &userData.Email,
		&userData.ImageURL, &userData.PasswordHash, &userData.CompanyId, &userData.DepartmentId,
		&userData.IsCompanyAdmin, &userData.IsDepartmentAdmin, &userData.VacationTimeStart, &userData.VacationTimeEnd)
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

func (u User) GetUser(userId int) (*core.UserDB, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) AddUser(user *core.AddUserDB) (*core.UserDB, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "user.go",
		"function": "AddUser",
	}

	query := `
				INSERT INTO users
				    (first_name, last_name, email, password_hash) 
				VALUES 
				    ($1, $2, $3, $4) 
				RETURNING 
					id`

	ctx := context.TODO()

	row := u.db.QueryRow(ctx, query, user.FirstName, user.LastName, user.Email, user.PasswordHash)

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
		PasswordHash: user.PasswordHash,
	}

	return &ret, nil
}
