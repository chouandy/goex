package sfnex

// Context context struct
type Context struct {
	Input     map[string]interface{}
	Output    map[string]interface{}
	Exception error

	// Extra
	Region string
}

// NewContext new context
func NewContext(input map[string]interface{}) (ctx *Context) {
	ctx = &Context{Input: input}

	return
}

// WrapTask wrap task
func (c *Context) WrapTask(task *Task) {
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
