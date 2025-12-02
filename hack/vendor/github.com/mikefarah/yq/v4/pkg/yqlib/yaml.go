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

package yqlib

type YamlPreferences struct {
	Indent                      int
	ColorsEnabled               bool
	LeadingContentPreProcessing bool
	PrintDocSeparators          bool
	UnwrapScalar                bool
	EvaluateTogether            bool
}

func NewDefaultYamlPreferences() YamlPreferences {
	return YamlPreferences{
		Indent:                      2,
		ColorsEnabled:               false,
		LeadingContentPreProcessing: true,
		PrintDocSeparators:          true,
		UnwrapScalar:                true,
		EvaluateTogether:            false,
	}
}

func (p *YamlPreferences) Copy() YamlPreferences {
	return YamlPreferences{
		Indent:                      p.Indent,
		ColorsEnabled:               p.ColorsEnabled,
		LeadingContentPreProcessing: p.LeadingContentPreProcessing,
		PrintDocSeparators:          p.PrintDocSeparators,
		UnwrapScalar:                p.UnwrapScalar,
		EvaluateTogether:            p.EvaluateTogether,
	}
}

var ConfiguredYamlPreferences = NewDefaultYamlPreferences()
