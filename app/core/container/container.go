package container

import "sync"

// SMap 定义一个全局键值对存储容器
var SMap sync.Map

// CreateContainersFactory 创建一个容器工厂
func CreateContainersFactory() *containers {
	return &containers{}
}

// 定义一个容器结构体
type containers struct{}
