package validation_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vbogretsov/go-validation"
)

func checkFieldName(tag, field string) error {
	validate := validation.Struct(&Address{}, tag, errorFields)(nil)

	err := validate(&Address{})
	if err == nil {
		return errors.New("expected error but got nil")
	}

	e := err.(validation.Errors)[0].(validation.StructError)
	if e.Field != field {
		return fmt.Errorf("expected field name `%s` but git `%s`", field, e.Field)
	}

	return nil
}

func checkValidatePanics(validate validation.Rule, v interface{}) error {
	err := validate(nil)(v)
	if err == nil {
		return errors.New("expected error but got nil")
	}
	if _, ok := err.(validation.Panic); !ok {
		return fmt.Errorf("expected validation.Panic but got %v", err)
	}
	return nil
}

func TestStruct(t *testing.T) {
	t.Run("ErrorIfNotPtr", func(t *testing.T) {
		v := validation.Struct(Address{}, ``, []validation.Field{})(nil)
		x := 10
		require.Error(t, v(&x))
	})
	t.Run("ErrorIfNotStruct", func(t *testing.T) {
		v := validation.Struct(&[]int{}, ``, []validation.Field{})(nil)
		x := 10
		require.Error(t, v(&x))
	})
}

func TestStructFieldName(t *testing.T) {
	t.Run("FieldNameIfTagEmpty", func(t *testing.T) {
		require.NoError(t, checkFieldName("", "Country"))
	})
	t.Run("TagIfTagNotEmpty", func(t *testing.T) {
		require.NoError(t, checkFieldName("json", "country"))
	})
	t.Run("FieldNameIfTagNotFound", func(t *testing.T) {
		require.NoError(t, checkFieldName("xxx", "Country"))
	})
}

func TestStructValidate(t *testing.T) {
	t.Run("PanicIfNotPtr", func(t *testing.T) {
		if err := checkValidatePanics(addressRule, Address{}); err != nil {
			t.Error(err)
		}
	})

	t.Run("PanicIfNotStruct", func(t *testing.T) {
		if err := checkValidatePanics(addressRule, &[]int{}); err != nil {
			t.Error(err)
		}
	})

	t.Run("PanicIfAttrNotPtr", func(t *testing.T) {
		if err := checkValidatePanics(addressRuleGetAttrNotPtr, &Address{}); err != nil {
			t.Error(err)
		}
	})

	t.Run("PanicIfAttrPanic", func(t *testing.T) {
		if err := checkValidatePanics(addressRuleValidatorPanic, &Address{}); err != nil {
			t.Error(err)
		}
	})

	for k, v := range addressFixtures {
		t.Run("ValidateAddress", func(t *testing.T) {
			err := addressRule(nil)(&k)
			require.Equal(t, v, err)
		})
	}

	for k, v := range userFixtures {
		t.Run("ValidateUser", func(t *testing.T) {
			err := userRule(&k)
			require.Equal(t, v, err)
		})
	}
}

func TestRulesValidate(t *testing.T) {
	fun := validation.Rules([]validation.Rule{
		validation.Func(stringRequired),
		validation.Func(email),
	})(nil)

	t.Run("TestPanicIfRulePanics", func(t *testing.T) {
		exp := validation.Panic{Err: eEpxectedStrPtr}
		act := fun("123")
		require.Equal(t, exp, act)
	})
	t.Run("TestErrorIfRuleFails", func(t *testing.T) {
		exp := validation.Errors([]error{errors.New(eEmail)})
		v := "user"
		act := fun(&v)
		require.Equal(t, exp, act)
	})
	t.Run("TestOkIfRulesOk", func(t *testing.T) {
		v := "user@mail.com"
		require.NoError(t, fun(&v))
	})
}

func TestCtxValidation(t *testing.T) {
	for u, e := range userRuleCtxFixtures {
		err := userRuleCtx(usersDB)(&u)
		require.Equal(t, e, err)
	}
}
