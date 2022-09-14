package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 将配置映射成 Config 对象

// LoadConfigFromToml 加载配置从文件
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path: %s, %s", filePath, err)
	}
	return nil
}

// LoadConfigFromEnv 加载配置从环境变量
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}
	return nil
}
