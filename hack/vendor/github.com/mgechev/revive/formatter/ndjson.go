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
	"bytes"
	"encoding/json"

	"github.com/mgechev/revive/lint"
)

// NDJSON is an implementation of the Formatter interface
// which formats the errors to NDJSON stream.
type NDJSON struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter.
func (*NDJSON) Name() string {
	return "ndjson"
}

// Format formats the failures gotten from the lint.
func (*NDJSON) Format(failures <-chan lint.Failure, config lint.Config) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for failure := range failures {
		obj := jsonObject{}
		obj.Severity = severity(config, failure)
		obj.Failure = failure
		err := enc.Encode(obj)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
