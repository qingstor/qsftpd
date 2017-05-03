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

// Package server provides all the tools to build your own FTP server: The core library and the driver.
package server

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/satori/go.uuid"
	"github.com/yunify/qsftpd/client"
	"github.com/yunify/qsftpd/context"
	"github.com/yunify/qsftpd/driver"
)

// FTPServer is where everything is stored.
// We want to keep it as simple as possible.
type FTPServer struct {
	Listener  net.Listener // Listener used to receive files
	StartTime time.Time    // Time when the s was started

	connectionsMutex sync.RWMutex // Connections map sync
	clientCounter    int          // Clients counter
}

// ListenAndServe simply chains the Listen and Serve method calls.
func (s *FTPServer) ListenAndServe() error {
	if err := s.Listen(); err != nil {
		return err
	}

	context.Logger.Infof("Starting...")
	s.Serve()

	// Note: At this precise time, the clients are still connected. We are just not accepting clients anymore.
	return nil
}

// Listen starts the listening. It's not a blocking call.
func (s *FTPServer) Listen() error {
	var err error

	s.Listener, err = net.Listen("tcp", fmt.Sprintf(
		"%s:%d", context.Settings.ListenHost, context.Settings.ListenPort,
	))
	if err != nil {
		context.Logger.Errorf("Cannot listen: %v", err)
		return err
	}

	context.Logger.Infof("Listening... %v", s.Listener.Addr())
	return err
}

// Serve accepts and process any new client coming.
func (s *FTPServer) Serve() {
	for {
		connection, err := s.Listener.Accept()
		if err != nil {
			if s.Listener != nil {
				context.Logger.Errorf("Accept error: %v", err)
			}
			break
		}

		s.connectionsMutex.Lock()

		id := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		go s.serveClient(id, connection)

		s.connectionsMutex.Unlock()
	}
}

// Stop closes the listener.
func (s *FTPServer) Stop() {
	if s.Listener != nil {
		l := s.Listener
		s.Listener = nil
		l.Close()
	}
}

func (s *FTPServer) serveClient(id string, connection net.Conn) {
	c := client.NewHandler(id, connection, func() client.Driver {
		return &driver.QSDriver{}
	})

	if err := s.clientArrival(id, connection); err != nil {
		c.WriteMessage(500, fmt.Sprintf("Can't accept you - %v", err.Error()))
		return
	}
	defer s.clientDeparture(id, connection)

	c.WriteMessage(220, "Welcome to QSFTP Server")
	context.Logger.Debugf("Accept client on: id: %s, IP: %v", id, connection.RemoteAddr())
	defer context.Logger.Debugf("Goodbye: id: %s, IP: %v", id, connection.RemoteAddr())

	c.HandleCommands()
}

// When a client connects, the s could refuse the connection.
func (s *FTPServer) clientArrival(id string, connection net.Conn) error {
	s.connectionsMutex.Lock()
	defer s.connectionsMutex.Unlock()

	if s.clientCounter+1 > context.Settings.MaxConnections {
		return fmt.Errorf("Too many clients %d > %d", s.clientCounter+1, context.Settings.MaxConnections)
	}

	s.clientCounter++
	context.Logger.Infof("FTP Client connected: ftp.connected, id: %s, RemoteAddr: %v, Total: %d", id, connection.RemoteAddr(), s.clientCounter)

	return nil
}

// When a client leaves.
func (s *FTPServer) clientDeparture(id string, connection net.Conn) {
	s.connectionsMutex.Lock()
	defer s.connectionsMutex.Unlock()

	s.clientCounter--
	context.Logger.Infof("FTP Client disconnected: ftp.disconnected, id: %s, RemoteAddr: %v, Total: %d", id, connection.RemoteAddr(), s.clientCounter)
}

// NewFTPServer creates a new FTPServer instance.
func NewFTPServer() *FTPServer {
	return &FTPServer{
		StartTime: time.Now().UTC(),
	}
}
