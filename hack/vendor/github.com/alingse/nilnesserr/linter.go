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

package nilnesserr

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

const (
	linterName = "nilnesserr"
	linterDoc  = `Reports constructs that checks for err != nil, but returns a different nil value error.
Powered by nilness and nilerr.`

	linterReturnMessage       = "return a nil value error after check error"
	linterCallMessage         = "call function with a nil value error after check error"
	linterVariadicCallMessage = "call variadic function with a nil value error after check error"
)

type LinterSetting struct{}

func NewAnalyzer(setting LinterSetting) (*analysis.Analyzer, error) {
	a, err := newAnalyzer(setting)
	if err != nil {
		return nil, err
	}

	return &analysis.Analyzer{
		Name: linterName,
		Doc:  linterDoc,
		Run:  a.run,
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}, nil
}

type analyzer struct {
	setting LinterSetting
}

func newAnalyzer(setting LinterSetting) (*analyzer, error) {
	a := &analyzer{setting: setting}

	return a, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	_, _ = a.checkNilnesserr(pass)

	return nil, nil
}
