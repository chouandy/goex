package api

import (
	"bytes"
	"strings"
	"text/template"
)

var tplConfig = template.Must(template.New("config").Parse(`
var scheme = os.Getenv("{{ .StructName }}_SCHEME")

var endpoint = os.Getenv("{{ .StructName }}_ENDPOINT")

var apiKey = os.Getenv("{{ .StructName }}_API_KEY")

// SetScheme set scheme
func SetScheme(s string) {
	scheme = s
}

// SetEndpoint set endpoint
func SetEndpoint(s string) {
	endpoint = s
}

// SetAPIKey set api key
func SetAPIKey(s string) {
	apiKey = s
}
`))

// ConfigGoCode config go code
func (a *API) ConfigGoCode() string {
	a.resetImports()
	a.imports = map[string]bool{
		"os": true,
	}

	var buf bytes.Buffer
	err := tplConfig.Execute(&buf, a)

	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + strings.TrimSpace(buf.String())
	return code
}
