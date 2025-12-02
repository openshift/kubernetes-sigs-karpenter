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

package sa1014

import (
	"fmt"
	"go/types"

	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/internal/passes/buildir"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA1014",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(checkUnmarshalPointerRules),
	},
	Doc: &lint.RawDocumentation{
		Title:    `Non-pointer value passed to \'Unmarshal\' or \'Decode\'`,
		Since:    "2017.1",
		Severity: lint.SeverityError,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var checkUnmarshalPointerRules = map[string]callcheck.Check{
	"encoding/xml.Unmarshal":                unmarshalPointer("xml.Unmarshal", 1),
	"(*encoding/xml.Decoder).Decode":        unmarshalPointer("Decode", 0),
	"(*encoding/xml.Decoder).DecodeElement": unmarshalPointer("DecodeElement", 0),
	"encoding/json.Unmarshal":               unmarshalPointer("json.Unmarshal", 1),
	"(*encoding/json.Decoder).Decode":       unmarshalPointer("Decode", 0),
}

func unmarshalPointer(name string, arg int) callcheck.Check {
	return func(call *callcheck.Call) {
		if !Pointer(call.Args[arg].Value) {
			call.Args[arg].Invalid(fmt.Sprintf("%s expects to unmarshal into a pointer, but the provided value is not a pointer", name))
		}
	}
}

func Pointer(v callcheck.Value) bool {
	switch v.Value.Type().Underlying().(type) {
	case *types.Pointer, *types.Interface:
		return true
	}
	return false
}
