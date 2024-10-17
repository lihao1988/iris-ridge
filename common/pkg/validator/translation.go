package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// translData translation data
type translData struct {
	Tag      string
	Msg      string
	TranslFn validator.TranslationFunc
}

// InitTranslation register translation
func InitTranslation(gValidator *validator.Validate, gVTranslator ut.Translator) error {
	var translFn validator.TranslationFunc
	for _, tData := range translations {
		// set default translation func
		translFn = translationFn
		if tData.TranslFn != nil {
			translFn = tData.TranslFn
		}

		// register translation
		_ = RegisterTranslation(
			gValidator,
			tData.Tag,
			gVTranslator,
			RegisterTranslationsFunc(tData.Tag, tData.Msg),
			translFn,
		)
	}

	return nil
}

// translationFn translate field value
func translationFn(ut ut.Translator, fe validator.FieldError) string {
	msg, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
	if err != nil {
		panic(fe.(error).Error())
	}

	return msg
}

// translFiledTestFn translate field value (demo test)
func translFiledTestFn(ut ut.Translator, fe validator.FieldError) string {
	msg, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
	if err != nil {
		panic(fe.(error).Error())
	}

	return msg
}
