package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"patients-golang-api/internal/model"

	"github.com/joho/godotenv"
)

func LoadConfig() (*model.Config, error) {
	f, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting wrking directory: %v", err)
	}

	envPath := filepath.Join(f, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		return nil, errors.New("MYSQL_PASSWORD is required but not set in environment variables")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		return nil, errors.New("MYSQL_USER is required but not set in environment variables")
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		return nil, errors.New("MYSQL_HOST is required but not set in environment variables")
	}

	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort == "" {
		mysqlPort = "8080"
	}

	mysqlDb := os.Getenv("MYSQL_DB")
	if mysqlDb == "" {
		return nil, errors.New("MYSQL_DB is required but not set in environment variables")
	}

	return &model.Config{
		MysqlPassword: mysqlPassword,
		MysqlUser:     mysqlUser,
		MysqlHost:     mysqlHost,
		MysqlPort:     mysqlPort,
		MysqlDb:       mysqlDb,
	}, nil
}
