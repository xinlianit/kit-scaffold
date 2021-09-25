package driver

import (
	"database/sql"
	"fmt"
	"github.com/xinlianit/kit-scaffold/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

// SqlLoggerConfig SQL 日志配置
type SqlLoggerConfig struct {
	// Enable SQL日志记录：true-开启、false-关闭
	Enable bool
	// LogFile SQL日志文件
	LogFile string
	// SlowThreshold 记录慢SQL阈值(单位：毫秒)
	SlowThreshold uint
	// LogLevel 日志记录级别(高到低): 1-Silent、2-Error、3-Warn、4-Info
	LogLevel logger.LogLevel
	// IgnoreRecordNotFoundError 忽略ErrRecordNotFound（记录未找到）错误
	IgnoreRecordNotFoundError bool
	// Colorful 是否彩色打印: false-否、true-是
	Colorful bool
}

var (
	SqlDB           *sql.DB
	SqlDbErr        error
	MysqlDB         *gorm.DB
	MysqlDbErr      error
	mysqlConfig     config.MySql
	sqlLoggerConfig SqlLoggerConfig
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

	// gorm 配置
	gormConfig := &gorm.Config{}

	// 解构配置到结构
	config.Config().UnmarshalKey("logger.sql", &sqlLoggerConfig)

	// 启用SQL日志记录
	if sqlLoggerConfig.Enable {
		// SQL 日志输出文件
		sqlLogFile, err := os.OpenFile(sqlLoggerConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Failed to open error logger file: " + err.Error())
		}

		// 数据库日志记录器 log.LstdFlags
		sqlLogger := logger.New(
			// io writer（日志输出的目标，前缀和日志包含的内容）
			log.New(io.MultiWriter(sqlLogFile, os.Stdout), "\r\n", log.LstdFlags),
			logger.Config{
				// 记录慢SQL阈值: 200 毫秒
				SlowThreshold: time.Duration(sqlLoggerConfig.SlowThreshold) * time.Millisecond,
				// 日志记录级别(高到低): 1-Silent、2-Error、3-Warn、4-Info
				LogLevel: sqlLoggerConfig.LogLevel,
				// 忽略ErrRecordNotFound（记录未找到）错误
				IgnoreRecordNotFoundError: sqlLoggerConfig.IgnoreRecordNotFoundError,
				// 是否彩色打印: false-否、true-是
				Colorful: sqlLoggerConfig.Colorful,
			})

		// 设置自定义SQL Logger
		gormConfig.Logger = sqlLogger
	}

	// GORM 数据库实例
	MysqlDB, MysqlDbErr = gorm.Open(mysql.New(mysql.Config{
		Conn: SqlDB,
	}), gormConfig)

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
