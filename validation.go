package validation

import (
	"errors"
	"reflect"
)

var (
	errorArgs = Panic{Err: errors.New("expected pointer to struct")}
	errorAttr = Panic{Err: errors.New("Attr must return a pointer")}
)

// Rule represents a validation function.
type Rule func(interface{}) error

// Attr represents an attribute getter of a struct.
type Attr func(interface{}) interface{}

// Context represents additional info that can be required for custom
// validation scenarios. The member Ptr contains pointer the value to be
// validated. The member Ctx contains additional information required for
// validation (e.g. database connection).
type Context struct {
	Ptr interface{}
	Ctx interface{}
}

// Field represents a schema field.
type Field struct {
	Attr  Attr
	Rules []Rule
}

// Rules combines several rules into single one.
func Rules(rules []Rule) Rule {
	return func(v interface{}) error {
		errs := []error{}
		for _, rule := range rules {
			if err := rule(v); err != nil {
				if _, ok := err.(Panic); ok {
					return err
				}
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return Errors(errs)
		}
		return nil
	}
}

type structRule struct {
	ftab   map[uintptr]string
	fields []Field
}

// Struct struct validation rule.
func Struct(v interface{}, tag string, fields []Field) (Rule, error) {
	tp := reflect.TypeOf(v)
	if tp.Kind() != reflect.Ptr {
		return nil, errorArgs
	}

	tp = tp.Elem()
	if tp.Kind() != reflect.Struct {
		return nil, errorArgs
	}

	ftab := map[uintptr]string{}
	for i := 0; i < tp.NumField(); i++ {
		ft := tp.Field(i)
		if tag == "" {
			ftab[ft.Offset] = ft.Name
		} else {
			fname := ft.Tag.Get(tag)
			if fname == "" {
				fname = ft.Name
			}
			ftab[ft.Offset] = fname
		}
	}

	s := structRule{
		ftab:   ftab,
		fields: fields,
	}

	return s.validate, nil
}

func (s structRule) validate(v interface{}) error {
	tp := reflect.TypeOf(v)
	if tp.Kind() != reflect.Ptr {
		return errorArgs
	}

	tp = tp.Elem()
	if tp.Kind() != reflect.Struct {
		return errorArgs
	}

	errs := []error{}
	for _, f := range s.fields {
		attr := f.Attr(v)

		name, value, err := s.nameOf(v, attr)
		if err != nil {
			return err
		}

		fe := []error{}
		for _, rule := range f.Rules {
			if err := rule(value); err != nil {
				if _, ok := err.(Panic); ok {
					return err
				}
				if e, ok := err.(Errors); ok {
					for _, i := range e {
						fe = append(fe, i)
					}
				} else {
					fe = append(fe, err)
				}
			}
		}

		if len(fe) > 0 {
			errs = append(errs, StructError{Field: name, Errors: fe})
		}
	}

	if len(errs) > 0 {
		return Errors(errs)
	}

	return nil
}

func (s structRule) nameOf(self, member interface{}) (string, interface{}, error) {
	var name string
	var ptr interface{}
	var val interface{}

	if ctx, ok := member.(Context); ok {
		ptr = ctx.Ptr
		val = ctx.Ctx
	} else {
		ptr = member
		val = member
	}

	fv := reflect.ValueOf(ptr)
	if fv.Kind() != reflect.Ptr {
		return name, nil, errorAttr
	}

	if ptr != self {
		name = s.ftab[fv.Pointer()-reflect.ValueOf(self).Pointer()]
	}

	return name, val, nil
}
