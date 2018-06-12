package validation

import (
	"fmt"
	"strings"
)

// Panic represents an internal error of schema package.
// TODO(vbogretsov): replace Err to error.
type Panic struct {
	Err string
}

// Error rerurns string representation of the internal error.
func (e Panic) Error() string {
	return e.Err
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
