package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

var Db *gorm.DB

// NewData .
func InitDataBase(Log *slog.Logger) error {
	dsn := "root:wsm1201..@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Error("open db error ,err = %v" + err.Error())
		panic("db creat error")
	}
	Db = db
	return nil
}
