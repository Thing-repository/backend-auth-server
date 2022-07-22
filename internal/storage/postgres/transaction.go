package postgres

import (
	"context"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const transaction = "transaction"

type db interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type Transaction struct {
	db db
}

func NewTransaction(db db) *Transaction {
	return &Transaction{db: db}
}

func (T *Transaction) InjectTx(ctx context.Context) (context.Context, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "InjectTx",
	}
	tx, err := T.db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error starting transaction")
		return nil, err
	}
	return context.WithValue(ctx, transaction, tx), nil
}

func (T *Transaction) ExtractTx(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(transaction).(pgx.Tx)
	return tx, ok
}

func (T *Transaction) CommitTx(ctx context.Context) error {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "CommitTx",
	}
	tx, ok := T.ExtractTx(ctx)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"base": logBase,
			"ok":   ok,
		}).Error(moduleErrors.ErrorDataBaseGetTransaction.Error())
		return moduleErrors.ErrorDataBaseGetTransaction
	}

	if err := tx.Commit(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error commit transaction")
		return err
	}

	return nil
}

func (T *Transaction) RollbackTx(ctx context.Context) error {
	logBase := logrus.Fields{
		"module":   "postgres",
		"function": "RollbackTx",
	}
	tx, ok := T.ExtractTx(ctx)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"base": logBase,
			"ok":   ok,
		}).Error(moduleErrors.ErrorDataBaseGetTransaction.Error())
		return moduleErrors.ErrorDataBaseGetTransaction
	}

	if err := tx.Rollback(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error rollback transaction")
		return err
	}

	return nil
}
