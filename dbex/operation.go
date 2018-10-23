package dbex

import (
	"database/sql"
	"fmt"
)

var (
	// CreateDatabaseStatement 新增資料庫語法
	CreateDatabaseStatement = "CREATE DATABASE `%s` DEFAULT CHARACTER SET = '%s' DEFAULT COLLATE '%s';"
	// DropDatabaseStatement 刪除資料庫語法
	DropDatabaseStatement = "DROP DATABASE `%s`;"
)

// CreateDatabase create database
func CreateDatabase(config *Config) error {
	// New db connection
	conn, err := sql.Open(config.Driver, config.DataSourceWithoutDatabase())
	if err != nil {
		return err
	}
	// defer close db connection
	defer conn.Close()
	// New statement
	statement := fmt.Sprintf(CreateDatabaseStatement,
		config.Database, config.Charset, config.DefaultCollate,
	)
	// Exec statement
	if _, err = conn.Exec(statement); err != nil {
		return err
	}

	return nil
}

// DropDatabase drop database
func DropDatabase(config *Config) error {
	conn, err := sql.Open(config.Driver, config.DataSourceWithoutDatabase())
	if err != nil {
		return err
	}
	// defer close db connection
	defer conn.Close()
	// New statement
	statement := fmt.Sprintf(DropDatabaseStatement, config.Database)
	// Exec statement
	if _, err = conn.Exec(statement); err != nil {
		return err
	}

	return nil
}
