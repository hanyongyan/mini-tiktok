package config

type GlobalConfig struct {
	DbConfig struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	} `mapstructure:"db"`
	MqConfig struct {
		Host string
		Port int
	} `mapstructure:"mq"`
	RedisConfig struct {
		Host string
		Port int
	} `mapstructure:"redis"`
	CosConfig struct {
		Url string
		SecretId string  `mapstructure:"secret_id"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"cos"`
	MongoDbConfig struct{
		Addr string `mapstructure:"addr"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DbName string `mapstructure:"dbname"`
	} `mapstructure:"mongo"`
}
