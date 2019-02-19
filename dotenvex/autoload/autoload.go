package autoload

import (
	"os"

	"github.com/chouandy/goex/dotenvex"
	"github.com/joho/godotenv"
)

func init() {
	// Load for go-gin
	ginMode := os.Getenv("GIN_MODE")
	if len(ginMode) > 0 {
		godotenv.Load(".env." + ginMode)
	}

	// Load by stage
	dotenvex.LoadByStage()

	// Load .env
	godotenv.Load()
}
