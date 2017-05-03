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

package client

import (
	"github.com/yunify/qsftpd/context"
)

// Handle the "USER" command.
func (c *Handler) handleUSER() {
	c.user = c.param
	_, ok := context.Settings.Users[c.user]
	if ok {
		c.WriteMessage(331, "User name okay, need password.")
	} else {
		c.WriteMessage(430, "Invalid username or password")
	}
}

// Handle the "PASS" command.
func (c *Handler) handlePASS() {
	c.driver = c.driverFactory()

	username := c.user
	password := c.param

	if username == "anonymous" {
		// User can log in as anonymous with any password
		_, ok := context.Settings.Users["anonymous"]
		if ok {
			c.WriteMessage(230, "Password ok, continue")
		} else {
			c.WriteMessage(430, "Invalid username or password")
		}
	} else {
		if password == context.Settings.Users[username] {
			c.WriteMessage(230, "Password ok, continue")
		} else {
			c.WriteMessage(430, "Invalid username or password")
		}
	}

}
