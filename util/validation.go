package util

import (
	"fmt"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func Validation(w http.ResponseWriter, data interface{}) (errs []string) {
	validate := validator.New()

	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)
	if err != nil {
		validatorErrs := err.(validator.ValidationErrors)

		for _, e := range validatorErrs {
			translatedErr := fmt.Errorf(e.Translate(trans))
			errs = append(errs, translatedErr.Error())
		}
	}

	return errs
}
