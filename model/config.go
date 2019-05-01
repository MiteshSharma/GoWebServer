package model

import (
	"time"

	"github.com/MiteshSharma/project/core/config"
)

type Config struct {
	LoggerConfig   config.LoggerConfig   `mapstructure:"logger"`
	DatabaseConfig config.DatabaseConfig `mapstructure:"database"`
	CacheConfig    config.CacheConfig    `mapstructure:"cache"`
	MqConfig       config.MqConfig       `mapstructure:"mq"`
	ServerConfig   ServerConfig          `mapstructure:"server"`
	ZipkinConfig   ZipkinConfig          `mapstructure:"zipkin"`
	AuthConfig     AuthConfig            `mapstructure:"auth"`
}

// ServerConfig has only server specific configuration
type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	CloseTimeout time.Duration `mapstructure:"closeTimeout"`
}

// ZipkinConfig has zipkin related configuration.
type ZipkinConfig struct {
	IsEnable bool   `mapstructure:"isEnable"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

// AuthConfig has logger related configuration.
type AuthConfig struct {
	HmacSecret string `mapstructure:"hmacSecret"`
}
