package main

import (
	"database/sql"
	"time"
)

const (
	maxOpenDbConn = 25
	maxIdleDbConn = 25
	maxDbLifetime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	return db, nil
}
