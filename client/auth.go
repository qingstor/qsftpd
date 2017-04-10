package client

import "github.com/yunify/qsftp/context"

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
