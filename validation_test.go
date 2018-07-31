package validation_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kr/pretty"

	"github.com/vbogretsov/go-validation"
)

// TODO: use require package for assertions.

func checkFieldName(tag, field string) error {
	validate, _ := validation.Struct(&Address{}, tag, errorFields)

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
	err := validate(v)
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
		_, err := validation.Struct(Address{}, ``, []validation.Field{})
		require.NotNil(t, err, "expected error but got nil")
	})
	t.Run("ErrorIfNotStruct", func(t *testing.T) {
		_, err := validation.Struct(&[]int{}, ``, []validation.Field{})
		require.NotNil(t, err, "expected error but got nil")
	})
}

func TestStructFieldName(t *testing.T) {
	t.Run("FieldNameIfTagEmpty", func(t *testing.T) {
		if err := checkFieldName("", "Country"); err != nil {
			t.Error(err)
		}
	})
	t.Run("TagIfTagNotEmpty", func(t *testing.T) {
		if err := checkFieldName("json", "country"); err != nil {
			t.Error(err)
		}
	})
	t.Run("FieldNameIfTagNotFound", func(t *testing.T) {
		if err := checkFieldName("xxx", "Country"); err != nil {
			t.Error(err)
		}
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
			err := addressRule(&k)
			if !reflect.DeepEqual(err, v) {
				t.Error(pretty.Diff(err, v))
			}
		})
	}

	for k, v := range userFixtures {
		t.Run("ValidateUser", func(t *testing.T) {
			err := userRule(&k)
			if !reflect.DeepEqual(err, v) {
				t.Error(pretty.Diff(err, v))
			}
		})
	}
}

func TestValidateStructWithCtx(t *testing.T) {
	for u, e := range userCtxFixtures {
		t.Run("ValidateCtx", func(t *testing.T) {
			err := userCtxRule(&u)
			require.Equal(t, e, err)
		})
	}
}

func TestRulesValidate(t *testing.T) {
	fun := validation.Rules([]validation.Rule{stringRequired, email})

	t.Run("TestPanicIfRulePanics", func(t *testing.T) {
		exp := validation.Panic{Err: eEpxectedStrPtr}
		act := fun("123")
		if !reflect.DeepEqual(exp, act) {
			t.Error(pretty.Diff(exp, act))
		}
	})
	t.Run("TestErrorIfRuleFails", func(t *testing.T) {
		exp := validation.Errors([]error{errors.New(eEmail)})
		v := "user"
		act := fun(&v)
		if !reflect.DeepEqual(exp, act) {
			t.Error(pretty.Diff(exp, act))
		}
	})
	t.Run("TestOkIfRulesOk", func(t *testing.T) {
		v := "user@mail.com"
		if err := fun(&v); err != nil {
			t.Errorf("expected error but got nil: %v", err)
		}
	})
}
