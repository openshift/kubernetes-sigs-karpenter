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

package cwe

import (
	"encoding/json"
	"fmt"
)

// Weakness defines a CWE weakness based on http://cwe.mitre.org/data/xsd/cwe_schema_v6.4.xsd
type Weakness struct {
	ID          string
	Name        string
	Description string
}

// SprintURL format the CWE URL
func (w *Weakness) SprintURL() string {
	return fmt.Sprintf("https://cwe.mitre.org/data/definitions/%s.html", w.ID)
}

// SprintID format the CWE ID
func (w *Weakness) SprintID() string {
	id := "0000"
	if w != nil {
		id = w.ID
	}
	return fmt.Sprintf("%s-%s", Acronym, id)
}

// MarshalJSON print only id and URL
func (w *Weakness) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	}{
		ID:  w.ID,
		URL: w.SprintURL(),
	})
}
