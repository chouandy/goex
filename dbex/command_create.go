package dbex

import (
	"fmt"
	"strings"
)

// CreateCommand the command struct
type CreateCommand struct{}

// Synopsis the synopsis of command
func (c *CreateCommand) Synopsis() string {
	return "Create database"
}

// Help the help of command
func (c *CreateCommand) Help() string {
	helpText := `
Usage: cmd db create
  Create database
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *CreateCommand) Run(args []string) int {
	/* New DB Config */
	fmt.Print("New DB Config...")
	config, err := NewConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	fmt.Print("Create Database...")
	if err := CreateDatabase(config); err != nil {
		if strings.Contains(err.Error(), "database exists") {
			fmt.Println("database already exists")
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
