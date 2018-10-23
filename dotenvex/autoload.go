package dotenvex

import (
	"github.com/chouandy/goex/deployex"
	"github.com/joho/godotenv"
)

func init() {
	stage := deployex.Stage()
	godotenv.Load(".env", ".env."+stage)
}
