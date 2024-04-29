package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

var Db *gorm.DB

// NewData .
func InitDataBase(Log *slog.Logger) error {
	dsn := "root:wsm1201..@tcp(127.0.0.1:3306)/kratos_try?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Error("open db error ,err = %v" + err.Error())
		panic("db creat error")
	}
	Db = db
	sqlDB, _ := Db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(50)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
