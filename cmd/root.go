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

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pengsrc/go-shared/check"
	"github.com/spf13/cobra"
	"github.com/yunify/qsftpd/constants"
	"github.com/yunify/qsftpd/context"
	"github.com/yunify/qsftpd/server"
)

var (
	versionFlag   bool
	cfgFileFlag   string
	currentConfig *context.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   constants.Name,
	Short: "A FTP server that persists all data to QingStor Object Storage.",
	Long:  "A FTP server that persists all data to QingStor Object Storage.",
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Fprintf(os.Stdout, "qsftpd version %s\n", constants.Version)
			return
		}

		loadConfig()

		ftpServer := server.NewFTPServer()
		go signalHandler(ftpServer)

		err := ftpServer.ListenAndServe()
		check.ErrorForExit(constants.Name, err)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		check.ErrorForExit(constants.Name, err)
		os.Exit(-1)
	}
}

func init() {
	currentConfig = context.NewConfig()

	RootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Show version")
	RootCmd.PersistentFlags().StringVarP(&cfgFileFlag, "config", "c", "qsftpd.yaml", "Specify config file")
}

func loadConfig() {
	err := currentConfig.LoadConfigFromFilepath(cfgFileFlag)
	check.ErrorForExit(constants.Name, err)

	err = context.SetupContext(currentConfig)
	check.ErrorForExit(constants.Name, err)
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
