package rule_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vbogretsov/go-validation"
	"github.com/vbogretsov/go-validation/rule"
)

func TestMin(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		msg := "ErrMin"
		min := 10
		fun := rule.Min(min, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: min,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := min - 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := min + 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Uint", func(t *testing.T) {
		msg := "ErrMin"
		min := uint(10)
		fun := rule.Min(min, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: min,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := min - 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := min + 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Float", func(t *testing.T) {
		msg := "ErrMin"
		min := 10.0
		fun := rule.Min(min, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: min,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := min - 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := min + 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Time", func(t *testing.T) {
		msg := "ErrMin"
		min := time.Date(2018, 2, 9, 0, 0, 0, 0, time.Local)
		fun := rule.Min(min, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: min,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := time.Date(2018, 2, 8, 0, 0, 0, 0, time.Local)
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := min
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfGt", func(t *testing.T) {
			v := time.Date(2018, 2, 10, 0, 0, 0, 0, time.Local)
			require.Nil(t, fun(&v))
		})
	})
}

func TestMax(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		msg := "ErrMax"
		max := 10
		fun := rule.Max(max, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMax: max,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := max + 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := max - 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Uint", func(t *testing.T) {
		msg := "ErrMax"
		max := uint(10)
		fun := rule.Max(max, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMax: max,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := max + 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := max - 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Float", func(t *testing.T) {
		msg := "ErrMax"
		max := 10.0
		fun := rule.Max(max, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMax: max,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := max + 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := max - 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Time", func(t *testing.T) {
		msg := "ErrMax"
		max := time.Date(2018, 2, 5, 0, 0, 0, 0, time.Local)
		fun := rule.Max(max, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMax: max,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := time.Date(2018, 2, 6, 0, 0, 0, 0, time.Local)
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEq", func(t *testing.T) {
			v := max
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfLt", func(t *testing.T) {
			v := time.Date(2018, 2, 4, 0, 0, 0, 0, time.Local)
			require.Nil(t, fun(&v))
		})
	})
}

func TestBetween(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		msg := "ErrBetween"
		a := 10
		b := 20
		fun := rule.Between(a, b, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: a,
				rule.ParamNumMax: b,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := a - 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := b + 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := a + 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Uint", func(t *testing.T) {
		msg := "ErrBetween"
		a := uint(10)
		b := uint(20)
		fun := rule.Between(a, b, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: a,
				rule.ParamNumMax: b,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := a - 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := b + 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := a + 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Float", func(t *testing.T) {
		msg := "ErrBetween"
		a := 10.0
		b := 20.0
		fun := rule.Between(a, b, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: a,
				rule.ParamNumMax: b,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := a - 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := b + 1
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := a + 1
			require.Nil(t, fun(&v))
		})
	})
	t.Run("Time", func(t *testing.T) {
		msg := "ErrBetween"
		a := time.Date(2018, 1, 10, 0, 0, 0, 0, time.Local)
		b := time.Date(2018, 1, 20, 0, 0, 0, 0, time.Local)
		fun := rule.Between(a, b, msg)(nil)
		exp := validation.Error{
			Message: msg,
			Params: validation.Params{
				rule.ParamNumMin: a,
				rule.ParamNumMax: b,
			},
		}

		t.Run("PanicIfNotPtr", func(t *testing.T) {
			assertPanic(t, fun(10))
		})
		t.Run("PanicIfInvalidType", func(t *testing.T) {
			v := ""
			assertPanic(t, fun(&v))
		})
		t.Run("ErrorIfLt", func(t *testing.T) {
			v := time.Date(2018, 1, 9, 0, 0, 0, 0, time.Local)
			require.Equal(t, exp, fun(&v))
		})
		t.Run("ErrorIfGt", func(t *testing.T) {
			v := time.Date(2018, 1, 21, 0, 0, 0, 0, time.Local)
			require.Equal(t, exp, fun(&v))
		})
		t.Run("OkIfEqL", func(t *testing.T) {
			v := a
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfEqG", func(t *testing.T) {
			v := b
			require.Nil(t, fun(&v))
		})
		t.Run("OkIfIn", func(t *testing.T) {
			v := time.Date(2018, 1, 11, 0, 0, 0, 0, time.Local)
			require.Nil(t, fun(&v))
		})
	})
}
