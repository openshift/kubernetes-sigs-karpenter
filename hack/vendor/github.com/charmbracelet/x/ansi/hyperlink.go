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

package ansi

import "strings"

// SetHyperlink returns a sequence for starting a hyperlink.
//
//	OSC 8 ; Params ; Uri ST
//	OSC 8 ; Params ; Uri BEL
//
// To reset the hyperlink, omit the URI.
//
// See: https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
func SetHyperlink(uri string, params ...string) string {
	var p string
	if len(params) > 0 {
		p = strings.Join(params, ":")
	}
	return "\x1b]8;" + p + ";" + uri + "\x07"
}

// ResetHyperlink returns a sequence for resetting the hyperlink.
//
// This is equivalent to SetHyperlink("", params...).
//
// See: https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
func ResetHyperlink(params ...string) string {
	return SetHyperlink("", params...)
}
