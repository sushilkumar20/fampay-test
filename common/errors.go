package common

import "fmt"

type NotFound struct {
	Message string
}

func (nf *NotFound) Error() string {
	return nf.Message
}

type IllegalState struct {
	Message string
}

func (is *IllegalState) Error() string {
	return is.Message
}

type ErrorResponse struct {
	StatusCode int
	Payload    string
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d: %s", r.StatusCode, r.Payload)
}

type CodeError struct {
	Code    string
	Message string
}

func (e *CodeError) Error() string { return e.Message }

type UnauthorizedError struct {
	Code    string
	Message string
}

func (e *UnauthorizedError) Error() string { return e.Message }

type BadStateError struct {
	Code    string
	Message string
}

func (e *BadStateError) Error() string { return e.Message }

type InternalError struct {
	Code    string
	Message string
}

func (e *InternalError) Error() string { return e.Message }

type AlreadyExist struct {
	Code    string
	Message string
}

func (e *AlreadyExist) Error() string { return e.Message }
