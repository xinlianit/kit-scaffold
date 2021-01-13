package enum

import scaffold "github.com/xinlianit/kit-scaffold"

const (
	// 公共错误码
	CodeSuccess   scaffold.EnumType = 0
	CodeFail      scaffold.EnumType = 1
	CodeErrParam  scaffold.EnumType = 2
	CodeErrAppId  scaffold.EnumType = 3
	CodeErrSign   scaffold.EnumType = 4
	CodeErrServer scaffold.EnumType = 500

	// 资源模块
	CodeErrResource          scaffold.EnumType = 10000
	CodeErrResourceInfo      scaffold.EnumType = 10001
	CodeErrResourceNotExists scaffold.EnumType = 10002
)

func init() {
	scaffold.EnumName[CodeSuccess] = "成功"
	scaffold.EnumName[CodeFail] = "失败"
	scaffold.EnumName[CodeErrParam] = "参数错误"
	scaffold.EnumName[CodeErrAppId] = "AppId无效"
	scaffold.EnumName[CodeErrSign] = "签名错误"
	scaffold.EnumName[CodeErrServer] = "系统异常，服务器内部错误"
	scaffold.EnumName[CodeErrResource] = "资源错误"
	scaffold.EnumName[CodeErrResourceInfo] = "资源详情错误"
	scaffold.EnumName[CodeErrResourceNotExists] = "资源不存在"
}
