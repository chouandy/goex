package dbex

import (
	"flag"
	"fmt"
	"strings"

	"github.com/chouandy/migrate"
)

// MigrateUpCommand the command struct
type MigrateUpCommand struct {
	Number int
}

// Synopsis the synopsis of command
func (c *MigrateUpCommand) Synopsis() string {
	return "Apply all or N up migrations"
}

// Help the help of command
func (c *MigrateUpCommand) Help() string {
	helpText := `
Usage: cmd db migrate up
	Apply all or N up migrations

Options:
  -n    The number of migrations
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateUpCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("db migrate up", flag.ContinueOnError)
	f.IntVar(&c.Number, "n", 0, "n")
	if err := f.Parse(args); err != nil {
		fmt.Println(err)
		return 1
	}

	/* New DB Config */
	fmt.Print("New DB Config...")
	config, err := NewConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	/* Migrate Up */
	fmt.Print("Migrate Up...")
	if err := MigrateUp(config, c.Number); err != nil {
		if strings.Contains(err.Error(), "file does not exist") {
			fmt.Println("no migrations")
			return 0
		}
		if err == migrate.ErrNoChange {
			fmt.Println(err)
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
