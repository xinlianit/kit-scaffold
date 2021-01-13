package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xinlianit/go-util/util"
	"github.com/xinlianit/kit-scaffold/common"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"path"
	"path/filepath"
	"sync"
)

var (
	once               sync.Once
	config             *viper.Viper             // 应用静态配置（运行时不可修改）
	dynamicConfig      *viper.Viper             // 应用动态配置（运行时可修改，如：通过配置中心修改）
	configDir          = common.GetConfigPath() // 配置目录
	dynamicConfigDir   string                   // 动态配置目录
	configBackupDir    string                   // 动态配置备份目录
	configBackupPrefix string                   // 动态配置备份前缀
)

// 配置初始化
func Init() {
	once.Do(func() {
		// 初始化静态配置
		initConfig()
	})

	// 动态配置目录
	dynamicConfigDir = Config().GetString("app.configCenter.configCacheDir")
	// 动态配置备份目录
	configBackupDir = dynamicConfigDir + "-backup"
	// 动态配置备份前缀
	configBackupPrefix = "backup-"
}

// 静态配置
func initConfig() {
	// 创建配置
	config = viper.New()

	// 自动加载环境变量
	config.AutomaticEnv()

	// 加载配置文件
	if err := loadDirAllConfig(config, configDir, true, false); err != nil {
		panic(err)
	}

	// 初始化默认配置
	initDefaultConfig(config, Default)
}

// 动态配置
func InitDynamicConfig() {
	// 创建配置
	dynamicConfig = viper.New()

	// 加载配置文件
	if err := loadDirAllConfig(dynamicConfig, dynamicConfigDir, false, true); err != nil {
		panic(err)
	}

	// 备份配置
	if err := util.FileUtil().CopyDirWithFix(dynamicConfigDir, configBackupDir, configBackupPrefix); err != nil {
		panic(err)
	}

	// 初始化默认配置
	initDefaultConfig(dynamicConfig, Default)

	// 监听配置
	go func() {
		dynamicConfig.OnConfigChange(func(e fsnotify.Event) {
			// todo 记录到日志
			fmt.Printf("[%v] Viper dynamic config changed: %v\n", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout), e.Name)

			if err := loadDirAllConfig(dynamicConfig, dynamicConfigDir, false, false); err != nil {
				// todo 记录到日志, 邮件告警【重要】
				fmt.Printf("[%v] Viper dynamic config reload error: %v\n", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout), err)

				// todo 重载备份配置
				if err := loadDirAllConfig(dynamicConfig, configBackupDir, false, false); err != nil {
					// todo 记录到日志, 邮件告警【重要】
					fmt.Printf("[%v] Viper dynamic backup config reload error: %v\n", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout), err)
				}
			} else {
				// 配置备份
				dstFile := filepath.Join(configBackupDir, path.Base(e.Name))
				util.FileUtil().CopyFileWithFix(e.Name, dstFile, configBackupPrefix)
			}
		})
	}()
}

// 初始化默认配置
// @param defaultConfig 默认配置
func initDefaultConfig(viperConfig *viper.Viper, defaultConfig map[string]interface{}) {
	for k, v := range defaultConfig {
		// 设置默认值
		viperConfig.SetDefault(k, v)
	}
}

// 加载目录所有配置
// @param config 配置实例
// @param configDir 配置目录
// @param loadChildDir 是否加载子目录配置: true-是、false-否
// @param watchEnable 是否监听配置变化
func loadDirAllConfig(config *viper.Viper, configDir string, loadChildDir bool, watchEnable bool) error {
	// 列出目录文件列表s
	fileList, err := util.FileUtil().LsDir(configDir, loadChildDir)
	if err != nil {
		return err
	}

	// 添加配置目录
	config.AddConfigPath(configDir)

	for _, file := range fileList {
		fileFullName := path.Base(file)
		fileExt := path.Ext(file)
		fileName := fileFullName[0 : len(fileFullName)-len(fileExt)]
		//fileName := strings.TrimSuffix(fileFullName, fileExt)

		// 获取子目录
		if childPath := filepath.Dir(file); childPath != "." && childPath != ".." {
			// 添加配置子目录
			config.AddConfigPath(filepath.Join(configDir, childPath))
		}

		// 设置配置
		config.SetConfigName(fileName)
		//config.SetConfigFile(fileFullName)
		//config.SetConfigType(strings.TrimLeft(fileExt, "."))

		// 合并配置文件
		if err := config.MergeInConfig(); err != nil {
			return err
		}

		// 监听配置变更
		if watchEnable {
			config.WatchConfig()
		}
	}

	return nil
}

// 静态配置
func Config() *viper.Viper {
	return config
}

// 动态配置
func DynamicConfig() *viper.Viper {
	return dynamicConfig
}
