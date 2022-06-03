package postgres

import (
	"context"
	"fmt"
	"github.com/Thing-repository/backend-server/pkg/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	maxAttemptsForConnect = 5
	attemptsInterval      = 5 * time.Second
	connectionTimeout     = 5 * time.Second
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
}

func NewPostgresDB(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	logBase := logrus.Fields{
		"module":   "postgres",
		"file":     "postgres.go",
		"function": "NewPostgresDB",
	}
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var pool *pgxpool.Pool

	err := utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, connectionTimeout)
		defer cancel()
		var err error
		pool, err = pgxpool.Connect(ctx, dns)
		return err
	}, maxAttemptsForConnect, attemptsInterval)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("error connect to db")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, connectionTimeout)
	defer cancel()
	err = pool.Ping(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("error ping db")
		return nil, err
	}

	return pool, nil

}
