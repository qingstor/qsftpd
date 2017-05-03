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
