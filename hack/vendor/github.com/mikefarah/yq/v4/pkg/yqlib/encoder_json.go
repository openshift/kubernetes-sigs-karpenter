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

//go:build !yq_nojson

package yqlib

import (
	"bytes"
	"io"

	"github.com/goccy/go-json"
)

type jsonEncoder struct {
	prefs        JsonPreferences
	indentString string
}

func NewJSONEncoder(prefs JsonPreferences) Encoder {
	var indentString = ""

	for index := 0; index < prefs.Indent; index++ {
		indentString = indentString + " "
	}

	return &jsonEncoder{prefs, indentString}
}

func (je *jsonEncoder) CanHandleAliases() bool {
	return false
}

func (je *jsonEncoder) PrintDocumentSeparator(_ io.Writer) error {
	return nil
}

func (je *jsonEncoder) PrintLeadingContent(_ io.Writer, _ string) error {
	return nil
}

func (je *jsonEncoder) Encode(writer io.Writer, node *CandidateNode) error {
	log.Debugf("I need to encode %v", NodeToString(node))
	log.Debugf("kids %v", len(node.Content))

	if node.Kind == ScalarNode && je.prefs.UnwrapScalar {
		return writeString(writer, node.Value+"\n")
	}

	destination := writer
	tempBuffer := bytes.NewBuffer(nil)
	if je.prefs.ColorsEnabled {
		destination = tempBuffer
	}

	var encoder = json.NewEncoder(destination)
	encoder.SetEscapeHTML(false) // do not escape html chars e.g. &, <, >
	encoder.SetIndent("", je.indentString)

	err := encoder.Encode(node)
	if err != nil {
		return err
	}
	if je.prefs.ColorsEnabled {
		return colorizeAndPrint(tempBuffer.Bytes(), writer)
	}
	return nil
}
