package errors

// ErrResponse is an error std implementation for response
type ErrResponse struct {
	Name string
	Code int
}

var _ error = &ErrResponse{}

func (e *ErrResponse) Error() string {
	return e.Name
}
