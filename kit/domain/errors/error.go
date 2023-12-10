package errors

import (
	"fmt"
	"runtime"
)

type ProjectError struct {
	object    string
	errorType string
	line      int
	cause     string
}

func NewProjectError(object, errorType, cause string) *ProjectError {
	_, _, line, _ := runtime.Caller(1)
	return &ProjectError{
		object:    object,
		errorType: errorType,
		line:      line,
		cause:     cause,
	}
}

func (e *ProjectError) Error() string {
	return fmt.Sprintf("Error in %s at line %d: %s - %s", e.object, e.line, e.errorType, e.cause)
}

func (e *ProjectError) ErrorType() string {
	return e.errorType
}
