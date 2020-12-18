package main

import (
	"fmt"
	"publisher/api"
	"publisher/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println("error")
	}
	services.InitRedis(viper.GetString("host"), viper.GetString("password"), viper.GetInt("db_index"))
	r.POST("/publisher", api.ApiPublisher)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
