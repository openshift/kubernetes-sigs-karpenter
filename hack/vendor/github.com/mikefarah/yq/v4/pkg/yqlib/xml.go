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

type XmlPreferences struct {
	Indent          int
	AttributePrefix string
	ContentName     string
	StrictMode      bool
	KeepNamespace   bool
	UseRawToken     bool
	ProcInstPrefix  string
	DirectiveName   string
	SkipProcInst    bool
	SkipDirectives  bool
}

func NewDefaultXmlPreferences() XmlPreferences {
	return XmlPreferences{
		Indent:          2,
		AttributePrefix: "+@",
		ContentName:     "+content",
		StrictMode:      false,
		KeepNamespace:   true,
		UseRawToken:     true,
		ProcInstPrefix:  "+p_",
		DirectiveName:   "+directive",
		SkipProcInst:    false,
		SkipDirectives:  false,
	}
}

func (p *XmlPreferences) Copy() XmlPreferences {
	return XmlPreferences{
		Indent:          p.Indent,
		AttributePrefix: p.AttributePrefix,
		ContentName:     p.ContentName,
		StrictMode:      p.StrictMode,
		KeepNamespace:   p.KeepNamespace,
		UseRawToken:     p.UseRawToken,
		ProcInstPrefix:  p.ProcInstPrefix,
		DirectiveName:   p.DirectiveName,
		SkipProcInst:    p.SkipProcInst,
		SkipDirectives:  p.SkipDirectives,
	}
}

var ConfiguredXMLPreferences = NewDefaultXmlPreferences()
