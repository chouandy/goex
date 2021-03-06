package dotenvex

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// DecryptCommand the command struct
type DecryptCommand struct {
	Password string
}

// Synopsis the synopsis of command
func (c *DecryptCommand) Synopsis() string {
	return "Decrypt dotenv files"
}

// Help the help of command
func (c *DecryptCommand) Help() string {
	helpText := `
Usage: cmd dotenv decrypt
	Decrypt dotenv files

Options:
  --password     The password for decrypt. It can be ENV["SECRETS_PASSWORD"]
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *DecryptCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("dotenv decrypt", flag.ContinueOnError)
	f.StringVar(&c.Password, "password", "", "password")
	if err := f.Parse(args); err != nil {
		fmt.Println(err)
		return 1
	}

	// Get options from env
	c.GetOptionsFromEnv()

	// Validate Options
	fmt.Print("Validate Options...")
	if err := c.ValidateOptions(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	// Decrypt Dotenv Files
	for _, stage := range Stages() {
		err := DecryptFile(stage, []byte(c.Password))
		if err != nil && strings.Contains(err.Error(), "no such file or directory") {
			continue
		}
		fmt.Print(`Decrypt "` + stage + `" Dotenv File...`)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Println("done")
	}

	return 0
}

// GetOptionsFromEnv get options from env
func (c *DecryptCommand) GetOptionsFromEnv() {
	if len(c.Password) == 0 {
		c.Password = os.Getenv("SECRETS_PASSWORD")
	}
}

// ValidateOptions validate options
func (c *DecryptCommand) ValidateOptions() error {
	if len(c.Password) == 0 {
		return errors.New("password can't be blank")
	}

	return nil
}
