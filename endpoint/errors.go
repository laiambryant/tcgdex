package endpoint

import "fmt"

type DecodeError struct {
	Resource string
	Err      error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("decode error for %s: %v", e.Resource, e.Err)
}

func (e *DecodeError) Unwrap() error {
	return e.Err
}
