package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) map[string]string {
	var validationErr *validator.ValidationErrors
	var unmarshalErr *json.UnmarshalTypeError

	if  (errors.As(err, &unmarshalErr)) {
		
	}

	if (errors.As(err, &validationErr)) {
		fmt.Println("validation error")
		os.Exit(1)
	}

	errors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())

			switch fe.Tag() {
			case "required":
				errors[field] = "is required"
			case "email":
				errors[field] = "must be a valid email address"
			default:
				errors[field] = "is invalid"
			}
		}
	}

	return errors
}