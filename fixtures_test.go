package validation_test

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/vbogretsov/go-validation"
)

const (
	eRequired        = "string cannot be blank"
	eStartsUpperCase = "should start with upper case"
	eEmail           = "invalid email"
	eDuplicate       = "email in use"
	eMinLen          = "should have at least %d characters"
	eZipCode         = "should contain only letters"
	ePwConfirmation  = "password confirmation does not math the password"
	minLen           = 10
)

var (
	eEpxectedStrPtr = errors.New("expected string pointer")
)

type Address struct {
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
}

type User struct {
	Email                string
	Password             string
	PasswordConfirmation string
	Address              Address
}

var (
	addressFixtures = map[Address]error{
		Address{}: validation.Errors([]error{
			validation.StructError{
				Field: "Country",
				Errors: []error{
					errors.New(eRequired),
					errors.New(eStartsUpperCase),
				},
			},
			validation.StructError{
				Field: "ZipCode",
				Errors: []error{
					errors.New(eRequired),
				},
			},
		}),
		Address{Country: "usa"}: validation.Errors([]error{
			validation.StructError{
				Field: "Country",
				Errors: []error{
					errors.New(eStartsUpperCase),
				},
			},
			validation.StructError{
				Field: "ZipCode",
				Errors: []error{
					errors.New(eRequired),
				},
			},
		}),
		Address{ZipCode: "aaa"}: validation.Errors([]error{
			validation.StructError{
				Field: "Country",
				Errors: []error{
					errors.New(eRequired),
					errors.New(eStartsUpperCase),
				},
			},
			validation.StructError{
				Field: "ZipCode",
				Errors: []error{
					errors.New(eZipCode),
				},
			},
		}),
		Address{Country: "Russia", ZipCode: "123"}: nil,
	}

	userFixtures = map[User]error{
		User{}: validation.Errors([]error{
			validation.StructError{
				Field: "Email",
				Errors: []error{
					errors.New(eRequired),
					errors.New(eEmail),
				},
			},
			validation.StructError{
				Field: "Password",
				Errors: []error{
					errors.New(eRequired),
					errors.New(fmt.Sprintf(eMinLen, minLen)),
				},
			},
			validation.StructError{
				Field: "Address",
				Errors: []error{

					validation.StructError{
						Field: "Country",
						Errors: []error{
							errors.New(eRequired),
							errors.New(eStartsUpperCase),
						},
					},
					validation.StructError{
						Field: "ZipCode",
						Errors: []error{
							errors.New(eRequired),
						},
					},
				},
			},
		}),
		User{
			Email:                "user",
			Password:             "123",
			PasswordConfirmation: "12",
			Address: Address{
				Country: "usa",
				ZipCode: "aaa",
			},
		}: validation.Errors([]error{
			validation.StructError{
				Field: "Email",
				Errors: []error{
					errors.New(eEmail),
				},
			},
			validation.StructError{
				Field: "Password",
				Errors: []error{
					errors.New(fmt.Sprintf(eMinLen, minLen)),
				},
			},
			validation.StructError{
				Field: "",
				Errors: []error{
					errors.New(ePwConfirmation),
				},
			},
			validation.StructError{
				Field: "Address",
				Errors: []error{
					validation.StructError{
						Field: "Country",
						Errors: []error{
							errors.New(eStartsUpperCase),
						},
					},
					validation.StructError{
						Field: "ZipCode",
						Errors: []error{
							errors.New(eZipCode),
						},
					},
				},
			},
		}),
	}
)

func stringRequired(v interface{}) error {
	s, ok := v.(*string)
	if !ok {
		return validation.Panic{Err: eEpxectedStrPtr}
	}
	if *s == "" {
		return errors.New(eRequired)
	}
	return nil
}

func startsUpperCase(v interface{}) error {
	s := []rune(*(v.(*string)))
	if len(s) == 0 || !unicode.IsUpper(s[0]) {
		return errors.New(eStartsUpperCase)
	}
	return nil
}

func email(v interface{}) error {
	s := v.(*string)
	if !strings.Contains(*s, "@") {
		return errors.New(eEmail)
	}
	return nil
}

func minlen(n int) func(v interface{}) error {
	return func(v interface{}) error {
		s := *(v.(*string))
		if len(s) < n {
			return fmt.Errorf(eMinLen, minLen)
		}
		return nil
	}
}

func zipCode(v interface{}) error {
	s := []rune(*(v.(*string)))
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return errors.New(eZipCode)
		}
	}
	return nil
}

var addressRule = validation.Struct(&Address{}, "", []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*Address).Country
		},
		Rules: []validation.Rule{
			validation.Func(stringRequired),
			validation.Func(startsUpperCase),
		},
	},
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*Address).ZipCode
		},
		Rules: []validation.Rule{
			validation.Func(stringRequired),
			validation.Func(zipCode),
		},
	},
})

var addressRuleGetAttrNotPtr = validation.Struct(&Address{}, ``, []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return v.(*Address).Country
		},
		Rules: []validation.Rule{
			validation.Func(stringRequired),
			validation.Func(email),
		},
	},
})

var addressRuleValidatorPanic = validation.Struct(&Address{}, ``, []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*Address).Country
		},
		Rules: []validation.Rule{
			validation.Func(stringRequired),
			validation.Func(func(v interface{}) error {
				return validation.Panic{Err: errors.New("test panic")}
			}),
		},
	},
})

var errorFields = []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*Address).Country
		},
		Rules: []validation.Rule{
			validation.Func(func(v interface{}) error {
				return errors.New("test error")
			}),
		},
	},
}

var userRule = validation.Struct(&User{}, ``, []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*User).Email
		},
		Rules: []validation.Rule{
			validation.Func(stringRequired),
			validation.Func(email),
		},
	},
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*User).Password
		},
		Rules: []validation.Rule{
			validation.Func(stringRequired),
			validation.Func(minlen(minLen)),
		},
	},
	{
		Attr: func(v interface{}) interface{} {
			return v
		},
		Rules: []validation.Rule{
			validation.Func(func(v interface{}) error {
				u := v.(*User)
				if u.Password != u.PasswordConfirmation {
					return errors.New(ePwConfirmation)
				}
				return nil
			}),
		},
	},
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*User).Address
		},
		Rules: []validation.Rule{validation.Func(addressRule(nil))},
	},
})(nil)

var usersDB = []string{"user1@mail.com", "user2@mail.com", "usermail.com"}

func emailUniq(ctx interface{}) func(interface{}) error {
	db := ctx.([]string)
	return func(v interface{}) error {
		s := v.(*string)
		fmt.Println("STR:", *s)
		fmt.Println("DB:", db)
		for _, u := range db {
			fmt.Println("ERROR")
			if u == *s {
				return errors.New(eDuplicate)
			} else {
				fmt.Printf("%s != %s\n", u, *s)
			}
		}
		fmt.Println("EXIT")
		return nil
	}
}

var userRuleCtx = validation.Struct(&User{}, ``, []validation.Field{
	{
		Attr: func(v interface{}) interface{} {
			return &v.(*User).Email
		},
		Rules: []validation.Rule{
			validation.Func(email),
			emailUniq,
		},
	},
})

var userRuleCtxFixtures = map[User]error{
	User{Email: "user1@mail.com"}: validation.Errors([]error{
		validation.StructError{
			Field: "Email",
			Errors: []error{
				errors.New(eDuplicate),
			},
		},
	}),
	User{Email: "user1mail.com"}: validation.Errors([]error{
		validation.StructError{
			Field: "Email",
			Errors: []error{
				errors.New(eEmail),
			},
		},
	}),
	User{Email: "usermail.com"}: validation.Errors([]error{
		validation.StructError{
			Field: "Email",
			Errors: []error{
				errors.New(eEmail),
				errors.New(eDuplicate),
			},
		},
	}),
	User{Email: "user3@mail.com"}: nil,
}
