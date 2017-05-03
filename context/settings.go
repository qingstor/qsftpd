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

package context

// ServerSettings define all the server settings.
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
