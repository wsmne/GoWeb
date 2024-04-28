package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

type Data struct {
	// TODO wrapped database client
	Db *gorm.DB
}

// NewData .
func NewData(Log *slog.Logger, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		Log.Info("closing the data resources")
	}
	dsn := "root:wsm1201..@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Error("open db error ,err = %v" + err.Error())
		panic("db creat error")
	}
	return &Data{
		Db: db,
	}, cleanup, nil
}
