package config

import (
	"database/sql"
	"qa/exception"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(engine string, address string) *sql.DB {
	db, err := sql.Open(engine, address)
	exception.CheckError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
