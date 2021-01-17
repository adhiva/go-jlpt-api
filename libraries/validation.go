package libraries

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func RequestValidation(iface interface{}) (bool, []string) {
	var (
		err           error
		errorMessages []string
	)

	// Define Validation
	validation := validator.New()

	// Define locale
	en := en.New()
	uni := ut.New(en, en)

	// Set translator locale validation to your language
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validation, trans)

	err = validation.Struct(iface)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, errMsg := range errors {
			errorMessages = append(errorMessages, errMsg.Translate(trans))
		}
		return false, errorMessages
	}

	return true, nil
}
