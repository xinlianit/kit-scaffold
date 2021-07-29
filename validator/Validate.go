package validator

// 仓库地址：https://github.com/go-playground/validator
// 参考资料：https://blog.csdn.net/guyan0319/article/details/105918559/
// 参考资料: https://www.cnblogs.com/zj420255586/p/13542395.html

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	global "github.com/xinlianit/kit-scaffold/common/gloabl"
	"reflect"
	"strings"
)

var (
	// validate 验证器
	validate *validator.Validate
	// translator 多语言翻译转换器
	translator ut.Translator
)

func init() {
	// 语言列表
	languages := map[string]string{
		"zh-cn": "zh", // 中文
		"en-us": "en", // 英文
	}

	// 当地地区
	locale := "zh"

	// 当地语言
	localLanguage := strings.ToLower(global.AcceptLanguage)

	// 获取请求头 Accept-Language
	if global.AcceptLanguage != "" {
		locale = languages[localLanguage]
	}

	// 默认中文语言包
	defaultTrans := zh.New()

	// 翻译语言包
	var translators []locales.Translator

	// 加载英文语言包
	if localLanguage == "en-us" {
		translators = append(translators, en.New())
	}

	// 创建翻译转换器
	uniTrans := ut.New(defaultTrans, translators...)

	// 获取翻译转换器
	translator, _ = uniTrans.GetTranslator(locale)

	// 验证器引擎
	validate = validator.New()

	// 注册翻译器
	switch locale {
	case "zh-cn":
		zhTranslations.RegisterDefaultTranslations(validate, translator)
	case "en-us":
		enTranslations.RegisterDefaultTranslations(validate, translator)
	default:
		zhTranslations.RegisterDefaultTranslations(validate, translator)
	}

	// 注册标签名称函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
}

// Validate 数据验证
// @param data 数据
// @return validateEngine 验证器
func Validate(data interface{}) validateEngine {
	return validateEngine{
		translator: translator,
		validate:   validate,
		err:        validate.Struct(data),
	}
}

// validateEngine 验证器
type validateEngine struct {
	// translator 中文翻译转换器
	translator ut.Translator
	// validate 多语言翻译转换器
	validate *validator.Validate
	// err 验证器错误信息
	err error
}

// GetAllValidateError 获取所有验证器错误
// @return []string 错误信息
func (v validateEngine) GetAllValidateError() []string {
	var allErrorMsg []string
	errs, ok := v.err.(validator.ValidationErrors)
	if ok {
		for _, err := range errs {
			allErrorMsg = append(allErrorMsg, err.Translate(v.translator))
		}
	}
	return allErrorMsg
}

// GetCurrentValidateError 获取当前验证错误
// @return string 错误信息
func (v validateEngine) GetCurrentValidateError() string {
	if errorsMsg := v.GetAllValidateError(); errorsMsg != nil && len(errorsMsg) > 0 {
		return errorsMsg[0]
	}

	return ""
}

// validateTranslation 注册验证器翻译
// @param tag 验证标签
// @param msg 提示信息
func (v *validateEngine) validateTranslation(tag string, msg string) {
	v.validate.RegisterTranslation(tag, v.translator, func(ut ut.Translator) error {
		return ut.Add(tag, msg, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field(), fe.Field())
		return t
	})
}
