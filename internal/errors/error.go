package errors

// ErrorResponse is a struct for error response
type ErrorResponse struct {
	Name string
	Code int
}

var _ error = &ErrorResponse{}

func (e *ErrorResponse) Error() string {
	return e.Name
}
