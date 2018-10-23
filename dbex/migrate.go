package dbex

import (
	"os"
	"time"
)

// MigrateNew migrate new
func MigrateNew(name string) error {
	// Make migrate dir
	if err := os.MkdirAll("db/migrate/", os.ModePerm); err != nil {
		return err
	}
	// New base
	base := "db/migrate/" + time.Now().Format("20060102150405") + "_" + name
	// Generate up and down migrations
	if _, err := os.Create(base + ".up.sql"); err != nil {
		return err
	}
	if _, err := os.Create(base + ".down.sql"); err != nil {
		return err
	}

	return nil
}
