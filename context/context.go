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
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
	"github.com/pengsrc/go-shared/check"
	"github.com/pengsrc/go-shared/log"

	qsConfig "github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/service"
)

var (
	// Logger is the global logger for qsftp
	Logger *log.ContextFreeLogger
	// Settings is the global settings for qsftp
	Settings *ServerSettings
	// Bucket is the global Bucket for qsftp
	Bucket *service.Bucket
)

//GetZone gets the zone from global
func GetZone(c *Config) (string, error) {
	url := fmt.Sprintf("%s://%s.%s:%d", c.QingStor.Protocol,
		c.BucketName,
		c.QingStor.Host,
		c.QingStor.Port)

	response, err := http.Head(url)
	check.ErrorForExit("Request error", err)
	if err != nil {
		return "", err
	}

	//URL for example: https://bucket.zone.example.com
	newURL := response.Request.URL.String()

	return strings.Split(newURL, ".")[1], nil
}

// SetupContext creates the server context.
func SetupContext(c *Config) error {
	var err error

	// Setup logger.
	l, err := log.NewTerminalLogger(c.LogLevel)
	if err != nil {
		return err
	}
	Logger = log.NewContextFreeLogger(l)

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
		CachePath: c.CachePath,
		Users:     c.Users,
	}

	// Setup bucket.
	curQingStorConfig, err := qsConfig.NewDefault()
	if err != nil {
		return err
	}

	curData, err := yaml.Marshal(c.QingStor)
	check.ErrorForExit("QingStor service settings encode error", err)

	err = curQingStorConfig.LoadConfigFromContent(curData)
	check.ErrorForExit("Load QingStor service settings error", err)

	qsService, err := service.Init(curQingStorConfig)
	check.ErrorForExit("Init QingStor config error", err)

	zone := c.Zone

	if zone == "" {
		zone, err = GetZone(c)
		check.ErrorForExit("Get zone error", err)
	}

	Bucket, err = qsService.Bucket(c.BucketName, zone)
	check.ErrorForExit("Create QingStor bucket error", err)

	return nil
}
