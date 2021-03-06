package helper

import (
	"app/api/config"
	"app/api/model"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/tr"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Validator(c *echo.Context, requestRules interface{}) string {
	var (
		uni   *ut.UniversalTranslator
		trans ut.Translator
	)

	validate := validator.New()
	eng := en.New()
	uni = ut.New(eng, eng)
	trans, _ = uni.GetTranslator("tr")
	_ = tr.RegisterDefaultTranslations(validate, trans)

	_ = (*c).Bind(requestRules)
	err := validate.Struct(requestRules)
	if err == nil {
		return ""
	}

	translateErrors := err.(validator.ValidationErrors).Translate(trans)
	translateErrorsString := ""
	counter := 0
	for key := range translateErrors {
		counter++
		translateErrorsString += convertFieldNames("Password", translateErrors[key])
		if counter == len(translateErrors) {
			translateErrorsString += "."
			break
		}
		translateErrorsString += ", "
	}
	return translateErrorsString

}

func convertFieldNames(field string, error string) string {
	for key, value := range map[string]string{
		"Email":    "E-posta",
		"Password": "Şifre",
	} {
		if key == field {
			return strings.Replace(error, key, value, 1)
		}
	}
	return field
}

func ValidateWithoutContext(data interface{}) (string, error) {
	var (
		uni   *ut.UniversalTranslator
		trans ut.Translator
	)

	validate := validator.New()
	eng := en.New()
	uni = ut.New(eng, eng)
	trans, _ = uni.GetTranslator("tr")
	_ = tr.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)
	if err == nil {
		return "", nil
	}

	translateErrors := err.(validator.ValidationErrors).Translate(trans)
	translateErrorsString := ""
	counter := 0
	for key := range translateErrors {
		counter++
		translateErrorsString += convertFieldNames("Password", translateErrors[key])
		if counter == len(translateErrors) {
			translateErrorsString += "."
			break
		}
		translateErrorsString += ", "
	}

	return translateErrorsString, err
}

func AuthId(c *echo.Context) uint32 {
	user := Auth(c)
	return uint32(user.ID)
}

func Auth(c *echo.Context) model.User {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JwtCustom)
	return claims.User
}
