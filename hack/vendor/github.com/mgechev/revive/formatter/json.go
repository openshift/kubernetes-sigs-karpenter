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

package formatter

import (
	"encoding/json"

	"github.com/mgechev/revive/lint"
)

// JSON is an implementation of the Formatter interface
// which formats the errors to JSON.
type JSON struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter.
func (*JSON) Name() string {
	return "json"
}

// jsonObject defines a JSON object of an failure.
type jsonObject struct {
	Severity     lint.Severity `json:"Severity"`
	lint.Failure `json:",inline"`
}

// Format formats the failures gotten from the lint.
func (*JSON) Format(failures <-chan lint.Failure, config lint.Config) (string, error) {
	var slice []jsonObject
	for failure := range failures {
		obj := jsonObject{}
		obj.Severity = severity(config, failure)
		obj.Failure = failure
		slice = append(slice, obj)
	}
	result, err := json.Marshal(slice)
	if err != nil {
		return "", err
	}
	return string(result), err
}
