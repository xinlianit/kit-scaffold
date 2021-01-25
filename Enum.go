package scaffold

// 枚举结构
type Enum struct {
	Value int
	Name  string
}

// 枚举类型
type EnumType int

// 枚举名称
var EnumName map[EnumType]string

// 包初始化
func init() {
	EnumName = make(map[EnumType]string)
}

// 获取枚举值
func (e EnumType) Value() int {
	return int(e)
}

// 获取枚举名称
func (e EnumType) Name() string {
	return EnumName[e]
}

