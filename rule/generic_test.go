package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vbogretsov/go-validation"
	"github.com/vbogretsov/go-validation/rule"
)

func TestNotNil(t *testing.T) {
	msg := "ErrNil"
	fun := rule.NotNil(msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("PanicIfValueCannotNil", func(t *testing.T) {
		v := 10
		assertPanic(t, fun(&v))
	})
	t.Run("ErrorIfNil", func(t *testing.T) {
		var n interface{} = nil
		require.Equal(t, validation.Error{Message: msg}, fun(&n))

	})
	t.Run("OkIfNotNil", func(t *testing.T) {
		v := []int{}
		require.Nil(t, fun(&v))
	})
}

func TestIn(t *testing.T) {
	msg := "ErrUnsupported"
	set := []interface{}{"a1", "b2", "c3"}
	fun := rule.In(set, msg)

	t.Run("PanicIfInvalidType", func(t *testing.T) {
		assertPanic(t, fun("xxx"))
	})
	t.Run("ErrorIfNotIn", func(t *testing.T) {
		val := "d4"
		exp := validation.Error{Message: msg, Params: validation.Params{
			rule.ParamInUnsupported: val,
			rule.ParamInSupported:   set,
		}}

		require.Equal(t, exp, fun(&val))
	})
	t.Run("OkIfIn", func(t *testing.T) {
		val := "b2"
		require.Nil(t, fun(&val))
	})
}
