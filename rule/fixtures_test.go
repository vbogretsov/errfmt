package rule_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"github.com/vbogretsov/go-validation"
	"github.com/vbogretsov/go-validation/rule"
)

const (
	eBlank     = "cannot be blank"
	eEmail     = "invalid email"
	eMinLen    = "cannot be shorther than %d"
	eDuplicate = "duplicated item"
	minLen     = 5
)

type User struct {
	Email    string
	Password string
}

func assertInternalError(t *testing.T, err error) {
	if err == nil {
		t.Error("expected validation.Panic but got nil")
	} else if _, ok := err.(validation.Panic); !ok {
		t.Errorf("expected validation.Panic but got %v", err)
	}
}

func assertError(t *testing.T, exp error, act error) {
	if !reflect.DeepEqual(exp, act) {
		t.Error(pretty.Diff(exp, act))
	}
}

func assertOk(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected nil but got error: %v", err)
	}
}

var userRule, _ = validation.Struct(&User{}, "", []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*User).Email
		},
		Rules: []validation.Rule{
			rule.StrRequired(eBlank),
			rule.StrEmail(eEmail),
		},
	},
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*User).Password
		},
		Rules: []validation.Rule{
			rule.StrMinLen(minLen, eMinLen),
		},
	},
})

func userIter(v interface{}, i int) interface{} {
	return &(*(v.(*[]User)))[i]
}

func userInvalidIter(v interface{}, i int) interface{} {
	return (*(v.(*[]User)))[i]
}

var users = []User{
	{
		Email:    "user1@mail.com",
		Password: "123456",
	},
	{
		Email:    "user2@mail.com",
		Password: "123456",
	},
	{
		Email:    "user3@mail.com",
		Password: "123456",
	},
}

var invalidUsers = []User{
	{
		Email:    "",
		Password: "",
	},
	{
		Email:    "user",
		Password: "1234",
	},
	{
		Email:    "user",
		Password: "123456",
	},
	{
		Email:    "user@mail.com",
		Password: "123456",
	},
}

var invalidUserErrors = validation.Errors([]error{
	validation.SliceError{
		Index: 0,
		Errors: validation.Errors([]error{
			validation.StructError{
				Field: "Email",
				Errors: validation.Errors([]error{
					errors.New(eBlank),
					errors.New(eEmail),
				}),
			},
			validation.StructError{
				Field: "Password",
				Errors: validation.Errors([]error{
					fmt.Errorf(eMinLen, minLen),
				}),
			},
		}),
	},
	validation.SliceError{
		Index: 1,
		Errors: validation.Errors([]error{
			validation.StructError{
				Field:  "Email",
				Errors: []error{errors.New(eEmail)},
			},
			validation.StructError{
				Field:  "Password",
				Errors: []error{fmt.Errorf(eMinLen, minLen)},
			},
		}),
	},
	validation.SliceError{
		Index: 2,
		Errors: validation.Errors([]error{
			validation.StructError{
				Field:  "Email",
				Errors: []error{errors.New(eEmail)},
			},
		}),
	},
})

var duplicatedUsers = []User{
	{
		Email:    "user0@mail.com",
		Password: "123456",
	},
	{
		Email:    "user1@mail.com",
		Password: "123456",
	},
	{
		Email:    "user2@mail.com",
		Password: "123456",
	},
	{
		Email:    "user1@mail.com",
		Password: "123456",
	},
	{
		Email:    "user3@mail.com",
		Password: "123456",
	},
	{
		Email:    "user0@mail.com",
		Password: "123456",
	},
}

var duplicatedUserErrors = validation.Errors([]error{
	validation.SliceError{
		Index:  3,
		Errors: validation.Errors([]error{errors.New(eDuplicate)}),
	},
	validation.SliceError{
		Index:  5,
		Errors: validation.Errors([]error{errors.New(eDuplicate)}),
	},
})
