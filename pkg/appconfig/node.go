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

package appconfig

import "github.com/corelayer/go-netscaler-nitro/pkg/client"

type Node struct {
	Name    string `json:"name" yaml:"name"`
	Address string `json:"address" yaml:"address"`
}

func (n Node) GetNodeUrl(scheme UrlSchemeReader) string {
	return scheme.GetUrlScheme() + n.Address
}

func (n Node) GetNitroSettings(settings ClientSettings, credentials Credentials) client.NitroSettings {
	return client.NitroSettings{
		BaseUrl:                   n.GetNodeUrl(settings.UrlScheme),
		Username:                  credentials.Username,
		Password:                  credentials.Password,
		Timeout:                   settings.Timeout,
		ValidateServerCertificate: settings.ValidateServerCertificate,
		LogTlsSecrets:             settings.LogTlsSecrets,
		LogTlsSecretsDestination:  settings.LogTlsSecretsDestination,
	}
}
