package rule

import (
	"errors"
	"reflect"
	"time"

	"github.com/vbogretsov/go-validation"
)

var (
	eTypeMismatch = errors.New("between parameters should have same type")
)

var (
	ParamNumMin = "min"
	ParamNumMax = "max"
)

func int64rule(fn func(int64) error) validation.Rule {
	return wrap(func(v interface{}) error {
		switch x := v.(type) {
		case *int:
			return fn(int64(*x))
		default:
			return unexpectedType(v)
		}
	})
}

func uint64rule(fn func(uint64) error) validation.Rule {
	return wrap(func(v interface{}) error {
		switch x := v.(type) {
		case *uint:
			return fn(uint64(*x))
		default:
			return unexpectedType(v)
		}
	})
}

func float64rule(fn func(v float64) error) validation.Rule {
	return wrap(func(v interface{}) error {
		switch x := v.(type) {
		case *float32:
			return fn(float64(*x))
		case *float64:
			return fn(float64(*x))
		default:
			return unexpectedType(v)
		}
	})
}

func timerule(fn func(time.Time) error) validation.Rule {
	return wrap(func(v interface{}) error {
		t, ok := v.(*time.Time)
		if !ok {
			return unexpectedType(v)
		}
		return fn(*t)
	})
}

func errorMin(min interface{}, msg string) validation.Error {
	return validation.Error{
		Message: msg,
		Params: validation.Params{
			ParamNumMin: min,
		},
	}
}

func errorMax(max interface{}, msg string) validation.Error {
	return validation.Error{
		Message: msg,
		Params: validation.Params{
			ParamNumMax: max,
		},
	}
}

func errorBetween(a, b interface{}, msg string) validation.Error {
	return validation.Error{
		Message: msg,
		Params: validation.Params{
			ParamNumMin: a,
			ParamNumMax: b,
		},
	}
}

// Min creates validator to check whether a number is not less than the
// value provided.
func Min(min interface{}, msg string) validation.Rule {
	switch x := min.(type) {
	case int:
		a := int64(x)
		return int64rule(func(v int64) error {
			if v < a {
				return errorMin(min, msg)
			}
			return nil
		})
	case uint:
		a := uint64(x)
		return uint64rule(func(v uint64) error {
			if v < a {
				return errorMin(min, msg)
			}
			return nil
		})
	case float32:
		a := float64(x)
		return float64rule(func(v float64) error {
			if v < a {
				return errorMin(min, msg)
			}
			return nil
		})
	case float64:
		a := float64(x)
		return float64rule(func(v float64) error {
			if v < a {
				return errorMin(min, msg)
			}
			return nil
		})
	case time.Time:
		a := time.Time(x)
		return timerule(func(v time.Time) error {
			if v.Sub(a) < 0 {
				return errorMin(min, msg)
			}
			return nil
		})
	default:
		return wrap(func(v interface{}) error { return unexpectedType(min) })
	}
}

// Max creates validator to check whether a number is not great than the
// value provided.
func Max(max interface{}, msg string) validation.Rule {
	switch x := max.(type) {
	case int:
		a := int64(x)
		return int64rule(func(v int64) error {
			if v > a {
				return errorMax(max, msg)
			}
			return nil
		})
	case uint:
		a := uint64(x)
		return uint64rule(func(v uint64) error {
			if v > a {
				return errorMax(max, msg)
			}
			return nil
		})
	case float32:
		a := float64(x)
		return float64rule(func(v float64) error {
			if v > a {
				return errorMax(max, msg)
			}
			return nil
		})
	case float64:
		a := float64(x)
		return float64rule(func(v float64) error {
			if v > a {
				return errorMax(max, msg)
			}
			return nil
		})
	case time.Time:
		a := time.Time(x)
		return timerule(func(v time.Time) error {
			if a.Sub(v) < 0 {
				return errorMax(max, msg)
			}
			return nil
		})
	default:
		return wrap(func(v interface{}) error { return unexpectedType(max) })
	}
}

// Between creates validator to check whether a number is the range provided.
func Between(a, b interface{}, msg string) validation.Rule {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta != tb {
		return wrap(func(v interface{}) error {
			return validation.Panic{Err: eTypeMismatch}
		})
	}

	switch a.(type) {
	case int:
		l := reflect.ValueOf(a).Int()
		h := reflect.ValueOf(b).Int()

		return int64rule(func(v int64) error {
			if v < l || v > h {
				return errorBetween(a, b, msg)
			}
			return nil
		})
	case uint:
		l := reflect.ValueOf(a).Uint()
		h := reflect.ValueOf(b).Uint()

		return uint64rule(func(v uint64) error {
			if v < l || v > h {
				return errorBetween(a, b, msg)
			}
			return nil
		})
	case float32, float64:
		l := reflect.ValueOf(a).Float()
		h := reflect.ValueOf(b).Float()

		return float64rule(func(v float64) error {
			if v < l || v > h {
				return errorBetween(a, b, msg)
			}
			return nil
		})
	case time.Time:
		l := a.(time.Time)
		h := b.(time.Time)

		return timerule(func(v time.Time) error {
			if v.Sub(l) < 0 || h.Sub(v) < 0 {
				return errorBetween(a, b, msg)
			}
			return nil
		})
	default:
		return wrap(func(v interface{}) error { return unexpectedType(v) })
	}
}
