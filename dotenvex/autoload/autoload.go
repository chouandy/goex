package autoload

import (
	"os"

	"github.com/chouandy/goex/dotenvex"
	"github.com/joho/godotenv"
)

func init() {
	// Load for go-gin
	mode := os.Getenv("GIN_MODE")
	if len(mode) == 0 {
		mode = "debug"
	}
	godotenv.Load(".env." + mode)

	// Load by stage
	dotenvex.LoadByStage()

	// Load .env
	godotenv.Load()
}
