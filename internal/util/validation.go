package util

import (
	"fmt"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
)

func Validate(req any) []string {
	var errs []string

	t := reflect.ValueOf(req)
	for i := range t.NumField() {
		field := t.Field(i)
		fieldName := t.Type().Field(i).Name

		tag := t.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}

		rules := strings.SplitSeq(tag, ",")
		for rule := range rules {
			err, hasError := applyRule(errs, rule, field.String(), fieldName)
			
			if hasError {
				errs = err
				break
			}
		}
	}

	return errs
}

func applyRule(errs []string, rule string, field string, fieldName string) ([]string, bool) {
	switch {
	case rule == "required":
		return validateRequired(errs, field, fieldName)
	case rule == "email":
		return validateEmail(errs, field)
	case strings.HasPrefix(rule, "min="):
		return validateMinLength(errs, rule, field, fieldName)
	default:
		return errs, false
	}	
}

func validateRequired(errs []string, field string, fieldName string) ([]string, bool) {
	if len(field) == 0 {
		return append(errs, fmt.Sprintf("%s is required!", fieldName)), true
	}

	return errs, false
}

func validateEmail(errs []string, field string) ([]string, bool) {
	if _, err := mail.ParseAddress(field); err != nil || !strings.HasSuffix(field, ".com") {
		return append(errs, "Email is invalid!"), true
	}

	return errs, false
}

func validateMinLength(errs []string, rule string, field string, fieldName string) ([]string, bool) {
	min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))				
	if len(field) < min {
		return append(errs, fmt.Sprintf("%s should be at least %d characters long!", fieldName, min)), true
	}

	return errs, false
}
