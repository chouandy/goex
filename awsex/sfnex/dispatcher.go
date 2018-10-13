package sfnex

// HandlerFunc handler func
type HandlerFunc func(ctx *Context)

// MiddlewareFunc middleware func
type MiddlewareFunc func(ctx *Context) error

// Task task struct
type Task struct {
	Handler     HandlerFunc
	Middlewares []MiddlewareFunc
}

// Dispatcher task dispatcher struct
type Dispatcher struct {
	Tasks map[string]*Task
}

// NewDispatcher new task dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Tasks: map[string]*Task{},
	}
}

// Add add task by state
func (d *Dispatcher) Add(state string, task *Task) {
	d.Tasks[state] = task
}

// Get get task by state
func (d *Dispatcher) Get(state string) *Task {
	task, ok := d.Tasks[state]
	if !ok {
		return &Task{
			Handler: NotFoundHandler,
		}
	}
	return task
}
