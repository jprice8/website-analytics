package shared

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to Database struct
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./../test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DB = db 
	return DB
}

// Use this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}