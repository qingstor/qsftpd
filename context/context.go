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

package context

import (
	"github.com/pengsrc/go-shared/check"
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
		PublicHost:     c.PublicHost,
		MaxConnections: c.MaxConnections,
		DataPortRange: &PortRange{
			Start: c.StartPort,
			End:   c.EndPort,
		},
		Users: c.Users,
	}

	// Setup bucket.
	curQingStorConfig, err := qsConfig.NewDefault()
	if err != nil {
		return err
	}

	curData, err := yaml.Encode(c.QingStor)
	check.ErrorForExit("QingStor service settings encode error", err)

	err = curQingStorConfig.LoadConfigFromContent(curData)
	check.ErrorForExit("Load QingStor service settings error", err)

	qsService, err := service.Init(curQingStorConfig)
	check.ErrorForExit("Init QingStor config error", err)

	Bucket, err = qsService.Bucket(c.BucketName, c.Zone)
	check.ErrorForExit("Create QingStor bucket error", err)

	return nil
}
