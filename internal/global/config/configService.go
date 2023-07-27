package config

import (
	"CTFe/internal/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	err          error
	buff         []byte
	GlobalConfig models.Config
)

func init() {
	// 读取配置文件
	buff, err = os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// 解析配置到结构体
	err = yaml.Unmarshal(buff, &GlobalConfig)
	if err != nil {
		fmt.Println(err)
	}
}
