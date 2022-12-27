package database

import (
	"user-generator/internal/domain/model"
)

func MySqlMigrations() {
	db := NewMySqlConfig()

	db.AutoMigrate(model.User{})
	logger.Trace("table created")
}
