package json_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"

	"github.com/vbogretsov/go-validation"
	jsonerr "github.com/vbogretsov/go-validation/json"
)

const (
	eBlank         = "cannot be blank"
	eEmail         = "invalid email"
	ePasswordShort = "password to short"
	eDigitsOnly    = "only digits are alowed"
	eLettersOnly   = "only letters are alowed"
)

type jsonError struct {
	Error string `json:"error"`
	Path  string `json:"path,omitempty"`
}

type fixture struct {
	err validation.Errors
	rep []jsonError
}

func (fx fixture) check() error {
	buf, err := json.Marshal(jsonerr.New(
		fx.err,
		jsonerr.DefaultFormatter,
		jsonerr.DefaultJoiner))

	if err != nil {
		return err
	}

	var v []jsonError
	if err := json.Unmarshal(buf, &v); err != nil {
		return err
	}

	if !reflect.DeepEqual(fx.rep, v) {
		return errors.New(strings.Join(pretty.Diff(fx.rep, v), " "))
	}

	return nil
}

var fixtures = []fixture{
	{
		err: validation.Errors([]error{
			validation.StructError{
				Field:  "username",
				Errors: []error{errors.New(eEmail)},
			},
			validation.StructError{
				Field:  "password",
				Errors: []error{errors.New(ePasswordShort)},
			},
		}),
		rep: []jsonError{
			{
				Error: eEmail,
				Path:  ".username",
			},
			{
				Error: ePasswordShort,
				Path:  ".password",
			},
		},
	},
	{
		err: validation.Errors([]error{
			validation.StructError{
				Field: "username",
				Errors: []error{
					errors.New(eBlank),
					errors.New(eEmail),
				},
			},
			validation.StructError{
				Field: "password",
				Errors: []error{
					errors.New(eBlank),
					errors.New(ePasswordShort),
				},
			},
			validation.StructError{
				Field: "addresses",
				Errors: []error{
					validation.SliceError{
						Index: 0,
						Errors: []error{
							validation.StructError{
								Field: "zipcode",
								Errors: []error{
									errors.New(eDigitsOnly),
								},
							},
							validation.StructError{
								Field: "country",
								Errors: []error{
									errors.New(eLettersOnly),
								},
							},
						},
					},
					validation.SliceError{
						Index: 3,
						Errors: []error{
							validation.StructError{
								Field: "zipcode",
								Errors: []error{
									errors.New(eBlank),
									errors.New(eDigitsOnly),
								},
							},
							validation.StructError{
								Field: "country",
								Errors: []error{
									errors.New(eBlank),
									errors.New(eLettersOnly),
								},
							},
						},
					},
				},
			},
		}),
		rep: []jsonError{
			{
				Error: eBlank,
				Path:  ".username",
			},
			{
				Error: eEmail,
				Path:  ".username",
			},
			{
				Error: eBlank,
				Path:  ".password",
			},
			{
				Error: ePasswordShort,
				Path:  ".password",
			},
			{
				Error: eDigitsOnly,
				Path:  ".addresses[0].zipcode",
			},
			{
				Error: eLettersOnly,
				Path:  ".addresses[0].country",
			},
			{
				Error: eBlank,
				Path:  ".addresses[3].zipcode",
			},
			{
				Error: eDigitsOnly,
				Path:  ".addresses[3].zipcode",
			},
			{
				Error: eBlank,
				Path:  ".addresses[3].country",
			},
			{
				Error: eLettersOnly,
				Path:  ".addresses[3].country",
			},
		},
	},
}

func TestMarshal(t *testing.T) {
	for _, fx := range fixtures {
		t.Run("", func(t *testing.T) {
			if err := fx.check(); err != nil {
				t.Error(err)
			}
		})
	}
}
