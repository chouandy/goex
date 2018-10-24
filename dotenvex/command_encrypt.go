package dotenvex

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// EncryptCommand the command struct
type EncryptCommand struct {
	Password string
	Stage    string
}

// Synopsis the synopsis of command
func (c *EncryptCommand) Synopsis() string {
	return "Encrypt dotenv files"
}

// Help the help of command
func (c *EncryptCommand) Help() string {
	helpText := `
Usage: cmd dotenv encrypt
	Encrypt dotenv files

Options:
  --password     The password for encrypt. It can be ENV["SECRETS_PASSWORD"]
  --stage        Encrypt the stage file.
`
	return strings.TrimSpace(helpText)
}

// Run the main execution of command
func (c *EncryptCommand) Run(args []string) int {
	// Init flag
	f := flag.NewFlagSet("dotenv encrypt", flag.ContinueOnError)
	f.StringVar(&c.Password, "password", "", "password")
	f.StringVar(&c.Stage, "stage", "", "stage")
	if err := f.Parse(args); err != nil {
		fmt.Println(err)
		return 1
	}

	// Get options from env
	c.GetOptionsFromEnv()

	/* Validate Options */
	fmt.Print("Validate Options...")
	if err := c.ValidateOptions(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("done")

	/* Encrypt Dotenv Files */
	for _, stage := range Stages() {
		err := EncryptFile(stage, []byte(c.Password))
		if err != nil && strings.Contains(err.Error(), "no such file or directory") {
			continue
		}
		fmt.Print(`Encrypt "` + stage + `" Dotenv File...`)
		if err != nil {
			fmt.Println(err)
			return 1
		}
		fmt.Println("done")
	}

	return 0
}

// GetOptionsFromEnv get options from env
func (c *EncryptCommand) GetOptionsFromEnv() {
	if len(c.Password) == 0 {
		c.Password = os.Getenv("SECRETS_PASSWORD")
	}
}

// ValidateOptions validate options
func (c *EncryptCommand) ValidateOptions() error {
	if len(c.Password) == 0 {
		return errors.New("password can't be blank")
	}

	return nil
}
