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
	"io"
	"log"
	"net/http"
)

type NitroClient struct {
	client   *http.Client
	settings NitroSettings
	log      *log.Logger
}

func NewNitroClient(settings NitroSettings, logger *log.Logger) (*NitroClient, error) {
	tlsLog, err := settings.GetTlsSecretLogWriter()
	if err != nil {
		return nil, err
	}

	if tlsLog != nil {
		logger.Printf("WARNING, exporting TLS Secrets to %s\n", settings.LogTlsSecretsDestination)
	}

	timeout, err := settings.GetTimeoutDuration()

	return &NitroClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					KeyLogWriter:       tlsLog,
					InsecureSkipVerify: !settings.ValidateServerCertificate,
				},
				Proxy: http.ProxyFromEnvironment,
			},
			Timeout: timeout,
		},
		settings: settings,
		log:      logger,
	}, nil
}

func (c NitroClient) createRequest(params NitroRequestParamsReader, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(params.GetMethod(), params.GetUrlPathAndQuery(), body)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("X-NITRO-USER", c.settings.Username)
	request.Header.Set("X-NITRO-PASS", c.settings.Password)

	return request, err
}
