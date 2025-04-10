package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	en       locales.Translator
	uniTrans *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
}

var val *Validator

const (
	walkingModeType = "walking"
	bicycleModeType = "bicycle"
)

func ValidateMode(fl validator.FieldLevel) bool {
	mode := strings.ToLower(fl.Field().String())
	return mode == walkingModeType || mode == bicycleModeType
}

func InitValidator() {
	val = &Validator{}
	val.en = en.New()
	val.uniTrans = ut.New(val.en, val.en)
	val.trans, _ = val.uniTrans.GetTranslator("en")
	val.validate = validator.New()
	en_translations.RegisterDefaultTranslations(val.validate, val.trans)
	val.validate.RegisterValidation("modeType", ValidateMode)
}

func Validate(v interface{}) error {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Slice {
		for i := 0; i < vValue.Len(); i++ {
			if err := val.validate.Struct(vValue.Index(i).Interface()); err != nil {
				return errors.New("validation error: " + err.Error())
			}
		}
	} else {
		if err := val.validate.Struct(v); err != nil {
			return errors.New("validation error: " + err.Error())
		}
	}
	return nil
}

func TranslateError(e error) validator.ValidationErrorsTranslations {
	valErr, ok := e.(validator.ValidationErrors)
	if !ok {
		return validator.ValidationErrorsTranslations{
			"error": e.Error(),
		}
	}
	return valErr.Translate(val.trans)
}