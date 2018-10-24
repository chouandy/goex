package dbex

import (
	"fmt"
	"strings"
)

// MigrateDropCommand the command struct
type MigrateDropCommand struct{}

// Synopsis the synopsis of command
func (c *MigrateDropCommand) Synopsis() string {
	return "Drop everyting inside database"
}

// Help the help of command
func (c *MigrateDropCommand) Help() string {
	helpText := `
Usage: cmd db migrate drop
	Drop everyting inside database
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *MigrateDropCommand) Run(args []string) int {
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
	if err := MigrateDrop(config); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	return 0
}
