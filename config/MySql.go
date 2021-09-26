package config

type MySql struct {
	Host               string `json:"host"`                 // 数据库主机地址
	Port               int    `json:"port"`                 // 数据库端口
	DatabaseName       string `json:"database_name"`        // 数据库名称
	User               string `json:"user"`                 // 数据库账号
	Password           string `json:"password"`             // 数据库密码
	Charset            string `json:"charset"`              // 数据库字符集编码
	Timeout            int    `json:"timeout"`              // 连接超时时间 (单位：毫秒)
	ReadTimeout        int    `json:"read_timeout"`         // 读超时时间 (单位：毫秒)
	WriteTimeout       int    `json:"write_timeout"`        // 写超时时间 (单位：毫秒)
	MaxOpenConnect     int    `json:"max_open_connect"`     // 连接池最大连接数
	MaxIdleConnect     int    `json:"max_idle_connect"`     // 连接池最大空闲数
	ConnectMaxLifetime int    `json:"connect_max_lifetime"` // 连接池链接最大生存期（单位：秒），0：永不过期
}
