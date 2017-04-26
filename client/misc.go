package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/yunify/qsftp/context"
)

func (c *Handler) handleAUTH() {
	c.WriteMessage(550, "Cannot get a TLS config")

	//FIXME: AUTH Has not been properly handled
	//c.WriteMessage(234, "AUTH command ok. Expecting TLS Negotiation.")
	//c.conn = tls.Server(c.conn, tlsConfig)
	//c.reader = bufio.NewReader(c.conn)
	//c.writer = bufio.NewWriter(c.conn)
}

func (c *Handler) handlePROT() {
	// P for Private, C for Clear
	c.transferTLS = c.param == "P"
	c.WriteMessage(200, "OK")
}

func (c *Handler) handlePBSZ() {
	c.WriteMessage(200, "Whatever")
}

func (c *Handler) handleSYST() {
	c.WriteMessage(215, "UNIX Type: L8")
}

func (c *Handler) handleSTAT() {
	// STAT is a bit tricky.

	if c.param == "" { // Without a file, it's the server stat.
		c.handleSTATServer()
	} else { // With a file/dir it's the file or the dir's files stat.
		c.handleSTATFile()
	}
}

func (c *Handler) handleSITE() {
	spl := strings.SplitN(c.param, " ", 2)
	if len(spl) > 1 {
		if strings.ToUpper(spl[0]) == "CHMOD" {
			c.handleCHMOD(spl[1])
		}
	}
}

func (c *Handler) handleSTATServer() {
	c.writeLine("213- FTP server status:")
	duration := time.Now().UTC().Sub(c.connectedAt)
	duration -= duration % time.Second
	c.writeLine(fmt.Sprintf(
		"Connected to %s:%d from %s for %s",
		context.Settings.ListenHost, context.Settings.ListenPort,
		c.conn.RemoteAddr(),
		duration,
	))
	if c.user != "" {
		c.writeLine(fmt.Sprintf("Logged in as %s", c.user))
	} else {
		c.writeLine("Not logged in yet")
	}
	c.writeLine("ftpserver - golang FTP server")
	defer c.WriteMessage(213, "End")
}

func (c *Handler) handleOPTS() {
	args := strings.SplitN(c.param, " ", 2)
	if strings.ToUpper(args[0]) == "UTF8" {
		c.WriteMessage(200, "I'm in UTF8 only anyway")
	} else {
		c.WriteMessage(500, "Don't know this option")
	}
}

func (c *Handler) handleNOOP() {
	c.WriteMessage(200, "OK")
}

func (c *Handler) handleFEAT() {
	c.writeLine("211- These are my features")
	defer c.WriteMessage(211, "End")

	features := []string{
		"UTF8",
		"SIZE",
		"MDTM",
		"REST STREAM",
	}

	for _, f := range features {
		c.writeLine(" " + f)
	}
}

func (c *Handler) handleTYPE() {
	switch c.param {
	case "I":
		c.WriteMessage(200, "Type set to binary")
	case "A":
		c.WriteMessage(200, "WARNING: ASCII isn't correctly supported")
	default:
		c.WriteMessage(500, "Not understood")
	}
}

func (c *Handler) handleQUIT() {
	c.WriteMessage(221, "Goodbye")
	c.disconnect()
	c.reader = nil
}
