package indoor

import (
	"ridge/common/global"
	"ridge/common/pkg/validator"
)

// AutoloadValidator autoload validator
func AutoloadValidator() error {
	var err error

	// new validator and translator
	global.GValidator, global.GVTranslator, err = validator.InitValidator()
	if err != nil {
		return nil
	}

	// validation register
	err = validator.InitValidation(global.GValidator)
	if err != nil {
		return err
	}

	// translation register
	err = validator.InitTranslation(global.GValidator, global.GVTranslator)
	if err != nil {
		return err
	}

	return nil
}
