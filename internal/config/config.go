package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"patients-golang-api/internal/model"

	"github.com/joho/godotenv"
)

var (
	configInstance *model.Config
	once           sync.Once
)

func GetConfig() *model.Config {
	once.Do(func() {
		configInstance = loadConfig()
	})

	return configInstance
}

func loadConfig() *model.Config {
	f, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	envPath := filepath.Join(f, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		log.Fatalf("MYSQL_PASSWORD is required but not set in environment variables")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		log.Fatalf("MYSQL_USER is required but not set in environment variables")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		log.Fatalf("MYSQL_USER is required but not set in environment variables")
	}

	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		mysqlPort = "3306"
	}

	mysqlDb := os.Getenv("MYSQL_DB")
	if mysqlDb == "" {
		log.Fatalf("MYSQL_DB is required but not set in environment variables")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	fromAddress := os.Getenv("FROM_ADDRESS")
	if fromAddress == "" {
		log.Fatalf("FROM_ADDRESS is required but not set in environment variables")
	}

	smtpServer := os.Getenv("SMTP_SERVER")
	if smtpServer == "" {
		log.Fatalf("SMTP_SERVER is required but not set in environment variables")
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		log.Fatalf("SMTP_PORT is required but not set in environment variables")
	}

	smtpUser := os.Getenv("SMTP_USER")
	if smtpUser == "" {
		log.Fatalf("SMTP_USER is required but not set in environment variables")
	}

	smtpPass := os.Getenv("SMTP_PASS")
	if smtpPass == "" {
		log.Fatalf("SMTP_PASS is required but not set in environment variables")
	}

	return &model.Config{
		MysqlPassword: mysqlPassword,
		MysqlUser:     mysqlUser,
		MysqlHost:     mysqlHost,
		MysqlPort:     mysqlPort,
		MysqlDb:       mysqlDb,
		ServerPort:    serverPort,
		FromAddress:   fromAddress,
		SmtpServer:    smtpServer,
		SmtpPort:      smtpPort,
		SmtpUser:      smtpUser,
		SmtpPass:      smtpPass,
	}
}
