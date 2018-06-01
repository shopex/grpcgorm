package model

import (
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB(dbinfo string) (err error) {
	db, err = gorm.Open("mysql", dbinfo)
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "CHARSET=utf8")
	// db.LogMode(true)
	return
}
func Db() *gorm.DB {
	return db
}
