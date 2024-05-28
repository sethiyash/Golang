package main

import (
	"fmt"
	"net"
)

func isValidIP(ip string) bool {
	parsedIp := net.ParseIP(ip)
	return parsedIp != nil
}

func main() {
	testIps := []string{
		"192.168.1.1",
		"test",
	}
	for _, ip := range testIps {
		if isValidIP(ip) {
			fmt.Println("Valid IP")
		} else {
			fmt.Println("IP is invalid")
		}
	}
}
