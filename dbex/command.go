package dbex

import "github.com/mitchellh/cli"

// Commands commands map
var Commands = map[string]cli.CommandFactory{
	"db create": func() (cli.Command, error) {
		return &CreateCommand{}, nil
	},
	"db drop": func() (cli.Command, error) {
		return &DropCommand{}, nil
	},
	"db migrate new": func() (cli.Command, error) {
		return &MigrateNewCommand{}, nil
	},
	"db migrate up": func() (cli.Command, error) {
		return &MigrateUpCommand{}, nil
	},
	"db migrate down": func() (cli.Command, error) {
		return &MigrateDownCommand{}, nil
	},
}
