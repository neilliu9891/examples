package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OverlayPort struct {
	L3vni             string `json:"l3vni" mapstructure:"l3vni"`
	L2vni             string `json:"l2vni" mapstructure:"l2vni"`
	SubNetGwMac       string `json:"subnet_gw_mac" mapstructure:"subnet_gw_mac"`
	SubNetGwIp        string `json:"subnet_gw_ip" mapstructure:"subnet_gw_ip"`
	SubNetGwNetmask   string `json:"subnet_mask" mapstructure:"subnet_mask"`
	SubNetGwIpv6      string `json:"subnet_gw_ipv6" mapstructure:"subnet_gw_ipv6"`
	SubNetGwNetmaskv6 string `json:"subnet_maskv6" mapstructure:"subnet_maskv6"`
	VmLportKey        string `json:"nic_index" mapstructure:"nic_index"`
	VmNet             string `json:"nic_name" mapstructure:"nic_name"`
	VmMac             string `json:"nic_mac" mapstructure:"nic_mac"`
	VmIp              string `json:"nic_ip" mapstructure:"nic_ip"`
	VmIpv6            string `json:"nic_ipv6" mapstructure:"nic_ipv6"`
	VmInCvkIp         string `json:"nic_cvk_ip" mapstructure:"nic_cvk_ip"`
	SelfCvkIp         string `json:"self_cvk_ip" mapstructure:"self_cvk_ip"`
	LocalBgpAs        string `json:"bgp_as_number" mapstructure:"bgp_as_number"`
	Azid              string `json:"rt" mapstructure:"rt"`
	OnlineType        string `json:"online_type" mapstructure:"online_type"` // 1:主机overlay 2:网络overlay
}

func main() {
	fmt.Println("vim-go")
	url := "http://10.254.4.6:22222/v1/network/vpc/addport"
	body := OverlayPort{L3vni: "123456"}
	bbody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("marshal body is error (%v)", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bbody))
	if err != nil {
		fmt.Printf("%v", err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("recieve code status is error %d", resp.StatusCode)
	}
}
