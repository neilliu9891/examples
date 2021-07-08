package main

import "fmt"

// 选项模式为了解决可选参数的问题，python可以指定不定参数且支持不通类型和默认值，golang需要通过option 模式进行处理
// 但是option模式会增加代码量，去方便调用者

type Connection struct{
	IP string
	Port int
	Retry bool
	Descript string
}

func New(ip string, port int)*Connection{
	return &Connection{
		IP:       ip,
		Port:     port,
		Retry:    false,
		Descript: "",
	}
}

type Option func(c *Connection)

func SetRetry(retry bool)Option{
	return func(c *Connection){
		c.Retry = retry
	}
}

func SetDescript(des string)Option{
	return func(c *Connection){
		c.Descript = des
	}
}

func (c *Connection) AddOptions(ops ...Option) {
	for _, f := range ops {
		f(c)
	}
}


func main(){
	conn := New("localhost", 1024)
	conn.AddOptions(SetRetry(true), SetDescript("description"))

	fmt.Println(conn)
}
