package validation

import (
	"errors"
	"fmt"
	"strings"
)

// Panic represents an internal error of schema package.
type Panic struct {
	Err error
}

// Error rerurns string representation of the internal error.
func (e Panic) Error() string {
	return e.Err.Error()
}

// Errors represents errors collection.
type Errors []error

// Error returns string representation of Errors.
func (e Errors) Error() string {
	errors := []string{}
	for _, i := range e {
		errors = append(errors, i.Error())
	}
	return strings.Join(errors, ", ")
}

// StructError represents a struct field validation error.
type StructError struct {
	Field  string
	Errors Errors
}

// Error returns string representation of a Struct.
func (e StructError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Errors.Error())
}

// SliceError represents a slice item validation error.
type SliceError struct {
	Index  int
	Errors Errors
}

// Error returns string representation of a SliceError.
func (e SliceError) Error() string {
	return fmt.Sprintf("%d: %s", e.Index, e.Errors.Error())
}

// MapError represents a map item validation error.
// NOTE(vbogretsov): map validation should be implemented later.
type MapError struct {
	Key    interface{}
	Errors Errors
}

// Error returns string representation of a MapError.
func (e MapError) Error() string {
	return fmt.Sprintf("%v: %s", e.Key, e.Errors.Error())
}

// Error creates validation errros from a single error.
func Error(text string) Errors {
	return Errors([]error{errors.New(text)})
}

// Errorf creates validation errros from a single error.
func Errorf(format string, args ...interface{}) Errors {
	return Errors([]error{fmt.Errorf(format, args...)})
}
