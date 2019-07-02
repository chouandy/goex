package gormex

const loggableTag = "gorm-loggable"

// LoggableInterface is used to get metadata from your models.
type LoggableInterface interface {
	isLoggable() bool
}

// LoggableModel loggable model struct
type LoggableModel struct {
	OriginalEntity interface{} `gorm:"-"`
	TriggerID      uint64      `gorm:"-"`
}

func (l LoggableModel) isLoggable() bool {
	return true
}

func isLoggable(value interface{}) bool {
	_, ok := value.(LoggableInterface)
	return ok
}
