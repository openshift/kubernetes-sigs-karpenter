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

package logger

import "os"

type Logger interface {
	Printf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

func DebugEnabled() bool {
	d := os.Getenv("SWAGGER_DEBUG")
	if d != "" && d != "false" && d != "0" {
		return true
	}
	d = os.Getenv("DEBUG")
	if d != "" && d != "false" && d != "0" {
		return true
	}
	return false
}
