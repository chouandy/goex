package dbex

import (
	"fmt"
	"strings"
)

// DBDropCommand the command struct
type DBDropCommand struct{}

// Synopsis the synopsis of command
func (c *DBDropCommand) Synopsis() string {
	return "Drop database"
}

// Help the help of command
func (c *DBDropCommand) Help() string {
	helpText := `
Usage: cmd db create
  Drop database
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *DBDropCommand) Run(args []string) int {
	/* New DB Config */
	fmt.Print("New DB Config...")
	config, err := NewConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	fmt.Print("Drop Database...")
	if err := DropDatabase(config); err != nil {
		if strings.Contains(err.Error(), "database doesn't exist") {
			fmt.Println("database doesn't exist")
			return 0
		}
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
