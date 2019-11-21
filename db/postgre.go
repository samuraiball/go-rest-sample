package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func init() {
	db, err := gorm.Open(
		"postgres",
		fmt.Printf("host=%s port=%s dbname=%s user%s password=%s sslmode=disable",
			"localhost",
			"5432",
			"todoDb",
			"password",
			"dotoUser"))

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
