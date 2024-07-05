package database

import (
	"bookswapper/utils/env"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	// get database configuration
	postgresHost := env.GetEnv("POSTGRES_HOST", "localhost")
	postgresPort := env.GetEnv("POSTGRES_PORT", "5432")
	postgresUser := env.GetEnv("POSTGRES_USER", "postgres")
	postgresPassword := env.GetEnv("POSTGRES_PASSWORD", "postgres")
	postgresDatabase := env.GetEnv("POSTGRES_DB", "bookswapper")

	// create db url
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresPort)

	// return database connection
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
