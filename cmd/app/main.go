package main

import (
	"github.com/Thing-repository/backend-auth-server/internal/transport/rest"
	rest_handler "github.com/Thing-repository/backend-auth-server/internal/transport/rest/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006.01.02 15:04:05",
		FullTimestamp:   true,
		DisableSorting:  true,
	})
	logBase := logrus.Fields{
		"module":   "main",
		"file":     "main",
		"function": "main",
	}

	initEnv()
	initConfig()

	httpHandler := rest_handler.NewHandler()

	httpServer := rest.NewHttpServer()
	err := httpServer.Run(os.Getenv("AUTH_SERVER_HTTP_PORT"), httpHandler.InitRoutes())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": "err",
		}).Error("httpServer run error")
	}
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
}
