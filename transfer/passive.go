// +-------------------------------------------------------------------------
// | Copyright (C) 2017 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

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
