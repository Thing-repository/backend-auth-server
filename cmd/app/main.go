package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var serverPort string
var tokenSecret string

func main() {
	logBase := logrus.Fields{
		"module":   "main",
		"file":     "main",
		"function": "main",
	}

	setupLogs()

	initEnv()
	initConfig()

	//httpHandler := restHandler.NewHandler()
	//httpServer := rest.NewHttpServer()

	//if err := httpServer.Run(serverPort, httpHandler.InitRoutes()); err != nil {
	//	logrus.WithFields(logrus.Fields{
	//		"base":  logBase,
	//		"error": "err",
	//	}).Error("httpServer run error")
	//}
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
		tokenSecret = os.Getenv("THINGS_REPOSITORY_TOKEN_SECRET")
	}

	if serverPort == "" {
		// for set serverPort in build add build parameter "-X main.serverPort=(token secret)"
		serverPort = os.Getenv("SERVER_HTTP_PORT")
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
