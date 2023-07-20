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
	buff, err = os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(buff, &GlobalConfig)
	if err != nil {
		fmt.Println(err)
	}
}
