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

package report

type Warning struct {
	Tag  string `json:",omitempty"`
	Text string
}

type LinterData struct {
	Name    string
	Enabled bool `json:",omitempty"`
}

type Data struct {
	Warnings []Warning    `json:",omitempty"`
	Linters  []LinterData `json:",omitempty"`
	Error    string       `json:",omitempty"`
}

func (d *Data) AddLinter(name string, enabled bool) {
	d.Linters = append(d.Linters, LinterData{
		Name:    name,
		Enabled: enabled,
	})
}
