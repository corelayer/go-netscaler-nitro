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

func TestNitroRequestParams_GetResourceName(t *testing.T) {
	var tests = []struct {
		r    NitroResourceSelector
		want string
	}{
		{UnknownResource, ""},
		{SystemBackup, "systembackup"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.r.GetNitroResourceName())
		t.Run(testName, func(t *testing.T) {
			p := NitroRequestParams{Resource: tt.r}
			result := p.GetResourceName()
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestNitroRequestParams_GetMethod(t *testing.T) {
	var tests = []struct {
		params NitroRequestParams
		want   string
	}{
		{NitroRequestParams{Method: NitroUnknownMethod}, ""},
		{NitroRequestParams{Method: NitroGetMethod}, "GET"},
		{NitroRequestParams{Method: NitroPostMethod}, "POST"},
		{NitroRequestParams{Method: 100}, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.params.Method.String())
		t.Run(testName, func(t *testing.T) {
			result := tt.params.GetMethod()
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestNitroRequestParams_GetUrl(t *testing.T) {
	var tests = []struct {
		params NitroRequestParamsReader
		want   string
	}{
		{NitroRequestParams{Resource: UnknownResource, Type: NitroUnknownRequest}, ""},
		{NitroRequestParams{Resource: UnknownResource, Type: 300}, ""},
		{NitroRequestParams{Resource: UnknownResource, Type: NitroConfigRequest}, "/nitro/v1/config/"},
		{NitroRequestParams{Resource: SystemBackup, Type: NitroConfigRequest}, "/nitro/v1/config/systembackup"},
		{NitroRequestParams{Resource: UnknownResource, Type: NitroStatsRequest}, "/nitro/v1/stats/"},
		{NitroRequestParams{Resource: SystemBackup, Type: NitroStatsRequest}, "/nitro/v1/stats/systembackup"},
		{NitroRequestParams{
			Resource:  SystemBackup,
			Type:      NitroConfigRequest,
			Arguments: map[string]string{"key1": "value1"}},
			"/nitro/v1/config/systembackup?args=key1:value1"}, {NitroRequestParams{
			Resource:  SystemBackup,
			Type:      NitroConfigRequest,
			Arguments: map[string]string{"fileLocation": "/var/ns_sys_backup"}},
			"/nitro/v1/config/systembackup?args=fileLocation:%2Fvar%2Fns_sys_backup"},
		{NitroRequestParams{
			Resource:  SystemBackup,
			Type:      NitroConfigRequest,
			Arguments: map[string]string{"key1": "value1", "key2": "value2"}},
			"/nitro/v1/config/systembackup?args=key1:value1,key2:value2"},
		{NitroRequestParams{
			Resource:  SystemBackup,
			Type:      NitroConfigRequest,
			Arguments: map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"}},
			"/nitro/v1/config/systembackup?args=key1:value1,key2:value2,key3:value3"},
		{NitroRequestParams{
			Resource:  SystemBackup,
			Type:      NitroConfigRequest,
			Arguments: map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"},
			Filter:    map[string]string{"key1": "value1"}},
			"/nitro/v1/config/systembackup?args=key1:value1,key2:value2,key3:value3&filter=key1:value1"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.params.GetResourceName())
		t.Run(testName, func(t *testing.T) {
			result := tt.params.GetUrlPathAndQuery()
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestGetNitroUrlPath(t *testing.T) {
	var tests = []struct {
		params NitroRequestParams
		want   string
	}{
		{NitroRequestParams{Resource: UnknownResource, Type: NitroUnknownRequest}, ""},
		{NitroRequestParams{Resource: UnknownResource, Type: NitroConfigRequest}, "/nitro/v1/config/"},
		{NitroRequestParams{Resource: UnknownResource, Type: NitroStatsRequest}, "/nitro/v1/stats/"},
		{NitroRequestParams{Resource: UnknownResource, Type: 300}, ""},
		{NitroRequestParams{Resource: SystemBackup, Type: NitroUnknownRequest}, ""},
		{NitroRequestParams{Resource: SystemBackup, Type: NitroConfigRequest}, "/nitro/v1/config/systembackup"},
		{NitroRequestParams{Resource: SystemBackup, Type: NitroStatsRequest}, "/nitro/v1/stats/systembackup"},
		{NitroRequestParams{Resource: SystemBackup, Type: 300}, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.params.Type.String())
		t.Run(testName, func(t *testing.T) {
			result := tt.params.GetUrlPath()
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestGetUrlQueryStringSeparator(t *testing.T) {
	var tests = []struct {
		length int
		want   string
	}{
		{0, "?"},
		{1, "&"},
		{100, "&"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.length)
		t.Run(testName, func(t *testing.T) {
			result := getUrlQueryStringSeparator(tt.length)
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestGetUrlQueryMapStringSeparator(t *testing.T) {
	var tests = []struct {
		index     int
		lastIndex int
		want      string
	}{
		{0, 10, ","},
		{10, 10, ""},
		{100, 10, ""},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.index, tt.lastIndex)
		t.Run(testName, func(t *testing.T) {
			result := getUrlQueryMapStringSeparator(tt.index, tt.lastIndex)
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}

func TestGetQueryMapEntriesAsString(t *testing.T) {
	var tests = []struct {
		entries map[string]string
		want    string
	}{
		{entries: map[string]string{"key1": "value1"}, want: "key1:value1"},
		{entries: map[string]string{"key1": "value1", "key2": "value2"}, want: "key1:value1,key2:value2"},
		{entries: map[string]string{"key3": "value3", "key1": "value1", "key2": "value2"}, want: "key1:value1,key2:value2,key3:value3"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.want)
		t.Run(testName, func(t *testing.T) {
			result := getQueryMapEntriesAsString(tt.entries)
			if result != tt.want {
				t.Errorf("result: %s - expected: %s", result, tt.want)
			}
		})
	}
}
