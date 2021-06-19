package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"viperconfig/model"
)

func initViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic("no config")
	}
}

type AgentConfig struct {
	Version string
	// HTTP port.
	Port string
	// Log dirname
	LogDirName string
	// Network log
	NetworkLog string
	SubscriberConfig model.SubscriberConfig
	Warn model.WarningSt
	Ovnnbdb string
	Ovnsbdb string
	BuildTime string
	GitHash string
	Route model.ToLeafRouter
	Az_ID string
	ForceSync bool
	OvnControllerURL string
	Monitor model.Monitor
}

var Configuration AgentConfig

func InitConfig(){
	initViper()
	if err := viper.Unmarshal(&Configuration); err != nil {
		fmt.Printf("Failed to parse configuration.Error msg +v", errors.WithStack(err))
	}
	go func(){
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("配置发生变更：", e.Name, e.Op)
			if err := viper.Unmarshal(&Configuration); err != nil {
				fmt.Printf("Failed to parse configuration.Error msg +v", errors.WithStack(err))
			}
			fmt.Println(Configuration)
		})
	}()
	data, _ := json.Marshal(Configuration)
	buffer := bytes.NewBuffer(data)
	viper.MergeConfig(buffer)
	viper.WriteConfig()
	var str string
	fmt.Scanf("%s", &str)
}

func SetWarnConfig(){

}

func Set

