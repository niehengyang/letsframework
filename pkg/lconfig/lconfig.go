package lconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

type LConfig struct {
	configFile string
	viper      *viper.Viper
}

func New(configFile string) *LConfig {
	viperInstance := viper.New()
	viperInstance.SetConfigFile(configFile)
	viperInstance.SetConfigType("toml")
	err := viperInstance.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取配置文件%v时发生错误，错误原因%v", configFile, err))
	}
	return &LConfig{configFile: configFile, viper: viperInstance}
}

func (l *LConfig) Unmarshal(configStruct interface{}) {
	l.viper.Unmarshal(configStruct)
}

func (l *LConfig) Get(key string) interface{} {
	return l.viper.Get(key)
}
