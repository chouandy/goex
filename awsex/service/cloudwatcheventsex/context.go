package cloudwatcheventsex

import "strings"

// Context context struct
type Context struct {
	Name      string                 `json:"name"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Exception error                  `json:"-"`
}

// Dispatch dispatch event
func (c *Context) Dispatch() {
	// Get event
	event := Dispatcher.Get(c.Name)
	// Middlewares
	for _, middleware := range event.Middlewares {
		if err := middleware(c); err != nil {
			c.Exception = err
			return
		}
	}
	// Handler
	event.Handler(c)
}

// GetCamelCaseName get camel case name
func (c *Context) GetCamelCaseName() string {
	return strings.Title(strings.Replace(c.Name, "_", " ", -1))
}
