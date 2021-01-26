package enum

import scaffold "github.com/xinlianit/kit-scaffold"

const (
	// 公共错误码
	CodeSuccess        scaffold.EnumType = 0
	CodeErrorServer    scaffold.EnumType = 500
	CodeErrorRateLimit scaffold.EnumType = 501
	CodeFail           scaffold.EnumType = 510
	CodeErrorParam     scaffold.EnumType = 512
	CodeErrorAppId     scaffold.EnumType = 513
	CodeErrorSign      scaffold.EnumType = 514
)

func init() {
	scaffold.EnumName[CodeSuccess] = "成功"
	scaffold.EnumName[CodeErrorServer] = "系统内部错误"
	scaffold.EnumName[CodeErrorRateLimit] = "系统繁忙"
	scaffold.EnumName[CodeFail] = "失败"
	scaffold.EnumName[CodeErrorParam] = "参数错误"
	scaffold.EnumName[CodeErrorAppId] = "AppId无效"
	scaffold.EnumName[CodeErrorSign] = "签名错误"
}
