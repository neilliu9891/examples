package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
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

type VMUpdate struct {
	op int `json:op`
}

func main() {
	flag.Parse()
	var client *redis.Client
	fmt.Printf("addr:%s , passwd:%s\n", *redisAddr, *passwd)
	opt := redisOptions()
	opt.MinIdleConns = 0
	opt.MaxConnAge = 0
	//opt.OnConnect = func(cn *redis.Conn) (err error) {
	//	clientID, err := cn.ClientID().Result()
	//	fmt.Println("clientID:", clientID, err)
	//	return err
	//}
	client = redis.NewClient(opt)
	defer client.Close()
	fmt.Println("new client")

	// client Subscribe msg
	pubsub := client.Subscribe("mychannel")
	defer pubsub.Close()

	fmt.Println("subscribe mychannel")
	for {
		//fmt.Println("start to selete")
		select {
		case msg := <-pubsub.Channel():
			sp := strings.SplitN(msg.Payload, ":", 2)
			//fmt.Println(sp)
			ot, _ := strconv.ParseInt(sp[0], 10, 64)
			nt := time.Now().UnixNano()
			fmt.Printf("time:%d ms, channel:%s, pattern:%s, payload:%s\n", (nt-ot)/1000/1000, msg.Channel, msg.Pattern, msg.Payload)
			//fmt.Println(ot)
			//fmt.Println(nt)
			if msg.Payload == "close" {
				return
			}
		}
	}
	//client := redis.NewClient(&redis.Options{
	//	Addr:     redisAddr,
	//	Password: passwd,
	//	DB:       15,
	//})
	//
	//pubsub := client.PSubscribe("*")
	//defer pubsub.Close()
	//for msg := range pubsub.Channel() {
	//	fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
	//}

	//fmt.Println("vim-go")
}
