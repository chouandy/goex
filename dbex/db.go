package dbex

import (
	"time"

	"github.com/jinzhu/gorm"
)

// DB gorm db
var DB *gorm.DB

// InitDB init db
func InitDB(config *Config, enableLogger bool) error {
	// New gorm db
	db, err := gorm.Open(config.Driver, config.DataSource())
	if err != nil {
		return err
	}

	// Setup db
	db.LogMode(enableLogger)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Hour)

	// Set db
	DB = db

	return nil
}
