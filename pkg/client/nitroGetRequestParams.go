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
	"sort"
	"strings"
)

type NitroGetRequestParams struct {
	Arguments  map[string]string
	Filter     map[string]string
	Attributes map[string]string
}

func (p NitroGetRequestParams) GetNitroRequestUrlQueryString() string {
	var output strings.Builder

	output.WriteString(buildUrlQueryMapString(output.Len(), "args=", p.Arguments))
	output.WriteString(buildUrlQueryMapString(output.Len(), "filter=", p.Filter))
	output.WriteString(buildUrlQueryMapString(output.Len(), "attrs=", p.Attributes))

	return output.String()
}

func buildUrlQueryMapString(urlQueryLength int, prefix string, queryMap map[string]string) string {
	if len(queryMap) == 0 {
		return ""
	}

	var output strings.Builder
	output.WriteString(getUrlQueryStringSeparator(urlQueryLength))
	output.WriteString(prefix)
	output.WriteString(getQueryMapEntriesAsString(queryMap))
	return output.String()
}

func getQueryMapSortedKeys(queryMap map[string]string) []string {
	keys := make([]string, 0, len(queryMap))
	for k, _ := range queryMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func getQueryMapEntriesAsString(queryMap map[string]string) string {
	var output strings.Builder

	keys := getQueryMapSortedKeys(queryMap)
	lastIndex := len(keys) - 1

	for index, key := range keys {
		value := queryMap[key]
		output.WriteString(key + ":" + value + getUrlQueryMapStringSeparator(index, lastIndex))
	}

	return output.String()
}

func getUrlQueryMapStringSeparator(index int, lastIndex int) string {
	if index < lastIndex {
		return ","
	} else {
		return ""
	}
}

func getUrlQueryStringSeparator(length int) string {
	if length == 0 {
		return "?"
	}
	return "&"
}
