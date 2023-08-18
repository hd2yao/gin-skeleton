package yml_config

import (
	"github.com/hd2yao/gin-skeleton/app/global/my_errors"
	"github.com/hd2yao/gin-skeleton/app/global/variable"
	"github.com/hd2yao/gin-skeleton/app/utils/yml_config/ymlconfig_interf"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

// CreateYamlFactory 创建一个 yaml 配置文件工厂
// 参数设置为可变参数的文件名，这样参数就可以不需要传递，如果传递了多个，我们只取第一个参数作为配置文件名
func CreateYamlFactory(fileName ...string) ymlconfig_interf.YmlConfigInterf {
	yamlConfig := viper.New()
	// 配置文件所在的目录
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名，默认为：config
	if len(fileName) == 0 {
		yamlConfig.SetConfigName("config")
	} else {
		yamlConfig.SetConfigName(fileName[0])
	}
	// 设置配置文件类型(后缀)为 yml
	yamlConfig.SetConfigType("yml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}

	return &ymlConfig{
		viper: yamlConfig,
		mu:    new(sync.Mutex),
	}
}

type ymlConfig struct {
	viper *viper.Viper
	mu    *sync.Mutex
}

func (y *ymlConfig) ConfigFileChangeListen() {

}

func (y *ymlConfig) Clone(fileName string) ymlconfig_interf.YmlConfigInterf {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) Get(keyName string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetString(keyName string) string {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetBool(keyName string) bool {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetInt(keyName string) int {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetInt32(keyName string) int32 {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetInt64(keyName string) int64 {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetFloat64(keyName string) float64 {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetDuration(keyName string) time.Duration {
	//TODO implement me
	panic("implement me")
}

func (y *ymlConfig) GetStringSlice(keyName string) []string {
	//TODO implement me
	panic("implement me")
}
