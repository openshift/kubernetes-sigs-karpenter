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

package config

type Formats struct {
	Text        Text         `mapstructure:"text"`
	JSON        SimpleFormat `mapstructure:"json"`
	Tab         Tab          `mapstructure:"tab"`
	HTML        SimpleFormat `mapstructure:"html"`
	Checkstyle  SimpleFormat `mapstructure:"checkstyle"`
	CodeClimate SimpleFormat `mapstructure:"code-climate"`
	JUnitXML    JUnitXML     `mapstructure:"junit-xml"`
	TeamCity    SimpleFormat `mapstructure:"teamcity"`
	Sarif       SimpleFormat `mapstructure:"sarif"`
}

func (f *Formats) IsEmpty() bool {
	formats := []SimpleFormat{
		f.Text.SimpleFormat,
		f.JSON,
		f.Tab.SimpleFormat,
		f.HTML,
		f.Checkstyle,
		f.CodeClimate,
		f.JUnitXML.SimpleFormat,
		f.TeamCity,
		f.Sarif,
	}

	for _, format := range formats {
		if format.Path != "" {
			return false
		}
	}

	return true
}

type SimpleFormat struct {
	Path string `mapstructure:"path"`
}

type Text struct {
	SimpleFormat    `mapstructure:",squash"`
	PrintLinterName bool `mapstructure:"print-linter-name"`
	PrintIssuedLine bool `mapstructure:"print-issued-lines"`
	Colors          bool `mapstructure:"colors"`
}

type Tab struct {
	SimpleFormat    `mapstructure:",squash"`
	PrintLinterName bool `mapstructure:"print-linter-name"`
	Colors          bool `mapstructure:"colors"`
}

type JUnitXML struct {
	SimpleFormat `mapstructure:",squash"`
	Extended     bool `mapstructure:"extended"`
}
