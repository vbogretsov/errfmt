package json

import (
	"encoding/json"
	"fmt"

	"github.com/vbogretsov/go-validation"
)

type jsonError struct {
	Path  string `json:"path,omitempty"`
	Error string `json:"error"`
}

// Formatter represents valdation error message formatter.
type Formatter func(validation.Error) string

// DefaultFormatter is the default validation error message formatter.
func DefaultFormatter(e validation.Error) string {
	return e.Message
}

// Joiner defines interface for building path to an item in the validation
// errors tree.
type Joiner interface {
	Struct(base, child string) string
	Slice(base string, index int) string
}

type joiner struct{}

func (joiner) Struct(base, child string) string {
	return fmt.Sprintf("%s.%s", base, child)
}

func (joiner) Slice(base string, index int) string {
	return fmt.Sprintf("%s[%d]", base, index)
}

// DefaultJoiner if the default implementation of the PathBuilder interface.
var DefaultJoiner = joiner{}

// Error represents a json serializable error.
type Error interface {
	error
	json.Marshaler
}

type marshaler struct {
	errors    validation.Errors
	formatter Formatter
	joiner    Joiner
}

// New creates new json serializable error from validation errors.
func New(errors validation.Errors, formatter Formatter, joiner Joiner) Error {

	return &marshaler{
		errors:    errors,
		formatter: formatter,
		joiner:    joiner,
	}
}

// Error is implemented to make Marshaler compatible with the standard error
// interface.
func (m *marshaler) Error() string {
	return m.errors.Error()
}

// MarshalJSON serializes validation errors into JSON.
func (m *marshaler) MarshalJSON() ([]byte, error) {
	path := ""
	errs := []jsonError{}

	for _, e := range m.errors {
		m.marshal(e, path, &errs)
	}

	return json.Marshal(errs)
}

func (m *marshaler) marshal(er error, path string, errs *[]jsonError) {
	switch x := er.(type) {
	case validation.Errors:
		for _, e := range []error(x) {
			m.marshal(e, path, errs)
		}
	case validation.StructError:
		v := validation.StructError(x)
		p := m.joiner.Struct(path, v.Field)

		for _, e := range []error(v.Errors) {
			m.marshal(e, p, errs)
		}
	case validation.SliceError:
		v := validation.SliceError(x)
		p := m.joiner.Slice(path, v.Index)

		for _, e := range []error(v.Errors) {
			m.marshal(e, p, errs)
		}
	case validation.Error:
		e := er.(validation.Error)
		*errs = append(*errs, jsonError{Path: path, Error: m.formatter(e)})
	default:
		*errs = append(*errs, jsonError{Path: path, Error: er.Error()})
	}
}
