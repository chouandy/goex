package dbex

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// DataSourceFormat data source format
var DataSourceFormat = "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true"

// DataSourceWithoutDatabaseFormat data source without database format
var DataSourceWithoutDatabaseFormat = "%s:%s@tcp(%s:%s)/?charset=%s"

// DatabaseURLFormat database url format
var DatabaseURLFormat = "%s://%s:%s@tcp(%s:%s)/%s?charset=%s"

// Config config struct
type Config struct {
	Driver         string
	Username       string
	Password       string
	Host           string
	Port           string
	Database       string
	Charset        string
	DefaultCollate string
	MaxIdleConns   int
	MaxOpenConns   int
}

// NewConfig new config
func NewConfig() (*Config, error) {
	// New config
	config := &Config{
		Driver:         os.Getenv("DB_DRIVER"),
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		Database:       os.Getenv("DB_DATABASE"),
		Username:       os.Getenv("DB_USERNAME"),
		Password:       os.Getenv("DB_PASSWORD"),
		Charset:        os.Getenv("DB_CHARSET"),
		DefaultCollate: os.Getenv("DB_DEFAULT_COLLATE"),
	}
	// Validate driver
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// Get max idle conns from env
	config.GetMaxIdleConnsFromEnv()
	// Get max open conns from env
	config.GetMaxOpenConnsFromEnv()
	// Load default
	config.LoadDefault()

	return config, nil
}

// Validate validate
func (c *Config) Validate() error {
	if len(c.Driver) == 0 {
		return errors.New("driver can't be blank")
	}
	if len(c.Host) == 0 {
		return errors.New("host can't be blank")
	}
	if len(c.Database) == 0 {
		return errors.New("database can't be blank")
	}

	return nil
}

// GetMaxIdleConnsFromEnv get max idle conns from env
func (c *Config) GetMaxIdleConnsFromEnv() {
	if maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS")); err == nil {
		c.MaxIdleConns = maxIdleConns
	}
}

// GetMaxOpenConnsFromEnv get max open conns from env
func (c *Config) GetMaxOpenConnsFromEnv() {
	if maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS")); err == nil {
		c.MaxOpenConns = maxOpenConns
	}
}

// LoadDefault load default
func (c *Config) LoadDefault() {
	// Set port default value
	if len(c.Port) == 0 {
		c.Port = "3306"
	}
	// Set charset default value
	if len(c.Charset) == 0 {
		c.Charset = "utf8"
	}
	// Set default collate default value
	if len(c.DefaultCollate) == 0 {
		c.DefaultCollate = "utf8_general_ci"
	}
	// Set max idle conns default value
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}
	// Set max open conns default value
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 100
	}
}

// DataSource data source
func (c *Config) DataSource() string {
	return fmt.Sprintf(DataSourceFormat,
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset,
	)
}

// DataSourceWithoutDatabase data source without database
func (c *Config) DataSourceWithoutDatabase() string {
	return fmt.Sprintf(DataSourceWithoutDatabaseFormat,
		c.Username, c.Password, c.Host, c.Port, c.Charset,
	)
}

// DatabaseURL return database url
func (c *Config) DatabaseURL() string {
	return fmt.Sprintf(DatabaseURLFormat,
		c.Driver, c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset,
	)
}
