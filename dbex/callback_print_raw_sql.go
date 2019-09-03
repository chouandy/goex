package dbex

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// PrintRawSQLCallback print raw sql callback
func PrintRawSQLCallback(scope *gorm.Scope) {
	fmt.Println(scope.SQL)
}

// RegisterPrintRawSQLCallback register print raw sql callback
func RegisterPrintRawSQLCallback() {
	DB.Callback().Query().Register("gorm:query_print_raw_sql", PrintRawSQLCallback)
	DB.Callback().Create().Register("gorm:create_print_raw_sql", PrintRawSQLCallback)
	DB.Callback().Update().Register("gorm:update_print_raw_sql", PrintRawSQLCallback)
	DB.Callback().Delete().Register("gorm:delete_print_raw_sql", PrintRawSQLCallback)
}
