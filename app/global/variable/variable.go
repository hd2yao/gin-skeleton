package variable

import "go.uber.org/zap"

var (
	BasePath string // 定义项目的根目录

	// 全局日志指针
	ZapLog *zap.Logger
)
