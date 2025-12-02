/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package misspell

import (
	"regexp"
)

// Regexp for URL https://mathiasbynens.be/demo/url-regex
//
// original @imme_emosol (54 chars) has trouble with dashes in hostname
// @(https?|ftp)://(-\.)?([^\s/?\.#-]+\.?)+(/[^\s]*)?$@iS.
var reURL = regexp.MustCompile(`(?i)(https?|ftp)://(-\.)?([^\s/?.#]+\.?)+(/\S*)?`)

// StripURL attempts to replace URLs with blank spaces, e.g.
//
//	"xxx http://foo.com/ yyy -> "xxx          yyyy".
func StripURL(s string) string {
	return reURL.ReplaceAllStringFunc(s, replaceWithBlanks)
}
