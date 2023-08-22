package variable

import "go.uber.org/zap"

var (
	BasePath        string      // 定义项目的根目录
	ConfigKeyPrefix = "Config_" // 配置文件键值缓存时，键的前缀

	// 全局日志指针
	ZapLog *zap.Logger
)
