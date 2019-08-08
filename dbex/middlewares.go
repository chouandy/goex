package dbex

import (
	"fmt"

	"github.com/chouandy/goex/awsex/service/cloudwatcheventsex"
)

// InitDBEventMiddleware init db middleware
func InitDBEventMiddleware(ctx *cloudwatcheventsex.Context) error {
	if DB == nil {
		// New DB Config
		config, err := NewConfig()
		if err != nil {
			fmt.Printf("[InitDBEventMiddleware][NewDBConfig] %s\n", err)
			return err
		}

		// Init DB
		if err := InitDB(config, false); err != nil {
			fmt.Printf("[InitDBEventMiddleware][[InitDB] %s\n", err)
			return err
		}
	}

	return nil
}
