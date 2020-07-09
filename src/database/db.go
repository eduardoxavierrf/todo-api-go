package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("failed to connect database")
		db.Close()
	}

	fmt.Println("Database connected")
	return db
}
