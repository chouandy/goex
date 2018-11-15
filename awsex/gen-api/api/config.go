package api

import (
	"bytes"
	"strings"
	"text/template"
)

var tplConfig = template.Must(template.New("config").Funcs(
	template.FuncMap{
		"ToUpper": strings.ToUpper,
	},
).Parse(`
var scheme = os.Getenv("{{ .StructName | ToUpper }}_SCHEME")

var endpoint = os.Getenv("{{ .StructName | ToUpper }}_ENDPOINT")

var apiKey = os.Getenv("{{ .StructName | ToUpper }}_API_KEY")

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
