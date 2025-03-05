package db

import (
	"database/sql"
	"fam/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDb(postgressConfig *config.PostgresConfig) *gorm.DB {

	sqlDB, err := connectToSqlServer(postgressConfig)
	if err != nil {
		panic("failed to connect to database")
	}
	exists, err := checkDatabaseExists(postgressConfig, sqlDB)
	if err != nil {
		panic("failed to checkDataBase")
	}

	if !exists {
		createDatabase(postgressConfig, sqlDB)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgressConfig.Host,
		postgressConfig.UserName,
		postgressConfig.Password,
		postgressConfig.DatabaseName,
		postgressConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database!, err: " + err.Error())
	}

	return db
}

func connectToSqlServer(postgresConfig *config.PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		postgresConfig.Host, postgresConfig.UserName, postgresConfig.Password, postgresConfig.Port)
	sqlDB, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	return sqlDB, nil
}

func checkDatabaseExists(postgresConfig *config.PostgresConfig, sqlDB *sql.DB) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", postgresConfig.DatabaseName)
	err := sqlDB.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func createDatabase(postgresConfig *config.PostgresConfig, sqlDB *sql.DB) {
	_, err := sqlDB.Exec(fmt.Sprintf("CREATE DATABASE %s", postgresConfig.DatabaseName))
	if err != nil {
		log.Panicf("Failed to create database: %v", err)
	}
	fmt.Printf("Database %s created successfully.\n", postgresConfig.DatabaseName)
}
