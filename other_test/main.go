package main

import "fmt"

const IPv4Prefix = "169.254."

func main() {
	fmt.Println("vim-go")
	vtepIP := "192.165.1.1"

	ip := IPv4Prefix + vtepIP[len(IPv4Prefix):]
	fmt.Println(ip)
}
