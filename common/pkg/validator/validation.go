package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
)

// validData validation data
type validData struct {
	Tag        string
	ValidFunc  validator.Func
	EvenIfNull bool // callValidationEvenIfNull
}

// InitValidation register validation
func InitValidation(gValidator *validator.Validate) error {
	for _, vData := range validations {
		// register validation
		_ = RegisterValidation(
			gValidator,
			vData.Tag,
			vData.ValidFunc,
			vData.EvenIfNull,
		)
	}

	return nil
}

// ValidPortFn validate port arrange
func ValidPortFn(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	if value < 1 || value > 65535 {
		return false
	}
	return true
}

// ValidGtFieldFn gt "field" value
func ValidGtFieldFn(fl validator.FieldLevel) bool {
	field := fl.Field()

	currentField, _, ok, _ := fl.GetStructFieldOK2()
	if !ok {
		return false
	}

	// gt
	fieldVal := cast.ToFloat64(field.String())
	currentFieldVal := cast.ToFloat64(currentField.String())
	return fieldVal >= currentFieldVal
}

// ValidIsUpperCaseFn validate "field" is upper case
func ValidIsUpperCaseFn(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()                // 获取字段值
	return fieldValue == strings.ToUpper(fieldValue) // 检查是否全大写
}
