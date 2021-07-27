package enum

import scaffold "github.com/xinlianit/kit-scaffold"

const (
	CodeSuccess        scaffold.EnumType = 0    // CodeSuccess 0-成功
	CodeErrorServer    scaffold.EnumType = 500  // CodeErrorServer 500-系统内部错误
	CodeErrorRateLimit scaffold.EnumType = 501  // CodeErrorRateLimit 501-系统繁忙
	CodeFail           scaffold.EnumType = 1000 // CodeFail 1000-失败
	CodeError          scaffold.EnumType = 1001 // CodeError 1001-错误
	CodeErrorParam     scaffold.EnumType = 1002 // CodeErrorParam 1002-参数错误
	CodeInvalidParams  scaffold.EnumType = 1003 // CodeInvalidParams 1003-无效认证参数
	CodeUnauthorized   scaffold.EnumType = 1004 // CodeUnauthorized 1004-应用未授权
	CodeAuthorizeFail  scaffold.EnumType = 1005 // CodeAuthorizeFail 1005-认证失败
	CodeErrorAppId     scaffold.EnumType = 1006 // CodeErrorAppId 1006-AppId无效
	CodeErrorSign      scaffold.EnumType = 1007 // CodeErrorSign 1007-签名错误
)

func init() {
	scaffold.EnumName[CodeSuccess] = "成功"
	scaffold.EnumName[CodeErrorServer] = "系统内部错误"
	scaffold.EnumName[CodeErrorRateLimit] = "系统繁忙"
	scaffold.EnumName[CodeFail] = "失败"
	scaffold.EnumName[CodeError] = "错误"
	scaffold.EnumName[CodeErrorParam] = "参数错误"
	scaffold.EnumName[CodeInvalidParams] = "无效认证参数"
	scaffold.EnumName[CodeUnauthorized] = "应用未授权"
	scaffold.EnumName[CodeAuthorizeFail] = "认证失败"
	scaffold.EnumName[CodeErrorAppId] = "AppId无效"
	scaffold.EnumName[CodeErrorSign] = "签名错误"
}
