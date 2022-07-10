package db

import (
	// "database/sql"
	// "fmt"
	// "gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "C:\\Users\\User\\Go\\go-class\\myapp1\\myapp.db"

// for Windows, use double slash.
// const dbName = "c:\\data\\myapp.db"

func DB() *gorm.DB {

	// dsn := "host=localhost user=terin password=secret dbname=terin_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	dsn := "sqlserver://sa:Terin2521+@localhost:1433?database=master"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate() {
	db := DB()
	db.AutoMigrate(&UserDB{})

}
