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
	"errors"
	"io/ioutil"
	"strings"

	"github.com/pengsrc/go-shared/check"
	"github.com/pengsrc/go-shared/yaml"
	"github.com/yunify/qsftpd/utils"
)

// A Config stores a configuration of qsftp.
type Config struct {
	QingStor struct {
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		Protocol        string `yaml:"protocol"`
		LogLevel        string `yaml:"log_level"`
	} `yaml:"qingstor"`

	ListenHost     string `yaml:"listen_host"`
	ListenPort     int    `yaml:"listen_port"`
	PublicHost     string `yaml:"public_host"`
	MaxConnections int    `yaml:"max_connections"`
	BucketName     string `yaml:"bucket_name"`
	Zone           string `yaml:"zone"`

	Users map[string]string `yaml:"users"`
}

// NewConfig creates a new Config
func NewConfig() *Config {
	return &Config{}
}

// LoadConfigFromFilepath loads configuration from a specified local path.
// It returns error if file not found or yaml decode failed.
func (c *Config) LoadConfigFromFilepath(p string) error {
	if strings.Index(p, "~/") == 0 {
		p = strings.Replace(p, "~/", utils.GetHome()+"/", 1)
	}

	configYAML, err := ioutil.ReadFile(p)
	check.ErrorForExit("File not found: "+p, err)

	return c.LoadConfigFromContent(configYAML)
}

// LoadConfigFromContent loads configuration from a given byte slice.
// It returns error if yaml decode failed.
func (c *Config) LoadConfigFromContent(content []byte) error {
	d := NewConfig()
	_, err := yaml.Decode(content, d)

	check.ErrorForExit("Config parse error: ", err)

	*c = *d
	err = c.Check()
	check.ErrorForExit("Config check error: ", err)

	return nil
}

// Check checks the configuration.
func (c *Config) Check() error {

	if c.ListenHost == "" {
		c.ListenHost = "0.0.0.0"
	}
	if c.ListenPort == 0 {
		// For the default value (0), We take the default port (2121).
		c.ListenPort = 2121
	} else if c.ListenPort == -1 {
		// For the automatic value, We let the system decide (0).
		c.ListenPort = 0
	}
	if c.PublicHost == "" {
		c.PublicHost = "127.0.0.1"
	}
	if c.MaxConnections == 0 {
		c.MaxConnections = 10000
	}
	if c.BucketName == "" {
		return errors.New("Bucket name not specified")
	}
	if c.Zone == "" {
		return errors.New("Bucket zone not specified")
	}
	if c.Users == nil {
		c.Users = make(map[string]string)
		c.Users["anonymous"] = ""
	}

	return nil
}
