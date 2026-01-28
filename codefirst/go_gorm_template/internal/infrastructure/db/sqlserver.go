package db

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectSQLServer(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
}
