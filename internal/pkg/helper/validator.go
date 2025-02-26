package helper

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
	"unicode"
)

type CustomValidator struct {
	Val *validator.Validate
}

var (
	GValidator CustomValidator
	enLang     = en.New()
	uni        = ut.New(enLang, enLang)
)

func InitValidator() {
	v := validator.New()
	trans, _ := getTranslator("en")

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	v.RegisterValidation("user_password", validateUserPassword)

	en_translations.RegisterDefaultTranslations(v, trans)

	GValidator = CustomValidator{Val: v}
}

func getTranslator(lang string) (ut.Translator, error) {
	trans, _ := uni.GetTranslator(lang)
	return trans, nil
}

func validateUserPassword(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	hasSpecialChar := false
	hasNumber := false
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			hasNumber = true
		}
		if strings.ContainsRune("!@#$%^&*()-_=+{}[]|;:,.<>?", ch) {
			hasSpecialChar = true
		}
		if hasNumber && hasSpecialChar {
			return true
		}
	}

	return false
}
