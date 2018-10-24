package dbex

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// MigrateCreateCommand the command struct
type MigrateCreateCommand struct {
	Name string
}

// Synopsis the synopsis of command
func (c *MigrateCreateCommand) Synopsis() string {
	return "Create a set of timestamped up/down migrations"
}

// Help the help of command
func (c *MigrateCreateCommand) Help() string {
	helpText := `
Usage: cmd db migrate create
	Create a set of timestamped up/down migrations

Options:
  --name    The migrations' name.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateCreateCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("db migrate create", flag.ContinueOnError)
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

	/* Migrate Create */
	fmt.Print("Migrate Create...")
	if err := MigrateCreate(c.Name); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}

// ValidateOptions validate options
func (c *MigrateCreateCommand) ValidateOptions() error {
	if len(c.Name) == 0 {
		return errors.New("name can't be blank")
	}

	return nil
}
