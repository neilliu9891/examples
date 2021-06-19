package model

// 告警功能
type WarningSt struct{
	Server_Addr string `json:"server_addr"`
	Server_URI string	`json:"server_uri"`
}

type ToLeafRouter struct{
	Cidr string `json:"cidr"`
	Gateway_Ip string `json:"gateway_ip"`
}

type SubscriberConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Channel  string `json:"channel"`
	Mode  string `json:"mode"`
}

type Monitor struct{
	Files []string
}