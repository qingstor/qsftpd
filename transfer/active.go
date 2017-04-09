package transfer

import (
	"fmt"
	"net"
)

// ActiveHandler handles active connection.
type ActiveHandler struct {
	RemoteAddr *net.TCPAddr // remote address of the client

	conn net.Conn
}

// Open opens connection.
func (a *ActiveHandler) Open() (net.Conn, error) {
	localAddr, _ := net.ResolveTCPAddr("tcp", ":20")

	// TODO: Support dialing with timeout
	// Issues:
	//	https://github.com/golang/go/issues/3097
	// 	https://github.com/golang/go/issues/4842
	conn, err := net.DialTCP("tcp", localAddr, a.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("Could not establish active connection due: %v", err)
	}

	// Keep connection as it will be closed by Close().
	a.conn = conn

	return a.conn, nil
}

// Close closes only if connection is established.
func (a *ActiveHandler) Close() error {
	if a.conn != nil {
		return a.conn.Close()
	}
	return nil
}
