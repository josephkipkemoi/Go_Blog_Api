package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// validationErrors returns found errors stored in a slice and true if errors are found empty slice and false otherwise
func validationErrors(err error) ([]string, bool) {
	errs := []string{}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errs, false
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println()
			fmt.Println(err)
			errs = append(errs, err.Field()+" Field is required")
		}
		return errs, false
	}
	return errs, true
}

func comparePasswords(str1, str2 string) bool {
	if str1 == str2 {
		return true
	} else {
		return false
	}
}
