package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := fmt.Sprintf("%s:%s", domain, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] Domain %s is down", domain)
	} else {
		status = fmt.Sprintf("[UP] Domain %s is reachable,\nFrom: %s\nTo: %s", domain, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
