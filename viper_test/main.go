package main

import (
	"time"
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

func initViper() {
	viper.SetConfigName("config") // name of config file (without extension)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("error (%v)", err)
	}
}

var wg sync.WaitGroup

func main() {
	initViper()

	azId := viper.GetInt("az_id")
	fmt.Printf("get azId(%d)", azId)
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			d := viper.GetInt("az_id")
			fmt.Printf("azid.%d\n", d)
			viper.Set("az_id", d+1)
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			d := viper.GetInt("az_id")
			fmt.Printf("azid..%d\n", d)
			viper.Set("az_id", d+1)
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}()

	wg.Wait()
}
