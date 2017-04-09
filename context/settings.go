package context

// Settings define all the server settings.
type ServerSettings struct {
	ListenHost     string     // Host to receive connections on
	ListenPort     int        // Port to listen on
	PublicHost     string     // Public IP to expose (only an IP address is accepted at this stage)
	MaxConnections int        // Max number of connections to accept
	DataPortRange  *PortRange // Port Range for data connections. Random one will be used if not specified
	Users          map[string]string
}

// PortRange is a range of ports.
type PortRange struct {
	Start int // Range start
	End   int // Range end
}
