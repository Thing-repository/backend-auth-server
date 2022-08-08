package main

import (
	"context"
	"github.com/Thing-repository/backend-server/internal/service"
	"github.com/Thing-repository/backend-server/internal/storage/postgres"
	"github.com/Thing-repository/backend-server/internal/transport/rest"
	restHandler "github.com/Thing-repository/backend-server/internal/transport/rest/handler"
	"github.com/Thing-repository/backend-server/pkg/generateToken"
	"github.com/Thing-repository/backend-server/pkg/userHash"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// @title Thing Repository API
// @version 0.0.1
// @description API server for thing repository service

// @contact.name Emil Islamov
// @contact.email emil.islamov110778@gmail.com

// @host thing-repository.emil110.keenetic.pro
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// base data
var httpPort string

// token data
var tokenSecret string

// hash data
var salt string

// db data
var postgresCfg postgres.Config
var postgresPassword string

func main() {
	logBase := logrus.Fields{
		"module":   "main",
		"file":     "main",
		"function": "main",
	}

	setupLogs()

	initEnv()
	initConfig()

	ctx := context.Background()

	postgresDb, err := postgres.NewPostgresDB(ctx, postgresCfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Fatal("error connect to postgres database")
	}

	transaction := postgres.NewTransaction(postgresDb)

	// DB modules
	userDb := postgres.NewUser(postgresDb, transaction)
	companyDb := postgres.NewCompanyDB(postgresDb, transaction)
	departmentDB := postgres.NewDepartmentDB(postgresDb, transaction)
	credentialsDB := postgres.NewCredentialsDB(postgresDb, transaction)

	// helper modules
	tokenGenerator := generateToken.NewToken([]byte(tokenSecret))
	hashGenerator := userHash.NewHash(salt)

	// service modules
	authService := service.NewAuth(tokenGenerator, userDb, hashGenerator, credentialsDB, transaction)
	companyService := service.NewCompany(userDb, companyDb, departmentDB, credentialsDB, transaction)
	userService := service.NewUser(userDb, transaction)

	httpHandler := restHandler.NewHandler(authService, companyService, tokenGenerator, userService, userDb)
	httpServer := rest.NewHttpServer()

	if err := httpServer.Run(httpPort, httpHandler.InitRoutes()); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": "err",
		}).Error("httpServer run error")
	}
}

func setupLogs() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006.01.02 15:04:05",
		FullTimestamp:   true,
		DisableSorting:  true,
	})
}

func initEnv() {
	logBase := logrus.Fields{
		"module":   "main",
		"file":     "main",
		"function": "initEnv",
	}
	hasEnv := os.Getenv("AUTH_SERVER_ENV")
	if hasEnv != "OK" {
		logrus.WithFields(logrus.Fields{
			"base":            logBase,
			"AUTH_SERVER_ENV": hasEnv,
		}).Warning("Env not found")
		if err := godotenv.Load(); err != nil {
			logrus.WithFields(logrus.Fields{
				"base":  logBase,
				"error": err,
			}).Fatal("error loading environment")
		}
	}

	if tokenSecret == "" {
		// for set tokenSecret in build add build parameter "-X main.tokenSecret=(token secret)"
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"build_cmd": "-X main.tokenSecret=(token secret)",
		}).Warning("no tokenSecret, use value from env, for add use cmd")
		tokenSecret = os.Getenv("THINGS_REPOSITORY_TOKEN_SECRET")
		if tokenSecret == "" {
			logrus.WithFields(logrus.Fields{
				"base":         logBase,
				"build_cmd":    "-X main.tokenSecret=(token secret)",
				"env variable": "THINGS_REPOSITORY_TOKEN_SECRET",
			}).Fatal("no token secret, for add use cmd or env")
		}
	}

	postgresCfg.Password = postgresPassword

	if postgresCfg.Password == "" {
		// for set tokenSecret in build add build parameter "-X main.postgresPassword=(postgres password)"
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"build_cmd": "-X main.postgresPassword=(postgres password)",
		}).Warning("no postgresPassword, use value from env, for add use cmd")
		postgresCfg.Password = os.Getenv("THINGS_REPOSITORY_SQL_DP_PASSWORD")
		if postgresCfg.Password == "" {
			logrus.WithFields(logrus.Fields{
				"base":         logBase,
				"build_cmd":    "-X main.postgresPassword=(postgres password)",
				"env variable": "THINGS_REPOSITORY_SQL_DP_PASSWORD",
			}).Fatal("no postgres password, for add use cmd or env")
		}
	}

	if salt == "" {
		// for set tokenSecret in build add build parameter "-X main.salt=(salt)"
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"build_cmd": "-X main.salt=(salt)",
		}).Warning("no salt, use value from env, for add use cmd")
		salt = os.Getenv("THINGS_REPOSITORY_HASH_SALT")
		if salt == "" {
			logrus.WithFields(logrus.Fields{
				"base":         logBase,
				"build_cmd":    "-X main.postgresPassword=(postgres password)",
				"env variable": "THINGS_REPOSITORY_HASH_SALT",
			}).Fatal("no salt, for add use cmd or env")
		}
	}

	httpPort = os.Getenv("THINGS_REPOSITORY_HTTP_PORT")
	if httpPort == "" {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"env variable": "THINGS_REPOSITORY_HTTP_PORT",
		}).Warning("no http port in env")
	}

	postgresCfg.Host = os.Getenv("THINGS_REPOSITORY_POSTGRES_HOST")
	if postgresCfg.Host == "" {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"env variable": "THINGS_REPOSITORY_POSTGRES_HOST",
		}).Warning("no postgres host in env")
	}

	postgresCfg.Port = os.Getenv("THINGS_REPOSITORY_POSTGRES_PORT")
	if postgresCfg.Port == "" {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"env variable": "THINGS_REPOSITORY_POSTGRES_PORT",
		}).Warning("no postgres port in env")
	}

	postgresCfg.DBName = os.Getenv("THINGS_REPOSITORY_POSTGRES_DB_NAME")
	if postgresCfg.DBName == "" {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"env variable": "THINGS_REPOSITORY_POSTGRES_DB_NAME",
		}).Warning("no postgres db name in env")
	}

	postgresCfg.Username = os.Getenv("THINGS_REPOSITORY_POSTGRES_USER")
	if postgresCfg.Username == "" {
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"env variable": "THINGS_REPOSITORY_POSTGRES_USER",
		}).Warning("no postgres username in env")
	}

}

func initConfig() {
	logBase := logrus.Fields{
		"module":   "main",
		"file":     "main",
		"function": "initConfig",
	}
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Fatal("error loading config")
	}

	if httpPort == "" {
		httpPort = viper.GetString("http_port")
		if httpPort == "" {
			logrus.WithFields(logrus.Fields{
				"base":        logBase,
				"config name": "http_port",
			}).Fatal("no http port in config, add http port to config file")
		}
		logrus.WithFields(logrus.Fields{
			"base":     logBase,
			"httpPort": httpPort,
		}).Warning("use http port from config")
	}

	if postgresCfg.Host == "" {
		postgresCfg.Host = viper.GetString("postgres.host")
		if postgresCfg.Host == "" {
			logrus.WithFields(logrus.Fields{
				"base":        logBase,
				"config name": "postgres.host",
			}).Fatal("no postgres host in config, add host to config file")
		}
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"postgresHost": postgresCfg.Host,
		}).Warning("use postgres host from config")
	}

	if postgresCfg.Port == "" {
		postgresCfg.Port = viper.GetString("postgres.port")
		if postgresCfg.Port == "" {
			logrus.WithFields(logrus.Fields{
				"base":        logBase,
				"config name": "postgres.port",
			}).Fatal("no postgres port in config, add port to config file")
		}
		logrus.WithFields(logrus.Fields{
			"base":         logBase,
			"postgresPort": postgresCfg.Port,
		}).Warning("use postgres port from config")
	}

	if postgresCfg.DBName == "" {
		postgresCfg.DBName = viper.GetString("postgres.name")
		if postgresCfg.DBName == "" {
			logrus.WithFields(logrus.Fields{
				"base":        logBase,
				"config name": "postgres.name",
			}).Fatal("no postgres database name in config, add name to config file")
		}
		logrus.WithFields(logrus.Fields{
			"base":           logBase,
			"postgresDBName": postgresCfg.DBName,
		}).Warning("use postgres userDB name from config")
	}

	if postgresCfg.Username == "" {
		postgresCfg.Username = viper.GetString("postgres.user")
		if postgresCfg.Username == "" {
			logrus.WithFields(logrus.Fields{
				"base":        logBase,
				"config name": "postgres.user",
			}).Fatal("no postgres username in config, add username to config file")
		}
		logrus.WithFields(logrus.Fields{
			"base":             logBase,
			"postgresUsername": postgresCfg.Username,
		}).Warning("use postgres username from config")
	}
}
