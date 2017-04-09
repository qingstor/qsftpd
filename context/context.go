package context

import (
	"github.com/pengsrc/go-utils/logger"
	"github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/service"
)

var (
	Logger   *logger.Logger
	Settings *ServerSettings
	Bucket   *service.Bucket
)

// SetupContext creates the server context.
func SetupContext() error {
	var err error

	// Setup logger.
	Logger, err = logger.NewTerminalLogger("debug")
	if err != nil {
		return err
	}

	// Setup settings.
	Settings = &ServerSettings{
		ListenHost:     "127.0.0.1",
		ListenPort:     2121,
		PublicHost:     "127.0.0.1",
		MaxConnections: 128,
		DataPortRange: &PortRange{
			Start: 6000,
			End:   7000,
		},
	}

	if Settings.ListenHost == "" {
		Settings.ListenHost = "0.0.0.0"
	}

	if Settings.ListenPort == 0 { // For the default value (0)
		// We take the default port (2121).
		Settings.ListenPort = 2121
	} else if Settings.ListenPort == -1 { // For the automatic value
		// We let the system decide (0).
		Settings.ListenPort = 0
	}
	if Settings.MaxConnections == 0 {
		Settings.MaxConnections = 10000
	}

	// Setup bucket.
	c, err := config.NewDefault()
	if err != nil {
		return err
	}
	err = c.LoadUserConfig()
	if err != nil {
		return err
	}

	qsService, err := service.Init(c)
	if err != nil {
		return err
	}
	Bucket, err = qsService.Bucket("example", "pek3a")
	if err != nil {
		return err
	}

	return nil
}
