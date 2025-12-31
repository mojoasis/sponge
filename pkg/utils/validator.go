package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义为包级变量，供下方的 Translate 函数使用
var trans ut.Translator

func InitTrans() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 1. 注册一个获取 json 标签名的函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 2. 初始化翻译器
		zhT := zh.New()
		uni := ut.New(zhT, zhT)
		var ok bool
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("uni.GetTranslator(zh) failed")
		}

		// 3. 注册默认的中文字典
		return zhtranslations.RegisterDefaultTranslations(v, trans)
	}
	return nil
}

// Translate 翻译错误信息
func Translate(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	var res []string
	for _, e := range errs.Translate(trans) {
		res = append(res, e)
	}
	return strings.Join(res, ", ")
}
