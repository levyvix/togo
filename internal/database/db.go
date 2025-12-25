package database

import (
	"levyvix/togo/schema"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	SQL_FILENAME = "tasks.db"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(SQL_FILENAME), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic("failed to connect database")
	}
	if err := DB.AutoMigrate(&schema.Task{}); err != nil {
		panic("failed to migrate database")
	}
}
