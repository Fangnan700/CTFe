package config

import (
	"CTFe/server/model/common"
	"CTFe/server/util/log"
	"gopkg.in/yaml.v3"
	"os"
)

var GlobalConfig common.CTFeConfig

func init() {
	buff, err := os.ReadFile("config.yaml")
	if err != nil {
		log.ErrorLogger.Println(err.Error())
	}

	err = yaml.Unmarshal(buff, &GlobalConfig)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
	}
}
