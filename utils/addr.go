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
