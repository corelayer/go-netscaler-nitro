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
)

func TestClient_getConfigUrl(t *testing.T) {
	c := Client{}

	var tests = []struct {
		baseUrl string
		want    string
	}{
		{"http://127.0.0.1", "http://127.0.0.1/nitro/v1/config/"},
		{"https://127.0.0.1", "https://127.0.0.1/nitro/v1/config/"},
		{"http://targetNode", "http://targetNode/nitro/v1/config/"},
		{"https://targetNode", "https://targetNode/nitro/v1/config/"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.baseUrl)
		t.Run(testName, func(t *testing.T) {
			c.baseUrl = tt.baseUrl
			result := c.getConfigUrl()
			if result != tt.want {
				t.Errorf("getConfigUrl(%s) = %s, expected: %s", tt.baseUrl, result, tt.want)
			}
		})
	}
}

func TestClient_getStatsUrl(t *testing.T) {
	c := Client{}

	var tests = []struct {
		baseUrl string
		want    string
	}{
		{"http://127.0.0.1", "http://127.0.0.1/nitro/v1/stats/"},
		{"https://127.0.0.1", "https://127.0.0.1/nitro/v1/stats/"},
		{"http://targetNode", "http://targetNode/nitro/v1/stats/"},
		{"https://targetNode", "https://targetNode/nitro/v1/stats/"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.baseUrl)
		t.Run(testName, func(t *testing.T) {
			c.baseUrl = tt.baseUrl
			result := c.getStatsUrl()
			if result != tt.want {
				t.Errorf("getStatsUrl(%s) = %s, expected: %s", tt.baseUrl, result, tt.want)
			}
		})
	}
}
