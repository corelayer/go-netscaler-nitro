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
	"fmt"
	"testing"
	"time"
)

func TestSettings_GetTimeoutDuration(t *testing.T) {
	s := Settings{
		Scheme:                   HTTP,
		InsecureSkipVerify:       false,
		Timeout:                  10,
		LogTlsSecrets:            false,
		LogTlsSecretsDestination: "",
	}

	var tests = []struct {
		timeout int
		want    time.Duration
	}{
		{
			timeout: 10,
			want:    10 * time.Second,
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.timeout)
		t.Run(testName, func(t *testing.T) {
			s.Timeout = tt.timeout
			result, _ := s.GetTimeoutDuration()

			if result != tt.want {
				t.Errorf("GetTimeoutDuration(%d) = %d, expected %d", tt.timeout, result, tt.want)
			}
		})
	}
}
