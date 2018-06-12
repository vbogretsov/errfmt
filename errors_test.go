package validation_test

import (
	"errors"
	"testing"

	"github.com/vbogretsov/go-validation"
)

func TestPanic(t *testing.T) {
	msg := "test"
	err := validation.Panic{Err: msg}

	if err.Error() != msg {
		t.Errorf("expected '%s' but was '%s'", msg, err.Error())
	}
}

func TestStructError(t *testing.T) {
	e := validation.StructError{
		Field:  "a",
		Errors: []error{errors.New("1"), errors.New("2")},
	}

	exp := "a: 1, 2"
	act := e.Error()

	if exp != act {
		t.Errorf("expected '%s' but got '%s", exp, act)
	}
}

func TestSliceError(t *testing.T) {
	e := validation.SliceError{
		Index:  1,
		Errors: []error{errors.New("1"), errors.New("2")},
	}

	exp := "1: 1, 2"
	act := e.Error()

	if exp != act {
		t.Errorf("expected '%s' but got '%s", exp, act)
	}
}

func TestMapError(t *testing.T) {
	e := validation.MapError{
		Key:    "k1",
		Errors: []error{errors.New("1"), errors.New("2")},
	}

	exp := "k1: 1, 2"
	act := e.Error()

	if exp != act {
		t.Errorf("expected '%s' but got '%s", exp, act)
	}
}
