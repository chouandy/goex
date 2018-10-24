package dbex

import (
	"os"
	"time"

	"github.com/golang-migrate/migrate"
)

var migrateDir = "db/migrate"

var migratePath = migrateDir + "/"

var migrateSource = "file://" + migrateDir

// SetMigrateDir set migrate dir
func SetMigrateDir(dir string) {
	migrateDir = dir
	migratePath = migrateDir + "/"
	migrateSource = "file://" + migrateDir
}

// MigrateCreate migrate create
func MigrateCreate(name string) error {
	// Make migrate dir
	if err := os.MkdirAll(migratePath, os.ModePerm); err != nil {
		return err
	}
	// New base
	base := migratePath + time.Now().Format("20060102150405") + "_" + name
	// Generate up and down migrations
	if _, err := os.Create(base + ".up.sql"); err != nil {
		return err
	}
	if _, err := os.Create(base + ".down.sql"); err != nil {
		return err
	}

	return nil
}

// MigrateUp migrate up
func MigrateUp(config *Config, n int) error {
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Check current version is dirty or not
	if v, dirty, err := migrater.Version(); err == nil && dirty {
		// Update the version's dirty to false
		if err := migrater.Force(int(v)); err != nil {
			return err
		}
		// Migrate down version
		if err := migrater.Steps(-1); err != nil {
			return err
		}
	}

	// Migrate up
	if n > 0 {
		if err := migrater.Steps(n); err != nil {
			return err
		}
	} else {
		if err := migrater.Up(); err != nil {
			return err
		}
	}

	return nil
}

// MigrateDown migrate down
func MigrateDown(config *Config, n int) error {
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Check current version is dirty or not
	if v, dirty, err := migrater.Version(); err == nil && dirty {
		// Update the version's dirty to false
		if err := migrater.Force(int(v)); err != nil {
			return err
		}
		// Migrate down version
		if err := migrater.Steps(-1); err != nil {
			return err
		}
	}

	// Migrate down
	if n > 0 {
		if err := migrater.Steps(-n); err != nil {
			return err
		}
	} else {
		if err := migrater.Steps(-1); err != nil {
			return err
		}
	}

	return nil
}

// MigrateDrop migrate drop
func MigrateDrop(config *Config) error {
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Migrate drop
	return migrater.Drop()
}

// MigrateForce migrate force
func MigrateForce(config *Config, version int) error {
	// New migrater
	migrater, err := migrate.New(migrateSource, config.DatabaseURL())
	if err != nil {
		return err
	}
	defer migrater.Close()

	// Migrate force
	if err := migrater.Force(version); err != nil {
		return err
	}

	return nil
}
