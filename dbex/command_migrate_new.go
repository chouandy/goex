package dbex

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// MigrateNewCommand the command struct
type MigrateNewCommand struct {
	Name string
}

// Synopsis the synopsis of command
func (c *MigrateNewCommand) Synopsis() string {
	return "New a set of timestamped up/down migrations"
}

// Help the help of command
func (c *MigrateNewCommand) Help() string {
	helpText := `
Usage: cmd db migrate new
	New a set of timestamped up/down migrations

Options:
  --name    The migrations' name.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateNewCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("db migrate new", flag.ContinueOnError)
	f.StringVar(&c.Name, "name", "", "name")
	if err := f.Parse(args); err != nil {
		fmt.Println(err)
		return 1
	}

	/* Validate Options */
	fmt.Print("Validate Options...")
	if err := c.ValidateOptions(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	/* Migrate New */
	fmt.Print("Migrate New...")
	if err := MigrateNew(c.Name); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}

// ValidateOptions validate options
func (c *MigrateNewCommand) ValidateOptions() error {
	if len(c.Name) == 0 {
		return errors.New("name can't be blank")
	}

	return nil
}
