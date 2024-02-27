package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationMsg(err error) []string {
	var errors []string
	if v, ok := err.(validator.ValidationErrors); ok {
		for _, e := range v {
			field := strings.ToLower(e.Field())
			message := fmt.Sprintf("Field '%s' failed validation for tag '%s %s'", field, e.Tag(), e.Param())
			errors = append(errors, message)
		}
	} else {
		errors = append(errors, err.Error())
	}

	return errors
}
