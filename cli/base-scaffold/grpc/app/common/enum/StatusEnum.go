package enum

import scaffold "github.com/xinlianit/kit-scaffold"

const (
	BusinessNotExists scaffold.EnumType = 20000001
)

func init()  {
	scaffold.EnumName[BusinessNotExists] = "商家信息不存在"
}
