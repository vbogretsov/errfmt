package jsonerr

import (
	"encoding/json"
	"fmt"

	"github.com/vbogretsov/go-validation"
)

// Errors represent JSON serializer for go-validation.Errors.
type Errors validation.Errors

// MarshalJSON serializes validation errors into JSON.
func (es Errors) MarshalJSON() ([]byte, error) {
	path := ""
	errs := []jsonError{}

	for _, e := range es {
		marshal(e, path, &errs)
	}

	return json.Marshal(errs)
}

type jsonError struct {
	Path  string `json:"path,omitempty"`
	Error string `json:"error"`
}

func marshal(er error, path string, errs *[]jsonError) {
	switch x := er.(type) {
	case validation.Errors:
		for _, e := range []error(x) {
			marshal(e, path, errs)
		}
	case validation.StructError:
		v := validation.StructError(x)
		p := fmt.Sprintf("%s/%s", path, v.Field)

		for _, e := range []error(v.Errors) {
			marshal(e, p, errs)
		}
	case validation.SliceError:
		v := validation.SliceError(x)
		p := fmt.Sprintf("%s/%d", path, v.Index)

		for _, e := range []error(v.Errors) {
			marshal(e, p, errs)
		}
	default:
		*errs = append(*errs, jsonError{Path: path, Error: er.Error()})
	}
}
