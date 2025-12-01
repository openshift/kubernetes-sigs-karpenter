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

// Package unqueryvet provides a Go static analysis tool that detects SELECT * usage
package unqueryvet

import (
	"golang.org/x/tools/go/analysis"

	"github.com/MirrexOne/unqueryvet/internal/analyzer"
	"github.com/MirrexOne/unqueryvet/pkg/config"
)

// Analyzer is the main unqueryvet analyzer instance
// This is the primary export that golangci-lint will use
var Analyzer = analyzer.NewAnalyzer()

// New creates a new instance of the unqueryvet analyzer
func New() *analysis.Analyzer {
	return Analyzer
}

// NewWithConfig creates a new analyzer instance with custom configuration
// This is the recommended way to use unqueryvet with custom settings
func NewWithConfig(cfg *config.UnqueryvetSettings) *analysis.Analyzer {
	if cfg == nil {
		return Analyzer
	}
	return analyzer.NewAnalyzerWithSettings(*cfg)
}
