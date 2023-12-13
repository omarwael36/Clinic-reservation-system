package config

import (
	"Clinic-Reservation-System/helper"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Mysql goalng driver

	"github.com/rs/zerolog/log"
)

var (
	host     = getEnv("DB_HOST", "databasecont")
	port     = getEnv("DB_PORT", "3306")
	user     = getEnv("DB_USER", "root")
	password = getEnv("DB_PASSWORD", "12345678")
	dbName   = getEnv("DB_NAME", "clinic_reservation_system")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to the database!!")

	return db
}
