package rule_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/vbogretsov/go-validation/rule"
)

func TestNotNil(t *testing.T) {
	msg := "cannot be nil"
	fun := rule.NotNil(msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertInternalError(t, fun(10))
	})
	t.Run("PanicIfValueCannotNil", func(t *testing.T) {
		v := 10
		assertInternalError(t, fun(&v))
	})
	t.Run("ErrorIfNil", func(t *testing.T) {
		var n interface{} = nil
		assertError(t, fun(&n), errors.New(msg))
	})
	t.Run("OkIfNotNil", func(t *testing.T) {
		v := []int{}
		assertOk(t, fun(&v))
	})
}

func TestIn(t *testing.T) {
	msg := "unsupported value %v, allowed values: %v"
	set := []interface{}{"a1", "b2", "c3"}
	fun := rule.In(set, msg)

	t.Run("ErrorIfNotIn", func(t *testing.T) {
		val := "d4"
		exp := fmt.Errorf(msg, val, set)
		assertError(t, exp, fun(val))
	})
	t.Run("OkIfIn", func(t *testing.T) {
		assertOk(t, fun("b2"))
	})
}
