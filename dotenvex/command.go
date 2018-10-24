package dotenvex

import "github.com/mitchellh/cli"

// Commands commands map
var Commands = map[string]cli.CommandFactory{
	"dotenv encrypt": func() (cli.Command, error) {
		return &EncryptCommand{}, nil
	},
}
