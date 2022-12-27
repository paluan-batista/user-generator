package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"user-generator/infra/config"
	"user-generator/infra/log"
)

var logger = log.NewLogger()

var (
	db           *gorm.DB
	runOnceCache sync.Once
)

type Database struct {
	mysqlDatabase *gorm.DB
}

func NewMySqlConfig() *gorm.DB {
	runOnceCache.Do(func() {
		db = getMySqlConfig()
	})

	return db
}

func getMySqlConfig() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.GetDatabaseConfig().MySqlConnection), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}

	return db
}
