package sfnex

// Context context struct
type Context struct {
	Input     map[string]interface{}
	Output    map[string]interface{}
	Exception error
}

// NewContext new context
func NewContext(input map[string]interface{}) (ctx *Context) {
	ctx = &Context{Input: input}

	return
}

// GetState get state
func (c *Context) GetState() string {
	// Get state
	if _, ok := c.Input["state"]; ok {
		return c.Input["state"].(string)
	}
	// Check exception
	if _, ok := c.Input["Exception"]; ok {
		return "Exception"
	}

	return ""
}

// WrapTask wrap task
func (c *Context) WrapTask() {
	// Find task
	task := Dispatcher.Get(c.GetState())
	// Middlewares
	for _, middleware := range task.Middlewares {
		if err := middleware(c); err != nil {
			c.Exception = err
			return
		}
	}
	// Handler
	task.Handler(c)
}
