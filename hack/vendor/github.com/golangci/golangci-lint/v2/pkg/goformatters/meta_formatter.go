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

package goformatters

import (
	"bytes"
	"fmt"
	"go/format"
	"slices"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/gci"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/gofmt"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/gofumpt"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/goimports"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/golines"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/swaggo"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type MetaFormatter struct {
	log        logutils.Log
	formatters []Formatter
}

func NewMetaFormatter(log logutils.Log, cfg *config.Formatters, runCfg *config.Run) (*MetaFormatter, error) {
	for _, formatter := range cfg.Enable {
		if !IsFormatter(formatter) {
			return nil, fmt.Errorf("invalid formatter %q", formatter)
		}
	}

	m := &MetaFormatter{log: log}

	if slices.Contains(cfg.Enable, gofmt.Name) {
		m.formatters = append(m.formatters, gofmt.New(&cfg.Settings.GoFmt))
	}

	if slices.Contains(cfg.Enable, gofumpt.Name) {
		m.formatters = append(m.formatters, gofumpt.New(&cfg.Settings.GoFumpt, runCfg.Go))
	}

	if slices.Contains(cfg.Enable, goimports.Name) {
		m.formatters = append(m.formatters, goimports.New(&cfg.Settings.GoImports))
	}

	if slices.Contains(cfg.Enable, swaggo.Name) {
		m.formatters = append(m.formatters, swaggo.New())
	}

	// gci is a last because the only goal of gci is to handle imports.
	if slices.Contains(cfg.Enable, gci.Name) {
		formatter, err := gci.New(&cfg.Settings.Gci)
		if err != nil {
			return nil, fmt.Errorf("gci: creating formatter: %w", err)
		}

		m.formatters = append(m.formatters, formatter)
	}

	// golines calls `format.Source()` internally so no need to format after it.
	if slices.Contains(cfg.Enable, golines.Name) {
		m.formatters = append(m.formatters, golines.New(&cfg.Settings.GoLines))
	}

	return m, nil
}

func (m *MetaFormatter) Format(filename string, src []byte) []byte {
	if len(m.formatters) == 0 {
		data, err := format.Source(src)
		if err != nil {
			m.log.Warnf("(fmt) formatting file %s: %v", filename, err)
			return src
		}

		return data
	}

	data := bytes.Clone(src)

	for _, formatter := range m.formatters {
		formatted, err := formatter.Format(filename, data)
		if err != nil {
			m.log.Warnf("(%s) formatting file %s: %v", formatter.Name(), filename, err)
			continue
		}

		data = formatted
	}

	return data
}

func IsFormatter(name string) bool {
	return slices.Contains([]string{gofmt.Name, gofumpt.Name, goimports.Name, gci.Name, golines.Name, swaggo.Name}, name)
}
