package dbex

import "github.com/mitchellh/cli"

// Commands commands map
var Commands = map[string]cli.CommandFactory{
	"db create": func() (cli.Command, error) {
		return &DBCreateCommand{}, nil
	},
	"db drop": func() (cli.Command, error) {
		return &DBDropCommand{}, nil
	},
	"db migrate new": func() (cli.Command, error) {
		return &DBMigrateNewCommand{}, nil
	},
	"db migrate up": func() (cli.Command, error) {
		return &DBMigrateUpCommand{}, nil
	},
	"db migrate down": func() (cli.Command, error) {
		return &DBMigrateDownCommand{}, nil
	},
}
