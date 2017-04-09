package transfer

import "net"

// Handler presents active/passive transfer connection handler.
type Handler interface {
	// Get the connection to transfer data on.
	Open() (net.Conn, error)

	// Close the connection (and any associated resource).
	Close() error
}
