package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Configuration interface {
	Init()
	GetString(key string) string
	GetListFileInfo() ListOfFilesInfo
}

type Config struct{}

type ListOfFilesInfo []FileInfo
type FileInfo struct {
	Type string
	Path string
}

func (c Config) Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (c Config) GetEnv() string {
	envVal := os.Getenv("LOGTOMONGO_ENV")
	if envVal == "" {
		envVal = "prod"
	}
	return envVal
}

func (c Config) GetListFileInfo() ListOfFilesInfo {
	var listOfFilesInfo ListOfFilesInfo
	logFiles := viper.Get(c.GetEnv() + ".log_files").([]interface{})
	for _, v := range logFiles {
		mapFile := v.(map[string]interface{})
		listOfFilesInfo = append(listOfFilesInfo, FileInfo{
			Type: mapFile["type"].(string),
			Path: mapFile["path"].(string),
		})
	}
	return listOfFilesInfo
}

func (c Config) GetString(key string) string {
	return viper.GetString(c.GetEnv() + "." + key)
}
