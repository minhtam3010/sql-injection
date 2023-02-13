package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// DBName is the name of the database
	DBName   = "sqlinjection"
	USERNAME = "root"
	PASSWORD = "quynhnhu2010"
)

type Querier struct {
	DB *sql.Tx
}

func Connect() (db *sql.DB) {
	var (
		err error
	)

	db, err = sql.Open("mysql", USERNAME+":"+PASSWORD+"@/"+DBName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func BeginTx(db *sql.DB) (tx *sql.Tx) {
	var err error
	tx, err = db.Begin()
	if err != nil {
		panic(err.Error())
	}

	return tx
}

func NewQuerier(db *sql.DB) *Querier {
	return &Querier{DB: BeginTx(db)}
}
