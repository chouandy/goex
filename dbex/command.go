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
	"db migrate create": func() (cli.Command, error) {
		return &MigrateCreateCommand{}, nil
	},
	"db migrate up": func() (cli.Command, error) {
		return &MigrateUpCommand{}, nil
	},
	"db migrate down": func() (cli.Command, error) {
		return &MigrateDownCommand{}, nil
	},
	"db migrate drop": func() (cli.Command, error) {
		return &MigrateDropCommand{}, nil
	},
}
