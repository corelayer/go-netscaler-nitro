/*
 * Copyright 2022 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package config

import (
	"encoding/base64"
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"io"
	"net/url"
	"strings"
)

type SystemBackupController struct {
	client service.NitroClient
	name   string
	level  SystemBackupLevel
}

func NewBackupController(client service.NitroClient, name string, level SystemBackupLevel) *SystemBackupController {
	c := SystemBackupController{
		client: service.NitroClient{},
		name:   name,
		level:  level,
	}

	return &c
}

// Create sends a request to the configured NitroClient to create a backup with the configured name and level
func (c *SystemBackupController) Create() error {
	// Filename must not have a filename extension
	data := system.Systembackup{
		Filename: strings.TrimSuffix(c.name, ".tgz"),
		Level:    c.level.String(),
	}

	err := c.client.ActOnResource(service.Systembackup.Type(), data, "create")
	return err
}

func (c *SystemBackupController) Get() (io.Reader, error) {
	params := service.FindParams{
		ArgsMap:                  map[string]string{"fileLocation": url.PathEscape("/var/ns_sys_backup")},
		ResourceType:             "systemfile",
		ResourceName:             c.name,
		ResourceMissingErrorCode: 0,
	}

	response, err := c.client.FindResourceArrayWithParams(params)
	var data string
	if err == nil {
		if response[0]["filecontent"] != "" {
			data = response[0]["filecontent"].(string)
		} else {
			data = ""
		}
	}

	output := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	return output, err
}

func (c *SystemBackupController) Delete() error {
	err := c.client.DeleteResource(service.Systembackup.Type(), c.name)
	return err
}
