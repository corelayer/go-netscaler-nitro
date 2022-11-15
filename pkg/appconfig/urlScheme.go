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

package appconfig

type UrlScheme int

const (
	UnknownUrlScheme UrlScheme = 0
	HTTP                       = 1
	HTTPS                      = 2
)

//go:generate stringer -type=UrlScheme -output=urlScheme_string.go

//GetUrlScheme Returns the URL prefix for the selected UrlScheme
func (s UrlScheme) GetUrlScheme() string {
	switch s {
	case UnknownUrlScheme:
		return ""
	case HTTP:
		return "http://"
	case HTTPS:
		return "https://"
	default:
		return ""
	}
}
