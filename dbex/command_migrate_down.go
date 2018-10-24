package dbex

import (
	"flag"
	"fmt"
	"strings"
)

// MigrateDownCommand the command struct
type MigrateDownCommand struct {
	Number int
}

// Synopsis the synopsis of command
func (c *MigrateDownCommand) Synopsis() string {
	return "Apply all or N down migrations"
}

// Help the help of command
func (c *MigrateDownCommand) Help() string {
	helpText := `
Usage: cmd db migrate down
	Apply all or N down migrations

Options:
  -n    The number of migrations
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateDownCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("db migrate down", flag.ContinueOnError)
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

	/* Migrate Down */
	fmt.Print("Migrate Down...")
	if err := MigrateDown(config, c.Number); err != nil {
		if strings.Contains(err.Error(), "file does not exist") {
			fmt.Println("no migrations")
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
