package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/logger"
	"go.uber.org/zap"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	nacosClientUtil *util.Nacos
	err             error
	once            sync.Once
	group string
	configSyncDir   string       // 配置同步目录
	nacosConfig     config.Nacos // Nacos 配置
)

// Init 初始化 Nacos
func Init() {
	once.Do(func() {
		// 配置同步目录
		configSyncDir = config.Config().GetString("app.configCenter.configCacheDir")

		// 解析配置到结构
		config.Config().UnmarshalKey("nacos", &nacosConfig)

		// 命令行指定地址
		var ipAddrSlice []string
		var portSlice []uint64
		// nacos 地址
		if address := config.Config().GetString("nacos.address"); address != "" {
			if addressSlice := strings.Split(address, ","); addressSlice != nil {
				for _, address := range addressSlice {
					ipPort := strings.Split(address, ":")

					ipAddrSlice = append(ipAddrSlice, ipPort[0])

					if len(ipPort) == 2 {
						port, _ := strconv.ParseUint(ipPort[1], 10, 64)
						portSlice = append(portSlice, port)
					}else{
						portSlice = append(portSlice, 80)
					}
				}
			}
		}

		//配置中心服务端配置
		var serverConfigs []constant.ServerConfig

		if ipAddrSlice != nil {
			// 命令行参数
			var scheme, contextPath string
			if len(nacosConfig.ServerConfig) > 0 {
				scheme = nacosConfig.ServerConfig[0].Scheme
				contextPath = nacosConfig.ServerConfig[0].ContextPath
			}

			for i, host := range ipAddrSlice {
				serverConfig := constant.ServerConfig{
					Scheme:      scheme,
					ContextPath: contextPath,
					IpAddr:      host,
					Port:        portSlice[i],
				}
				serverConfigs = append(serverConfigs, serverConfig)
			}
		}else{
			// 配置文件
			for _, cfg := range nacosConfig.ServerConfig {
				serverConfig := constant.ServerConfig{
					Scheme:      cfg.Scheme,
					ContextPath: cfg.ContextPath,
					IpAddr:      cfg.IpAddr,
					Port:        cfg.Port,
				}
				serverConfigs = append(serverConfigs, serverConfig)
			}
		}

		// 配置中心客户端配置
		clientConfig := constant.ClientConfig{
			NamespaceId:         nacosConfig.ClientConfig.NamespaceId,
			TimeoutMs:           nacosConfig.ClientConfig.Timeout,
			NotLoadCacheAtStart: nacosConfig.ClientConfig.NotLoadCacheAtStart,
			LogDir:              nacosConfig.ClientConfig.LogDir,
			CacheDir:            nacosConfig.ClientConfig.CacheDir,
			RotateTime:          nacosConfig.ClientConfig.RotateTime,
			MaxAge:              nacosConfig.ClientConfig.MaxAge,
			LogLevel:            nacosConfig.ClientConfig.LogLevel,
		}

		// 配置中心属性
		properties := map[string]interface{}{
			"serverConfigs": serverConfigs,
			"clientConfig":  clientConfig,
		}

		// 创建配置中心客户端
		configClient, err := clients.CreateConfigClient(properties)
		if err != nil {
			panic(err)
		}

		// 创建服务发现客户端
		namingClient, err := clients.CreateNamingClient(properties)
		if err != nil {
			panic(err)
		}

		// nacos 分组
		if group = config.Config().GetString("app.configCenter.nacosGroup"); group == "" {
			group = strings.ToUpper(config.Config().GetString("env"))
		}
		nacosClientUtil = util.NacosUtil().WithConfigClient(configClient).WithNamingClient(namingClient).Group(group)
	})
}

// NacosClientUtil Nacos 客户端工具
func NacosClientUtil() *util.Nacos {
	return nacosClientUtil
}

// ListenSyncConfig 同步配置
func ListenSyncConfig() {
	syncDataIds := config.Config().GetString("app.configCenter.syncConfigDataIds")
	// 同步配置文件列表
	configFileList := strings.Split(syncDataIds, ",")
	for _, configFile := range configFileList {
		// 获取配置到文件
		getConfigToFile(group, configFile)

		// 监听配置
		NacosClientUtil().ListenConfig(configFile, syncConfigToFile)
	}
}

// getConfigToFile 获取配置到文件
func getConfigToFile(group string, dataId string) {
	// 名称空间
	var namespace string
	if namespace = config.Config().GetString("nacos.namespace"); namespace == "" {
		namespace = config.Config().GetString("nacos.clientConfig.namespaceId")
	}

	// 配置目录检测
	if !util.FileUtil().FileExist(configSyncDir) {
		if err := util.DirUtil().CreateDir(configSyncDir, true); err != nil {
			logger.ZapLogger.Error(err.Error())
			panic(err)
		}
	}

	// 配置文件
	configFile := filepath.Join(configSyncDir, fmt.Sprintf("%v-%v-%v", namespace, group, dataId))

	// 获取配置
	configData := NacosClientUtil().Group(group).GetConfig(dataId, "")
	// 配置数据为空，且配置文件不存在
	if configData == "" && !util.FileUtil().FileExist(configFile) {
		// 创建配置文件
		if _, err := util.FileUtil().Write(configFile, configData); err != nil {
			logger.ZapLogger.Error(err.Error())
			panic(err)
		}
	}

	// 写入配置数据到配置文件
	if _, err := util.FileUtil().Write(configFile, configData); err != nil {
		logger.ZapLogger.Error(err.Error())
		panic(err)
	}

	fields := []zap.Field{
		zap.String("namespace", namespace),
		zap.String("group", group),
		zap.String("data_id", dataId),
		zap.String("data", configData),
	}
	logger.ZapLogger.Info(fmt.Sprintf("Nacos 获取并同步: namespace: %v, group: %v, dataId: %v", namespace, group, dataId), fields...)
}

// syncConfigToFile 同步配置到文件
func syncConfigToFile(namespace string, group string, dataId string, data string) {
	fields := []zap.Field{
		zap.String("namespace", namespace),
		zap.String("group", group),
		zap.String("data_id", dataId),
		zap.String("data", data),
	}
	logger.ZapLogger.Info(fmt.Sprintf("Nacos 监听并同步: namespace: %v, group: %v, dataId: %v", namespace, group, dataId), fields...)

	// 配置防空保护，配置为空时，跳过同步
	if data == "" {
		logger.ZapLogger.Sugar().Infof("Nacos 跳过同步: namespace: %v, group: %v, dataId: %v", namespace, group, dataId)
		return
	}

	// 目录检测
	if !util.FileUtil().FileExist(configSyncDir) {
		if err := util.DirUtil().CreateDir(configSyncDir, true); err != nil {
			logger.ZapLogger.Error(err.Error())
			return
		}
	}

	// 写入配置缓存
	cacheFile := filepath.Join(configSyncDir, fmt.Sprintf("%v-%v-%v", namespace, group, dataId))
	_, err := util.FileUtil().Write(cacheFile, data)

	if err != nil {
		logger.ZapLogger.Error(err.Error())
	}
}
