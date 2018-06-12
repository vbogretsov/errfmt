package rule

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"

	"github.com/vbogretsov/go-validation"
)

var (
	// StrEmail creates validator to check whether a string is an email.
	StrEmail = fromfn(govalidator.IsEmail)
	// StrIPv4 creates validator to check whether a string is an IPv4.
	StrIPv4 = fromfn(govalidator.IsIPv4)
	// StrIPv6 creates validator to check whether a string is an IPv6.
	StrIPv6 = fromfn(govalidator.IsIPv6)
	// StrIP creates validator to check whether a string is an IP.
	StrIP = fromfn(govalidator.IsIP)
	// StrIsURL creates validator to check whether a string is an URL.
	StrIsURL = fromfn(govalidator.IsURL)
	// StrIsUpperCase creates validator to check whether a string is in upper case.
	StrIsUpperCase = fromfn(govalidator.IsUpperCase)
	// StrIsLowerCase creates validator to check whether a string is in lower case.
	StrIsLowerCase = fromfn(govalidator.IsLowerCase)
	// StrIsJSON creates validator to check whether a string is a JSON.
	StrIsJSON = fromfn(govalidator.IsJSON)
	// TODO(vbogretsov): import other string rules.
)

// TODO(vbogretsov): maybe avoid error formatting.

func strrule(fn func(*string) error) validation.Rule {
	return func(v interface{}) error {
		s, ok := v.(*string)
		if !ok {
			return unexpectedType(v)
		}
		return fn(s)
	}
}

func fromfn(fn func(string) bool) func(string) validation.Rule {
	return func(msg string) validation.Rule {
		return strrule(func(s *string) error {
			if !fn(*s) {
				return errors.New(msg)
			}
			return nil
		})
	}
}

// StrLen creates validator to check whether length of a string is in the range
// provided. The 'msg' parameter should be a format string with 2 slots for int.
func StrLen(min, max int, msg string) validation.Rule {
	return strrule(func(s *string) error {
		n := len(*s)
		if n < min || n > max {
			return fmt.Errorf(msg, min, max)
		}
		return nil
	})
}

// StrRequired creates validator to check whether a string is blank.
func StrRequired(msg string) validation.Rule {
	return strrule(func(s *string) error {
		if *s == "" {
			return errors.New(msg)
		}
		return nil
	})
}

// StrMinLen creates validator to check whether length of a string is not less
// than the value provided. The 'msg' parameter should be a format string with
// 1 slot for int.
func StrMinLen(min int, msg string) validation.Rule {
	return strrule(func(s *string) error {
		n := len(*s)
		if n < min {
			return fmt.Errorf(msg, min)
		}
		return nil
	})
}

// StrMaxLen creates validator to check whether length of a string is not great
// than the value provided. The 'msg' parameter should be a format string with
// 1 slot for int.
func StrMaxLen(max int, msg string) validation.Rule {
	emsg := fmt.Sprintf(msg, max)
	return strrule(func(s *string) error {
		n := len(*s)
		if max < n {
			return errors.New(emsg)
		}
		return nil
	})
}

// StrMatch creates validator to check whether a string matches the regular
// expression provided.
func StrMatch(pattern *regexp.Regexp, msg string) validation.Rule {
	return strrule(func(s *string) error {
		if !pattern.MatchString(*s) {
			return errors.New(msg)
		}
		return nil
	})
}
