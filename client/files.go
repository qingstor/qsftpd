package client

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func (c *Handler) handleSTOR() {
	c.handleStoreAndAppend(false)
}

func (c *Handler) handleAPPE() {
	c.handleStoreAndAppend(true)
}

// Handles both the "STOR" and "APPE" commands.
func (c *Handler) handleStoreAndAppend(append bool) {

	path := c.absPath(c.param)

	if tr, err := c.TransferOpen(); err == nil {
		defer c.TransferClose()
		if _, err := c.storeOrAppend(tr, path, append); err != nil && err != io.EOF {
			c.WriteMessage(550, err.Error())
		}
	} else {
		c.WriteMessage(550, err.Error())
	}
}

func (c *Handler) handleRETR() {

	path := c.absPath(c.param)

	if tr, err := c.TransferOpen(); err == nil {
		defer c.TransferClose()
		if _, err := c.download(tr, path); err != nil && err != io.EOF {
			c.WriteMessage(550, err.Error())
		}
	} else {
		c.WriteMessage(550, err.Error())
	}
}

func (c *Handler) download(conn net.Conn, name string) (int64, error) {
	file, err := c.driver.OpenFile(c, name, os.O_RDONLY)

	if err != nil {
		return 0, err
	}

	if c.ctxRest != 0 {
		file.Seek(c.ctxRest, 0)
		c.ctxRest = 0
	}

	defer file.Close()
	return io.Copy(conn, file)
}

func (c *Handler) handleCHMOD(params string) {
	spl := strings.SplitN(params, " ", 2)
	modeNb, err := strconv.ParseUint(spl[0], 10, 32)

	mode := os.FileMode(modeNb)
	path := c.absPath(spl[1])

	if err == nil {
		err = c.driver.ChmodFile(c, path, mode)
	}

	if err != nil {
		c.WriteMessage(550, err.Error())
		return
	}

	c.WriteMessage(200, "SITE CHMOD command successful")
}

func (c *Handler) storeOrAppend(conn net.Conn, name string, append bool) (int64, error) {
	flag := os.O_WRONLY
	if append {
		flag |= os.O_APPEND
	}

	file, err := c.driver.OpenFile(c, name, flag)

	if err != nil {
		return 0, err
	}

	if c.ctxRest != 0 {
		file.Seek(c.ctxRest, 0)
		c.ctxRest = 0
	}

	defer file.Close()
	return io.Copy(file, conn)
}

func (c *Handler) handleDELE() {
	path := c.absPath(c.param)
	if err := c.driver.DeleteFile(c, path); err == nil {
		c.WriteMessage(250, fmt.Sprintf("Removed file %s", path))
	} else {
		c.WriteMessage(550, fmt.Sprintf("Couldn't delete %s: %v", path, err))
	}
}

func (c *Handler) handleRNFR() {
	path := c.absPath(c.param)
	if _, err := c.driver.GetFileInfo(c, path); err == nil {
		c.WriteMessage(350, "Sure, give me a target")
		c.ctxRnfr = path
	} else {
		c.WriteMessage(550, fmt.Sprintf("Couldn't access %s: %v", path, err))
	}
}

func (c *Handler) handleRNTO() {
	dst := c.absPath(c.param)
	if c.ctxRnfr != "" {
		if err := c.driver.RenameFile(c, c.ctxRnfr, dst); err == nil {
			c.WriteMessage(250, "Done !")
			c.ctxRnfr = ""
		} else {
			c.WriteMessage(550, fmt.Sprintf("Couldn't rename %s to %s: %s", c.ctxRnfr, dst, err.Error()))
		}
	}
}

func (c *Handler) handleSIZE() {
	path := c.absPath(c.param)
	if info, err := c.driver.GetFileInfo(c, path); err == nil {
		c.WriteMessage(213, fmt.Sprintf("%d", info.Size()))
	} else {
		c.WriteMessage(550, fmt.Sprintf("Couldn't access %s: %v", path, err))
	}
}

func (c *Handler) handleSTATFile() {
	path := c.absPath(c.param)

	c.writeLine("213-Status follows:")
	if info, err := c.driver.GetFileInfo(c, path); err == nil {
		if info.IsDir() {
			if files, err := c.driver.ListFiles(c, path); err == nil {
				for _, f := range files {
					c.writeLine(fileStat(f))
				}
			}
		} else {
			c.writeLine(fileStat(info))
		}
	}
	c.writeLine("213 End of status")
}

func (c *Handler) handleALLO() {
	// We should probably add a method in the driver
	if size, err := strconv.Atoi(c.param); err == nil {
		if ok, err := c.driver.CanAllocate(c, size); err == nil {
			if ok {
				c.WriteMessage(202, "OK, we have the free space")
			} else {
				c.WriteMessage(550, "NOT OK, we don't have the free space")
			}
		} else {
			c.WriteMessage(500, fmt.Sprintf("Driver issue: %v", err))
		}
	} else {
		c.WriteMessage(501, fmt.Sprintf("Couldn't parse size: %v", err))
	}
}

func (c *Handler) handleREST() {
	if size, err := strconv.ParseInt(c.param, 10, 0); err == nil {
		c.ctxRest = size
		c.WriteMessage(350, "OK")
	} else {
		c.WriteMessage(550, fmt.Sprintf("Couldn't parse size: %v", err))
	}
}

func (c *Handler) handleMDTM() {
	path := c.absPath(c.param)
	if info, err := c.driver.GetFileInfo(c, path); err == nil {
		c.WriteMessage(250, info.ModTime().UTC().Format("20060102150405"))
	} else {
		c.WriteMessage(550, fmt.Sprintf("Couldn't access %s: %s", path, err.Error()))
	}
}
