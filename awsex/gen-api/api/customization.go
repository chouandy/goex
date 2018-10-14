package api

import (
	"bytes"
	"strings"
	"text/template"
)

var tplCustomization = template.Must(template.New("customization").Parse(`
func init() {
	initClient = func(s *{{ .StructName }}) {
		// Set default request headers
		s.Handlers.Build.PushBack(func(r *aws.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
			r.HTTPRequest.Header.Add("X-Api-Key", apiKey)
		})
		// Set url
		var url string
		if len(scheme) > 0 {
			url = scheme + "://" + endpoint
		} else {
			url = "https://" + endpoint
		}
		s.Client.Config.EndpointResolver = aws.ResolveWithEndpointURL(url)
	}
}
`))

// CustomizationGoCode customization go code
func (a *API) CustomizationGoCode() string {
	a.resetImports()
	a.imports = map[string]bool{
		"github.com/aws/aws-sdk-go-v2/aws": true,
	}

	var buf bytes.Buffer
	err := tplCustomization.Execute(&buf, a)

	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + strings.TrimSpace(buf.String())
	return code
}
