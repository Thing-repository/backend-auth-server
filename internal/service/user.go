package service

import "C"
import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/sirupsen/logrus"
)

type userDBUser interface {
	PathUser(ctx context.Context, user *core.UserDB, userId int) error
	GetUser(ctx context.Context, userId int) (*core.UserDB, error)
	GetUsersFilter(ctx context.Context, filter string, limit int, offset int) ([]core.User, error)
}

type transactionDBUser interface {
	InjectTx(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error
	RollbackTxDefer(ctx context.Context)
}

type User struct {
	userDB        userDBUser
	transactionDB transactionDBUser
}

func NewUser(userDB userDBUser, transactionDB transactionDBUser) *User {
	return &User{
		userDB:        userDB,
		transactionDB: transactionDB,
	}
}

func (U *User) FindUsersForInvite(ctx context.Context, filter string, limit int, offset int) ([]core.User, error) {
	logBase := logrus.Fields{
		"module":   "user",
		"function": "FindUsersForInvite",
		"filter":   filter,
		"limit":    limit,
		"offset":   offset,
		"context":  *core.LogContext(ctx),
	}

	users, err := U.userDB.GetUsersFilter(ctx, filter, limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get users from database")
		return nil, err
	}

	return users, nil
}
