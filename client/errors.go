package client

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type HTTPError struct {
	Status int
	URL    string
	Body   string
	Cause  error
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http %d for %s: %s", e.Status, e.URL, e.Body)
}

func (e *HTTPError) Unwrap() error {
	return e.Cause
}

type RequestError struct {
	Op  string
	Err error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request error during %s: %v", e.Op, e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}
