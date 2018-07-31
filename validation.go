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
type Rule func(interface{}) func(interface{}) error

// Attr represents an attribute getter of a struct.
type Attr func(interface{}) interface{}

// Field represents a schema field.
type Field struct {
	Attr  Attr
	Rules []Rule
}

// Func creates a Rule from function.
func Func(r func(interface{}) error) Rule {
	return func(interface{}) func(interface{}) error { return r }
}

// Rules combines several rules into single one.
func Rules(rules []Rule) Rule {
	return func(ctx interface{}) func(v interface{}) error {
		return func(v interface{}) error {
			errs := []error{}
			for _, rule := range rules {
				if err := rule(ctx)(v); err != nil {
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
}

func panicRule(err error) Rule {
	return func(interface{}) func(interface{}) error {
		return func(interface{}) error {
			return err
		}
	}
}

type structRule struct {
	ftab   map[uintptr]string
	fields []Field
}

// Struct struct validation rule.
func Struct(v interface{}, tag string, fields []Field) Rule {
	tp := reflect.TypeOf(v)
	if tp.Kind() != reflect.Ptr {
		return panicRule(errorArgs)
	}

	tp = tp.Elem()
	if tp.Kind() != reflect.Struct {
		return panicRule(errorArgs)
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

	return s.validate
}

func (s structRule) validate(ctx interface{}) func(interface{}) error {
	return func(v interface{}) error {
		tp := reflect.TypeOf(v)
		if tp.Kind() != reflect.Ptr {
			return errorArgs
		}

		tp = tp.Elem()
		if tp.Kind() != reflect.Struct {
			return errorArgs
		}

		self := reflect.ValueOf(v).Pointer()

		errs := []error{}
		for _, f := range s.fields {
			attr := f.Attr(v)

			fv := reflect.ValueOf(attr)
			if fv.Kind() != reflect.Ptr {
				return errorAttr
			}

			name := ""
			if attr != v {
				name = s.ftab[fv.Pointer()-self]
			}

			fe := []error{}
			for _, rule := range f.Rules {
				if err := rule(ctx)(attr); err != nil {
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
}
