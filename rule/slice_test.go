package rule_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vbogretsov/go-validation"
	"github.com/vbogretsov/go-validation/rule"
)

func TestSliceLen(t *testing.T) {
	min := 2
	max := 8
	msg := "ErrLen"
	fun := rule.SliceLen(min, max, msg)
	exp := validation.Error{
		Message: msg,
		Params: validation.Params{
			rule.ParamSliceMinLen: min,
			rule.ParamSliceMaxLen: 8,
		},
	}

	t.Run("PanicIfNotPtr", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("PanicIfInvalidType", func(t *testing.T) {
		v := 10
		assertPanic(t, fun(&v))
	})
	t.Run("ErrorIfMin", func(t *testing.T) {
		v := []int{}
		require.Equal(t, exp, fun(&v))

	})
	t.Run("ErrorIfMax", func(t *testing.T) {
		v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfInRange", func(t *testing.T) {
		v := []int{1, 2, 3, 4}
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEqMin", func(t *testing.T) {
		v := []int{1, 2}
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEqMax", func(t *testing.T) {
		v := []int{1, 2, 3, 4, 5, 6, 7, 8}
		require.Nil(t, fun(&v))
	})
}

func TestSliceMinLen(t *testing.T) {
	min := 2
	msg := "ErrMinLen"
	fun := rule.SliceMinLen(min, msg)
	exp := validation.Error{
		Message: msg,
		Params: validation.Params{
			rule.ParamSliceMinLen: min,
		},
	}

	t.Run("PanicIfNotPtr", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("PanicIfInvalidType", func(t *testing.T) {
		v := 10
		assertPanic(t, fun(&v))
	})
	t.Run("ErrorIfMin", func(t *testing.T) {
		v := []string{}
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfLenGt", func(t *testing.T) {
		v := []string{"1", "2", "3", "4"}
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEq", func(t *testing.T) {
		v := []string{"1", "2"}
		require.Nil(t, fun(&v))
	})
}

func TestSliceMaxLen(t *testing.T) {
	max := 8
	msg := "len must be not great than %d"
	fun := rule.SliceMaxLen(max, msg)
	exp := validation.Error{
		Message: msg,
		Params: validation.Params{
			rule.ParamSliceMaxLen: max,
		},
	}

	t.Run("PanicIfNotPtr", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("PanicIfInvalidType", func(t *testing.T) {
		v := 10
		assertPanic(t, fun(&v))
	})
	t.Run("ErrorIfMax", func(t *testing.T) {
		v := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}
		require.Equal(t, exp, fun(&v))
	})
	t.Run("OkIfLenLt", func(t *testing.T) {
		v := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0}
		require.Nil(t, fun(&v))
	})
	t.Run("OkIfLenEq", func(t *testing.T) {
		v := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
		require.Nil(t, fun(&v))
	})
}

func TestSliceEach(t *testing.T) {
	t.Run("PanicIfValidatorPanics", func(t *testing.T) {
		failed := rule.SliceEach(userIter, []validation.Rule{
			func(v interface{}) error {
				return validation.Panic{Err: errors.New("test")}
			},
		})
		v := []User{{}}
		assertPanic(t, failed(&v))
	})

	fun := rule.SliceEach(userIter, []validation.Rule{userRule})

	t.Run("PanicIfNotPtr", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("PanicIfInvalidType", func(t *testing.T) {
		v := 10
		assertPanic(t, fun(&v))
	})
	t.Run("ErrorIfErrors", func(t *testing.T) {
		require.Equal(t, invalidUserErrors, fun(&invalidUsers))
	})
	t.Run("OkIfNoErorrs", func(t *testing.T) {
		require.Nil(t, fun(&users))
	})
}

func TestSliceUnique(t *testing.T) {
	t.Run("PanicIfItemNotPtr", func(t *testing.T) {
		r := rule.SliceUnique(userInvalidIter, eDuplicate)
		assertPanic(t, r(&users))
	})

	fun := rule.SliceUnique(userIter, eDuplicate)

	t.Run("PanicIfNotPtr", func(t *testing.T) {
		assertPanic(t, fun(10))
	})
	t.Run("PanicIfInvalidType", func(t *testing.T) {
		v := 10
		assertPanic(t, fun(&v))
	})
	t.Run("ErrorIfErorrs", func(t *testing.T) {
		require.Equal(t, duplicatedUserErrors, fun(&duplicatedUsers))
	})
	t.Run("OkIfNoDuplicates", func(t *testing.T) {
		require.Nil(t, fun(&users))
	})
}
