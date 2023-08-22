package variable

import (
    "go.uber.org/zap"

    "github.com/hd2yao/gin-skeleton/app/utils/yml_config/ymlconfig_interf"
)

var (
    BasePath        string      // 定义项目的根目录
    ConfigKeyPrefix = "Config_" // 配置文件键值缓存时，键的前缀

    // 全局日志指针
    ZapLog *zap.Logger
    // 全局配置文件
    ConfigYml       ymlconfig_interf.YmlConfigInterf // 全局配置文件指针
    ConfigGormv2Yml ymlconfig_interf.YmlConfigInterf // 全局配置文件指针
)
