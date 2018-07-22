package rule_test

import (
	"testing"

	"github.com/vbogretsov/go-validation"
	"github.com/vbogretsov/go-validation/rule"
)

const (
	eBlank     = "ErrEmailBlank"
	eEmail     = "ErrEmailInvalid"
	eMinLen    = "ErrMinLen"
	eDuplicate = "ErrDuplicated"
	minLen     = 5
)

type User struct {
	Email    string
	Password string
}

func assertPanic(t *testing.T, err error) {
	if err == nil {
		t.Error("expected validation.Panic but got nil")
	} else if _, ok := err.(validation.Panic); !ok {
		t.Errorf("expected validation.Panic but got %v", err)
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
		Errors: []error{
			validation.StructError{
				Field: "Email",
				Errors: []error{
					validation.Error{Message: eBlank},
					validation.Error{Message: eEmail},
				},
			},
			validation.StructError{
				Field: "Password",
				Errors: []error{
					validation.Error{Message: eMinLen, Params: validation.Params{
						rule.ParamStrMinLen: minLen,
					}},
				},
			},
		},
	},
	validation.SliceError{
		Index: 1,
		Errors: []error{
			validation.StructError{
				Field:  "Email",
				Errors: []error{validation.Error{Message: eEmail}},
			},
			validation.StructError{
				Field: "Password",
				Errors: []error{
					validation.Error{Message: eMinLen, Params: validation.Params{
						rule.ParamStrMinLen: minLen,
					}},
				},
			},
		},
	},
	validation.SliceError{
		Index: 2,
		Errors: []error{
			validation.StructError{
				Field:  "Email",
				Errors: []error{validation.Error{Message: eEmail}},
			},
		},
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
		Errors: []error{validation.Error{Message: eDuplicate}},
	},
	validation.SliceError{
		Index:  5,
		Errors: []error{validation.Error{Message: eDuplicate}},
	},
})
