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

func TestNode_GetNodeUrl(t *testing.T) {
	n := Node{
		Name:    "node",
		Address: "",
	}

	var tests = []struct {
		address string
		scheme  Scheme
		want    string
	}{
		{
			address: "node",
			scheme:  HTTP,
			want:    "http://node",
		}, {
			address: "node",
			scheme:  HTTPS,
			want:    "https://node",
		},
		{
			address: "127.0.0.1",
			scheme:  HTTP,
			want:    "http://127.0.0.1",
		},
		{
			address: "127.0.0.1",
			scheme:  HTTPS,
			want:    "https://127.0.0.1",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s, %v", tt.address, tt.scheme)
		t.Run(testName, func(t *testing.T) {
			n.Address = tt.address
			result := n.GetNodeUrl(tt.scheme)
			want := tt.want

			if result != want {
				t.Errorf("GetConnectionString(%s, %v) = %s, expected: %s", tt.address, tt.scheme, result, want)
			}
		})
	}
}
