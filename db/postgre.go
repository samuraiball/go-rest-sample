package db

import (
	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB

func init() {
	/*
		db, err := gorm.Open(
			"postgres",
			"host=127.0.0.1 port=5432 user=todoUser dbname=todo_db password=password")*/

	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")

	if err != nil {
		panic(err)
	}
	conn = db
}

func DB() *gorm.DB {
	return conn
}

func CloseDB() error {
	return conn.Close()
}
