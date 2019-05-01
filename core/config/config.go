package config

// DatabaseConfig has database related configuration.
type DatabaseConfig struct {
	Type             string `mapstructure:"type"`
	Host             string `mapstructure:"host"`
	DbName           string `mapstructure:"dbName"`
	UserName         string `mapstructure:"userName"`
	Password         string `mapstructure:"password"`
	ConnectionString string `mapstructure:"connectionString"`
}

// CacheConfig has cache related configuration.
type CacheConfig struct {
	Type     string `mapstructure:"type"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type MqConfig struct {
	Type              string `mapstructure:"type"`
	QueueUrl          string `mapstructure:"queueUrl"`
	VisibilityTimeout int64  `mapstructure:"visibilityTimeout"`
	Region            string `mapstructure:"region"`
	AccessKey         string `mapstructure:"accessKey"`
	Secret            string `mapstructure:"secret"`
}

// LoggerConfig has logger related configuration.
type LoggerConfig struct {
	LogFilePath string `mapstructure:"file"`
}
