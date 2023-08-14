package common

type CTFeConfig struct {
	ServerConfig serverConfig `yaml:"server_config"`
	MysqlConfig  mysqlConfig  `yaml:"mysql_config"`
	RedisConfig  redisConfig  `yaml:"redis_config"`
}

// 服务配置
type serverConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// mysql数据库配置
type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// redis缓存配置
type redisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database int    `yaml:"database"`
	Password string `yaml:"password"`
}
