package scaffold

// Enum 枚举结构
type Enum struct {
	Value int
	Name  string
}

// EnumType 枚举类型
type EnumType int

// EnumName 枚举名称
var EnumName map[EnumType]string

// 初始化
func init() {
	EnumName = make(map[EnumType]string)
}

// Value 获取枚举值
func (e EnumType) Value() int {
	return int(e)
}

// Name 获取枚举名称
func (e EnumType) Name() string {
	return EnumName[e]
}

