package common

var (
	RootPath string	// 根路径
	ConfigPath string // 配置路径
)

// ArgsToMap args 参数转map
func ArgsToMap(args []string, keys []string) map[string]string {
	// 参数解析
	argsMap := make(map[string]string)
	for i:=0; i<len(keys); i++ {
		if i < len(args) {
			argsMap[keys[i]] = args[i]
		}
	}

	return argsMap
}

