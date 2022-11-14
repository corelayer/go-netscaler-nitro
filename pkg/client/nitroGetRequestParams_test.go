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
