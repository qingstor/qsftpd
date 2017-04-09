package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/pengsrc/go-utils/check"
	"github.com/pengsrc/qsftp/context"
	"github.com/pengsrc/qsftp/server"
	"github.com/spf13/cobra"
)

var (
	cfgFile   string
	curConfig *context.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "qsftp",
	Short: "A ftp server that uses QingStor Object Storage as storage backend.",
	Long:  `A ftp server that uses QingStor Object Storage as storage backend.`,
	Run: func(cmd *cobra.Command, args []string) {
		reloadConfig()
		curConfig.WatchConfig(cfgFile, func(e fsnotify.Event) {
			reloadConfig()
			context.Logger.Debug("Config file reloaded.")
		})

		ftpServer := server.NewFTPServer()
		go signalHandler(ftpServer)

		err := ftpServer.ListenAndServe()
		check.ErrorForExit("qsftp", err)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		check.ErrorForExit("qsftp", err)
		os.Exit(-1)
	}
}

func init() {
	curConfig = context.NewConfig()

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "qsftp.yaml", "Specify config file")
}

func reloadConfig() {
	err := curConfig.LoadConfigFromFilepath(cfgFile)
	check.ErrorForExit("qsftp", err)

	err = context.SetupContext(curConfig)
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
