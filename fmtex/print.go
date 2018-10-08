package fmtex

import "fmt"

// StructPrintln struct println
func StructPrintln(a ...interface{}) (n int, err error) {
	return fmt.Println(fmt.Sprintf("%+v", a...))
}
