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

package filters

import "fmt"

// invalidFilter indicates that the provided filter or its value is invalid
type invalidFilter struct {
	Filter string
	Value  []string
}

func (e invalidFilter) Error() string {
	msg := "invalid filter"
	if e.Filter != "" {
		msg += " '" + e.Filter
		if e.Value != nil {
			msg = fmt.Sprintf("%s=%s", msg, e.Value)
		}
		msg += "'"
	}
	return msg
}

// InvalidParameter marks this error as ErrInvalidParameter
func (e invalidFilter) InvalidParameter() {}
