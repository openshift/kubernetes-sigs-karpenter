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

import (
	"fmt"
	"slices"
	"strings"

	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type LogWrapper struct {
	rd      *Data
	tags    []string
	origLog logutils.Log
}

func NewLogWrapper(log logutils.Log, reportData *Data) *LogWrapper {
	return &LogWrapper{
		rd:      reportData,
		origLog: log,
	}
}

func (lw LogWrapper) Fatalf(format string, args ...any) {
	lw.origLog.Fatalf(format, args...)
}

func (lw LogWrapper) Panicf(format string, args ...any) {
	lw.origLog.Panicf(format, args...)
}

func (lw LogWrapper) Errorf(format string, args ...any) {
	lw.origLog.Errorf(format, args...)
	lw.rd.Error = fmt.Sprintf(format, args...)
}

func (lw LogWrapper) Warnf(format string, args ...any) {
	lw.origLog.Warnf(format, args...)
	w := Warning{
		Tag:  strings.Join(lw.tags, "/"),
		Text: fmt.Sprintf(format, args...),
	}

	lw.rd.Warnings = append(lw.rd.Warnings, w)
}

func (lw LogWrapper) Infof(format string, args ...any) {
	lw.origLog.Infof(format, args...)
}

func (lw LogWrapper) Child(name string) logutils.Log {
	c := lw
	c.origLog = lw.origLog.Child(name)
	c.tags = slices.Clone(lw.tags)
	c.tags = append(c.tags, name)
	return c
}

func (lw LogWrapper) SetLevel(level logutils.LogLevel) {
	lw.origLog.SetLevel(level)
}

func (lw LogWrapper) GoString() string {
	return fmt.Sprintf("lw: %+v, orig log: %#v", lw, lw.origLog)
}
