package config

import (
	"Clinic-Reservation-System/helper"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Mysql goalng driver

	"github.com/rs/zerolog/log"
)

const (
	host     = "databasecont"
	port     = 3306
	user     = "root"
	password = "12345678"
	dbName   = "clinic_reservation_system"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database!!")

	return db
}
