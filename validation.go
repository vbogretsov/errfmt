package validation

import (
	"reflect"
)

var (
	errorArgs = Panic{Err: "expected pointer to struct"}
	errorAttr = Panic{Err: "Attr must return a pointer"}
)

// Rule represents a validation function.
type Rule func(interface{}) error

// Attr represents an attribute getter of a struct.
type Attr func(interface{}) interface{}

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
			if err := rule(attr); err != nil {
				if _, ok := err.(Panic); ok {
					return err
				}
				fe = append(fe, err)
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
