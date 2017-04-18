package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// ParseRemoteAddr parses remote address of the client from param. This address
// is used for establishing a connection with the client.
//
// Param Format: 192,168,150,80,14,178
// Host: 192.168.150.80
// Port: (14 * 256) + 148
func ParseRemoteAddr(param string) *net.TCPAddr {
	params := strings.Split(param, ",")
	ip := ""
	p1 := 0
	p2 := 0
	if len(params) > 5 {
		ip = strings.Join(params[0:4], ".")
		p1, _ = strconv.Atoi(params[4])
		p2, _ = strconv.Atoi(params[5])
	}

	port := (p1 * 256) + p2

	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	return addr
}
