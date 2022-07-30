package infra

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ProvideMySQL ...
func ProvideMySQL() *sql.DB {
	// Load config, currently hardcoded
	dbUser := "user"
	dbPass := "12345"
	host := "localhost"
	port := "3306"
	dbName := "user"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&parseTime=true", dbUser, dbPass, host, port, dbName))
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("Got nil db")
	}

	db.SetConnMaxLifetime(120 * time.Second)
	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(64)
	db.SetConnMaxIdleTime(60 * time.Second)
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
