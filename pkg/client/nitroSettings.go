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
	"io"
	"os"
	"strconv"
	"time"
)

type NitroSettings struct {
	BaseUrl                   string
	Username                  string
	Password                  string
	Timeout                   int
	ValidateServerCertificate bool
	LogTlsSecrets             bool
	LogTlsSecretsDestination  string
}

func (s *NitroSettings) GetTlsSecretLogWriter() (io.Writer, error) {
	if !s.LogTlsSecrets {
		return nil, nil
	}

	tlsLog, err := os.OpenFile(s.LogTlsSecretsDestination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	return tlsLog, nil

}

//GetTimeoutDuration Returns a time.Duration based on the set timout in seconds
func (s *NitroSettings) GetTimeoutDuration() (time.Duration, error) {
	return time.ParseDuration(strconv.Itoa(s.Timeout) + "s")
}
