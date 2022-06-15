package database

import (
	"fmt"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/object_service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const dsnFormat = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC"

var connection *gorm.DB
var connectionError error

func init() {
	logger.Info("connecting to database")
	user, password := os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")
	databaseName, port := os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")

	dsn := fmt.Sprintf(dsnFormat, host, user, password, databaseName, port)
	connection, connectionError = gorm.Open(postgres.Open(dsn))
	if connectionError != nil {
		logger.Fatalf("cannot connect to database: %e", connectionError)
	}
	logger.Info("successful connected to database")

	migrator()
}

func migrator() {
	if err := connection.AutoMigrate(&model.Object{}); err != nil {
		logger.Fatalf("cannot apply migration: %e", err)
	}
}

func GetConnection() *gorm.DB {
	return connection
}

func GetConnectionError() error {
	return connectionError
}
