package autoload

import (
	"github.com/chouandy/goex/dotenvex"
	"github.com/joho/godotenv"
)

func init() {
	stage := dotenvex.Stage()
	godotenv.Load(".env", ".env."+stage)
}
