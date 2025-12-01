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

package reporters

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/onsi/ginkgo/v2/types"
)

// GenerateJSONReport produces a JSON-formatted report at the passed in destination
func GenerateJSONReport(report types.Report, destination string) error {
	if err := os.MkdirAll(path.Dir(destination), 0770); err != nil {
		return err
	}
	f, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode([]types.Report{
		report,
	})
	if err != nil {
		return err
	}
	return nil
}

// MergeJSONReports produces a single JSON-formatted report at the passed in destination by merging the JSON-formatted reports provided in sources
// It skips over reports that fail to decode but reports on them via the returned messages []string
func MergeAndCleanupJSONReports(sources []string, destination string) ([]string, error) {
	messages := []string{}
	allReports := []types.Report{}
	for _, source := range sources {
		reports := []types.Report{}
		data, err := os.ReadFile(source)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Could not open %s:\n%s", source, err.Error()))
			continue
		}
		err = json.Unmarshal(data, &reports)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Could not decode %s:\n%s", source, err.Error()))
			continue
		}
		os.Remove(source)
		allReports = append(allReports, reports...)
	}

	if err := os.MkdirAll(path.Dir(destination), 0770); err != nil {
		return messages, err
	}
	f, err := os.Create(destination)
	if err != nil {
		return messages, err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	err = enc.Encode(allReports)
	if err != nil {
		return messages, err
	}
	return messages, nil
}
