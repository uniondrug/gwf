// author: wsfuyibing <websearch@163.com>
// date: 2021-02-01

package xapp

import (
	"errors"
	"reflect"

	i18n "github.com/go-playground/locales/zh"
	i18nTranslator "github.com/go-playground/universal-translator"
	i18nValidator "github.com/go-playground/validator/v10"
	i18nTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 校验结构体.
//
//   req := &ExampleStruct {
//       Id int `validate:"required,gte=3,lte=10"`
//   }
//
//   if err := xapp.Validate(req); err != nil {
//       println("Validate error:", err.Error())
//   }
//
func Validate(s interface{}) error {
	// translate manager.
	trans := i18n.New()
	translator := i18nTranslator.New(trans, trans)
	translation, _ := translator.GetTranslator("zh")
	// tag filter accepted.
	validator := i18nValidator.New()
	validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		if label := field.Tag.Get("label"); label != "" {
			return label
		}
		return field.Name
	})
	// register transaction.
	if err := i18nTranslations.RegisterDefaultTranslations(validator, translation); err != nil {
		return err
	}
	// return first error
	if e0 := validator.Struct(s); e0 != nil {
		for _, e1 := range e0.(i18nValidator.ValidationErrors) {
			err := errors.New(e1.Translate(translation))
			return err
		}
	}
	return nil
}
