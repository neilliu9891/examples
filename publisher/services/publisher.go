package services

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisAddr = flag.String("addr", "10.254.7.1:6379", "redis addr")
var passwd = flag.String("passwd", "Moove", "redis passwd")

//var redisAddr = flag.String("addr", "10.252.146.111:6379", "redis addr")
//var passwd = flag.String("passwd", "unic-moove", "redis passwd")

var client *redis.Client

var vmupdate = `{"header":{"send_time":"11111", "request_id":"22222","operation":"addPort"},"overlayport":{"nic_name":"vnet5","nic_ip":"172.16.255.130","nic_mac":"02:ac:10:ff:01:30","nic_index":"3","cvk_ip":"10.254.100.10","l3vni":"100","l2vni":"10","subnet_gw_ip":"172.16.255.129","subnet_mask":"255.255.255.0","subnet_gw_mac":"02:ac:10:ff:01:29","Reserved_vni":"11","bgp_as_number":"65000"}}`

func InitRedis(host, passwd string, db_index int) {
	fmt.Println(host, passwd, db_index)
	opt := &redis.Options{
		Addr:     host,
		DB:       db_index,
		Password: passwd,

		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,

		//MaxRetries: -1,

		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
	opt.MinIdleConns = 0
	opt.MaxConnAge = 0

	client = redis.NewClient(opt)
}

func Publish(channel, msg string) (int64, error) {
	fmt.Println(channel, msg)
	return client.Publish(channel, msg).Result()
}
