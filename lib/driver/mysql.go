package driver

import (
	"database/sql"
	"fmt"
	"github.com/xinlianit/kit-scaffold/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	SqlDB       *sql.DB
	SqlDbErr    error
	MysqlDB     *gorm.DB
	MysqlDbErr  error
	mysqlConfig config.MySql
)

// InitMySql 初始化 MySql 数据库
func InitMySql() {
	// 解构配置到结构
	config.DynamicConfig().UnmarshalKey("datasource.mysql", &mysqlConfig)

	// 数据源名称
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%ds",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DatabaseName,
		mysqlConfig.Charset,
		mysqlConfig.Timeout,
	)

	// 连接并打开数据库
	SqlDB, SqlDbErr = sql.Open("mysql", dbDSN)
	// 连接错误
	if SqlDbErr != nil {
		panic("database data source name error: " + SqlDbErr.Error())
	}

	// GORM 数据库实例
	MysqlDB, MysqlDbErr = gorm.Open(mysql.New(mysql.Config{
		Conn: SqlDB,
	}), &gorm.Config{})

	// 连接错误
	if MysqlDbErr != nil {
		panic("mysql db instance error: " + MysqlDbErr.Error())
	}

	// 数据库连接池
	// 设置最大连接数
	SqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConnect)

	// 设置最大空闲连接数
	SqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConnect)

	// 设置连接最大生存期，0：永不过期
	//SqlDB.SetConnMaxLifetime(time.Duration(MySqlConfig.ConnectMaxLifetime * 1000000000))
	SqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnectMaxLifetime) * time.Second)

	// 检测连接是否成功
	if SqlDbErr = SqlDB.Ping(); SqlDbErr != nil {
		panic("database connect failed: " + SqlDbErr.Error())
	}
}
