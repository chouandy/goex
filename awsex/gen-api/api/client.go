package api

import (
	"bytes"
	"strings"
	"text/template"
)

// A tplClient defines the template for the service client.
var tplClient = template.Must(template.New("client").Parse(`
// Client {{ .PackageName }} client
var Client *{{ .StructName }}

// InitClient init {{ .PackageName }} client
func InitClient() error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	cfg.Region = os.Getenv("REGION")
	Client = New(cfg)

	return nil
}

// InitClientMiddleware init {{ .PackageName }} client middleware
func InitClientMiddleware(ctx *apigatewayex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init {{ .StructName }} Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return httpex.NewError(500, "", "Failed to init {{ .PackageName }} client")
		}
		fmt.Println("done")
	}

	return nil
}

// InitClientTaskMiddleware init {{ .PackageName }} client task middleware
func InitClientTaskMiddleware(ctx *sfnex.Context) error {
	if Client == nil {
		fmt.Print("[Middleware] Init {{ .StructName }} Client...")
		if err := InitClient(); err != nil {
			fmt.Println(err)
			return errors.New("Failed to init {{ .PackageName }} client")
		}
		fmt.Println("done")
	}

	return nil
}
`))

// ClientGoCode client go code
func (a *API) ClientGoCode() string {
	a.resetImports()
	a.imports = map[string]bool{
		"errors": true,
		"fmt":    true,
		"os":     true,
		"github.com/aws/aws-sdk-go-v2/aws/external":           true,
		"github.com/chouandy/goex/awsex/service/apigatewayex": true,
		"github.com/chouandy/goex/awsex/service/sfnex":        true,
		"github.com/chouandy/goex/httpex":                     true,
	}

	var buf bytes.Buffer
	err := tplClient.Execute(&buf, a)

	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + strings.TrimSpace(buf.String())
	return code
}
