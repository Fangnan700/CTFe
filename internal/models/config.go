package models

type Config struct {
	MySqlConfig mysqlConfig `yaml:"mysql_config"`
}

type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
