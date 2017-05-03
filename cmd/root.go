package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pengsrc/go-shared/check"
	"github.com/spf13/cobra"
	"github.com/yunify/qsftpd/context"
	"github.com/yunify/qsftpd/server"
)

var (
	cfgFile   string
	curConfig *context.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "qsftpd",
	Short: "A FTP server that persists all data to QingStor Object Storage.",
	Long:  "A FTP server that persists all data to QingStor Object Storage.",
	Run: func(cmd *cobra.Command, args []string) {
		reloadConfig()

		ftpServer := server.NewFTPServer()
		go signalHandler(ftpServer)

		err := ftpServer.ListenAndServe()
		check.ErrorForExit("qsftpd", err)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		check.ErrorForExit("qsftpd", err)
		os.Exit(-1)
	}
}

func init() {
	curConfig = context.NewConfig()

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "qsftpd.yaml", "Specify config file")
}

func reloadConfig() {
	err := curConfig.LoadConfigFromFilepath(cfgFile)
	check.ErrorForExit("qsftpd", err)

	err = context.SetupContext(curConfig)
	check.ErrorForExit("qsftpd", err)
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
