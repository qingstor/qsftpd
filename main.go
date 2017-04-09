package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pengsrc/go-utils/check"
	"github.com/pengsrc/qsftp/context"
	"github.com/pengsrc/qsftp/server"
)

func main() {
	err := context.SetupContext()
	check.ErrorForExit("qsftp", err)

	ftpServer := server.NewFTPServer()
	go signalHandler(ftpServer)

	err = ftpServer.ListenAndServe()
	check.ErrorForExit("qsftp", err)
}

func signalHandler(ftpServer *server.FTPServer) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	for {
		switch <-ch {
		case syscall.SIGTERM:
			ftpServer.Stop()
			break
		}
	}
}
