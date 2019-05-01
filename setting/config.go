package setting

import (
	"fmt"
	"os"

	"github.com/MiteshSharma/project/model"
	"github.com/spf13/viper"
)

func GetConfig() *model.Config {
	appConfig := &model.Config{}
	GetConfigFromFile1("default", appConfig)
	return appConfig
}

func GetConfigFromFile(file string) *model.Config {
	appConfig := &model.Config{}
	GetConfigFromFile1(file, appConfig)
	return appConfig
}

func GetConfigFromFile1(fileName string, config *model.Config) {
	if fileName == "" {
		fileName = "default"
	}
	viper.SetConfigName(fileName)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&config)
	fmt.Println(config)
	fmt.Println(config.ServerConfig)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
}
