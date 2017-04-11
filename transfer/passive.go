package transfer

import (
	"net"
	"time"
)

// PassiveHandler handles passive connection.
type PassiveHandler struct {
	TCPListener *net.TCPListener // TCP Listener (only keeping it to define a deadline during the accept)
	Listener    net.Listener     // TCP or SSL Listener
	Port        int              // TCP Port we are listening on
	connection  net.Conn         // TCP Connection established
}

// Open opens connection.
func (p *PassiveHandler) Open() (net.Conn, error) {
	return p.ConnectionWait(time.Minute)
}

// Close only the client connection is not supported at that time.
func (p *PassiveHandler) Close() error {
	if p.TCPListener != nil {
		p.TCPListener.Close()
	}
	if p.connection != nil {
		p.connection.Close()
	}
	return nil
}

// ConnectionWait wait for connection time out
func (p *PassiveHandler) ConnectionWait(wait time.Duration) (net.Conn, error) {
	if p.connection == nil {
		p.TCPListener.SetDeadline(time.Now().Add(wait))
		var err error
		p.connection, err = p.Listener.Accept()

		if err != nil {
			return nil, err
		}
	}

	return p.connection, nil
}
