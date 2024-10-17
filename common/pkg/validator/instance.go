package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translZh "github.com/go-playground/validator/v10/translations/zh"
)

// InitValidator init validator instance
func InitValidator() (*validator.Validate, ut.Translator, error) {
	// create "zh" translator
	zhTrans := zh.New()
	uni := ut.New(zhTrans, zhTrans)

	// get "zh" translator
	zhTransl, _ := uni.GetTranslator("zh")

	// validator instance
	var err error
	validate := validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// register default translations
	err = translZh.RegisterDefaultTranslations(validate, zhTransl)
	if err != nil {
		return nil, zhTransl, err
	}

	return validate, zhTransl, nil
}

// ValidStruct validate struct (div validator)
func ValidStruct(v *validator.Validate, t ut.Translator, s interface{}) error {
	return GetValidError(v.Struct(s), t)
}

// ValidVariable validate variable (div validator)
func ValidVariable(v *validator.Validate, t ut.Translator, field interface{}, tag string) error {
	return GetValidError(v.Var(field, tag), t)
}

// RegisterValidation register validation (div validator)
func RegisterValidation(v *validator.Validate, tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

// RegisterTranslation register translation (div validator)
func RegisterTranslation(v *validator.Validate, tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) (err error) {
	return v.RegisterTranslation(tag, trans, registerFn, translationFn)
}

// RegisterTranslationsFunc create "RegisterTranslationsFunc" func
func RegisterTranslationsFunc(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, true); err != nil {
			return err
		}

		return nil
	}
}

// GetValidError get the validator error
func GetValidError(errs error, t ut.Translator) error {
	var vErrs validator.ValidationErrors
	ok := errors.As(errs, &vErrs)
	if ok {
		for _, errInfo := range vErrs.Translate(t) {
			return errors.New(errInfo)
		}
	}

	return errs
}
