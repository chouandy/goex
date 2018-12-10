package sfnex

// Dispatcher dispatcher instance
var Dispatcher *dispatcher

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
type dispatcher struct {
	Tasks map[string]*Task
}

// NewDispatcher new task dispatcher
func NewDispatcher() {
	Dispatcher = &dispatcher{
		Tasks: map[string]*Task{},
	}
}

// Add add task by state
func (d *dispatcher) Add(state string, task *Task) {
	d.Tasks[state] = task
}

// Get get task by state
func (d *dispatcher) Get(state string) *Task {
	task, ok := d.Tasks[state]
	if !ok {
		return &Task{
			Handler: NotFoundHandler,
		}
	}
	return task
}
