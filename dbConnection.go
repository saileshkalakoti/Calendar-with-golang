package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDb() (db *sqlx.DB) {
	db = sqlx.MustConnect("mysql", "root:root@tcp(127.0.0.1:7701)/testDb")
	// if err != nil {
	// 	panic(err.Error())
	// }
	return db
}
