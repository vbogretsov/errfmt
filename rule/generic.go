package rule

import (
	"fmt"
	"reflect"

	"github.com/vbogretsov/go-validation"
)

var (
	ParamInUnsupported = "unsupported"
	ParamInSupported   = "supported"
)

func unexpectedType(v interface{}) validation.Panic {
	return validation.Panic{
		Err: fmt.Errorf("unexpected type: %v", reflect.TypeOf(v)),
	}
}

func wrap(r func(interface{}) error) validation.Rule {
	return func(interface{}) func(interface{}) error {
		return r
	}
}

// NotNil creates validator to check whether a value is nil.
func NotNil(msg string) validation.Rule {
	return wrap(func(v interface{}) error {
		t := reflect.TypeOf(v)
		if t.Kind() != reflect.Ptr {
			return unexpectedType(v)
		}

		switch t.Elem().Kind() {
		case reflect.Interface,
			reflect.Ptr,
			reflect.Slice,
			reflect.Func,
			reflect.Map,
			reflect.Chan:

			if reflect.ValueOf(v).Elem().IsNil() {
				return validation.Error{Message: msg}
			}
		default:
			return unexpectedType(v)
		}

		return nil
	})
}

// In creates a validator to chech wheter an item belongs to the set provided.
func In(values []interface{}, msg string) validation.Rule {
	set := map[interface{}]bool{}
	for _, v := range values {
		set[v] = true
	}

	return wrap(func(v interface{}) error {
		vl := reflect.ValueOf(v)
		if vl.Type().Kind() != reflect.Ptr {
			return unexpectedType(v)
		}
		vl = vl.Elem()

		k := vl.Interface()
		if !set[k] {
			return validation.Error{Message: msg, Params: validation.Params{
				ParamInUnsupported: k,
				ParamInSupported:   values,
			}}
		}
		return nil
	})
}
