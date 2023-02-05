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
	FtpConfig struct {
		Host     string
		Port     int
		Username string
		Password string
	} `mapstructure:"ftp"`
	StaticConfig struct {
		Url     string
		TmpPath string `mapstructure:"local_tmp_path"`
	} `mapstructure:"static"`
}
