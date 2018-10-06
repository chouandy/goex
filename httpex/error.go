package httpex

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
