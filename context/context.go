package context

import (
	"github.com/pengsrc/go-shared/logger"
	"github.com/pengsrc/go-shared/yaml"
	qsConfig "github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/service"
)

var (
	// Logger is the global logger for qsftp
	Logger *logger.Logger
	// Settings is the global settings for qsftp
	Settings *ServerSettings
	// Bucket is the global Bucket for qsftp
	Bucket *service.Bucket
)

// SetupContext creates the server context.
func SetupContext(c *Config) error {
	var err error

	// Setup logger.
	Logger, err = logger.NewTerminalLogger("debug")
	if err != nil {
		return err
	}

	// Setup settings.
	Settings = &ServerSettings{
		ListenHost:     c.ListenHost,
		ListenPort:     c.ListenPort,
		PublicHost:     c.ListenHost,
		MaxConnections: c.MaxConnections,
		DataPortRange: &PortRange{
			Start: 6000,
			End:   7000,
		},
		Users: c.Users,
	}

	// Setup bucket.
	curQingStorConfig, err := qsConfig.NewDefault()
	if err != nil {
		return err
	}

	curData, err := yaml.Encode(c.QingStor)
	if err != nil {
		return err
	}

	err = curQingStorConfig.LoadConfigFromContent(curData)
	if err != nil {
		return err
	}

	qsService, err := service.Init(curQingStorConfig)
	if err != nil {
		return err
	}
	Bucket, err = qsService.Bucket(c.BucketName, c.Zone)
	if err != nil {
		return err
	}

	return nil
}
