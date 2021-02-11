package main

import (
	"encoding/binary"
	"fmt"
)

type IPMsg struct {
	IpLen uint32
	Ip    [4]byte
	//Ipv6  [6]byte
}

type RecvType2Msg struct {
	Mac [6]byte
	IPMsg
	L2vni       uint32
	L3vni       uint32
	IsGw        uint8
	VtepIp      [4]byte
	Rmac        [6]byte
	LocalBgpAs  uint32
	RemoteBgpAs uint32
}

func main() {
	fmt.Printf("%d", binary.Size(RecvType2Msg{}))
	fmt.Println("vim-go")
}