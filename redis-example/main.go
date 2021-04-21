package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisAddr = flag.String("addr", "10.252.146.111:6379", "redis addr")
var passwd = flag.String("passwd", "unic-moove", "redis passwd")

func redisOptions() *redis.Options {
	return &redis.Options{
		Addr:     *redisAddr,
		DB:       15,
		Password: *passwd,

		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,

		//MaxRetries: -1,

		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func redisFailoverOptions(redisAddr, passwd string) *redis.FailoverOptions {
	return &redis.FailoverOptions{
		MasterName:         "mymaster",
		SentinelAddrs:      []string{redisAddr},
		OnConnect:          nil,
		Password:           passwd,
		DB:                 1,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 100 * time.Millisecond,
		TLSConfig:          nil,
	}
}

type SubRequest struct {
	Header string `json:"header"`
	Mapp   map[string]interface{}
}

type VM struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	flag.Parse()
	var client *redis.Client
	fmt.Printf("addr:%s , passwd:%s\n", *redisAddr, *passwd)
	opt := redisOptions()
	//opt := redisFailoverOptions(*redisAddr, *passwd)
	opt.MinIdleConns = 0
	opt.MaxConnAge = 0
	//opt.onconnect = func(cn *redis.conn) (err error) {
	//	clientid, err := cn.clientid().result()
	//	fmt.println("clientid:", clientid, err)
	//	return err
	//}
	client = redis.Newclient(opt)
	//client = redis.NewFailoverClient(opt)
	defer client.Close()
	fmt.Println("new client")

	// client Subscribe msg
}
