package rule_test

import (
	"errors"
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
