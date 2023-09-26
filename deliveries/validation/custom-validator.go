package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/thedevsaddam/govalidator"
)

const CustomChar string = "^[-a-zA-Z0-9., ]+$"

var regexChar = regexp.MustCompile(CustomChar)

func isCustom(str string) bool {
	return regexChar.MatchString(str)
}

func toString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		str = fmt.Sprintf("%v", v)
	}
	return str
}

func AddCustomValidator() {
	govalidator.AddCustomRule("custom_char", func(field string, rule string, message string, value interface{}) error {
		str := toString(value)
		if str == "" {
			return nil
		}

		err := fmt.Errorf("the %s field must be a contains alpha numeric, space, dot or comma", field)
		if message != "" {
			err = errors.New(message)
		}

		if !isCustom(str) {
			return err
		}

		return nil
	})
}
