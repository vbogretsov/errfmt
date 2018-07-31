package rule_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vbogretsov/go-validation"
	"github.com/vbogretsov/go-validation/rule"
)

func TestStrRequired(t *testing.T) {
	msg := "ErrBlank"
	fun := rule.StrRequired(msg)(nil)
	exp := validation.Error{Message: msg}

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("ErrorIfEmpty", func(t *testing.T) {
		v := ""
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfLenEqMax", func(t *testing.T) {
		v := "123"
		require.Nil(t, fun(&v))
	})
}

func TestStrLen(t *testing.T) {
	min := 2
	max := 8
	msg := "ErrLen"
	fun := rule.StrLen(min, max, msg)(nil)
	exp := validation.Error{
		Message: msg,
		Params: validation.Params{
			rule.ParamStrMinLen: min,
			rule.ParamStrMaxLen: max,
		},
	}

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("ErrorIfMin", func(t *testing.T) {
		v := ""
		require.Equal(t, exp, fun(&v))
	})
	t.Run("ErrorIfMax", func(t *testing.T) {
		v := "123456789"
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfInRange", func(t *testing.T) {
		v := "1234"
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEqMin", func(t *testing.T) {
		v := "12"
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEqMax", func(t *testing.T) {
		v := "12345678"
		require.Nil(t, fun(&v))
	})
}

func TestStrMinLen(t *testing.T) {
	min := 2
	msg := "ErrMinLen"
	fun := rule.StrMinLen(min, msg)(nil)
	exp := validation.Error{
		Message: msg,
		Params: validation.Params{
			rule.ParamStrMinLen: min,
		},
	}

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("ErrorIfMin", func(t *testing.T) {
		v := ""
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfLenGt", func(t *testing.T) {
		v := "1234"
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEq", func(t *testing.T) {
		v := "12"
		require.Nil(t, fun(&v))
	})
}

func TestStrMaxLen(t *testing.T) {
	max := 8
	msg := "ErrMinLen"
	fun := rule.StrMaxLen(max, msg)(nil)
	exp := validation.Error{
		Message: msg,
		Params: validation.Params{
			rule.ParamStrMaxLen: max,
		},
	}

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("ErrorIfMax", func(t *testing.T) {
		v := "123456789"
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfLenLt", func(t *testing.T) {
		v := "1234567"
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEq", func(t *testing.T) {
		v := "12345678"
		require.Nil(t, fun(&v))
	})
}

func TestStrMatch(t *testing.T) {
	msg := "ErrMatch"
	fun := rule.StrMatch(regexp.MustCompile(`\d+`), msg)(nil)
	exp := validation.Error{Message: msg}

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("ErrorIfNotMatch", func(t *testing.T) {
		v := "abcd"
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfMatch", func(t *testing.T) {
		v := "1234"
		require.Nil(t, fun(&v))
	})
}
