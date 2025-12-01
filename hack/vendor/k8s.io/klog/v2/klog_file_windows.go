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

//go:build windows
// +build windows

package klog

import (
	"os"
	"strings"
)

func getUserName() string {
	userNameOnce.Do(func() {
		// On Windows, the Go 'user' package requires netapi32.dll.
		// This affects Windows Nano Server:
		//   https://github.com/golang/go/issues/21867
		// Fallback to using environment variables.
		u := os.Getenv("USERNAME")
		if len(u) == 0 {
			return
		}
		// Sanitize the USERNAME since it may contain filepath separators.
		u = strings.Replace(u, `\`, "_", -1)

		// user.Current().Username normally produces something like 'USERDOMAIN\USERNAME'
		d := os.Getenv("USERDOMAIN")
		if len(d) != 0 {
			userName = d + "_" + u
		} else {
			userName = u
		}
	})

	return userName
}
