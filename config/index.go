package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 系统相关配置
type Config struct {
	RedisAdmin RedisAdmin
}

// RedisAdmin Redis连接
type RedisAdmin struct {
	Addr     string
	Password string
	DB       int
}

// GinVueAdminconfig GinVueAdminconfig
var GinVueAdminconfig Config

func init() {
	v := viper.New()
	v.SetConfigName("config")   //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./config") // 第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := v.Unmarshal(&GinVueAdminconfig); err != nil {
		fmt.Println(err)
	}
}
