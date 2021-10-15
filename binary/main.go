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

const GetGuestPing = "{\"execute\":\"guest-ping\"}"

func main() {
	fmt.Printf("%d", binary.Size(RecvType2Msg{}))
	fmt.Println("vim-go")
	instanceId := "12333"
	getQgaStatusCmd := fmt.Sprintf("virsh qemu-agent-command %s '%s'", instanceId, GetGuestPing)
	fmt.Println(getQgaStatusCmd)
	fmt.Println([]byte("vim-go"))
}
