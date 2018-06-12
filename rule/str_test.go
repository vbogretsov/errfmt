package rule_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/vbogretsov/go-validation/rule"
)

func TestStrRequired(t *testing.T) {
	msg := "cannot be nlank"
	fun := rule.StrRequired(msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertInternalError(t, fun(10))
	})
	t.Run("ErrorIfEmpty", func(t *testing.T) {
		v := ""
		assertError(t, errors.New(msg), fun(&v))
	})
	t.Run("OkIfLenEqMax", func(t *testing.T) {
		v := "123"
		assertOk(t, fun(&v))
	})
}

func TestStrLen(t *testing.T) {
	min := 2
	max := 8
	msg := "len must be in range [%d, %d]"
	fun := rule.StrLen(min, max, msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertInternalError(t, fun(10))
	})
	t.Run("ErrorIfMin", func(t *testing.T) {
		v := ""
		assertError(t, fmt.Errorf(msg, min, max), fun(&v))
	})
	t.Run("ErrorIfMax", func(t *testing.T) {
		v := "123456789"
		assertError(t, fmt.Errorf(msg, min, max), fun(&v))
	})
	t.Run("OkIfInRange", func(t *testing.T) {
		v := "1234"
		assertOk(t, fun(&v))
	})
	t.Run("OkIfLenEqMin", func(t *testing.T) {
		v := "12"
		assertOk(t, fun(&v))
	})
	t.Run("OkIfLenEqMax", func(t *testing.T) {
		v := "12345678"
		assertOk(t, fun(&v))
	})
}

func TestStrMinLen(t *testing.T) {
	min := 2
	msg := "len must be not less than %d"
	fun := rule.StrMinLen(min, msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertInternalError(t, fun(10))
	})
	t.Run("ErrorIfMin", func(t *testing.T) {
		v := ""
		assertError(t, fmt.Errorf(msg, min), fun(&v))
	})
	t.Run("OkIfLenGt", func(t *testing.T) {
		v := "1234"
		assertOk(t, fun(&v))
	})
	t.Run("OkIfLenEq", func(t *testing.T) {
		v := "12"
		assertOk(t, fun(&v))
	})
}

func TestStrMaxLen(t *testing.T) {
	max := 8
	msg := "len must be not great than %d"
	fun := rule.StrMaxLen(max, msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertInternalError(t, fun(10))
	})
	t.Run("ErrorIfMax", func(t *testing.T) {
		v := "123456789"
		assertError(t, fmt.Errorf(msg, max), fun(&v))
	})
	t.Run("OkIfLenLt", func(t *testing.T) {
		v := "1234567"
		assertOk(t, fun(&v))
	})
	t.Run("OkIfLenEq", func(t *testing.T) {
		v := "12345678"
		assertOk(t, fun(&v))
	})
}

func TestStrMatch(t *testing.T) {
	msg := "does not match the pattern"
	fun := rule.StrMatch(regexp.MustCompile(`\d+`), msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertInternalError(t, fun(10))
	})
	t.Run("ErrorIfNotMatch", func(t *testing.T) {
		v := "abcd"
		assertError(t, errors.New(msg), fun(&v))
	})
	t.Run("OkIfMatch", func(t *testing.T) {
		v := "1234"
		assertOk(t, fun(&v))
	})
}
