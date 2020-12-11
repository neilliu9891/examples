package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

//var redisAddr = flag.String("addr", "10.254.7.1:6379", "redis addr")
//var passwd = flag.String("passwd", "Moove", "redis passwd")
var redisAddr = flag.String("addr", "10.252.146.111:6379", "redis addr")
var passwd = flag.String("passwd", "unic-moove", "redis passwd")

func redisOptions() *redis.Options {
	return &redis.Options{
		Addr:     *redisAddr,
		DB:       15,
		Password: *passwd,

		//DialTimeout:  10 * time.Second,
		//ReadTimeout:  30 * time.Second,
		//WriteTimeout: 30 * time.Second,
		//
		////MaxRetries: -1,
		//
		//PoolSize:           10,
		//PoolTimeout:        30 * time.Second,
		//IdleTimeout:        time.Minute,
		//IdleCheckFrequency: 100 * time.Millisecond,
	}
}

var vmupdate = `{"nic_name":"vnet5","nic_ip":"172.16.255.130","nic_mac":"02:ac:10:ff:01:30","nic_index":"3","cvk_ip":"10.254.100.10","l3vni":"100","l2vni":"10","subnet_gw_ip":"172.16.255.129","subnet_mask":"255.255.255.0","subnet_gw_mac":"02:ac:10:ff:01:29","Reserved_vni":"11","bgp_as_number":"65000"}`

func main() {
	flag.Parse()
	var client *redis.Client
	opt := redisOptions()
	//opt.MinIdleConns = 0
	//opt.MaxConnAge = 0
	//opt.OnConnect = func(cn *redis.Conn) (err error) {
	//	clientID, err := cn.ClientID().Result()
	//	fmt.Println("clientID:", clientID, err)
	//	return err
	//}
	fmt.Printf("addr: %s, passwd: %s \n", *redisAddr, *passwd)
	client = redis.NewClient(opt)
	defer client.Close()
	fmt.Println("new client")

	// client Subscribe msg
 	for i:=0; i< 1000; i++{
		msg := strconv.FormatInt(time.Now().UnixNano(), 10) + ":" + vmupdate

		if num, err := client.Publish("mychannel", msg).Result(); err == nil {
			fmt.Println("publish num:", num)
		}
		time.Sleep(time.Duration(time.Millisecond))
	}


	//defer pubsub.Close()

	//fmt.Println("publish mychannel")
	//for {
	////fmt.Println("start to selete")
	//select {
	//case msg := <-pubsub.Channel():
	//fmt.Printf("channel:%s, pattern:%s, payload:%s timers:%s\n", msg.Channel, msg.Pattern, msg.Payload, time.Now().String())
	//if msg.Payload == "close" {
	//return
	//}
	//}
	//}
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
