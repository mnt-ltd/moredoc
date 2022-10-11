package validate

import (
	"errors"
	"strings"

	localesZH "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	transZH "github.com/go-playground/validator/v10/translations/zh"
)

var (
	zh       = localesZH.New()
	uni      = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
)

func init() {
	transZH.RegisterDefaultTranslations(validate, trans)
}

// ValidateStruct 验证结构体
func ValidateStruct(s interface{}, fieldZH ...map[string]string) error {
	errValidate := validate.Struct(s)
	if errValidate != nil {
		var (
			errs          []string
			fieldReplaces []string
		)
		for _, field := range fieldZH {
			for k, v := range field {
				fieldReplaces = append(fieldReplaces, k, v)
			}
		}
		replacer := strings.NewReplacer(fieldReplaces...)
		for _, v := range errValidate.(validator.ValidationErrors).Translate(trans) {
			errs = append(errs, replacer.Replace(v))
		}
		return errors.New(strings.Join(errs, "；"))
	}
	return nil
}
