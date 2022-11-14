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

func TestNitroResource_GetNitroResourceName(t *testing.T) {
	var tests = []struct {
		resource NitroResource
		want     string
	}{
		{UnknownResource, ""},
		{SystemBackup, "systembackup"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.resource.GetNitroResourceName())
		t.Run(testName, func(t *testing.T) {
			result := tt.resource.GetNitroResourceName()
			if result != tt.want {
				t.Errorf("result: %s, expected: %s", result, tt.want)
			}
		})
	}
}
