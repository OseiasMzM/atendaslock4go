package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importação do driver PostgreSQL
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "1234",
		DBName:   "postgres",
	}
}

func ConnectDB(config *DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Connect() *sql.DB {
	db, _ := ConnectDB(NewDatabaseConfig())
	return db
}
