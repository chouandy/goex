package gormex

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

const loggableLogsTag = "gorm-loggable-logs"

// LoggableLogModel loggable log model
type LoggableLogModel struct {
	ID          uint64
	TriggerID   uint64
	AuditableID uint64
	Action      string
	Changes     []byte
	CreatedAt   time.Time
}

// NewLoggableLog new loggable log
func NewLoggableLog(scope *gorm.Scope, action string) (interface{}, error) {
	// Get loggable log type
	var loggableLogType reflect.Type
	for _, f := range scope.Fields() {
		if _, ok := f.Tag.Lookup(loggableLogsTag); ok {
			loggableLogType = f.Struct.Type.Elem()
			break
		}
	}
	if loggableLogType == nil {
		return nil, errors.New("loggable log type not found")
	}

	// New changes
	changes, err := NewChanges(scope, action)
	if err != nil {
		return nil, err
	}

	// New loggable log instance
	newDB := scope.NewDB()
	loggableLog := reflect.New(loggableLogType).Interface()
	newScope := newDB.NewScope(loggableLog)
	// Set columns
	if f, ok := scope.FieldByName("TriggerID"); ok {
		newScope.SetColumn("TriggerID", f.Field.Interface())
	}
	newScope.SetColumn("AuditableID", scope.PrimaryKeyValue())
	newScope.SetColumn("Action", action)
	newScope.SetColumn("Changes", changes.MustMarshal())

	return loggableLog, nil
}

// Changes changes
type Changes struct {
	Was map[string]interface{} `json:"was"`
	Is  map[string]interface{} `json:"is"`
}

// NewChanges new changes
func NewChanges(scope *gorm.Scope, action string) (*Changes, error) {
	changes := &Changes{
		Was: make(map[string]interface{}),
		Is:  make(map[string]interface{}),
	}

	switch action {
	case actionCreate:
		// Get loggable fields
		for _, f := range scope.Fields() {
			if _, ok := f.Tag.Lookup(loggableTag); ok && !f.IsBlank {
				changes.Is[f.DBName] = f.Field.Interface()
			}
		}
	case actionUpdate:
		// Get original entity field
		originalEntityField, ok := scope.FieldByName("OriginalEntity")
		if !ok {
			return nil, errors.New("original entity not found")
		}
		originalEntityScope := scope.NewDB().NewScope(
			originalEntityField.Field.Elem().Interface(),
		)
		// Get loggable fields
		for _, nf := range scope.Fields() {
			if _, ok := nf.Tag.Lookup(loggableTag); ok {
				if of, ok := originalEntityScope.FieldByName(nf.Name); ok {
					// New old value and new value variables
					var ov, nv interface{}
					// Check filed is pointer or not
					if nf.Field.Kind() == reflect.Ptr {
						// of is blank ? nil : value
						if of.IsBlank {
							ov = nil
						} else {
							ov = of.Field.Elem().Interface()
						}
						// nf is blank ? nil : value
						if nf.IsBlank {
							nv = nil
						} else {
							nv = nf.Field.Elem().Interface()
						}
					} else {
						ov = of.Field.Interface()
						nv = nf.Field.Interface()
					}
					// Check ov and nv is the same or not
					if ov != nv {
						changes.Was[nf.DBName] = ov
						changes.Is[nf.DBName] = nv
					}
				}
			}
		}
	case actionDelete:
		// Get loggable fields
		for _, f := range scope.Fields() {
			if _, ok := f.Tag.Lookup(loggableTag); ok && !f.IsBlank {
				changes.Was[f.DBName] = f.Field.Interface()
			}
		}
	}

	return changes, nil
}

// MustMarshal must marshal
func (c *Changes) MustMarshal() []byte {
	value, _ := json.Marshal(c)
	return value
}
