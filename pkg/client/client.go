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

package client

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	baseUrl string
	client  *http.Client
	log     *log.Logger
}

func (c *Client) getConfigUrl() string {
	return strings.Trim(c.baseUrl, " /") + "/nitro/v1/config/"
}

func (c *Client) getStatsUrl() string {
	return strings.Trim(c.baseUrl, " /") + "/nitro/v1/stats/"
}

func newClient(node NodeReader, settings Settings, logger *log.Logger) (*Client, error) {
	baseUrl := node.GetNodeUrl(settings.Scheme)

	tlsLog, err := settings.GetTlsSecretLogWriter()
	if err != nil {
		return nil, err
	}

	if tlsLog != nil {
		logger.Printf("WARNING, exporting TLS Secrets to %s\n", settings.LogTlsSecretsDestination)

	}

	timeout, err := settings.GetTimeoutDuration()

	return &Client{
		baseUrl: baseUrl,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					KeyLogWriter:       tlsLog,
					InsecureSkipVerify: settings.InsecureSkipVerify,
				},
				Proxy: http.ProxyFromEnvironment,
			},
			Timeout: timeout,
		},
		log: logger,
	}, nil
}
