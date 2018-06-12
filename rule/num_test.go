package rule_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vbogretsov/go-validation/rule"
)

func TestMin(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		msg := "cannot be less than %v"
		min := 10
		fun := rule.Min(min, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := min - 1
			assertError(t, fmt.Errorf(msg, min), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			assertOk(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := min + 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Uint", func(t *testing.T) {
		msg := "cannot be less than %v"
		min := uint(10)
		fun := rule.Min(min, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := min - 1
			assertError(t, fmt.Errorf(msg, min), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			assertOk(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := min + 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Float", func(t *testing.T) {
		msg := "cannot be less than %v"
		min := 10.0
		fun := rule.Min(min, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := min - 1
			assertError(t, fmt.Errorf(msg, min), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			assertOk(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := min + 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Time", func(t *testing.T) {
		msg := "cannot be less than %v"
		min := time.Date(2018, 2, 9, 0, 0, 0, 0, time.Local)
		fun := rule.Min(min, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := time.Date(2018, 2, 8, 0, 0, 0, 0, time.Local)
			assertError(t, fmt.Errorf(msg, min), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			assertOk(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := time.Date(2018, 2, 10, 0, 0, 0, 0, time.Local)
			assertOk(t, fun(&v))
		})
	})
}

func TestNumMax(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		msg := "cannot be great than %v"
		max := 10
		fun := rule.Max(max, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := max + 1
			assertError(t, fmt.Errorf(msg, max), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			assertOk(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := max - 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Uint", func(t *testing.T) {
		msg := "cannot be great than %v"
		max := uint(10)
		fun := rule.Max(max, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := max + 1
			assertError(t, fmt.Errorf(msg, max), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			assertOk(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := max - 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Float", func(t *testing.T) {
		msg := "cannot be great than %v"
		max := 10.0
		fun := rule.Max(max, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := max + 1
			assertError(t, fmt.Errorf(msg, max), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			assertOk(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := max - 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Time", func(t *testing.T) {
		msg := "cannot be great than %d"
		max := time.Date(2018, 2, 5, 0, 0, 0, 0, time.Local)
		fun := rule.Max(max, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := time.Date(2018, 2, 6, 0, 0, 0, 0, time.Local)
			assertError(t, fmt.Errorf(msg, max), fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			assertOk(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := time.Date(2018, 2, 4, 0, 0, 0, 0, time.Local)
			assertOk(t, fun(&v))
		})
	})
}

func TestBetween(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		msg := "must be in [%v, %v]"
		a := 10
		b := 20
		fun := rule.Between(a, b, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := a - 1
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := b + 1
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			assertOk(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			assertOk(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := a + 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Uint", func(t *testing.T) {
		msg := "must be in [%v, %v]"
		a := uint(10)
		b := uint(20)
		fun := rule.Between(a, b, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := a - 1
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := b + 1
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			assertOk(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			assertOk(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := a + 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Float", func(t *testing.T) {
		msg := "must be in [%v, %v]"
		a := 10.0
		b := 20.0
		fun := rule.Between(a, b, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := a - 1
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := b + 1
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			assertOk(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			assertOk(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := a + 1
			assertOk(t, fun(&v))
		})
	})
	t.Run("Time", func(t *testing.T) {
		msg := "must be in [%v, %v]"
		a := time.Date(2018, 1, 10, 0, 0, 0, 0, time.Local)
		b := time.Date(2018, 1, 20, 0, 0, 0, 0, time.Local)
		fun := rule.Between(a, b, msg)

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertInternalError(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertInternalError(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := time.Date(2018, 1, 9, 0, 0, 0, 0, time.Local)
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := time.Date(2018, 1, 21, 0, 0, 0, 0, time.Local)
			assertError(t, fmt.Errorf(msg, a, b), fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			assertOk(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			assertOk(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := time.Date(2018, 1, 11, 0, 0, 0, 0, time.Local)
			assertOk(t, fun(&v))
		})
	})
}
