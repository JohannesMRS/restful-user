package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field) // field kosong
			case "email":
				errorsMap[field] = "Invalid email format" // format email salah
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exist", field) // data sudah ada
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param()) // nilai tidak cukup / kurang
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be most %s characters", field, fieldError.Param()) // nilai lebih
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field) // nilai harus angka
			default:
				errorsMap[field] = "Invalid value" // untuk validasi error lainnya
			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate Entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exist"
			}

			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exist"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorsMap["Error"] = "Record not found"
		}
	}
	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate Entry")
}
