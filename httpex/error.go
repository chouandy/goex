package httpex

import (
	"fmt"
	"regexp"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcErrorInlineExpression grpc error inline expression
const GrpcErrorInlineExpression = "statusCode: ([0-9]{3})(?:|, code: ([0-9.].*)), message: (.*)"

// Error error struct
type Error interface {
	// Satisfy the generic error interface.
	error

	// The status code of the HTTP response.
	StatusCode() int

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the inline representation of the error.
	ErrorInline() string

	// Returns the inline representation of the error.
	GrpcErrorInline() string
}

// NewError new error
func NewError(statusCode int, code, message string) Error {
	return newBaseError(statusCode, code, message)
}

type baseError struct {
	statusCode int
	code       string
	message    string
}

func newBaseError(statusCode int, code, message string) *baseError {
	b := &baseError{
		statusCode: statusCode,
		code:       code,
		message:    message,
	}

	return b
}

// StatusCode returns the wrapped status code for the error
func (b baseError) StatusCode() int {
	return b.statusCode
}

// Code returns the short phrase depicting the classification of the error.
func (b baseError) Code() string {
	return b.code
}

// Message returns the error details message.
func (b baseError) Message() string {
	return b.message
}

// Error returns the json representation of the error.
func (b baseError) Error() string {
	if len(b.code) == 0 {
		return `{"message":"` + b.message + `"}`
	}
	return `{"code":"` + b.code + `","message":"` + b.message + `"}`
}

// ErrorInline returns the inline representation of the error.
func (b baseError) ErrorInline() string {
	if len(b.code) == 0 {
		return b.message
	}
	return `code: ` + b.code + `, message: ` + b.message
}

// String returns the json representation of the error.
// Alias for Error to satisfy the stringer interface.
func (b baseError) String() string {
	return b.Error()
}

func (b baseError) MarshalJSON() ([]byte, error) {
	return []byte(b.Error()), nil
}

// NewGrpcError new grpc error
func NewGrpcError(c codes.Code, statusCode int, code, message string) error {
	return status.Errorf(c, newBaseError(statusCode, code, message).GrpcErrorInline())
}

// GrpcErrorInline returns the inline representation of the error for grpc.
func (b baseError) GrpcErrorInline() string {
	if len(b.code) == 0 {
		return fmt.Sprintf(`statusCode: %d, message: %s`, b.statusCode, b.message)
	}
	return fmt.Sprintf(`statusCode: %d, code: %s, message: %s`, b.statusCode, b.code, b.message)
}

// ParseGrpcErrorInline parse grpc error inline
func ParseGrpcErrorInline(s string) (httpErr Error, ok bool) {
	// New regexp
	reg, err := regexp.Compile(GrpcErrorInlineExpression)
	if err != nil {
		return
	}

	// Check match string
	if !reg.MatchString(s) {
		return
	}

	// Find match string
	values := reg.FindStringSubmatch(s)
	if len(values) != 4 {
		return
	}

	// convert status code to int
	statusCode, err := strconv.Atoi(values[1])
	if err != nil {
		return
	}

	httpErr = NewError(statusCode, values[2], values[3])
	ok = true
	return
}
