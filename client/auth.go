package client

// Handle the "USER" command.
func (c *Handler) handleUSER() {
	c.user = c.param
	c.WriteMessage(331, "OK")
}

// Handle the "PASS" command.
func (c *Handler) handlePASS() {
	c.driver = c.driverFactory()

	c.WriteMessage(230, "Password ok, continue")

	//c.WriteMessage(530, fmt.Sprintf("Authentication problem: %v", err))
	//c.disconnect()

	//c.WriteMessage(530, "I can't deal with you (nil driver)")
	//c.disconnect()
}
