package enum

import scaffold "github.com/xinlianit/kit-scaffold"

const (
	StatusEnable  scaffold.EnumType = 1
	StatusDisable scaffold.EnumType = 2
	StatusDelete  scaffold.EnumType = 99
)

func init() {
	scaffold.EnumName[StatusEnable] = "启用"
	scaffold.EnumName[StatusDisable] = "禁用"
	scaffold.EnumName[StatusDelete] = "删除"
}

// 状态枚举
func StatusEnums() []scaffold.Enum {
	return []scaffold.Enum{
		{Value: StatusEnable.Value(), Name: StatusEnable.Name()},
		{Value: StatusDisable.Value(), Name: StatusDisable.Name()},
		{Value: StatusDelete.Value(), Name: StatusDelete.Name()},
	}
}
