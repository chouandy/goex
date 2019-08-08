package cloudwatcheventsex

// Dispatcher dispatcher
var Dispatcher *dispatcher

// HandlerFunc handler func
type HandlerFunc func(ctx *Context)

// MiddlewareFunc middleware func
type MiddlewareFunc func(ctx *Context) error

// Event event struct
type Event struct {
	Handler     HandlerFunc
	Middlewares []MiddlewareFunc
}

// dispatcher event dispatcher struct
type dispatcher struct {
	Events map[string]*Event
}

// Add add event by name
func (d *dispatcher) Add(name string, event *Event) {
	d.Events[name] = event
}

// Get get event by name
func (d *dispatcher) Get(name string) *Event {
	event, ok := d.Events[name]
	if !ok {
		return &Event{
			Handler: NotFoundHandler,
		}
	}
	return event
}

func init() {
	Dispatcher = &dispatcher{
		Events: make(map[string]*Event),
	}
}
