package models

// Config 全局配置
type Config struct {
	MySqlConfig mysqlConfig `yaml:"mysql_config"`
	RedisConfig redisConfig `yaml:"redis_config"`
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
