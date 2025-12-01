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

package migrate

import (
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/ptr"
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versionone"
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versiontwo"
)

func toLinters(old *versionone.Config) versiontwo.Linters {
	enable, disable := ProcessEffectiveLinters(old.Linters)

	return versiontwo.Linters{
		Default:    getDefaultName(old.Linters),
		Enable:     onlyLinterNames(convertStaticcheckLinterNames(enable)),
		Disable:    onlyLinterNames(convertDisabledStaticcheckLinterNames(disable)),
		FastOnly:   nil,
		Settings:   toLinterSettings(old.LintersSettings),
		Exclusions: toExclusions(old),
	}
}

func getDefaultName(old versionone.Linters) *string {
	switch {
	case ptr.Deref(old.DisableAll):
		return ptr.Pointer("none")
	case ptr.Deref(old.EnableAll):
		return ptr.Pointer("all")
	default:
		return nil // standard is the default
	}
}
