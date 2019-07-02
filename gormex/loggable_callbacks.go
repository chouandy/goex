package gormex

import (
	"github.com/jinzhu/gorm"
)

const (
	actionCreate = "create"
	actionUpdate = "update"
	actionDelete = "delete"
)

func init() {
	gorm.DefaultCallback.Query().After("gorm:after_query").Register("loggable:query", loggableQuery)
	gorm.DefaultCallback.Create().After("gorm:after_create").Register("loggable:create", loggableCreate)
	gorm.DefaultCallback.Update().After("gorm:after_update").Register("loggable:update", loggableUpdate)
	gorm.DefaultCallback.Delete().After("gorm:after_delete").Register("loggable:delete", loggableDelete)
}

func loggableQuery(scope *gorm.Scope) {
	if !isLoggable(scope.Value) {
		return
	}

	if f, ok := scope.FieldByName("OriginalEntity"); ok {
		f.Set(scope.IndirectValue())
	}
}

func loggableCreate(scope *gorm.Scope) {
	// Check is loggable or not
	if !isLoggable(scope.Value) {
		return
	}
	// New loggable log
	loggableLog, err := NewLoggableLog(scope, actionCreate)
	if err != nil {
		return
	}
	// Save loggable log
	scope.Err(scope.NewDB().Save(loggableLog).Error)
}

func loggableUpdate(scope *gorm.Scope) {
	// Check is loggable or not
	if !isLoggable(scope.Value) {
		return
	}
	// New loggable log
	loggableLog, err := NewLoggableLog(scope, actionUpdate)
	if err != nil {
		return
	}
	// Save loggable log
	scope.Err(scope.NewDB().Save(loggableLog).Error)
}

func loggableDelete(scope *gorm.Scope) {
	// Check is loggable or not
	if !isLoggable(scope.Value) {
		return
	}
	// New loggable log
	loggableLog, err := NewLoggableLog(scope, actionDelete)
	if err != nil {
		return
	}
	// Save loggable log
	scope.Err(scope.NewDB().Save(loggableLog).Error)
}
